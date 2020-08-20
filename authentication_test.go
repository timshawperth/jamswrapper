package jamswrapper

import "testing"

func TestAuthentication(t *testing.T) {
	auth := &Authentication{Username: "tim.shaw", Password: "1Lamentations2:13"}

	resp, ok := auth.Login("10.1.1.31")
	if ok != nil {
		t.Errorf("%v", ok)
	} else {
		t.Logf("Response: %+v", resp)
	}
}
