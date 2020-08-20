package jamswrapper

import (
	"fmt"
	"time"
)

type Agent struct {
	AgentName        string     `json:"agentName"`
	AgentTypeName    string     `json:"agentTypeName"`
	AgentType        string     `json:"agentType,omitempty"`
	PlatformTypeName string     `json:"platformTypeName"`
	AgentPlatform    string     `json:"agentPlatform,omitempty"`
	AgentID          int        `json:"agentID,omitempty"`
	Description      string     `json:"description,omitempty"`
	LastChangeUTC    time.Time  `json:"lastChangeUTC,omitempty"`
	LastChangedBy    string     `json:"lastChangedBy,omitempty"`
	JobCount         int        `json:"jobCount,omitempty"`
	JobLimit         int        `json:"jobLimit"`
	LicenseAllocated bool       `json:"licenseAllocated,omitempty"`
	Online           bool       `json:"online,omitempty"`
	Version          string     `json:"version,omitempty"`
	TimeOffset       string     `json:"timeOffset,omitempty"`
	AgentState       string     `json:"agentState,omitempty"`
	Properties       []PropertyValue `json:"properties,omitempty"`
	IsExecutionAgent bool       `json:"isExecutionAgent,omitempty"`
	ACL              ACL        `json:"acl,omitempty"`
}

func (a *Agent) InsertInto(connection *AuthResponse) error {
	target := fmt.Sprintf("%s/agent", destinationServer)
	request := setupClient(connection)
	response, err := request.SetBody(a).Post(target)
	err = manageResponse(response, err)
	return err
}
