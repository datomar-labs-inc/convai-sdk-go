package byob

import (
	uuid "github.com/satori/go.uuid"
)

// BrainSignal is the collection of information that the brain requires to process a request
type BrainSignal struct {
	EnvironmentID   uuid.UUID       `json:"environmentId" mapstructure:"environmentId" msgpack:"environmentId"`
	ContextModifier ContextModifier `json:"contextModifier" mapstructure:"contextModifier" msgpack:"contextModifier"`
	PlatformID      string          `json:"platformId" mapstructure:"platformId" msgpack:"platformId"`
	OriginPlatform  string          `json:"originPlatform" mapstructure:"originPlatform" msgpack:"originPlatform"`
	Text            string          `json:"text" mapstructure:"text" msgpack:"text"`
	RawRequest      interface{}     `json:"rawRequest" mapstructure:"rawRequest" msgpack:"rawRequest"`
}
