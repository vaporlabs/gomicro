package ports

import (
	"context"

	"github.com/vaporlabs/gomicro/payment/internal/application/core/domain"
)

type DBPort interface {
	Save(context.Context, *domain.Payment) error
}
