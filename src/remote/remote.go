package remote

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-errors/errors"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/koding/multiconfig"
	"golang.org/x/crypto/ssh"

	"github.com/elastifile/emanage-go/src/github.com/pkg/sftp"
	"github.com/elastifile/emanage-go/src/helputils"
	"github.com/elastifile/emanage-go/src/ioutil2"
	"github.com/elastifile/emanage-go/src/logging"
	terrors "github.com/elastifile/emanage-go/src/tools/errors"
	"github.com/elastifile/emanage-go/src/types"
)

var logger = logging.NewLogger("remote")

var (
	DefaultUser = "root"
	password    = "123456"
)

type Remote struct {
	*ssh.Client
	address string
}

// Tesla private key
const privateKeyTesla = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxiAgU+TBpjMZP+bpYDAd9inRmkux2zWKJ0x2mIF9ePTOqKyO
U4b4qtAChNL1/WFW087myW/ngUevaDhDCl13Q+gPTNdKryQ5tSZt7mM+gCQ5wa8+
TJ7upOMt8ewQfFqu/nvfA4sP7wwXGAyac7e0V/Z/CYtEac/86EevzvMAvSGtyOsv
VguYbTvZ7yBAiX9KivvS8iiQJFs/vg2O3W5i9KCXJhMxJbpZK4VFXLxnqP3SUCoR
3xCCimnyJVcifoJrLNY5pLIdZCUWvYvB4sZ2++iM73FSylVW5AlGjEJY3RfSREin
GXdmwGTJw7Qf9Zl34TigFsqF74Rs37TiamvBlwIDAQABAoIBADEvnJb3PZQTL3uL
yfhVObr5Gs4haKAxJIEpHHA0BYBX1k5NuG/IHEfbqCRtcyBaYAHpZaYwn9qaf1Ny
VEJofclf/RxofmQIrELqrXP3M3cAh04sQV0oP7qzo55hqp4UrUkEW5M3nNcNu1X/
GwELZDxKN7OY3X0fY3wO//hyn4ZUO+D67hy/7jWeqBSI5ldhOUlsy2tH07nTdfJA
aNg6FDXTGhdrGTTqmF/nx+JrbN82TP3Z8rF1fH6vw8tcR1wLgmGYJ03rRZARM8mV
KOl+HjZKn4wFvPPi94LZPiQjlImthIrcYzItrGCOV2IMPgXAjp/oXR7f+w6b+OdB
quA5bdECgYEA9Y8tGjqqcC41UBWOiuCK1gWUAcGSygEm4U1fhs0Ta4AJcB+Y4sP1
Wj7HSCH4pBle5ptQGXr/D7Aba5RKv9nG0GDB+rbUqDWrMBZUvvBNbJdPCuc/d5an
CBPig+TT0+NhqSo0ikLKQjyv3ffC0JeltohsH6l9MWvI5uFDq3YNvXMCgYEAzoym
hsci7ZBL1n1Rd7mDsP6HPcm5kd9/BFFMtlQ4J7Lo6P3mRC+MVNPA2y1OjZz6BSmW
JWadRBZQX++8we0pJDRr95Z/h/KMF6NUTSQFLTDyKJZJ0fzE0VFZ4cFNt9dInRcr
NcDKrulPBooFYtisCBTvqvQZB6Bqf54rciY9ok0CgYBlKIRWB9Dqwb5orVi5UWMg
Y8/jIVFNMkhIDE3MBHN5l5dLv1iGIl/Su4Xw6z24rFLyAQ1wUcD9P3ksSjy46AgF
E7cVh3f3i5RoDOsAdZIPqA2B2l3E8CUDKN3po7y3zzYOMDrQsk8MFmAEuupgb/hu
R6jzKilO0K5Enov3+HwUjQKBgQCSdctzUGtk9ue/vPbypVpLPACcAJmUZV1Eq31w
4EE9bTCYXclU2j5wvAJJqFPGOORUmst708p6SDRQAsfe3xNIN9/DHTVrKcLK1Cbn
2D3PcKCVxROtUcYiAsRwP0ijroUv1wjtwbo20vsVUPNGxwsdY/MaTqIc9VPvEDIQ
bCtSFQKBgF5IyCNjBLN0/oqChmIQRh7MtQcZpRiRuKHsveiCIun4FFfzxZbDr/Uh
ci9OaxGKYxGJm5LBZW67FWjZ6XE2jNmCxw04Rlro0nWmesUu28mnbzPWBsF1UCUE
fr42DvL1bAC0WzUvWcant5FyRF3l6GEk0Z8bDryzHOB2lHlTaPXu
-----END RSA PRIVATE KEY-----`

// Elastifile cloud pivate key
const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAlamnqgnlYAZQtb+12bNUE0EbE0imUbUGbgZCAsoEtr3SGXn+lVZ/HfezH68h
uGGsOIas4zcd0D4JkAh6SZI8bYRg/xyGLdBH0Aeda2XYSMBjcJpYWi+zbBlnrEahC0ZkLCyfBR9D
iQ7sgGWA01LHuR+F5W3qzP8bRWqdtoZBqQzOodeGbE/usBxExuc+aQjqQ/Myy9qlzqyfvJma/dPl
GOPT/y+pyiE4jMmutenoFqULyblbDjsR5rn62XO++MsEeJpVXas0SmdtCB6wNV2I/pGhAwKQNDCb
sAlyxDn+DY6HreyCrEs/2UzE2R12BzPT5icdGF4GIHSMMSbVsI1QIwIDAQABAoIBAHhs4Sqjd9wC
eOMl9a/M6nPryVuE879/SEMz6d3QrGZW6z3wufd0eCim7Y0rIJj6jteVFnxmFyk/Ifaxu6SDdPgy
stHUI+utUnuEBRksBBLWK62ppDo9c8Qh1Rf8ZQ1rgKNC7Fn0xCT+yzKtejGQxICjatwny+1TwJ/k
oC8EbMUW9SdM4dstLiolg89LyR47z9fjRdOz7B+wUo2L3i2aO8dASz6+s75Pw6Q1vfvV9aa5tAxv
YHVj/CwPE/8GiFoAhYzK9Klc6yr2gykIBuexEBbX9n/GiC65jfdLYLoE5mARUAVgX+nzdtpRxrr+
AJ4O4odddLsCbtnph/bExBjFaVECgYEA8VBTn7oTtM1MaQ9kq8dI2ONc/o9KDRhfsepg6H780URR
rS8hzWsBH3i8IzzrUeaB+wfDuvnyJxNDnQH0osH1b5/EkTnCBpoPxIqhT5aRLVGj5bob5p7AZ9re
nbY9NvSoXTHTc/YfnfrWBe/YcW03m/PxLPgK5CA5DNkySNcLBqcCgYEAnsVnKBY28BcgTGWTrZ+A
wSILgsdMqaqFbNOeFz1eoRmwjuEr6vgavR+elBirs9iVWPunXZZp8Kk2REvvbZzc8W+2SrnynPcp
hRfPa3AHmCxjuEa2IvOAWuXju3vZBuGi6RmsvIW+k2JR9moYnsCwKg0a4NUI4NTrEVt8YFrlFiUC
gYEA3svnKWXHo/qYPbe0ntTT8UuaJH1vRT/nuoBHEOGN0jhdoKQPdXUCCGiUa4RnT6qe/4JfuFXk
uScUGfYgQ9UIIhNEjUDiXu4G+NN1s+G0ED98LzHxaK/RvRFa0And4stoVwMV0869ykI5MCoNkEvi
N962IWqxw4Xspmwx/QZua6kCgYEAmlhQxfmCgrgYn+wHOjva0OWlxwa0EiRA/OBjyTfRrE7rApsH
NK7oX35DjLjHAljEJvLNav90EY6NwKsJIMfjSZpN3isSHWFxLX9OFqWb3GUas+OZ4DRTTQo67uGg
wz3KP0zETai+8UwdXa0aLgont1lyDatLrDWQJxmGuazjEzkCgYBRDeCBACb+6S7BtqaHAdv5jB3o
zLFo4fmUKJCwvcZ2NJ79Pnp7Ddh7NXtqyva6ka0QkBQpxeFia0NvkhQ7vQ36DNOMw4PoMhzypdl+
9LG4qI+g/pM28YpY4tGXJ5J19hGxl8N3rluT4JIY6BufeOEdSS5pBwJsCPRDYV11PTDg+g==
-----END RSA PRIVATE KEY-----`

