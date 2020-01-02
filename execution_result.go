package byob

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
	LinkSourceID    *int64           `json:"lsid" mapstructure:"lsid" msgpack:"lsid"`
	LinkDestID      *int64           `json:"ldid" mapstructure:"ldid" msgpack:"ldid"`
	ExecutionType   int16            `json:"ety" mapstructure:"ety" msgpack:"ety"`
	ExecutionTime   int64            `json:"eti" mapstructure:"eti" msgpack:"eti"`
	ContextModifier *ContextModifier `json:"cmo" mapstructure:"cmo" msgpack:"cmo"`
}
