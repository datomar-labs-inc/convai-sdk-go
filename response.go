package convai

import "encoding/xml"

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
	Data        XMLResponse     `json:"data" mapstructure:"data" msgpack:"data"`
	Seq         int             `json:"seq" mapstructure:"seq" msgpack:"seq"`
}

type MessageConfig struct {
	Text       string  `json:"text" mapstructure:"text" msgpack:"text"`
	TypingTime float64 `json:"typingTime" mapstructure:"typingTime" msgpack:"typingTime"`
}

type XMLResponse struct {
	XMLName         xml.Name            `xml:"response" json:"-" msgpack:"-" mapstructure:"-"`
	InnerXML        string              `xml:",innerxml" json:"-" msgpack:"-" mapstructure:"-"`
	Message         string              `json:"message" msgpack:"message" mapstructure:"message"`
	QuickReplies    []XMLQR             `xml:"qr" json:"quickReplies" msgpack:"quickReplies" mapstructure:"quickReplies"`
	Phone           []XMLPhone          `xml:"phone" json:"phone" msgpack:"phone"`
	Cards           []XMLCard           `xml:"card" json:"cards" msgpack:"cards"`
	CardCollections []XMLCardCollection `xml:"cards" json:"cardCollections" msgpack:"cardCollections"`
	Images          []XMLImage          `xml:"image" json:"images" msgpack:"images"`
}

type XMLQR struct {
	XMLName  xml.Name `xml:"qr" json:"-" msgpack:"-" mapstructure:"-"`
	Text     string   `xml:",innerxml" json:"text" msgpack:"text" mapstructure:"text"`
	Value    *string  `xml:"value,attr,omitempty" json:"value,omitempty" msgpack:"value,omitempty" mapstructure:"value,omitempty"`
	Phone    bool     `xml:"phone,attr,omitempty" json:"phone,omitempty" msgpack:"phone,omitempty" mapstructure:"phone,omitempty"`
	Email    bool     `xml:"email,attr,omitempty" json:"email,omitempty" msgpack:"email,omitempty" mapstructure:"email,omitempty"`
	Image    *string  `xml:"image,attr,omitempty" json:"image,omitempty" msgpack:"image,omitempty" mapstructure:"image,omitempty"`
	ImageURL *string  `xml:"imageURL,attr,omitempty" json:"imageUrl,omitempty" msgpack:"imageUrl,omitempty" mapstructure:"imageUrl,omitempty"`
}

type XMLPhone struct {
	XMLName xml.Name `xml:"phone" json:"-" msgpack:"-" mapstructure:"-"`
	Number  string   `xml:",innerxml" json:"number" msgpack:"number" mapstructure:"number"`
	Display *string  `xml:"display,omitempty" json:"display,omitempty" msgpack:"display,omitempty"`
}

type XMLCardCollection struct {
	XMLName xml.Name  `xml:"cards" json:"-" msgpack:"-" mapstructure:"-"`
	Cards   []XMLCard `xml:"card" json:"cards" msgpack:"cards"`
}

type XMLCard struct {
	XMLName  xml.Name    `xml:"card" json:"-" msgpack:"-" mapstructure:"-"`
	Title    string      `xml:"title" json:"title" msgpack:"title" mapstructure:"title"`
	Subtitle *string     `xml:"subtitle,omitempty" json:"subtitle,omitempty" msgpack:"subtitle,omitempty" mapstructure:"subtitle,omitempty"`
	Image    *XMLImage   `xml:"image,omitempty" json:"image,omitempty" msgpack:"image,omitempty"`
	Buttons  []XMLButton `xml:"button" json:"buttons" msgpack:"button"`
}

type XMLImage struct {
	XMLName xml.Name `xml:"image" json:"-" msgpack:"-" mapstructure:"-"`
	ID      string   `xml:"id,attr" json:"id" msgpack:"id"`
	Width   *uint64  `xml:"width,attr,omitempty" json:"width" msgpack:"width"`
	Height  *uint64  `xml:"height,attr,omitempty" json:"height" msgpack:"height"`
	X       *uint64  `xml:"x,attr,omitempty" json:"x" msgpack:"x"`
	Y       *uint64  `xml:"y,attr,omitempty" json:"y" msgpack:"y"`
}

type XMLButton struct {
	XMLName xml.Name `xml:"button" json:"-" msgpack:"-" mapstructure:"-"`
	Text    string   `xml:",innerxml" json:"text" msgpack:"text" mapstructure:"text"`
	Value   *string  `xml:"value,attr,omitempty" json:"value,omitempty" msgpack:"value,omitempty" mapstructure:"value,omitempty"`
	URL     *string  `xml:"url,attr,omitempty" json:"url,omitempty" msgpack:"url,omitempty" mapstructure:"url,omitempty"`
}