func CopySSHPublicKeyToHostsBootstrap(host string, hosts ...types.Host) error {
	r, e := NewRemote(host)
	if e != nil {
		return e
	}

	pubKeyAuth, err := AuthMethodTeslaPubKey()
	if err != nil {
		return err
	}

	opts := &RemoteOpts{
		User:        "root",
		AuthMethods: []ssh.AuthMethod{pubKeyAuth},
	}
	// test ssh connection
	var targets []types.Host
	for _, host := range hosts {
		if host == "" {
			panic(fmt.Errorf("empty host address specified, hosts=%v", hosts))
		}

		opts.Address = string(host)
		if client, err := NewRemoteWithCustomAuth(opts); err == nil {
			if _, e := client.Run("echo test"); e == nil {
				continue
			}
		}
		logger.Info("Found host not familiar with tesla key, will try to fix", "host", host)
		targets = append(targets, host)
	}
	if len(targets) == 0 {
		logger.Debug("Nothing to do (all hosts are already familiar with tesla key).")
		return nil
	}

	return r.CopySSHPublicKey(targets...)
}

func NewRemote(address string) (*Remote, error) {
	pubKeys, err := AuthMethodAllPubKeys()
	if err != nil {
		return nil, errors.WrapPrefix(err, "Failed to get public keys for auth", 0)
	}

	config := &ssh.ClientConfig{
		User: DefaultUser,
		Auth: []ssh.AuthMethod{
			pubKeys,
			ssh.Password(password),
		},
	}

	// ssh.Dial uses `net.Dial', but even if it used
	// `net.DialTimeout', it would still hang forever because when it
	// resolves the port, it doesn't use timeout at all.
	logger.Debug("Dialing SSH", "address", address, "config", *config)
	client, err := ssh.Dial("tcp", net.JoinHostPort(address, "22"), config)
	if err != nil {
		return nil, errors.Errorf("(SSH) Couldn't connect to %s: %s", address, err)
	}

	return &Remote{client, address}, nil
}

