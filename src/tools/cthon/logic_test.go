package cthon

// func makeDummy() (types.ExternalProperties, error) {
// 	result, err := NewTool(&types.ToolParams{
// 		TargetLoaders: []string{"func13-loader1"},
// 		Config: types.Config{
// 			Elfs: struct {
// 				Frontend     string `doc:"ELFS frontend address, e.g. '192.168.0.1'"`
// 				Export       string `doc:"ELFS export to use, e.g. 'my_fs0/root'"`
// 				MountOptions string `default:"-o soft" doc:"ELFS export mount options, e.g. '-o soft'"`
// 			}{
// 				Frontend: "192.168.0.1",
// 				Export:   "my_fs0/root",
// 			},
// 			Cthon: cthon_config.Config{
// 				Tests: cthon_config.CthonTests{
// 					Basic:   true,
// 					Special: true,
// 				},
// 			},
// 			Setup: struct {
// 				Vheads        []string `doc:"Comma-separated list of vheads to use"`
// 				Loaders       []string `doc:"Comma-separated list of loaders to use"`
// 				CurrentLoader string
// 			}{
// 				Loaders:       []string{"func13-loader1"},
// 				CurrentLoader: "func13-loader1",
// 			},
// 		},
// 	})

// 	if err == nil {
// 		if lresult, ok := result.(types.ExternalProperties); !ok {
// 			return nil, fmt.Errorf("Type conversion failed")
// 		} else {
// 			return lresult, nil
// 		}
// 	}
// 	return nil, err
// }

// func TestInterfaceImplementation(t *testing.T) {
// 	var maybeTool *tool
// 	_, ok := (interface{})(maybeTool).(types.Tool)
// 	if !ok {
// 		t.Fatal("Interface types.Tool is not implemented by tool")
// 	}
// }

// func TestName(t *testing.T) {
// 	dummy, err := makeDummy()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	name := dummy.Name()
// 	if name != "cthon" {
// 		t.Fatalf("Wrong tool name. Expected 'cthon', got '%v'", name)
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
// 	expected := `mkdir -p /mnt/elfs && \
//          mount -o soft 192.168.0.1:my_fs0/root /mnt/elfs && \
// 	     ./runtests -b -t /mnt/elfs/func13-loader1 && \
// 	     ./runtests -s -t /mnt/elfs/func13-loader1 && \
//          true`
// 	re := regexp.MustCompile("\\s{2,}|\\t")
// 	expected = re.ReplaceAllString(expected, "")
// 	given := re.ReplaceAllString(cmd[len(cmd)-1], "")
// 	if given != expected {
// 		t.Fatalf(`Wrong command line.
//                   Expected: '%v'
//                   Got: '%v'`,
// 			expected, given)
// 	}
// }

// func TestGetImage(t *testing.T) {
// 	dummy, err := makeDummy()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	image := dummy.(types.LoggingProperties).GetImage()
// 	if image.String() == "" {
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
