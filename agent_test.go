// agent_test
package jamswrapper

import (
	"errors"
	"testing"
)

func TestAgent(t *testing.T) {
	var perr *RequestError

	auth := &Authentication{Username: "tim.shaw", Password: "1Lamentations2:13"}
	resp, _ := auth.Login("10.1.1.31")

	agent := &Agent{AgentName: "Agent99",
		AgentTypeName:    "Outgoing",
		PlatformTypeName: "Windows",
		JobLimit:         99999}

	err := agent.InsertInto(resp)
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
