package gateway

import "github.com/gsilvasouza/ms-waller/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
