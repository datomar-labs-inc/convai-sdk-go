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

type TriggerRequest struct {
	ContextModifier *ContextModifier `json:"contextModifier,omitempty"`
	ChannelID       string           `json:"channelId"`
	Text            string           `json:"text"`
	IsStart         bool             `json:"isStart"`
	IsTrigger       bool             `json:"isTrigger"`
	Source          interface{}      `json:"source,omitempty"`
}

type BroadcastResult struct {
	Status string `json:"status" msgpack:"status" mapstructure:"status"`
	Users  int    `json:"users" msgpack:"users" mapstructure:"users"`
}

type MergeUsersRequest struct {
	SuperUserIDs []uuid.UUID `json:"superUserIds" binding:"required"`
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

type Execution struct {
	ID                uuid.UUID              `json:"id" msgpack:"id" mapstructure:"id"`
	UserID            uuid.UUID              `json:"userId" msgpack:"userId" mapstructure:"userId"`
	ChannelUserID     string                 `json:"channelUserId" msgpack:"channelUserId" mapstructure:"channelUserId"`
	SessionID         uuid.UUID              `json:"sessionId" msgpack:"sessionId" mapstructure:"sessionId"`
	EnvironmentID     uuid.UUID              `json:"environmentId" msgpack:"environmentId" mapstructure:"environmentId"`
	BlueprintID       uuid.UUID              `json:"blueprintId" msgpack:"blueprintId" mapstructure:"blueprintId"`
	Data              map[string]interface{} `json:"data" msgpack:"data" mapstructure:"data"`
	UserData          map[string]interface{} `json:"userData" msgpack:"userData" mapstructure:"userData"`
	SessionData       map[string]interface{} `json:"sessionData" msgpack:"sessionData" mapstructure:"sessionData"`
	Text              string                 `json:"text" msgpack:"text" mapstructure:"text"`
	Channel           string                 `json:"channel" msgpack:"channel" mapstructure:"channel"`
	Source            interface{}            `json:"source" msgpack:"source" mapstructure:"source"`
	IsStart           bool                   `json:"isStart" msgpack:"isStart" mapstructure:"isStart"`
	IsTrigger         bool                   `json:"isTrigger" msgpack:"isTrigger" mapstructure:"isTrigger"`
	Errors            []ExecError            `json:"errors" msgpack:"errors" mapstructure:"errors"`
	Response          *Response              `json:"response" msgpack:"response" mapstructure:"response"`
	Logs              []ExecutionLog         `json:"logs" msgpack:"logs" mapstructure:"logs"`
	ExecutionDuration int64                  `json:"executionDuration" msgpack:"executionDuration" mapstructure:"executionDuration"`
	StartTime         time.Time              `json:"startTime" msgpack:"startTime" mapstructure:"startTime"`
}

type ExecutionQueryResult struct {
	Executions []Execution `json:"executions" msgpack:"e" mapstructure:"executions"`
	Total      int         `json:"total" msgpack:"t" mapstructure:"total"`
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
	Limit  int          `json:"limit" msgpack:"limit" mapstructure:"limit"`
	Offset int          `json:"offset" msgpack:"offset" mapstructure:"offset"`
}

type UQBuilder struct {
	mode         int64
	checks       []QueryCheck
	currentCheck *QueryCheck
	limit        int
	offset       int
}

func UserQueryBuilder() *UQBuilder {
	return &UQBuilder{
		limit:  10,
		offset: 0,
	}
}

func (u *UQBuilder) All() *UQBuilder {
	u.mode = UQMAll
	return u
}

func (u *UQBuilder) Any() *UQBuilder {
	u.mode = UQMAny
	return u
}

func (u *UQBuilder) None() *UQBuilder {
	u.mode = UQMNone
	return u
}

func (u *UQBuilder) Where(field string) *UQBuilder {
	if u.currentCheck != nil {
		u.checks = append(u.checks, *u.currentCheck)
	}

	u.currentCheck = &QueryCheck{
		Field:     field,
		Operation: UQEquals,
		Values:    []string{},
	}

	return u
}

func (u *UQBuilder) Equals(values ...string) *UQBuilder {
	if u.currentCheck == nil {
		panic("cannot call equals without calling where first")
	}

	u.currentCheck.Operation = UQEquals
	u.currentCheck.Values = values

	return u
}

func (u *UQBuilder) NotEquals(values ...string) *UQBuilder {
	if u.currentCheck == nil {
		panic("cannot call not equals without calling where first")
	}

	u.currentCheck.Operation = UQNotEquals
	u.currentCheck.Values = values

	return u
}

func (u *UQBuilder) StartsWith(values ...string) *UQBuilder {
	if u.currentCheck == nil {
		panic("cannot call starts with without calling where first")
	}

	u.currentCheck.Operation = UQStartsWith
	u.currentCheck.Values = values

	return u
}

func (u *UQBuilder) GreaterThan(values ...string) *UQBuilder {
	if u.currentCheck == nil {
		panic("cannot call greater than without calling where first")
	}

	u.currentCheck.Operation = UQGreaterThan
	u.currentCheck.Values = values

	return u
}

func (u *UQBuilder) LessThan(values ...string) *UQBuilder {
	if u.currentCheck == nil {
		panic("cannot call less than without calling where first")
	}

	u.currentCheck.Operation = UQLessThan
	u.currentCheck.Values = values

	return u
}

func (u *UQBuilder) Exists() *UQBuilder {
	if u.currentCheck == nil {
		panic("cannot call exists without calling where first")
	}

	u.currentCheck.Operation = UQExists

	return u
}

func (u *UQBuilder) NotExists() *UQBuilder {
	if u.currentCheck == nil {
		panic("cannot call not exists without calling where first")
	}

	u.currentCheck.Operation = UQNotExists
	return u
}

func (u *UQBuilder) Limit(limit int) *UQBuilder {
	u.limit = limit
	return u
}

func (u *UQBuilder) Offset(offset int) *UQBuilder {
	u.offset = offset
	return u
}

func (u *UQBuilder) Build() *UserQuery {
	if u.currentCheck != nil {
		u.checks = append(u.checks, *u.currentCheck)
	}

	return &UserQuery{
		Checks: u.checks,
		Mode:   u.mode,
		Limit:  u.limit,
		Offset: u.offset,
	}
}
