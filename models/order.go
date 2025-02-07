package models

import "time"

type Order struct {
	ID string `json:"id" bson:"_id,omitempty"`
	TenantID string `json:"tenant_id" bson:"tenant_id"`
	SellerID string `json:"seller_id" bson:"seller_id"`
	HubID    string `json:"hub_id" bson:"hub_id"`
	SKU       string    `json:"sku" bson:"sku"`
	Quantity  int       `json:"quantity" bson:"quantity"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type MessageOutput struct {
	Message string
}
type OrderEvent struct {
	FilePath string `json:"file_path"`
	UserID   string `json:"user_id"`
}
