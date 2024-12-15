package store

import (
	"errors"

	"github.com/dills122/receipt-rack/models"
)

// RedisStore implements the Store interface using Redis (Placeholder)
type RedisStore struct {
	// Redis client setup here (for now, leave it unimplemented)
}

func NewRedisStore() *RedisStore {
	return &RedisStore{}
}

func (s *RedisStore) SaveReceipt(id string, receipt models.Receipt) error {
	return errors.New("Redis store not implemented")
}

func (s *RedisStore) GetReceipt(id string) (models.Receipt, bool) {
	return models.Receipt{}, false
}

func (s *RedisStore) SavePoints(id string, points int) error {
	return errors.New("Redis store not implemented")
}

func (s *RedisStore) GetPoints(id string) (models.Points, bool) {
	return models.Points{}, false
}
