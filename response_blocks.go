package convai

type ResponseBlock struct {
	Type     string `json:"type" mapstructure:"type" msgpack:"type"`
	Data     string `json:"data" mapstructure:"data" msgpack:"data"`
	Position int    `json:"pos" mapstructure:"pos" msgpack:"pos"`
}
