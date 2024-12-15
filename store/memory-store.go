package store

import (
	"sync"

	"github.com/dills122/receipt-rack/models"
)

type MemoryStore struct {
	receipts map[string]models.Receipt
	points   map[string]models.Points
	mu       sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		receipts: make(map[string]models.Receipt),
		points:   make(map[string]models.Points),
	}
}

func (s *MemoryStore) SaveReceipt(id string, receipt models.Receipt) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.receipts[id] = receipt
	return nil
}

func (s *MemoryStore) GetReceipt(id string) (models.Receipt, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	receipt, exists := s.receipts[id]
	return receipt, exists
}

func (s *MemoryStore) SavePoints(id string, points int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.points[id] = models.Points{Points: points, Id: id}
	return nil
}

func (s *MemoryStore) GetPoints(id string) (models.Points, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	points, exists := s.points[id]
	return points, exists
}
