// resource
package jamswrapper

import (
	"fmt"
	"time"
)

type Resource struct {
	LastChangeUTC     time.Time         `json:"lastChangeUTC,omitempty"`
	LastChangedBy     string            `json:"lastChangedBy,omitempty"`
	ResourceName      string            `json:"resourceName"`
	ResourceID        int               `json:"resourceID,omitempty"`
	QuantityAvailable int               `json:"quantityAvailable"`
	QuantityInUse     int               `json:"quantityInUse,omitempty"`
	AgentSpecific     bool              `json:"agentSpecific,omitempty"`
	Description       string            `json:"description,omitempty"`
	ResourceDetails   []ResourceDetails `json:"resourceDetails,omitempty"`
	ACL               ACL               `json:"acl,omitempty"`
}

type ResourceDetails struct {
	AgentName         string `json:"agentName"`
	AgentID           int    `json:"agentId,omitempty"`
	QuantityAvailable int    `json:"quantityAvailable"`
	QuantityInUse     int    `json:"quantityInUse,omitempty"`
}

func (r *Resource) InsertInto(connection *AuthResponse) error {
	target := fmt.Sprintf("%s/resource", destinationServer)
	request := setupClient(connection)
	response, err := request.SetBody(r).Post(target)
	err = manageResponse(response, err)
	return err
}

func (r *Resource) ResourceCopy(resourceName string) error {
	r.ResourceName = resourceName
	return r.InsertInto(authenticationResponse)
}
