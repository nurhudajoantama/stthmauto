package bbolt

import (
	"context"

	"github.com/nurhudajoantama/stthmauto/internal/config"
	bolt "go.etcd.io/bbolt"
)

func InitializeBolt(conf config.KVBolt) *bolt.DB {
	db, err := bolt.Open(conf.Path, 0666, nil)
	if err != nil {
		panic(err)
	}
	return db
}

func CloseBolt(ctx context.Context, db *bolt.DB) {
	db.Close()
}
