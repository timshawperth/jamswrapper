package jamswrapper

import (
	"errors"
	"testing"
)

func TestProperty(t *testing.T) {
	var perr *RequestError

	auth := &Authentication{Username: "tim.shaw", Password: "1Lamentations2:13"}
	resp, _ := auth.Login("10.1.1.31")

	r, err := GetPropertiesFor(resp, "Folder")
	if err != nil {
		if errors.As(err, &perr) {
			if perr.Message != nil {
				t.Logf("%s\n", *perr.Message)
			}

			if perr.Messages != nil {
				for _, v := range perr.Messages {
					t.Logf("%s\n", v)
				}
			}
		}
		t.Fail()
	}

	for _, v := range r {
		t.Log(v)
	}
}
