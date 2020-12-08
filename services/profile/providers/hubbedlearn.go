package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"in-backend/services/profile/configs"
	"in-backend/services/profile/interfaces"
	"in-backend/services/profile/models"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

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

	name := c.FirstName + " " + c.LastName
	alpha, err := regexp.Compile("[^a-zA-Z\\s]+")
	if err != nil {
		return nil, err
	}
	name = alpha.ReplaceAllString(name, "")[:19]

	password, err := password.Generate(10, 2, 2, false, false)
	if err != nil {
		return nil, err
	}

	contact := c.ContactNumber
	if contact == "" {
		rand.Seed(time.Now().UnixNano())
		min := 1000000
		max := 9999999
		contact = "888" + strconv.Itoa(rand.Intn(max-min+1)+min)
	}

	reqBody, err := json.Marshal(map[string]string{
		"email":        c.Email,
		"phone":        contact,
		"password":     password,
		"name":         name,
		"business_key": p.config.HubbedLearn.ApiKey,
		"classroom":    "HUBBEDIN-BATCH-3",
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

	err = getHubbedLearnRequestError(bytes.NewBuffer(reqBody), res.Body)
	if err != nil {
		return nil, err
	}

	return &password, nil
}

func getHubbedLearnRequestError(in *bytes.Buffer, r io.Reader) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	jsonBody := make(map[string]string)
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return err
	}

	msg := jsonBody["message"]
	if msg != "Created Successfully" {
		return fmt.Errorf("%s: %+v", msg, in)
	}
	return nil
}