func NewCancelableRemote(address string, timeout time.Duration) (r *Remote, e error) {
	watcher := make(chan struct{})
	go func() {
		r, e = NewRemote(address)
		watcher <- struct{}{}
	}()
	select {
	case <-time.After(timeout):
		return nil, errors.Errorf("Couldn't create Remote in %s", timeout)
	case <-watcher:
		return r, e
	}
}

func AuthMethodTeslaPubKey() (ssh.AuthMethod, error) {
	signer, err := ssh.ParsePrivateKey([]byte(privateKeyTesla))
	if err != nil {
		return nil, errors.Errorf("Couldn't parse private key: %s", err)
	}
	return ssh.PublicKeys(signer), nil
}

func AuthMethodAllPubKeys() (ssh.AuthMethod, error) {
	signerElastifile, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, errors.Errorf("(SSH) Couldn't parse Elastifile private key: %s", err)
	}

	signerTesla, err := ssh.ParsePrivateKey([]byte(privateKeyTesla))
	if err != nil {
		return nil, errors.Errorf("(SSH) Couldn't parse Tesla private key: %s", err)
	}

	return ssh.PublicKeys(signerElastifile, signerTesla), nil
}

type RemoteOpts struct {
	Address     string
	User        string
	AuthMethods []ssh.AuthMethod
}

func NewRemoteWithCustomAuth(opts *RemoteOpts) (*Remote, error) {
	config := &ssh.ClientConfig{
		User: opts.User,
		Auth: opts.AuthMethods,
	}
	// ssh.Dial uses `net.Dial', but even if it used
	// `net.DialTimeout', it would still hang forever because when it
	// resolves the port, it doesn't use timeout at all.
	client, err := ssh.Dial("tcp", net.JoinHostPort(opts.Address, "22"), config)
	if err != nil {
		return nil, errors.Errorf("Couldn't connect to %s: %s", opts.Address, err)
	}
	return &Remote{client, opts.Address}, nil
}

func extractPublicKey(pkey []byte) ([]byte, error) {
	p, _ := pem.Decode([]byte(pkey))

	key, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}

	pub, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		return nil, err
	}

	return ssh.MarshalAuthorizedKey(pub), nil
}

func (rem *Remote) sshCopyId(publicKeyFileName string, host types.Host, timeout time.Duration) error {
	cmdHack := fmt.Sprintf("ssh -o StrictHostKeyChecking=no root@%s 'echo \"\" >> ~/.ssh/authorized_keys'", host)
	cmdOrig := fmt.Sprintf("SSH_OPTS='-F /dev/null' ssh-copy-id -o StrictHostKeyChecking=no -i %s root@%s", publicKeyFileName, host)
	cmd := fmt.Sprintf("%s&&%s", cmdHack, cmdOrig)
	ctx, _ := ioutil2.WithDeadline(nil, time.Now().Add(timeout))
	if out, err := rem.RunWithContext(ctx, cmd); err != nil {
		return errors.Errorf("error running %s %v %v", cmd, err, out)
	}

	return nil
}

const teslaKeyDir = "/elastifile/emanage/"
const teslaKeyName = "tesla"

func (rem *Remote) CopySSHPublicKey(hosts ...types.Host) (err error) {
	key := teslaKeyDir + teslaKeyName
	privateKeyFileName := key
	publicKeyFileName := key + ".pub"

	privReader := strings.NewReader(privateKey)
	pubReaderBytes, err := extractPublicKey([]byte(privateKey))
	if err != nil {
		return err
	}
	readers := []io.Reader{privReader, bytes.NewReader(pubReaderBytes)}

	if _, e := rem.RunWithTimeout(fmt.Sprintf("mkdir -p %s", teslaKeyDir), 20*time.Second); e != nil {
		return err
	}

	for i, fileName := range []string{privateKeyFileName, publicKeyFileName} {
		if e := rem.Upload(readers[i], fileName); e != nil {
			err = multierror.Append(err, e)
		} else {
			if out, e := rem.Run(fmt.Sprintf("chmod 666 %s*", key)); e != nil {
				err = multierror.Append(err, errors.Errorf(
					"Couldn't change permissions: %s, because: %s",
					e, out,
				))
			}
		}
	}
	if err != nil {
		return err
	}

	for _, host := range hosts {
		if e := rem.sshCopyId(publicKeyFileName, host, time.Minute); e != nil {
			err = multierror.Append(err, e)
		}
	}

	return err
}

