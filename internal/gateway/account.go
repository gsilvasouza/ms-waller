package gateway

import "github.com/gsilvasouza/ms-waller/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
