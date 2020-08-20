// common
package jamswrapper

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Parameter struct {
	AllowEntry             bool            `json:"allowEntry"`
	DataType               string          `json:"dataType"`
	DefaultFormat          string          `json:"defaultFormat"`
	DefaultValue           string          `json:"defaultValue"`
	HelpText               string          `json:"helpText"`
	Hide                   bool            `json:"hide"`
	JobID                  int             `json:"jobID"`
	LastChangeUTC          time.Time       `json:"lastChangeUTC"`
	Length                 int             `json:"length"`
	MustFill               bool            `json:"mustFill"`
	ParameterOrigin        string          `json:"parameterOrigin"`
	ParamName              string          `json:"paramName"`
	Prompt                 string          `json:"prompt"`
	Required               bool            `json:"required"`
	Sequence               int             `json:"sequence"`
	Uppercase              bool            `json:"uppercase"`
	ValidationData         string          `json:"validationData"`
	ValidationType         string          `json:"validationType"`
	VariableID             int             `json:"variableID"`
	VariableName           string          `json:"variableName"`
	VariableNameIsRelative bool            `json:"variableNameIsRelative"`
	Properties             []PropertyValue `json:"properties"`
	EncryptionID           int             `json:"encryptionId"`
	NewEncryptedValue      string          `json:"newEncryptedValue"`
}

type Tags struct {
	TagID   int    `json:"tagId"`
	TagName string `json:"tagName"`
}

type GenericACL struct {
	Identifier string   `json:"identifier"`
	Inherited  bool     `json:"inherited"`
	Flags      string   `json:"flags"`
	AccessList []string `json:"accessList"`
}

type ACL struct {
	GenericACL []GenericACL `json:"genericACL,omitempty"`
}

type RequestError struct {
	Message    *string     `json:"message,omitempty"`
	Stacktrace interface{} `json:"stackTrace,omitempty"`
	Messages   []string    `json:"messages,omitempty"`
}

type JAMS interface {
	InsertInto(connection *AuthResponse) error
}

func (e *RequestError) Error() string {
	var retval string

	if e.Message != nil {
		return *e.Message
	} else {
		for _, v := range e.Messages {
			retval += fmt.Sprintf("%s", v)
		}
		return retval
	}
}

func setupClient(connection *AuthResponse) *resty.Request {
	client := resty.New()
	return client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", connection.AuthHeader()).
		SetError(&RequestError{}).
		SetJSONEscapeHTML(false)
}

func manageResponse(r *resty.Response, err error) error {
	if err != nil {
		return err
	}

	if r.Error() != nil {
		err = r.Error().(*RequestError)
	}

	return err
}
