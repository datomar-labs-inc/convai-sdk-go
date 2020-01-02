package byob

import "errors"

type WebhookRequest struct {
	Name    string          `json:"name" mapstructure:"name" msgpack:"name"`
	Context *RequestContext `json:"ctx" mapstructure:"ctx" msgpack:"ctx"`
}

type WebhookHandler func(name string, ctx *RequestContext, cm *ContextModifier) error

type WebhookManager struct {
	handlers map[string]WebhookHandler
	catch    WebhookHandler
}

var ErrNoValidHandlers = errors.New("no valid handlers existed, and no catch handler was defined")

func NewWebhookManager() *WebhookManager {
	return &WebhookManager{
		handlers: make(map[string]WebhookHandler),
		catch:    nil,
	}
}

// Handle registers a handler that will be called when a webhook request is received with a matching name field
func (w *WebhookManager) Handle(name string, handler WebhookHandler) {
	w.handlers[name] = handler
}

// Catch will register a handler that will be called when no other handlers are matched
func (w *WebhookManager) Catch(handler WebhookHandler) {
	w.catch = handler
}

func (w *WebhookManager) Process(req *WebhookRequest) (*ContextModifier, error) {
	h, ok := w.handlers[req.Name]
	if !ok {
		if w.catch == nil {
			return nil, ErrNoValidHandlers
		}

		h = w.catch
	}

	cm := NewContextModifier()

	err := h(req.Name, req.Context, cm)
	if err != nil {
		return nil, err
	}

	return cm, nil
}
