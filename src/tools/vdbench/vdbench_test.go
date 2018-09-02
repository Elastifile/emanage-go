package vdbench

import (
	"fmt"

	"github.com/elastifile/emanage-go/src/types"
)

func makeDummy() (types.Tool, error) {
	result, err := NewTool(&types.ToolParams{
	// Add whatever paramters you need for the tool
	// to function properly
	})
	if err == nil {
		if _, ok := result.(types.LoggingProperties); !ok {
			return nil, fmt.Errorf("Type conversion failed")
		} else {
			return result, nil
		}
	}
	return nil, err
}

// func TestValidateResults(t *testing.T) {
// 	dummy, err := makeDummy()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	opts := &types.ResultOpts{}
// 	err = dummy.(types.LoggingProperties).GetResults(opts)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestGetResultFilesPatterns(t *testing.T) {
// 	dummy, err := makeDummy()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	files := dummy.(types.LoggingProperties).GetResultFilesPatterns()
// 	if len(files) > 0 {
// 		t.Fatalf("vdbench should have no result files, but had: %v", files)
// 	}
// }

// func TestName(t *testing.T) {
// 	dummy, err := makeDummy()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	name := dummy.Name()
// 	if name != "vdbench" {
// 		t.Fatalf("Wrong tool name. Expected 'vdbench', got '%v'", name)
// 	}
// }

// func TestGetCommand(t *testing.T) {
// 	dummy, err := makeDummy()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	cmd, err := dummy.(types.LoggingProperties).GetMasterCommand()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	expected := `echo "Hello World!"`
// 	re := regexp.MustCompile("\\s{2,}|\\t")
// 	expected = re.ReplaceAllString(expected, "")
// 	given := re.ReplaceAllString(cmd[len(cmd)-1], "")
// 	if given != expected {
// 		t.Fatalf(`Wrong command line.
//                   Expected: '%v'
//                   Got: '%v'`,
// 			expected, given)
// 	}
// 	cmd, err = dummy.(types.LoggingProperties).GetSlaveCommand("localhost")
// 	if err == nil {
// 		fmt.Println(cmd)
// 		t.Fatal("Empty strings should not be accepted")
// 	}
// }

// func TestGetImage(t *testing.T) {
// 	dummy, err := makeDummy()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	image := dummy.(types.LoggingProperties).GetImage()
// 	if image == "" {
// 		t.Fatal("Image name must not be empty")
// 	}
// }

// func TestWaitFor(t *testing.T) {
// 	dummy, err := makeDummy()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	hosts := dummy.WaitFor()
// 	if len(hosts) == 0 {
// 		t.Fatal("Must have at least one host to wait for")
// 	}
// }
