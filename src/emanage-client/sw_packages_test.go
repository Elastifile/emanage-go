package emanage_test

// func TestPackageUpload(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("skipping test in short mode.")
// 	}
// 	const tarLink = "TESLA_TESTS_COLD_UPGRADE_UPGRADE_TAR_LINK"
// 	uri := os.Getenv(tarLink)
// 	if uri == "" {
// 		t.Fatalf("Environment variable %v not set", tarLink)
// 	}
// 	const host = "TESLA_EMANAGE_SERVER"
// 	ip := os.Getenv(host)
// 	if ip == "" {
// 		t.Fatalf("Environment variable %v not set", host)
// 	}
// 	mgmt := startEManage(t)

// 	_, err := mgmt.SwPackages.UploadByCli(uri, ip)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// }

// func TestPackageInfo(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("skipping test in short mode.")
// 	}
// 	mgmt := startEManage(t)

// 	info, err := mgmt.SwPackages.Info()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	spew.Dump(info)
// }
// func TestPackageVerify(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("skipping test in short mode.")
// 	}
// 	mgmt := startEManage(t)

// 	pack, err := mgmt.SwPackages.Verify()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	spew.Dump(pack)
// 	if pack.Status != emanage.SwPackageVerifySuccess {
// 		t.Fatal(err)
// 	}
// }
// func TestPackageDelete(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("skipping test in short mode.")
// 	}
// 	mgmt := startEManage(t)

// 	err := mgmt.SwPackages.Delete()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
