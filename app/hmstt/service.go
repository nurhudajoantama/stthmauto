package hmstt

import "context"

type hmsttService struct {
	store *hmsttStore
}

func NewService(hmsttStore *hmsttStore) *hmsttService {
	return &hmsttService{
		store: hmsttStore,
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
