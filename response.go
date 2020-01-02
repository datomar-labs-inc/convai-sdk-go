package byob

type Response struct {
	Messages []Message `json:"messages" mapstructure:"messages" msgpack:"messages"`
}

type ResponseConfig struct {
	Basic           []MessageConfig            `json:"messages" mapstructure:"messages" msgpack:"messages"`
	SendNow         bool                       `json:"sendNow" mapstructure:"sendNow" msgpack:"sendNow"`
	CustomResponses map[string][]MessageConfig `json:"customResponses" mapstructure:"customResponses" msgpack:"customResponses"`
}

type Message struct {
	Text        string          `json:"text" mapstructure:"text" msgpack:"text"`
	TypingTime  float64         `json:"typingTime" mapstructure:"typingTime" msgpack:"typingTime"`
	ShouldBatch bool            `json:"shouldBatch" mapstructure:"shouldBatch" msgpack:"shouldBatch"`
	GraphID     *int64          `json:"graphId" mapstructure:"graphId" msgpack:"graphId"`
	NodeID      *int64          `json:"nodeId" mapstructure:"nodeId" msgpack:"nodeId"`
	Blocks      []ResponseBlock `json:"blocks" mapstructure:"blocks" msgpack:"blocks"`
}

type MessageConfig struct {
	Text       string  `json:"text" mapstructure:"text" msgpack:"text"`
	TypingTime float64 `json:"typingTime" mapstructure:"typingTime" msgpack:"typingTime"`
}

