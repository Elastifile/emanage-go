package filestore_api

// func init() {
// 	logging.Setup(logging_config.ConfigForUnitTest())
// }

// //// Integration tests ////
// const filestoreEndpoint = types.Host("172.17.0.1") // change to your minio live instance address (tesla slave uses conf.CurrentLoader)

// func TestStoreRetrieveFiles(t *testing.T) {
// 	// Integration test - skipping by default
// 	if true {
// 		t.SkipNow()
// 	}

// 	msg := messaging.NewClient(filestoreEndpoint)
// 	fs, err := NewClient(msg, filestoreEndpoint, false)
// 	check(t, err)

// 	var count int

// 	count++
// 	t.Logf("%d) Create a file in dir (simulating tool's config file)", count)
// 	path := "jobid/erun"

// 	f, err := createTempFile(path, "testfile")
// 	check(t, err)
// 	defer func() { _ = os.RemoveAll(path) }()

// 	bucket := filestore_types.Bucket(path).ToBucketName()
// 	check(t, fs.filestore.MakeBucket(bucket))

// 	// upload file to minio before telling slaves to sync
// 	check(t, fs.filestore.Upload(bucket, types.NewNamedReader(f.Name(), f)))

// 	count++
// 	t.Logf("%d) CCStore files (simulating config file being sent to slave on tool start)", count)
// 	err = fs.RequestStoreFiles(bucket, path, []types.Host{types.Host("localhost")})
// 	check(t, err)

// 	count++
// 	t.Logf("%d) Retrieve files (simulating results file being sent from slave to filestore "+
// 		"and then being downloaded here)", count)
// 	filesByHost, err := fs.RequestRetrieveFiles(path, []string{filepath.Base(f.Name())}, []types.Host{types.Host(filestoreEndpoint)})
// 	check(t, err)

// 	//// Validations ////
// 	if len(filesByHost) == 0 {
// 		t.Fatal("Map 'filesByHost' is empty! len(filesByHost) == 0!")
// 	}

// 	if files, ok := filesByHost[filestoreEndpoint]; ok {
// 		if len(files) == 0 {
// 			t.Fatal("No files found! len(files) == 0!")
// 		}
// 		for _, f := range files {
// 			data, err := ioutil.ReadAll(f)
// 			check(t, err)
// 			t.Logf("Got file from host: %v ,content: %v", filestoreEndpoint, string(data))
// 		}
// 	} else {
// 		t.Fatalf("Missing key on filesByHost. key='%v'", filestoreEndpoint)
// 	}
// }

// func createTempFile(dir, fileName string) (*os.File, error) {
// 	text := fmt.Sprintf("####FILECONTENT: This test file was created on: %v####\n",
// 		time.Now().Local())

// 	if err := os.MkdirAll(dir, 0755); err != nil {
// 		return nil, err
// 	}

// 	f, err := ioutil.TempFile(dir, fileName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = ioutil.WriteFile(f.Name(), []byte(text), os.ModePerm)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return f, nil
// }

// func check(t *testing.T, err error) {
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
