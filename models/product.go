package models

import (
	"time"

	"github.com/google/uuid"
)

// Product represents the Product Table
type Product struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey" json:"product_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Quantity     int       `json:"quantity"`
	Image        string    `json:"image"`
	SellerID     uuid.UUID `gorm:"type:char(36)" json:"seller_id"`
	CreationDate time.Time `gorm:"autoCreateTime" json:"creation_date"`
}

// Order represents the Order Table
type Order struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"order_id"`
	UserID      uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	OrderDate   time.Time `gorm:"autoCreateTime" json:"order_date"`
	TotalAmount float64   `json:"total_amount"`
}

// OrderItem represents the OrderItem Table (Many-to-Many Relationship)
type OrderItem struct {
	OrderID   uuid.UUID `gorm:"type:char(36);primaryKey" json:"order_id"`
	ProductID uuid.UUID `gorm:"type:char(36);primaryKey" json:"product_id"`
	Quantity  int       `json:"quantity"`
}
