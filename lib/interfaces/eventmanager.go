package interfaces

import "github.com/Felamande/wails/lib/messages"

// EventManager is the event manager interface
type EventManager interface {
	PushEvent(*messages.EventData)
	Emit(eventName string, optionalData ...interface{})
	On(eventName string, callback func(...interface{}))
	Start(Renderer)
	Shutdown()
}
