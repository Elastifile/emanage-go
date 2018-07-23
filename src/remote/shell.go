package remote

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-errors/errors"
	"golang.org/x/crypto/ssh"

	"helputils"
)

type ShellOpts struct {
	Host     string // IP address of the remote host
	TermCols int
	TermRows int
}

type ShellSession struct {
	remote       *Remote
	session      *ssh.Session
	stdinWriter  io.WriteCloser
	stdoutReader io.Reader
	stderrReader io.Reader
	stdoutChan   chan string
	stderrChan   chan string
	backChan     chan error
	prompts      []string
}

func terminalModes() ssh.TerminalModes {
	return ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
}

// NewShellSession opens a shell on a remote host and prepares it for accepting commands
func NewShellSession(opts ShellOpts) (s ShellSession, err error) {
	const buffSize = 32 * 1024
	if opts.TermCols == 0 {
		opts.TermCols = 400 // Default terminal width
	}

	if opts.TermRows == 0 {
		opts.TermRows = 80 // Default terminal height
	}

	s.stdoutChan = make(chan string)
	s.stderrChan = make(chan string)
	s.backChan = make(chan error)

	// Create remote session
	if s.remote, err = NewRemote(opts.Host); err != nil {
		return ShellSession{}, errors.Errorf("failed to open connection to host %v. err=%v", opts.Host, err)
	}

	if s.session, err = s.remote.NewSession(); err != nil {
		return ShellSession{}, errors.Errorf("failed to create session on host %v. err=%v", opts.Host, err)
	}

	// Configure the terminal
	termModes := terminalModes()
	if err = s.session.RequestPty("xterm", opts.TermRows, opts.TermCols, termModes); err != nil {
		err = errors.Errorf("failed to request PTY on host %v (opts: %v, modes: %v). err=%v",
			opts.Host, opts, termModes, err)
		s.session.Close()
		return ShellSession{}, err
	}

	// Bind to the process' stdin/out/err
	if s.stdinWriter, err = s.session.StdinPipe(); err != nil {
		return ShellSession{}, errors.Errorf("failed to assign session's stdinWriter")
	}
	if s.stdoutReader, err = s.session.StdoutPipe(); err != nil {
		return ShellSession{}, errors.Errorf("failed to assign session's stdoutReader")
	}
	if s.stderrReader, err = s.session.StderrPipe(); err != nil {
		return ShellSession{}, errors.Errorf("failed to assign session's stderr")
	}

	// Start shell
	if err = s.session.Shell(); err != nil {
		return ShellSession{}, errors.Errorf("failed to start shell on host %v. err=%v", opts.Host, err)
	}

	// Start reading from stdout/stderr continuously (should start before the first WaitFor*Prompt*)
	go s.streamFromPipeRoutine(&s.stdoutReader, s.stdoutChan, buffSize)
	go s.streamFromPipeRoutine(&s.stderrReader, s.stderrChan, buffSize)

	if _, _, err = s.WaitForDefaultPrompt(); err != nil {
		return ShellSession{}, errors.Errorf("didn't get prompt after login on host %v. err=%v",
			opts.Host, err)
	}

	// Set custom, session-unique prompt
	s.prompts = make([]string, 1)
	s.prompts[0] = fmt.Sprintf("PROMPT-%s# ", helputils.RandString(5))
	// Replace hash with its special representation, so the command is not recognized as the prompt by WaitForPrompt
	cmd := fmt.Sprintf("export PS1='%s'", strings.Replace(s.prompts[0], "#", "\\$", -1))
	if _, _, _, err := s.Run(cmd); err != nil {
		return ShellSession{}, errors.Errorf("failed to set custom prompt (%v) on host %v. err=%v",
			s.prompts, opts.Host, err)
	}

	return s, err
}

// Close closes the shell session and frees the resources
func (s *ShellSession) Close() error {
	if s.session != nil {
		if err := s.session.Close(); err != nil {
			return errors.Errorf("failed to close session on host %v", s.remote.address)
		}
	}

	if s.remote != nil {
		if err := s.remote.Close(); err != nil {
			return errors.Errorf("failed to close remote connection to host %v", s.remote.address)
		}
	}

	return nil
}

// streamFromPipeRoutine reads from a pipe until EOF and feeds the results back via channels
// streamFromPipeRoutine is intended to be used as a goroutine running for the duration of the session
// Example: go streamFromPipeRoutine(&s.Stdout, s.stdoutChan, 32*1024)
func (s *ShellSession) streamFromPipeRoutine(pipe *io.Reader, ch chan string, buffSize int) {
	bufReader := bufio.NewReader(*pipe)
	buff := make([]byte, buffSize)
	var (
		bytesRead int
		err       error
	)

	for done := false; !done; {
		bytesRead, err = bufReader.Read(buff)
		if bytesRead > 0 {
			ch <- string(buff[:bytesRead])
		}
		if err != nil {
			if err != io.EOF {
				logger.Info("Read returned an unexpected error", "err", err, "bytesRead", bytesRead)
				logger.Debug("Data received up to read error", "buff", string(buff[:bytesRead]))
				s.backChan <- err
			}
			done = true
		}
	}
}

