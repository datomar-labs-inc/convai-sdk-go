package convai

import (
	uuid "github.com/satori/go.uuid"
)

// Signal is the collection of information that the brain requires to process a request
type Signal struct {
	EnvironmentID   uuid.UUID       `json:"environmentId" mapstructure:"environmentId" msgpack:"environmentId"`
	ContextModifier ContextModifier `json:"contextModifier" mapstructure:"contextModifier" msgpack:"contextModifier"`
	ChannelID       string          `json:"channelId" mapstructure:"channelId" msgpack:"channelId"`
	Channel         string          `json:"channel" channel:"channel" msgpack:"channel"`
	Text            string          `json:"text" mapstructure:"text" msgpack:"text"`
	RawRequest      interface{}     `json:"rawRequest" mapstructure:"rawRequest" msgpack:"rawRequest"`
}
