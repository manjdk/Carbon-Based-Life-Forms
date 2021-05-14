package domain

const (
	Solid     MineralState = "solid"
	Liquid    MineralState = "liquid"
	Fractured MineralState = "fractured"

	Melt     MineralAction = "melt"
	Fracture MineralAction = "fracture"
	Condense MineralAction = "condense"
)

type MineralState string
type MineralAction string

type Mineral struct {
	ID        string       `json:"id"`
	ClientID  string       `json:"clientId"`
	Name      string       `json:"name"`
	State     MineralState `json:"state"`
	Fractures int          `json:"fractures"`
}

type QueueMessage struct {
	MineralID string        `json:"mineralId"`
	Action    MineralAction `json:"action"`
	TraceID   string        `json:"traceId"`
}

func NewQueueMessage(mineralID, action, traceID string) *QueueMessage {
	return &QueueMessage{
		MineralID: mineralID,
		Action:    NewMineralAction(action),
		TraceID:   traceID,
	}
}

func NewMineralAction(action string) MineralAction {
	return MineralAction(action)
}

func NewMineralState(state string) MineralState {
	return MineralState(state)
}

func (m *Mineral) SetState(state MineralState) {
	m.State = state
}

func (m *Mineral) SetFractureNumber(number int) {
	m.Fractures = number
}
