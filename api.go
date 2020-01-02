package byob

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%d: %s", a.Code, a.Message)
}

func NewAPIClient(apiKey string) *Client {
	return &Client{
		baseURL: "https://api.convai.dev/api/v1",
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func NewCustomAPIClient(apiKey string, baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *Client) makeRequestWithBody(method, url string, body interface{}, out interface{}) (*APIError, error) {
	jsb, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.baseURL, url), bytes.NewReader(jsb))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	rsb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var resErr APIError

		err = json.Unmarshal(rsb, &resErr)
		if err != nil {
			return nil, err
		}

		return &resErr, nil
	} else {
		err = json.Unmarshal(rsb, out)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func (c *Client) GetReachableUsers(query *UserQuery) (*ReachableUserResult, error) {
	var res ReachableUserResult

	apiErr, err := c.makeRequestWithBody("POST", "/reachable-users", query, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) Broadcast(input *BroadcastInput) (*BroadcastResult, error) {
	var res BroadcastResult

	apiErr, err := c.makeRequestWithBody("POST", "/broadcast", input, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) UpdateUserData(superUserId string, input *UpdateUserDataInput) (map[string]interface{}, error) {
	var res map[string]interface{}

	apiErr, err := c.makeRequestWithBody("PUT", fmt.Sprintf("/users/%s/data", superUserId), input, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return res, nil
}
