package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"in-backend/services/profile/configs"
	"in-backend/services/profile/interfaces"
	"in-backend/services/profile/models"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// KlentyProvider describes the methods to interact with the Klenty CRM platform
type KlentyProvider interface {
	CreateProspect(c *models.User) error
	UpdateProspect(c *models.User) error
	StartCadence(email, role string) error
}

type klentyProvider struct {
	config configs.Config
	client interfaces.HTTPClient
}

var (
	klentyURL string = "https://app.klenty.com/apis/v1/user/agnes@hubbedin.com"
)

// NewKlenty creates and returns a new KlentyProvider
func NewKlenty(cfg configs.Config, client interfaces.HTTPClient) KlentyProvider {
	return &klentyProvider{
		config: cfg,
		client: client,
	}
}

func (p *klentyProvider) CreateProspect(u *models.User) error {
	url := klentyURL + "/prospects"
	reqBody, err := json.Marshal(map[string]interface{}{
		"Email":     u.Email,
		"FirstName": u.FirstName,
		"LastName":  u.LastName,
	})
	if err != nil {
		return err
	}

	req, err := p.newRequest("POST", url, bytes.NewBuffer(reqBody))
	res, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = getKlentyRequestError(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (p *klentyProvider) UpdateProspect(u *models.User) error {
	url := auth0URL + "/prospects/" + url.QueryEscape(u.Email)
	reqBody, err := json.Marshal(map[string]string{
		"FirstName":   u.FirstName,
		"LastName":    u.LastName,
		"Phone":       u.ContactNumber,
		"Location":    u.Candidate.ResidenceCity,
		"Country":     u.Candidate.Nationality,
		"LinkedinURL": u.Candidate.LinkedInURL,
	})
	if err != nil {
		return err
	}

	req, err := p.newRequest("POST", url, bytes.NewBuffer(reqBody))
	res, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = getKlentyRequestError(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (p *klentyProvider) StartCadence(email, role string) error {
	var cadence string
	if role == "Candidate" {
		cadence = p.config.Klenty.CandidateSignupCadence
	}
	if role == "Company" {
		cadence = p.config.Klenty.CompanySignupCadence
	}

	url := klentyURL + "/startCadence"
	reqBody, err := json.Marshal(map[string]string{
		"Email":       email,
		"cadenceName": cadence,
	})
	if err != nil {
		return err
	}

	req, err := p.newRequest("POST", url, bytes.NewBuffer(reqBody))
	res, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = getKlentyRequestError(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (p *klentyProvider) newRequest(method, url string, buf *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-api-key", p.config.Klenty.APIKey)
	req.Header.Add("content-type", "application/json")
	return req, nil
}

func getKlentyRequestError(r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	jsonBody := make(map[string]bool)
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return err
	}

	if jsonBody["status"] != true {
		return errors.New("Request failed")
	}
	return nil
}
