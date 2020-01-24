package convai

import (
	"time"
)

type ExecutionResult struct {
	FinalContext  RequestContext `json:"finalContext" mapstructure:"finalContext" msgpack:"finalContext"`
	ExecutionLogs []ExecutionLog `json:"execLogs" mapstructure:"execLogs" msgpack:"execLogs"`
	ExecutionTime time.Duration  `json:"execTime" mapstructure:"execTime" msgpack:"execTime"`
}

// ExecutionLog logs a single execution of a node or link
type ExecutionLog struct {
	GraphID         int64            `json:"gid" mapstructure:"gid" msgpack:"gid"`
	NodeID          *int64           `json:"nid" mapstructure:"nid" msgpack:"nid"`
	LinkID    *int64           `json:"linkId" mapstructure:"linkId" msgpack:"linkId"`
	ExecutionType   int16            `json:"ety" mapstructure:"ety" msgpack:"ety"`
	ExecutionTime   int64            `json:"eti" mapstructure:"eti" msgpack:"eti"`
	ContextModifier *ContextModifier `json:"cmo" mapstructure:"cmo" msgpack:"cmo"`
}
