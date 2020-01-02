package byob

import (
	"encoding/json"
	"time"
)

const (
	CMOContext = iota
	CMOSession
	CMOUser
	CMOEnvironment
)

const (
	CMOPSet = iota
	CMOPDelete
	CMOPClear
)

type ContextModifier struct {
	ContextChanges []ContextChange `json:"changes" mapstructure:"changes" msgpack:"changes"`
	Logs           []LogEntry      `json:"logs" mapstructure:"logs" msgpack:"logs"`
	Errors         []ExecError     `json:"errors" mapstructure:"errors" msgpack:"errors"`
}

type ContextChange struct {
	Type      int         `json:"type" mapstructure:"type" msgpack:"type"`
	Operation int         `json:"op" mapstructure:"op" msgpack:"op"`
	Key       string      `json:"key" mapstructure:"key" msgpack:"key"`
	Data      interface{} `json:"data" mapstructure:"data" msgpack:"data"`
}

type LogEntry struct {
	Level   int       `json:"level" mapstructure:"level" msgpack:"level"`
	Message string    `json:"message" mapstructure:"message" msgpack:"message"`
	Time    time.Time `json:"time" mapstructure:"time" msgpack:"time"`
}

type ExecError struct {
	// NotationType should be one of "node", "link", "other"
	ErrorType  string `json:"et" mapstructure:"et" msgpack:"et"`
	GraphID    int64  `json:"graphId" mapstructure:"graphId" msgpack:"graphId"`
	GraphName  string `json:"graphName" mapstructure:"graphName" msgpack:"graphName"`
	NodeID     *int64 `json:"nodeId" mapstructure:"nodeId" msgpack:"nodeId"`
	LinkSource *int64 `json:"linkSource" mapstructure:"linkSource" msgpack:"linkSource"`
	LinkDest   *int64 `json:"linkDest" mapstructure:"linkDest" msgpack:"linkDest"`
	Message    string `json:"message" mapstructure:"message" msgpack:"message"`
}

func NewContextModifier() *ContextModifier {
	return &ContextModifier{
		ContextChanges: []ContextChange{},
		Logs:           []LogEntry{},
		Errors:         []ExecError{},
	}
}

// MergeContextModifiers creates a new context modifier consisting of the parts of a and b
func MergeContextModifiers(a, b *ContextModifier) *ContextModifier {
	n := NewContextModifier()

	n.Logs = append(n.Logs, a.Logs...)
	n.Logs = append(n.Logs, b.Logs...)

	n.ContextChanges = append(n.ContextChanges, a.ContextChanges...)
	n.ContextChanges = append(n.ContextChanges, b.ContextChanges...)

	n.Errors = append(n.Errors, a.Errors...)
	n.Errors = append(n.Errors, b.Errors...)

	return n
}

func (cm *ContextModifier) AddOperation(ccType int, op int, key string, data interface{}) *ContextModifier {
	// Sanitize data for json
	if data != nil {
		jsb, err := json.Marshal(data)
		if err == nil {
			var newData map[string]interface{}

			err = json.Unmarshal(jsb, &newData)
			if err == nil {
				data = newData
			}
		}
	}

	cm.ContextChanges = append(cm.ContextChanges, ContextChange{
		Type:      ccType,
		Operation: op,
		Key:       key,
		Data:      data,
	})

	return cm
}

func (cm *ContextModifier) Apply(context *RequestContext) {
	for _, change := range cm.ContextChanges {
		switch change.Type {
		case CMOContext:
			switch change.Operation {
			case CMOPSet:
				context.Set(change.Key, change.Data)
			case CMOPDelete:
				context.Del(change.Key)
			case CMOPClear:
				context.Clear()
			}

		case CMOSession:
			switch change.Operation {
			case CMOPSet:
				context.Session.Set(change.Key, change.Data)
			case CMOPDelete:
				context.Session.Del(change.Key)
			case CMOPClear:
				context.Session.Clear()
			}

		case CMOUser:
			switch change.Operation {
			case CMOPSet:
				context.User.Set(change.Key, change.Data)
			case CMOPDelete:
				context.User.Del(change.Key)
			case CMOPClear:
				context.User = RequestUser{}
			}

		case CMOEnvironment:
			switch change.Operation {
			case CMOPSet:
				context.EnvironmentData[change.Key] = change.Data
			case CMOPDelete:
				delete(context.EnvironmentData, change.Key)
			case CMOPClear:
				context.EnvironmentData = nil
			}
		}
	}

	if cm.Errors != nil && len(cm.Errors) > 0 {
		context.Errors = append(context.Errors, cm.Errors...)
		context.LastError = &cm.Errors[len(cm.Errors)-1]
	}
}

func (cm *ContextModifier) Error(err ExecError) *ContextModifier {
	cm.Errors = append(cm.Errors, err)
	return cm
}

func (cm *ContextModifier) SetSession(key string, data interface{}) *ContextModifier {
	return cm.AddOperation(CMOSession, CMOPSet, key, data)
}

func (cm *ContextModifier) Set(key string, data interface{}) *ContextModifier {
	return cm.AddOperation(CMOContext, CMOPSet, key, data)
}

func (cm *ContextModifier) SetUser(key string, data interface{}) *ContextModifier {
	return cm.AddOperation(CMOUser, CMOPSet, key, data)
}

func (cm *ContextModifier) SetEnvironment(key string, data interface{}) *ContextModifier {
	return cm.AddOperation(CMOEnvironment, CMOPSet, key, data)
}

func (cm *ContextModifier) DeleteSession(key string) *ContextModifier {
	return cm.AddOperation(CMOSession, CMOPDelete, key, nil)
}

func (cm *ContextModifier) Delete(key string) *ContextModifier {
	return cm.AddOperation(CMOContext, CMOPDelete, key, nil)
}

func (cm *ContextModifier) DeleteUser(key string) *ContextModifier {
	return cm.AddOperation(CMOUser, CMOPDelete, key, nil)
}

func (cm *ContextModifier) DeleteEnvironment(key string) *ContextModifier {
	return cm.AddOperation(CMOEnvironment, CMOPDelete, key, nil)
}

func (cm *ContextModifier) ClearSession() *ContextModifier {
	return cm.AddOperation(CMOSession, CMOPClear, "", nil)
}

func (cm *ContextModifier) ClearUser() *ContextModifier {
	return cm.AddOperation(CMOUser, CMOPClear, "", nil)
}

func (cm *ContextModifier) ClearEnvironment() *ContextModifier {
	return cm.AddOperation(CMOEnvironment, CMOPClear, "", nil)
}

func (cm *ContextModifier) ClearContext() *ContextModifier {
	return cm.AddOperation(CMOContext, CMOPClear, "", nil)
}

func (cm *ContextModifier) Log(level int, message string) *ContextModifier {
	cm.Logs = append(cm.Logs, LogEntry{
		Level:   level,
		Message: message,
		Time:    time.Now().UTC(),
	})

	return cm
}

func (cm *ContextModifier) LogTrace(message string) *ContextModifier {
	return cm.Log(0, message)
}

func (cm *ContextModifier) LogDebug(message string) *ContextModifier {
	return cm.Log(5, message)
}

func (cm *ContextModifier) LogInfo(message string) *ContextModifier {
	return cm.Log(10, message)
}

func (cm *ContextModifier) LogWarning(message string) *ContextModifier {
	return cm.Log(15, message)
}

func (cm *ContextModifier) LogError(message string) *ContextModifier {
	return cm.Log(20, message)
}
