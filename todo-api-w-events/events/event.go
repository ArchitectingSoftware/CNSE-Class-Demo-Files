package events

type EventIDType int

const (
	ToDoQueryEvent EventIDType = iota
	ToDoAddEvent
	ToDoUpdateEvent
	ToDoDeleteEvent
	ToDoErrorEvent
)

type ToDoEvent struct {
	EventID   EventIDType
	EventData map[string]any
}

func NewEvent(id EventIDType, key string, value any) *ToDoEvent {
	return &ToDoEvent{
		EventID: id,
		EventData: map[string]any{
			key: value,
		},
	}
}
