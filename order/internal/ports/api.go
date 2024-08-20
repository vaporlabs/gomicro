package ports

import "github.com/vaporlabs/gomicro/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
