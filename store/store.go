package store

import "github.com/dills122/receipt-rack/models"

type Store interface {
	// Methods for managing points
	SavePoints(id string, points int) error
	GetPoints(id string) (models.Points, bool)

	// Methods for managing receipts
	SaveReceipt(id string, receipt models.Receipt) error
	GetReceipt(id string) (models.Receipt, bool)
}

func NewStore(useRedis bool) Store {
	return NewMemoryStore()
}
