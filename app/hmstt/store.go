package hmstt

import (
	"context"

	"gorm.io/gorm"
)

type hmsttStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *hmsttStore {
	// Auto migrate the hmsttState model
	db.AutoMigrate(&hmsttState{})

	return &hmsttStore{db: db}
}

func (s *hmsttStore) SetState(ctx context.Context, key string, value string) error {
	state := &hmsttState{
		Key:   generateKey(key),
		Value: value,
	}
	return s.db.WithContext(ctx).Save(state).Error
}

func (s *hmsttStore) GetState(ctx context.Context, key string) (string, error) {
	var state hmsttState
	err := s.db.WithContext(ctx).First(&state, "key = ?", generateKey(key)).Error
	if err != nil {
		return "", err
	}
	return state.Value, nil
}

func generateKey(parts ...string) string {
	key := PREFIX_HMSTT + "_"
	for _, part := range parts {
		key += part + "_"
	}
	return key
}
