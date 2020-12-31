package providers

import (
	"bytes"
	"encoding/json"
	"in-backend/services/profile/configs"
	"in-backend/services/profile/interfaces"
	"in-backend/services/profile/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Auth0Provider describes the methods to interact with the Auth0 identity provider
type Auth0Provider interface {
	GetToken() (map[string]interface{}, error)
	UpdateUser(token string, u *models.User) error
	SetUserRole(token, authID string, roles []string) error
}

type auth0Provider struct {
	config configs.Config
	client interfaces.HTTPClient
}

var (
	auth0URL        string = "https://hubbed-in.au.auth0.com"
	candidateRoleID string = "rol_zlZ3Ha3n1E7WIbNI"
	companyRoleID   string = "rol_I8Ol4fIrKZrFph2p"
	adminRoleID     string = "rol_NjsiJ7p3Z6IEhlRm"
)

// NewAuth0 creates and returns a new Auth0Provider
func NewAuth0(cfg configs.Config, client interfaces.HTTPClient) Auth0Provider {
	return &auth0Provider{
		config: cfg,
		client: client,
	}
}

func (p *auth0Provider) GetToken() (map[string]interface{}, error) {
	url := auth0URL + "/oauth/token"
	reqBody, err := json.Marshal(map[string]string{
		"client_id":     p.config.Auth0.MgmtClientID,
		"client_secret": p.config.Auth0.MgmtClientSecret,
		"audience":      auth0URL + "/api/v2/",
		"grant_type":    "client_credentials",
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonBody := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func (p *auth0Provider) UpdateUser(token string, u *models.User) error {
	url := auth0URL + "/api/v2/users/" + url.QueryEscape(u.AuthID)

	data := make(map[string]string)
	data["id"] = strconv.FormatUint(u.ID, 10)
	if u.CandidateID > 0 {
		data["candidateId"] = strconv.FormatUint(u.CandidateID, 10)
	}
	if u.JobCompanyID > 0 {
		data["companyId"] = strconv.FormatUint(u.JobCompanyID, 10)
	}

	reqBody, err := json.Marshal(map[string](map[string]string){
		"app_metadata": data,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("content-type", "application/json")

	res, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (p *auth0Provider) SetUserRole(token, authID string, roles []string) error {
	var val []string
	for _, role := range roles {
		switch role {
		case "Candidate":
			val = append(val, candidateRoleID)
		case "Company":
			val = append(val, companyRoleID)
		case "Admin":
			val = append(val, adminRoleID)
		}
	}

	url := auth0URL + "/api/v2/users/" + url.QueryEscape(authID) + "/roles"
	reqBody, err := json.Marshal(map[string]([]string){
		"roles": val,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}
