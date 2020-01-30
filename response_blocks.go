package convai

type ResponseBlock struct {
	Type     string `json:"t" mapstructure:"t" msgpack:"t"`
	Data     string `json:"d" mapstructure:"d" msgpack:"d"`
	Position int    `json:"p" mapstructure:"p" msgpack:"p"`
}
