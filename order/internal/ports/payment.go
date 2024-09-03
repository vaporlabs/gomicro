package ports

import (
	"context"

	"github.com/vaporlabs/gomicro/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
