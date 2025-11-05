package hmstt

import "context"

type hmsttService struct {
	store *hmsttStore
	event *hmsttEvent
}

func NewService(hmsttStore *hmsttStore, hmsttEvent *hmsttEvent) *hmsttService {
	return &hmsttService{
		store: hmsttStore,
		event: hmsttEvent,
	}
}

type GetStateResponse struct {
	States string `json:"state"`
}

func (s *hmsttService) GetState(ctx context.Context, state ...string) (map[string]string, error) {
	result := make(map[string]string)

	for _, st := range state {
		val, err := s.store.GetState(ctx, st)
		if err != nil {
			return nil, err
		}
		result[st] = val
	}

	return result, nil
}

func (s *hmsttService) SetState(ctx context.Context, key string, value string) error {
	err := s.store.SetState(ctx, key, value)
	if err != nil {
		return err
	}

	err = s.event.StateChange(ctx, key, value)
	if err != nil {
		return err
	}

	return nil
}
