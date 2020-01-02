package byob

type ReachableUserResult struct {
	Count uint64 `json:"count" msgpack:"count" mapstructure:"count"`
}

type BroadcastInput struct {
	BroadcastType   string           `json:"broadcastType" msgpack:"broadcastType" mapstructure:"broadcastType"`
	ContextModifier *ContextModifier `json:"contextModifier,omitempty" msgpack:"contextModifier" mapstructure:"contextModifier"`
	Platform        string           `json:"platform" msgpack:"platform" mapstructure:"platform" binding:"required"`
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