func (rem *Remote) Address() string {
	return rem.address
}

func (rem *Remote) Run(cmds ...string) (string, error) {
	cmd := strings.Join(cmds, " ")

	session, err := rem.Client.NewSession()
	if err != nil {
		return "", errors.New(err)
	}
	defer func() { _ = session.Close() }()

	logger.Debug("Running command", "host", rem.address, "cmd", cmd)
	output, err := session.CombinedOutput(cmd)
	return string(output), err
}

func (rem *Remote) Start(cmds ...string) (*ssh.Session, error) {
	cmd := strings.Join(cmds, " ")

	session, err := rem.Client.NewSession()
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	logger.Debug("Starting command", "host", rem.address, "cmd", cmd)
	err = session.Start(cmd)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return session, nil
}

type blockingBytesReadWriter struct {
	sync.Mutex
	tick    time.Duration
	timeout time.Duration
	buf     *bytes.Buffer
}

func (r *blockingBytesReadWriter) Read(p []byte) (int, error) {
	start := time.Now()
	for t := range time.Tick(r.tick) {
		if r.buf.Len() >= len(p) || t.Sub(start) > r.timeout {
			break
		}
	}
	r.Lock()
	defer r.Unlock()
	return r.buf.Read(p)
}

func (r *blockingBytesReadWriter) Write(p []byte) (int, error) {
	r.Lock()
	defer r.Unlock()
	return r.buf.Write(p)
}

func (r *blockingBytesReadWriter) Close() error {
	return nil
}

type lazyGzipReader struct {
	sync.Mutex
	brw *blockingBytesReadWriter
	z   *gzip.Reader
}

func (l *lazyGzipReader) Read(p []byte) (int, error) {
	l.Lock()
	if l.z == nil {
		z, err := gzip.NewReader(l.brw)
		if err != nil {
			return 0, err
		}
		l.z = z
	}
	l.Unlock()
	return l.z.Read(p)
}

func (l *lazyGzipReader) Close() error {
	return l.z.Close()
}

type blockingFileReader struct {
	sync.Mutex
	f       *sftp.File
	info    os.FileInfo
	timeout time.Duration
	pos     int64
	client  *sftp.Client
}

func (l *blockingFileReader) Read(p []byte) (int, error) {
	l.Lock()
	defer l.Unlock()
	if l.info == nil {
		info, err := l.f.Stat()
		if err != nil {
			return 0, errors.Errorf("Cannot stat %s because %s", l.f.Name(), err)
		}
		l.info = info
	}
	start := time.Now()
	n := int64(0)
	buf := make([]byte, len(p))
	done := false
	var (
		combined bytes.Buffer
		err      error
	)

	for t := range time.Tick(l.timeout / 100) {
		if t.Sub(start) > l.timeout {
			done = true
			err = errors.Errorf("Couldn't read file in %s", l.timeout)
			break
		}
		_, _ = l.f.Seek(l.pos+n, 0)
		m, e := l.f.Read(buf)
		if l.pos+n+int64(m) == l.info.Size() {
			done = true
			err = io.EOF
		} else if e != nil && e != io.EOF {
			done = true
			err = e
		}
		if m > 0 {
			_, _ = combined.Write(buf[:m])
		}
		if len(buf) == m {
			done = true
		}
		buf = buf[m:]
		n += int64(m)
		if done {
			break
		}
	}
	copy(p, combined.Bytes())
	l.pos += n
	return int(n), err
}

func (l *blockingFileReader) Close() error {
	err := l.f.Close()
	if e := l.client.Close(); e != nil {
		if err != nil {
			err = multierror.Append(err, e)
		} else {
			err = e
		}
	}
	return err
}

// There's a package https://godoc.org/github.com/tmc/scp
// Don't use it.  It simply calls shell which in turn calls scp
// with some bizarre arguments.
func (rem *Remote) OpenCompressedFile(filename string) (io.ReadCloser, error) {
	session, err := rem.Client.NewSession()
	if err != nil {
		return nil, errors.New(err)
	}

	brw := &blockingBytesReadWriter{
		tick:    10 * time.Millisecond,
		timeout: time.Second,
		buf:     &bytes.Buffer{},
	}
	session.Stdout = brw
	err = session.Start(fmt.Sprintf("cat '%s' | gzip -f", filename))
	if err != nil {
		go func() { _ = session.Wait() }()
	}
	return &lazyGzipReader{
		brw: brw,
	}, err
}

