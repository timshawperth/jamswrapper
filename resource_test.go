// agent_test
package jamswrapper

import (
	"errors"
	"testing"
)

func TestResource(t *testing.T) {
	var perr *RequestError

	auth := &Authentication{Username: "tim.shaw", Password: "1Lamentations2:13"}
	resp, _ := auth.Login("10.1.1.31")

	resource := &Resource{ResourceName: "TimsNewResource",
		QuantityAvailable: 99}

	err := resource.InsertInto(resp)
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
}
