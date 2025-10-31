package hmstt

import (
	"context"

	"github.com/rs/zerolog/log"

	bolt "go.etcd.io/bbolt"
)

type hmsttStore struct {
	db *bolt.DB
}

const (
	KV_BUCKET_HMSTT = "hmstt"
	PREFIX_HMSTT    = "hmstt"
)

func NewStore(db *bolt.DB) *hmsttStore {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(KV_BUCKET_HMSTT))
		log.Err(err).Msgf("Initialized bucket %s for HMSTT store", KV_BUCKET_HMSTT)
		return err
	})
	log.Info().Msg("HMSTT store initialized")
	return &hmsttStore{db: db}
}

func (s *hmsttStore) SetState(ctx context.Context, key string, value string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(KV_BUCKET_HMSTT))
		return b.Put(generateKey(key), []byte(value))
	})
}

func (s *hmsttStore) GetState(ctx context.Context, key string) (string, error) {
	var result string
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(KV_BUCKET_HMSTT))
		val := b.Get(generateKey(key))
		if val != nil {
			result = string(val)
		}
		return nil
	})

	return result, err
}

func generateKey(parts ...string) []byte {
	key := PREFIX_HMSTT + "_"
	for _, part := range parts {
		key += part + "_"
	}
	return []byte(key)
}