func (rem *Remote) OpenFile(filename string) (io.ReadCloser, error) {
	session, err := rem.Client.NewSession()
	if err != nil {
		return nil, errors.New(err)
	}

	session.Stdout = &blockingBytesReadWriter{
		tick:    10 * time.Millisecond,
		timeout: time.Second,
		buf:     &bytes.Buffer{},
	}
	err = session.Start("cat '" + filename + "'")
	if err != nil {
		go func() { _ = session.Wait() }()
	}
	return session.Stdout.(io.ReadCloser), err
}

// Even though the test seems to pass, there are some strange issues
// with this method when using it concurrently / with multiple files /
// same file, but multiple connections (it hangs).  Not only this, the
// SFTP library doesn't have an option to compress files while
// sending.  So, sending uncompressed text files using this function
// is a bad idea.
func (rem *Remote) OpenSftpFile(remotePath string) (io.ReadCloser, error) {
	sftp, err := sftp.NewClient(rem.Client)
	if err != nil {
		return nil, errors.New(err)
	}

	remoteFile, err := sftp.Open(remotePath)
	if err != nil {
		return nil, errors.New(err)
	}
	return &blockingFileReader{
		f:       remoteFile,
		timeout: time.Second,
		client:  sftp,
	}, nil
}

func (rem *Remote) RunWithTimeout(cmd string, timeout time.Duration) (string, error) {
	type res struct {
		out string
		err error
	}
	watcher := make(chan res)
	go func() {
		out, err := rem.Run(cmd)
		if err != nil {
			err = errors.Errorf(
				"'%s' on %s failed due to %s, with output: %s",
				cmd, rem.address, err, out,
			)
		}
		watcher <- res{out: out, err: err}
	}()
	select {
	case <-time.After(timeout):
		return "", errors.Errorf(
			"'%s' on %s timed out after %s",
			cmd, rem.address, timeout,
		)
	case r := <-watcher:
		return r.out, r.err
	}
}

func (rem *Remote) RunWithContext(ctx ioutil2.Context, cmd string) (string, error) {
	type res struct {
		out string
		err error
	}
	watcher := make(chan res)
	go func() {
		out, err := rem.Run(cmd)
		if err != nil {
			err = errors.Errorf(
				"'%s' on %s failed due to %s, with output: %s",
				cmd, rem.address, err, out,
			)
		}
		watcher <- res{out: out, err: err}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case r := <-watcher:
		return r.out, r.err
	}
}

func (rem *Remote) Mkdir(remoteDir string) error {
	if _, err := rem.Run("mkdir -p " + remoteDir); err != nil {
		logger.Error(err.Error(), "host", rem.Address(), "dir", remoteDir)
		return errors.New(err)
	}

	logger.Debug("Made", "dir", remoteDir)
	return nil
}

func (rem *Remote) Upload(r io.Reader, remotePath string) error {
	logger.Debug("Uploading", "host", rem.address, "path", remotePath)

	sftp, err := sftp.NewClient(rem.Client)
	if err != nil {
		logger.Error(err.Error())
		return errors.New(err)
	}
	defer func() { _ = sftp.Close() }()

	remoteFile, err := sftp.Create(remotePath)
	if err != nil {
		logger.Error(err.Error() + " " + remotePath)
		return errors.New(err)
	}
	defer func() { _ = remoteFile.Close() }()

	_, err = io.Copy(remoteFile, r)
	if err != nil {
		logger.Error(err.Error())
		return errors.New(err)
	}

	return nil
}

func (rem *Remote) UploadFile(localPath, remotePath string) error {
	if localPath == "" {
		return errors.New("nothing to upload")
	}

	logger.Debug("Uploading file", "localPath", localPath, "host", rem.address, "remotePath", remotePath)

	localFile, err := os.Open(localPath)
	if err != nil {
		return errors.Errorf("Couldn't open %s because %s", localPath, err)
	}
	defer func() { _ = localFile.Close() }()

	return rem.Upload(localFile, remotePath)
}

func RunString(script string, timeout time.Duration, addresses ...string) (outs map[string]string, _ error) {
	var err error
	scriptPath, err := helputils.TmpRandomFileName("tesla-remote.sh")
	if err != nil {
		return outs, errors.Wrap(err, 0)
	}
	if err := helputils.WriteFile(scriptPath, []byte(script), os.ModePerm); err != nil {
		return outs, errors.Wrap(err, 0)
	}

	outs = make(map[string]string)
	mtx := sync.Mutex{}

	runScript := fmt.Sprintf("bash -c 'source %v'", scriptPath)
	if DefaultUser != "root" {
		runScript = fmt.Sprintf("sudo " + runScript)
	}
	errChan := make(chan error)
	for _, addr := range addresses {
		addr := addr
		go func() {
			if rem, err := NewCancelableRemote(addr, timeout); err != nil {
				errChan <- err
			} else if err := rem.Mkdir(filepath.Dir(scriptPath)); err != nil {
				errChan <- errors.Errorf("Failed making remote dir for: %s, err: %s", scriptPath, err)
			} else if err := rem.UploadFile(scriptPath, scriptPath); err != nil {
				errChan <- errors.Errorf("Failed uploading file: %s, err: %s", scriptPath, err)
			} else if out, err := rem.RunWithTimeout(runScript, timeout); err != nil {
				errChan <- errors.Errorf("Failed remote run: %s, err: %s, out:\n%s", scriptPath, err, out)
			} else {
				mtx.Lock()
				outs[addr] = string(out)
				mtx.Unlock()
				errChan <- nil
				logger.Debug("succeeded remote run", "path", scriptPath, "host", addr)
			}
		}()
	}
	if err := terrors.CollectAll(errChan, len(addresses)); err != nil {
		return outs, errors.Wrap(err, 0)
	}

	if err := os.Remove(scriptPath); err != nil {
		return outs, errors.Wrap(err, 0)
	}

	return outs, nil
}

