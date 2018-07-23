package shadow

// func TestShadowErrorNilNil(t *testing.T) {
// 	err := checkShadow("", nil, nil)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestShadowErrorErrErr(t *testing.T) {
// 	testShadow(t, errors.New("refr"), errors.New("elfs"))
// }

// func TestShadowErrorNilErr(t *testing.T) {
// 	testShadow(t, nil, errors.New("elfs"))
// }

// func TestShadowErrorErrNil(t *testing.T) {
// 	testShadow(t, errors.New("elfs"), nil)
// }

// func TestShadowErrorSameValueDifferentInstances(t *testing.T) {
// 	er := &nfsx.NfsError{Status: nfsx.V3ERR_ACCES}
// 	ee := &nfsx.NfsError{Status: nfsx.V3ERR_ACCES}

// 	err := checkShadow("", er, ee)
// 	if err != ee {
// 		t.Error(err)
// 	}
// }

// func TestShadowErrorDifferentValueDifferentInstances(t *testing.T) {
// 	er := &nfsx.NfsError{Status: nfsx.V3ERR_ACCES}
// 	ee := &nfsx.NfsError{Status: nfsx.V3ERR_BADHANDLE}

// 	testShadow(t, er, ee)
// }

// func testShadow(t *testing.T, er, ee error) {
// 	err := checkShadow("", er, ee)

// 	errShadow, ok := err.(*ShadowError)
// 	if !ok {
// 		t.Errorf("Expected ShadowError, got %v", err)
// 	}

// 	if errShadow.ErrRefr != er || errShadow.ErrElfs != ee {
// 		t.Error(err)
// 	}
// }
