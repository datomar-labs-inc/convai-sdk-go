package byob

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ResponseBlock struct {
	Type string `json:"t" mapstructure:"t" msgpack:"t"`
	Data string `json:"d" mapstructure:"d" msgpack:"d"`
}

func NewResponseBlock(t string, data interface{}) *ResponseBlock {
	jsb, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return &ResponseBlock{Type: t, Data: string(jsb)}
}

func (rb *ResponseBlock) Encode() string {
	jsb, err := json.Marshal(rb)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("!!cblk!!%s!!end!!", string(jsb))
}

func (rb *ResponseBlock) Decode(out interface{}) error {
	err := json.Unmarshal([]byte(rb.Data), out)
	return err
}

func ExtractBlocks(msg string) ([]ResponseBlock, string) {
	var blocks []ResponseBlock

	msg = strings.TrimSpace(msg)
	reconstructedMsg := ""

	ignoreFirst := !strings.HasPrefix(msg, "!!cblk!!")
	blockStrings := strings.Split(msg, "!!cblk!!")

	for i, bs := range blockStrings {
		if i == 0 && ignoreFirst {
			reconstructedMsg += bs
			continue
		}

		blockParts := strings.Split(bs, "!!end!!")

		if len(blockParts) == 2 {
			reconstructedMsg += blockParts[1]
		}

		var block ResponseBlock

		err := json.Unmarshal([]byte(blockParts[0]), &block)
		if err != nil {
			fmt.Println("decode error", err)
			continue
		}

		blocks = append(blocks, block)
	}

	return blocks, reconstructedMsg
}