func PingFromAllHosts(hosts []string, targets ...string) error {
	var multi error

	for _, host := range hosts {
		for _, target := range targets {
			_, err := PingFromRemoteServer(host, target)
			if err != nil {
				multi = multierror.Append(multi, err)
			}
		}
	}
	return multi
}

func PingFromRemoteServer(remoteServerIP string, IPToPing string) (gotPing bool, err error) {
	const NUMBER_OF_PINGS = 10
	const NUMBER_OF_CONCURRENT_PINGS = 10
	// remote client
	remote, err := NewRemote(remoteServerIP)
	if err != nil {
		return false, errors.Errorf("Couldn't create remote ssh client. Error: %s", err)
	}
	// run ping command

	commandString := fmt.Sprintf("ping -c %v -l %v %v", NUMBER_OF_PINGS, NUMBER_OF_CONCURRENT_PINGS, IPToPing)
	if DefaultUser != "root" {
		commandString = fmt.Sprintf("sudo " + commandString)
	}
	cmdOutput, err := remote.Run(commandString)
	if err != nil {
		return false, errors.Errorf("Error running ping command on remote host (%v). Command: %v Error: '%v', Actual Output: %v",
			remoteServerIP, commandString, err, cmdOutput)
	}
	// verify output
	if strings.Contains(cmdOutput, "100% packet loss") {
		return false, errors.Errorf("%v out of %v pings from host (%v) to (%v) got lost, Actual Output: %v",
			NUMBER_OF_PINGS, NUMBER_OF_PINGS, remoteServerIP, IPToPing, cmdOutput)
	}
	if !strings.Contains(cmdOutput, " 0% packet loss") {
		logger.Warn(fmt.Sprintf("Some of the %v pings attempted from host (%v) to (%v) got lost, Actual Output: %v",
			NUMBER_OF_PINGS, remoteServerIP, IPToPing, cmdOutput))
	}
	return true, nil
}

func ValidateHostname(host string) error {
	timeout := time.Minute

	r, err := NewCancelableRemote(host, timeout)
	if err != nil {
		return err
	}

	hostname, err := r.RunWithTimeout("hostname", time.Minute)
	if err != nil {
		return err
	}
	if strings.Contains(hostname, "localhost") {
		return errors.Errorf("hostname of '%v' contains string 'localhost'  (hostname: %v)", host, hostname)
	}

	ip, err := r.RunWithTimeout("hostname -i", time.Minute)
	if err != nil {
		return err
	}

	loopback := "127.0.0.1"
	if strings.Contains(ip, loopback) {
		return errors.Errorf("resolving hostname of '%v' directs to loopback '%v'", host, loopback)
	}

	gotPing, err := PingFromRemoteServer(host, hostname)
	if err != nil || !gotPing {
		return errors.Errorf("ping failed from host '%v' to its hostname '%v' (original error: %v)", host, hostname, err)
	}

	return nil
}

func GetMemoryStats(host string) (result map[string]string, err error) {
	rh, err := NewCancelableRemote(host, time.Minute)
	if err != nil {
		return result, err
	}

	out, err := rh.RunWithTimeout("free", time.Minute)
	if err != nil {
		return result, err
	}

	result, err = parseMemoryStats(out)
	if err != nil {
		return result, err
	}

	return result, nil
}

// ParseFreeCmd parses linux 'free' command and returns a map[keys]values.
// NOTE: It currently parses the 'Mem:' line only, skipping the 'Swap:' line.
// Example
// input:
//               total        used        free      shared  buff/cache   available
// Mem:            47G         38G        5.4G        2.2G        3.6G        6.5G
// Swap:            0B          0B          0B
//
// output:
// map[total:47G used:38G free:5.4G shared:2.2G buff/cache:3.6G available:6.5G]
func parseMemoryStats(memStatsOutput string) (fieldsMap map[string]string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(memStatsOutput))
	var processedHeader bool
	var header, fields []string

	logger.Debug("Parsing linux 'free' cmd...", "cmd", memStatsOutput)
	for scanner.Scan() {
		line := scanner.Text()
		if !processedHeader {
			header = strings.Fields(line)
			processedHeader = true
			continue
		}

		fields = strings.Fields(line)[1:]
		break
	}

	if len(header) != len(fields) {
		return fieldsMap, errors.Errorf("failed parsing linux 'free' cmd. len(header) != len(fields), %v != %v", len(header), len(fields))
	}

	fieldsMap = make(map[string]string)
	for i := 0; i < len(header); i++ {
		fieldsMap[header[i]] = fields[i]
	}

	return fieldsMap, nil
}

