package models

import (
	"time"

	"github.com/google/uuid"
)

type StripeSubscription struct {
	Title                 string `json:"title_subscription"`
	SubscriptionId        string `json:"subscription_id"`
	PriceSubscriptionId   string `json:"price_id"`
	ProductSubscriptionId string `json:"product_id"`
}
type Contract struct {
	Record             uuid.UUID          `json:"record"`
	Status             string             `json:"status" validate:"oneof=CURRENTLY CANCELED EXPIRED"`
	BackUrl            string             `json:"back_url"`
	Driver             Driver             `json:"driver"`
	School             School             `json:"school"`
	Child              Child              `json:"child"`
	StripeSubscription StripeSubscription `json:"stripe"`
	CreatedAt          time.Time          `json:"created_at"`
	ExpireAt           time.Time          `json:"expire_at"`
	Amount             int64              `json:"amount"`
	Months             int64              `json:"months"`
}
