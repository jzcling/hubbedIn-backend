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
)

// KlentyProvider describes the methods to interact with the Klenty CRM platform
type KlentyProvider interface {
	GetProspect(c *models.User) (map[string]string, error)
	UpdateOrCreateProspect(c *models.User) error
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

type Prospect struct {
	ID            string `json:"id"`
	FullName      string `json:"FullName"`
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	MiddleName    string `json:"MiddleName"`
	Account       string `json:"Account"`
	Department    string `json:"Department"`
	Company       string `json:"Company"`
	CompanyDomain string `json:"CompanyDomain"`
	Title         string `json:"Title"`
	Location      string `json:"Location"`
	Phone         string `json:"Phone"`
	Email         string `json:"Email"`
	TwitterID     string `json:"TwitterId"`
	CompanyEmail  string `json:"CompanyEmail"`
	CompanyPhone  string `json:"CompanyPhone"`
	City          string `json:"City"`
	Country       string `json:"Country"`
	LinkedinURL   string `json:"LinkedinURL"`
	Tags          string `json:"Tags"`
	List          string `json:"List"`
	AssignTo      string `json:"assignTo"`
}

// NewKlenty creates and returns a new KlentyProvider
func NewKlenty(cfg configs.Config, client interfaces.HTTPClient) KlentyProvider {
	return &klentyProvider{
		config: cfg,
		client: client,
	}
}

func (p *klentyProvider) GetProspect(u *models.User) (map[string]string, error) {
	url := klentyURL + "/prospects"
	req, err := p.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("Email", u.Email)
	req.URL.RawQuery = q.Encode()
	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var prospect []map[string]string
	err = json.Unmarshal(body, &prospect)
	if err != nil {
		return nil, err
	}
	if len(prospect) > 0 {
		return prospect[0], nil
	}
	return nil, nil
}

func (p *klentyProvider) UpdateOrCreateProspect(u *models.User) error {
	prospect, err := p.GetProspect(u)
	if err != nil {
		return err
	}
	if prospect != nil {
		return p.UpdateProspect(u)
	} else {
		return p.CreateProspect(u)
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
	if err != nil {
		return err
	}
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
	url := klentyURL + "/prospects/" + u.Email
	reqBody, err := json.Marshal(map[string]string{
		"FirstName": u.FirstName,
		"LastName":  u.LastName,
		"Phone":     u.ContactNumber,
	})
	if err != nil {
		return err
	}

	req, err := p.newRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
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

	if len(cadence) > 0 {
		url := klentyURL + "/startCadence"
		reqBody, err := json.Marshal(map[string]string{
			"Email":       email,
			"cadenceName": cadence,
		})
		if err != nil {
			return err
		}

		req, err := p.newRequest("POST", url, bytes.NewBuffer(reqBody))
		if err != nil {
			return err
		}
		res, err := p.client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		err = getKlentyRequestError(res.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *klentyProvider) newRequest(method, url string, buf *bytes.Buffer) (*http.Request, error) {
	var req *http.Request
	var err error
	if buf != nil {
		req, err = http.NewRequest(method, url, buf)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
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