func ReadDebugfsCounters(host string, counters []string) (result map[string]string, err error) {
	result = make(map[string]string)

	var (
		debugfsLocation = "/tmp/debug-fs"
		mountDebugfs    = "/elastifile/tools/debug_fs_mount.sh 127.0.0.1"
		umount          = fmt.Sprintf("umount %v", debugfsLocation)
		getVarCmd       = fmt.Sprintf("cat %v/127.0.0.1:1112/var/%%v", debugfsLocation)
		timeout         = time.Minute
	)

	rh, err := NewCancelableRemote(host, timeout)
	if err != nil {
		return result, err
	}

	// Disable 'require tty' in sudoers file
	// to avoid error: 'sudo: sorry, you must have a tty to run sudo'
	ttySwitch := make(map[string]string)
	ttySwitch["disable"] = "sed -i -e \"s/Defaults    requiretty.*/#Defaults    requiretty/g\" /etc/sudoers"
	ttySwitch["enable"] = "sed -i -e \"s/#Defaults    requiretty.*/Defaults    requiretty/g\" /etc/sudoers"
	defer func() { _, _ = rh.RunWithTimeout(ttySwitch["enable"], timeout) }()
	if _, err := rh.RunWithTimeout(ttySwitch["disable"], timeout); err != nil {
		return result, err
	}

	// mount debugfs
	if _, err := rh.RunWithTimeout(mountDebugfs, timeout); err != nil {
		return result, err
	}

	// read counters
	for _, counter := range counters {
		out, err := rh.RunWithTimeout(fmt.Sprintf(getVarCmd, counter), timeout)
		if err != nil {
			return result, err
		}
		result[counter] = out
	}

	// umount debugfs
	if _, err := rh.RunWithTimeout(umount, timeout); err != nil {
		return result, err
	}

	return result, nil
}

type TelegrafDeployOpts struct {
	TargetHost string
	DBServer   string
	DBName     string
	ConfTags   []string
}

func (opts *TelegrafDeployOpts) Update(upd TelegrafDeployOpts) {
	if upd.TargetHost != "" {
		opts.TargetHost = upd.TargetHost
	}
	if upd.DBServer != "" {
		opts.DBServer = upd.DBServer
	}
	if upd.DBName != "" {
		opts.DBName = upd.DBName
	}
	if upd.ConfTags != nil {
		opts.ConfTags = upd.ConfTags[:]
	}
}

func DeployTelegraf(opts TelegrafDeployOpts) error {
	const (
		rpm          = "telegraf-1.4.4-1.x86_64.rpm"
		telegrafConf = "/etc/telegraf/telegraf.conf"
		timeout      = time.Minute
	)
	deleteRpm := fmt.Sprintf("rm -f %v*", rpm)
	wget := fmt.Sprintf("wget https://dl.influxdata.com/telegraf/releases/%v", rpm)
	yumRemove := fmt.Sprintf("rpm -qa | grep telegraf | xargs yum remove -y ; echo $?")
	yumInstall := fmt.Sprintf("yum install -y %v", rpm)
	kickstart := "systemctl start telegraf"
	updateConfDbAddress := fmt.Sprintf("sed -i -- 's/localhost:8086/%v/g' %v", opts.DBServer, telegrafConf)                   // catches something like: "urls = ["http://localhost:8086"]"
	updateConfDbName := fmt.Sprintf("sed -i -- 's/database = \"telegraf\"/database =\"%v\"/g' %v", opts.DBName, telegrafConf) // catches something like: 'database = "telegraf" # required'
	updateConfProcStat := fmt.Sprintf(`
		cat <<- EOF >> %v
		[[inputs.procstat]]
		exe = "elfs"
		[[inputs.procstat]]
		exe = "telegraf"
		EOF`, telegrafConf)

	// to debug telegraf, run it manually: "usr/bin/telegraf --config etc/telegraf/telegraf.conf --test"
	commands := []string{
		deleteRpm, wget, yumRemove, yumInstall,
		updateConfDbAddress,
		updateConfDbName,
		updateConfProcStat,
	}

	for _, tag := range opts.ConfTags {
		updateTags := fmt.Sprintf("sed -i '/global_tags/a %v' %v", tag, telegrafConf) // appends after global_tags
		commands = append(commands, updateTags)
	}
	commands = append(commands, kickstart)

	rh, err := NewCancelableRemote(opts.TargetHost, timeout)
	if err != nil {
		return err
	}

	for _, cmd := range commands {
		if _, err := rh.RunWithTimeout(cmd, timeout); err != nil {
			return err
		}
	}

	logger.Info("Deployed telegraf", "opts", opts)
	return nil
}

