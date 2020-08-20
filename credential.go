package jamswrapper

import (
	"fmt"
	"time"
)

type Credential struct {
	Description         string    `json:"description,omitempty"`
	Fingerprint         string    `json:"fingerprint,omitempty"`
	LastChangeUTC       time.Time `json:"lastChangeUTC,omitempty"`
	LogonUserName       string    `json:"logonUserName,omitempty"`
	Password            string    `json:"password,omitempty"`
	PrivateKey          string    `json:"privateKey,omitempty"`
	EncryptedPrivateKey string    `json:"encryptedPrivateKey,omitempty"`
	EncryptedPassword   string    `json:"encryptedPassword,omitempty"`
	PublicKey           string    `json:"publicKey,omitempty"`
	CredentialID        int       `json:"credentialID,omitempty"`
	CredentialName      string    `json:"credentialName"`
	KeyFileContent      string    `json:"keyFileContent,omitempty"`
	ACL                 ACL       `json:"acl,omitempty"`
}

func (a *Credential) InsertInto(connection *AuthResponse) error {
	target := fmt.Sprintf("%s/credential", destinationServer)
	request := setupClient(connection)
	response, err := request.SetBody(a).Post(target)
	err = manageResponse(response, err)
	return err
}

func (a *Credential) CredentialCopy(credential string) error {
	a.CredentialName = credential
	return a.InsertInto(authenticationResponse)
}
