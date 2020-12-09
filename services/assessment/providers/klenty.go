package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"in-backend/services/assessment/configs"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// KlentyProvider describes the methods to interact with the Klenty CRM platform
type KlentyProvider interface {
	CreateProspect(c *models.Candidate, pw string) error
	UpdateProspect(c *models.Candidate) error
	StartCadence(email string) error
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

func (p *klentyProvider) CreateProspect(c *models.Candidate, pw string) error {
	url := klentyURL + "/prospects"
	reqBody, err := json.Marshal(map[string]interface{}{
		"Email":     c.Email,
		"FirstName": c.FirstName,
		"LastName":  c.LastName,
		"CustomFields": []map[string]string{
			0: {"key": "hubbedlearn password", "value": pw},
		},
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

func (p *klentyProvider) UpdateProspect(c *models.Candidate) error {
	url := auth0URL + "/prospects/" + url.QueryEscape(c.Email)
	reqBody, err := json.Marshal(map[string]string{
		"FirstName":   c.FirstName,
		"LastName":    c.LastName,
		"Phone":       c.ContactNumber,
		"Location":    c.ResidenceCity,
		"Country":     c.Nationality,
		"LinkedinURL": c.LinkedInURL,
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

func (p *klentyProvider) StartCadence(email string) error {
	url := klentyURL + "/startCadence"
	reqBody, err := json.Marshal(map[string]string{
		"Email":       email,
		"cadenceName": p.config.Klenty.SignupCadence,
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
	req.Header.Add("x-api-key", p.config.Klenty.ApiKey)
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