type SshTunnelConfig struct {
	SourceIp   string
	TargetIp   string
	FrontendIp string `default:"172.16.0.1"`
	MountPort  int    `default:"644"`
	NfsPort    int    `default:"2049"`
	PortOffset int    `default:"10000"`
}

func NewSshTunnelConfig(upd ...SshTunnelConfig) (conf SshTunnelConfig) {
	tagLoader := &multiconfig.TagLoader{}
	if err := tagLoader.Load(&conf); err != nil {
		panic(err)
	}

	switch len(upd) {
	case 0:
		// pass
	case 1:
		conf.Update(upd[0])
	default:
		panic("may specify one update opts")
	}

	return conf
}

func (conf *SshTunnelConfig) Update(upd SshTunnelConfig) {
	if upd.SourceIp != "" {
		conf.SourceIp = upd.SourceIp
	}
	if upd.TargetIp != "" {
		conf.TargetIp = upd.TargetIp
	}
	if upd.FrontendIp != "" {
		conf.FrontendIp = upd.FrontendIp
	}
	if upd.MountPort > 0 {
		conf.MountPort = upd.MountPort
	}
	if upd.NfsPort > 0 {
		conf.NfsPort = upd.NfsPort
	}
	if upd.PortOffset > 0 {
		conf.PortOffset = upd.PortOffset
	}
}

// Create SSH tunnel:
// ssh -n -N -L 10644:$fe_ip:644 -L 12049:$fe_ip:2049 $target_ip
// $fe_ip=172.16.0.1
func OpenSshTunnel(opts SshTunnelConfig) (closer func() error, err error) {
	opts.TargetIp = helputils.ResolvedHost(opts.TargetIp)
	if opts.TargetIp == "" {
		return nil, errors.New("invalid target IP: " + opts.TargetIp)
	}

	cmdStr := fmt.Sprintf("ssh -n -N -L %d:%s:%d -L %d:%s:%d %s",
		opts.PortOffset+opts.MountPort, opts.FrontendIp, opts.MountPort,
		opts.PortOffset+opts.NfsPort, opts.FrontendIp, opts.NfsPort,
		opts.TargetIp,
	)

	var (
		cmd  *exec.Cmd
		rem  *Remote
		sess *ssh.Session
	)

	if opts.SourceIp != "" {
		opts.SourceIp = helputils.ResolvedHost(opts.SourceIp)
		if opts.SourceIp == "" {
			return nil, errors.New("invalid source IP: " + opts.SourceIp)
		}

		rem, err = NewCancelableRemote(opts.SourceIp, time.Minute)
		if err != nil {
			return nil, err
		}

		for err == nil {
			err = rem.KillPgrep(cmdStr)
		}

		sess, err = rem.Start(cmdStr)
		if err != nil {
			return nil, err
		}

		time.Sleep(time.Second)
		_, err = rem.Pgrep(cmdStr)
		if err != nil {
			return nil, err
		}
	} else {
		cmds := strings.Split(cmdStr, " ")
		cmd = exec.Command(cmds[0], cmds[1:]...)

		err := cmd.Start()
		if err != nil {
			return nil, err
		}

		if cmd.Process == nil {
			return nil, errors.Errorf("unexpected nil process for command: %v", cmd)
		}
	}

	closer = func() error {
		if opts.SourceIp == "" {
			return cmd.Process.Kill()
		} else {
			defer sess.Close()
			err := rem.KillPgrep(cmdStr)
			if err != nil {
				pid, e := rem.Pgrep(cmdStr)
				if e == nil && pid > 0 {
					err = errors.WrapPrefix(err, fmt.Sprintf("Failed killing pid: %d, cmd: %s", pid, cmdStr), 0)
				}
				return err
			}
		}
		return nil
	}

	return closer, nil
}

func (rem *Remote) KillPgrep(pgrep interface{}) error {
	pid, err := rem.Pgrep(pgrep)
	if err != nil {
		return err
	}

	out, err := rem.Run(fmt.Sprintf("kill %d", pid))
	if err != nil {
		return errors.WrapPrefix(err, out, 0)
	}

	return nil
}

func (rem *Remote) Pgrep(pgrep interface{}) (int, error) {
	psStr := fmt.Sprintf("ps ax | grep '%v' | grep -v grep | xargs | cut -d' ' -f1", pgrep)
	out, err := rem.Run(psStr)
	if err != nil {
		return -1, errors.WrapPrefix(err, out, 0)
	}

	pidStr := strings.Split(out, "\n")[0]
	if pidStr == "" {
		return -1, errors.Errorf("Process ID not found for; %v", pgrep)
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		err = errors.WrapPrefix(err, "Non-numeric process ID: "+pidStr, 0)
	}

	return pid, err
}
