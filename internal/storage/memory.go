package storage

import (
	"sync"

	"receipt-processor/internal/models"
)

type MemoryStore struct {
	mu       sync.RWMutex
	receipts map[string]models.Receipt
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{receipts: make(map[string]models.Receipt)}
}

func (ms *MemoryStore) SaveReceipt(receipt models.Receipt) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.receipts[receipt.ID] = receipt
}

func (ms *MemoryStore) GetReceipt(id string) (models.Receipt, bool) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	receipt, exists := ms.receipts[id]
	return receipt, exists
}
