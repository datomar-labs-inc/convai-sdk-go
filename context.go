package byob

import uuid "github.com/satori/go.uuid"

// RequestContext stores the context of a single byob request
type RequestContext struct {
	Flaggable
	ID              uuid.UUID              `json:"id" mapstructure:"id" msgpack:"id"`
	User            RequestUser            `json:"user" mapstructure:"user" msgpack:"user"`
	Session         Session                `json:"session" mapstructure:"session" msgpack:"session"`
	EnvironmentData map[string]interface{} `json:"envData" mapstructure:"envData" msgpack:"envData"`
	Text            string                 `json:"text" mapstructure:"text" msgpack:"text"`
	OriginPlatform  string                 `json:"originPlatform" mapstructure:"originPlatform" msgpack:"originPlatform"`
	OriginalRequest interface{}            `json:"originalRequest" mapstructure:"originalRequest" msgpack:"originalRequest"`
	IsStart         bool                   `json:"isStart" mapstructure:"isStart" msgpack:"isStart"`
	IsTrigger       bool                   `json:"isTrigger" mapstructure:"isTrigger" msgpack:"isTrigger"`
	Errors          []ExecError            `json:"errors" mapstructure:"errors" msgpack:"errors"`
	Response        *Response              `json:"response" mapstructure:"response" msgpack:"response"`

	// LastError is the error that ocurred during execution of the last node
	LastError *ExecError `json:"lastError" mapstructure:"lastError" msgpack:"lastError"`
}

// Request user is holds the information for a user making a bot request
type RequestUser struct {
	Flaggable
	ID         uuid.UUID `json:"id" mapstructure:"id" msgpack:"id"`
	PlatformID string    `json:"platformId" mapstructure:"platformId" msgpack:"platformId"`
	Name       string    `json:"name" mapstructure:"name" msgpack:"name"`
}

func (r *RequestContext) Error(err ExecError) {
	r.Errors = append(r.Errors, err)
	r.LastError = &err
}