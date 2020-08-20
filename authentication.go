package jamswrapper

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshToken struct {
	GrantType    string `json:"grant_Type"`
	RefreshToken string `json:"refresh_Token"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type AuthError struct {
	ID      int
	Message string
}

var (
	destinationServer      string
	authenticationResponse *AuthResponse
)

func (e *AuthError) Error() string {
	return fmt.Sprintf("ID: %d, Message: %s", e.ID, e.Message)
}

func (a *Authentication) Login(jamsWebServer string) (*AuthResponse, error) {
	destinationServer = fmt.Sprintf("http://%s/jams/api", jamsWebServer)
	client := resty.New()

	resp, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(a).
		SetError(AuthError{}).
		SetResult(&AuthResponse{}).
		SetJSONEscapeHTML(false).
		Post(fmt.Sprintf("%s/authentication/login", destinationServer))

	if resp.Error() != nil {
		authError := resp.Error().(*AuthError)
		if authError.ID == 0 {
			authError.ID = resp.StatusCode()
			authError.Message = resp.Status()
		}
		err = authError
	}
	authenticationResponse = (resp.Result()).(*AuthResponse)
	return (resp.Result()).(*AuthResponse), err
}

func (a *AuthResponse) AuthHeader() string {
	return fmt.Sprintf("bearer %s", a.AccessToken)
}