// getPromptMatchId looks for prompt in text and returns the matching prompt's id
// match will be -1, and an error will be set if there was no match
func getPromptMatchId(text string, prompts []string) (match int, err error) {
	match = -1
	for i, prompt := range prompts {
		if strings.Contains(text, string(prompt)) {
			match = i
			break
		}
	}
	if match == -1 {
		err = errors.New("Prompt not found in text")
	}
	return
}

// matchPrompt returns true if one of the prompts is found in the string, false otherwise
func matchPrompt(text string, prompts []string) bool {
	_, err := getPromptMatchId(text, prompts)
	return err == nil
}

// WaitForPromptWithTimeout reads from ssh stdout and waits for one of the prompts to appear or until timeout is reached
func (s *ShellSession) WaitForPromptWithTimeout(prompts []string, timeout time.Duration) (stdout string,
	stderr string, err error) {

	var received string
	timer := time.After(timeout)
	for done := false; !done; {
		select {
		case received = <-s.stdoutChan:
			stdout += received

			if matchPrompt(stdout, prompts) {
				logger.Debug("Prompt matched", "prompts", prompts)
				done = true
			}
		case received = <-s.stderrChan:
			stderr += received
		case err = <-s.backChan:
			err = errors.WrapPrefix(err, "I/O error while waiting for prompt", 0)
			done = true
		case <-timer:
			var promptsStr string
			for _, v := range prompts {
				promptsStr += fmt.Sprintf("'%s' ", string(v))
			}
			err = errors.Errorf("Timed out (%v) waiting for channel. "+
				"Prompts: %v. Output: %v", timeout, promptsStr, stdout)
			done = true
		}
	}

	return
}

// WaitForPrompt reads from ssh stdout and waits for one of the prompts to appear or until default timeout is reached
func (s *ShellSession) WaitForPrompt(prompts []string) (stdout string, stderr string, err error) {
	return s.WaitForPromptWithTimeout(prompts, 30*time.Second)
}

// WaitForDefaultPrompt reads from ssh stdout and waits for one of the standard prompts to appear
func (s *ShellSession) WaitForDefaultPrompt() (output string, stdout string, err error) {
	prompts := []string{"# ", "$ "}
	return s.WaitForPrompt(prompts)
}

// In order not to interfere with the command's exit codes, special Run error codes
// are negative numbers starting with -1, e.g. -1, -2 etc.
const (
	ecRunErrorGeneric = iota*-1 - 1
	ecRunErrorNoPromptAfterExitCodeCommand
	ecRunErrorNoExitCodeIdentifierFound
)

// Run executes a command in an existing shell session
// It's intended to be used in scenarios that need to keep the context in-between command executions
// Make sure to close the session once done, i.e. the context is not needed anymore
// If you only need to run a single command, use remote.Run() instead
func (s *ShellSession) Run(cmd string) (exitCode int, stdout string, stderr string, err error) {
	exitCode = ecRunErrorGeneric
	var (
		byteCount   int
		exitCodeStr string
	)

	// Send the command
	logger.Info("Running shell command", "host", s.remote.address, "cmd", cmd)
	if byteCount, err = s.stdinWriter.Write([]byte(cmd + "\n")); err != nil {
		err = errors.WrapPrefix(err, fmt.Sprintf("Failed to send command to host %v. "+
			"byteCount=%v, cmd='%v'", s.remote.address, byteCount, cmd), 0)
		return
	}

	// Wait for prompt
	if stdout, stderr, err = s.WaitForPrompt(s.prompts); err != nil {
		err = errors.WrapPrefix(err, fmt.Sprintf("Didn't get prompt after command on host %v. cmd='%v', "+
			"prompts=%v, stdoutReader='%v', stderr='%v'", s.remote.address, cmd, s.prompts, stdout, stderr), 0)
		return
	}

	// Get exit code
	exitCodeIdentifier := "EXITCODE=" // Prevents getting exit code 0 from non-relevant output
	exitCodeRE := fmt.Sprintf(`(?m:^)%s([0-9]+)`, exitCodeIdentifier)
	re := regexp.MustCompile(exitCodeRE)

	ecCmd := fmt.Sprintf("/usr/bin/echo %s$?\n", exitCodeIdentifier)
	if byteCount, err = s.stdinWriter.Write([]byte(ecCmd)); err != nil {
		err = errors.WrapPrefix(err, fmt.Sprintf("Failed to send command to host %v. "+
			"cmd='%v', byteCount=%v", s.remote.address, ecCmd, byteCount), 0)
		return
	}

	if exitCodeStr, _, err = s.WaitForPrompt(s.prompts); err != nil {
		exitCode = ecRunErrorNoPromptAfterExitCodeCommand
		err = errors.WrapPrefix(err, fmt.Sprintf("Didn't get exit code (prompt not seen) on host %v. "+
			"exitCodeStr=%v", s.remote.address, exitCodeStr), 0)
		return
	} else {
		match := re.FindStringSubmatch(exitCodeStr)
		exitCodeStr = match[len(match)-1]
		if exitCode, err = strconv.Atoi(exitCodeStr); err != nil {
			exitCode = ecRunErrorNoExitCodeIdentifierFound
			err = errors.WrapPrefix(err, fmt.Sprintf("Didn't get exit code (Atoi failed) on host %v",
				s.remote.address), 0)
			return
		}
	}

	return
}

func (s *ShellSession) Address() string {
	return s.remote.Address()
}
