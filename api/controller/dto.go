package controller

import (
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
)

type mineralStateUpdateRequest struct {
	MineralID string `json:"mineralId"`
	Action    string `json:"action"`
}

type mineral struct {
	ID        string `json:"id"`
	ClientID  string `json:"clientId"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Fractures int    `json:"fractures"`
}

func (m *mineral) toDomain() *domain.Mineral {
	return &domain.Mineral{
		ID:        m.ID,
		ClientID:  m.ClientID,
		Name:      m.Name,
		State:     domain.NewMineralState(m.State),
		Fractures: m.Fractures,
	}
}

func (m *mineralStateUpdateRequest) toQueueMessage(traceID string) *domain.QueueMessage {
	return domain.NewQueueMessage(
		m.MineralID,
		m.Action,
		traceID,
	)
}
