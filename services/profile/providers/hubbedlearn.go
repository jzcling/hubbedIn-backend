package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"in-backend/services/profile/interfaces"
	"in-backend/services/profile/configs"
	"in-backend/services/profile/models"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sethvargo/go-password/password"
)

// HubbedLearnProvider describes the methods to interact with the HubbedLearn platform
type HubbedLearnProvider interface {
	CreateUser(c *models.Candidate) (*string, error)
}

type hubbedlearnProvider struct {
	config configs.Config
	client interfaces.HTTPClient
}

var (
	hubbedlearnURL string = "http://staging2.hubbedlearn.com/api"
)

// NewHubbedLearn creates and returns a new HubbedLearnProvider
func NewHubbedLearn(cfg configs.Config, client interfaces.HTTPClient) HubbedLearnProvider {
	return &hubbedlearnProvider{
		config: cfg,
		client: client,
	}
}

func (p *hubbedlearnProvider) CreateUser(c *models.Candidate) (*string, error) {
	url := hubbedlearnURL + "/auth/student/create"
	password, err := password.Generate(10, 2, 2, false, false)
	if err != nil {
		return nil, err
	}
	reqBody, err := json.Marshal(map[string]string{
		"email":        c.Email,
		"phone":        c.ContactNumber,
		"password":     password,
		"name":         c.FirstName + " " + c.LastName,
		"business_key": p.config.HubbedLearn.ApiKey,
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

	err = getHubbedLearnRequestError(res.Body)
	if err != nil {
		return nil, err
	}

	return &password, nil
}

func getHubbedLearnRequestError(r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	jsonBody := make(map[string]string)
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return err
	}

	if jsonBody["message"] != "Created Successfully" {
		return errors.New("Request failed")
	}
	return nil
}
