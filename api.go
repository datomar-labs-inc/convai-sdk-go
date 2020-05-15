package convai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
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

// Function to create a Super User and associated channel users
func (c *Client) CreateSuperUser(request *CreateCombinedUserRequest) (*CreateCombinedUserResult, *APIError) {
	var res CreateCombinedUserResult

	apiErr, err := c.makeRequestWithBody("POST", "/users/super/create", request, &res)
	if err != nil {
		return nil, &APIError{Code:500, Message: fmt.Sprintf("SDK error: %s", err.Error())}
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

// Function to create Channel Users for an existing super user
func (c *Client) CreateChannelUsers(request *CreateChannelUsersRequest) (*CreateChannelUsersResult, *APIError) {
	var res CreateChannelUsersResult

	apiErr, err := c.makeRequestWithBody("POST", "/users/channel/create", request, &res)
	if err != nil {
		return nil, &APIError{Code:500, Message: fmt.Sprintf("SDK error: %s", err.Error())}
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) QueryExecutions(matcher *ExecutionMatcher) (*ExecutionQueryResult, error) {
	var res ExecutionQueryResult

	apiErr, err := c.makeRequestWithBody("POST", "/executions/query", matcher, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) Trigger(req *TriggerRequest) (*Execution, error) {
	var res Execution

	apiErr, err := c.makeRequestWithBody("POST", "/executions/trigger", req, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) Broadcast(input *BroadcastInput) (*BroadcastResult, error) {
	var res BroadcastResult

	apiErr, err := c.makeRequestWithBody("POST", "/executions/broadcast", input, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) QueryUsers(query *UserQuery) (*UserQueryResult, error) {
	var res UserQueryResult

	apiErr, err := c.makeRequestWithBody("POST", "/users/super/query", query, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) QueryUsersReachable(query *UserQuery) (*ReachableUserResult, error) {
	var res ReachableUserResult

	apiErr, err := c.makeRequestWithBody("POST", "/users/super/query/reachable", query, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) MergeUsers(req *MergeUsersRequest) (*SuperUser, error) {
	var res SuperUser

	apiErr, err := c.makeRequestWithBody("POST", "/users/super/merge", req, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) DeleteSuperUser(id uuid.UUID) (*SuperUser, error) {
	var res SuperUser

	apiErr, err := c.makeRequestWithBody("DELETE", fmt.Sprintf("/users/super/%s", id.String()), nil, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) UpdateUserData(superUserId string, input *UpdateUserDataInput) (*SuperUser, error) {
	var res SuperUser

	apiErr, err := c.makeRequestWithBody("PUT", fmt.Sprintf("/users/super/%s", superUserId), input, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) DeleteChannelUser(userID string) (*ChannelUser, error) {
	var res ChannelUser

	apiErr, err := c.makeRequestWithBody("DELETE", fmt.Sprintf("/users/channel/%s", userID), nil, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) UpdateSession(userID string, input *UpdateUserDataInput) (*Session, error) {
	var res Session

	apiErr, err := c.makeRequestWithBody("PUT", fmt.Sprintf("/users/session/%s", userID), input, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}

func (c *Client) DeleteSession(userID string) (*Session, error) {
	var res Session

	apiErr, err := c.makeRequestWithBody("DELETE", fmt.Sprintf("/users/session/%s", userID), nil, &res)
	if err != nil {
		return nil, err
	} else if apiErr != nil {
		return nil, apiErr
	}

	return &res, nil
}