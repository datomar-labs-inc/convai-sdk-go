package convai

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ReachableUserResult struct {
	Count uint64 `json:"count" msgpack:"count" mapstructure:"count"`
}

type BroadcastInput struct {
	BroadcastType   string           `json:"broadcastType" msgpack:"broadcastType" mapstructure:"broadcastType"`
	ContextModifier *ContextModifier `json:"contextModifier,omitempty" msgpack:"contextModifier" mapstructure:"contextModifier"`
	Channel         string           `json:"channel" msgpack:"channel" mapstructure:"channel" binding:"required"`
	UserQuery       UserQuery        `json:"userQuery" msgpack:"userQuery" mapstructure:"userQuery" binding:"required"`
}

type BroadcastResult struct {
	Status string `json:"status" msgpack:"status" mapstructure:"status"`
	Users  int    `json:"users" msgpack:"users" mapstructure:"users"`
}

const (
	UQEquals = iota
	UQExists
	UQNotExists
	UQNotEquals
	UQStartsWith
	UQGreaterThan
	UQLessThan
)

const (
	UQMAny = iota
	UQMAll
	UQMNone
)

type QueryCheck struct {
	Field     string   `json:"field" msgpack:"field" mapstructure:"field"`
	Operation int64    `json:"operation" msgpack:"operation" mapstructure:"operation"`
	Values    []string `json:"values" msgpack:"values" mapstructure:"values"`
}

type UserQuery struct {
	Checks []QueryCheck `json:"checks" msgpack:"checks" mapstructure:"checks"`
	Mode   int64        `json:"mode" msgpack:"mode" mapstructure:"mode"`
}

type UpdateUserDataInput struct {
	// A list of keys and their values to be set on the user data
	// If the key exists it will be overwritten
	Set map[string]interface{} `json:"set" msgpack:"set" mapstructure:"set"`

	// A list of keys to be remove from the user data, if the keys did not exist, nothing happens
	Delete []string `json:"delete" msgpack:"delete" mapstructure:"delete"`
}

type UserQueryResult struct {
	Users []SuperUser
	Count uint64
}

type SuperUser struct {
	ID            uuid.UUID              `json:"id" msgpack:"id"`
	EnvironmentID uuid.UUID              `json:"environment_id" msgpack:"environment_id"`
	Data          map[string]interface{} `json:"data" msgpack:"data"`
	CreatedAt     *time.Time             `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time             `json:"updatedAt,omitempty"`
	ChannelUsers  []ChannelUser          `json:"channelUsers"`
}

type ChannelUser struct {
	ChannelId     string                 `json:"channelId"`
	EnvironmentID uuid.UUID              `json:"environmentId"`
	Channel       string                 `json:"channel"`
	Data          map[string]interface{} `json:"data"`
	SuperUserID   uuid.UUID              `json:"superUserId"`
	Session       *Session               `json:"session"`
	CreatedAt     *time.Time             `json:"createdAt"`
	UpdatedAt     *time.Time             `json:"updatedAt"`
}
