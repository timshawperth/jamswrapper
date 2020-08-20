package jamswrapper

import (
	"fmt"
	"time"
)

type Folder struct {
	Description      string          `json:"description,omitempty"`
	Parameters       []Parameter     `json:"parameters,omitempty"`
	FolderID         int             `json:"folderID,omitempty"`
	FolderName       string          `json:"folderName"`
	ParentFolderID   int             `json:"parentFolderID,omitempty"`
	ParentFolderName string          `json:"parentFolderName"`
	QualifiedName    string          `json:"qualifiedName,omitempty"`
	Elements         []Element       `json:"elements,omitempty"`
	Properties       []PropertyValue `json:"properties,omitempty"`
	TagIds           []int           `json:"tagIds,omitempty"`
	Tags             []Tags          `json:"tags,omitempty"`
	InheritedTagIds  []int           `json:"inheritedTagIds,omitempty"`
	LastChangeUTC    time.Time       `json:"lastChangeUTC,omitempty"`
	LastChangedBy    string          `json:"lastChangedBy,omitempty"`
	ACL              ACL             `json:"acl,omitempty"`
}

func (f *Folder) InsertInto(connection *AuthResponse) error {
	target := fmt.Sprintf("%s/folder", destinationServer)
	request := setupClient(connection)
	response, err := request.SetBody(f).Post(target)
	err = manageResponse(response, err)
	return err
}

func (f *Folder) FolderCopy(datacenter string, foldername string) error {
	f.FolderName = foldername
	f.ParentFolderName = datacenter
	return f.InsertInto(authenticationResponse)
}
