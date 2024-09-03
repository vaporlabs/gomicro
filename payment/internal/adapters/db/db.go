package db

import (
	"context"
	"fmt"

	"github.com/vaporlabs/gomicro/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderId    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) Save(ctx context.Context, payment *domain.Payment) error {
	paymentModel := Payment{
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		OrderId:    payment.OrderId,
		TotalPrice: payment.TotalPrice,
	}
	res := a.db.WithContext(ctx).Create(&paymentModel)
	if res.Error == nil {
		payment.ID = int64(paymentModel.ID)
	}
	return res.Error
}
