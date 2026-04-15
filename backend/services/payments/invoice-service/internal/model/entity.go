package model

import "time"

type InvoiceStatus string

const (
	InvoiceStatusGenerated InvoiceStatus = "generated"
	InvoiceStatusSent      InvoiceStatus = "sent"
	InvoiceStatusPaid      InvoiceStatus = "paid"
)

type Invoice struct {
	ID            string        `json:"id" db:"id"`
	BookingID     string        `json:"booking_id" db:"booking_id"`
	UserID        string        `json:"user_id" db:"user_id"`
	InvoiceNumber string        `json:"invoice_number" db:"invoice_number"`
	Items         []InvoiceItem `json:"items" db:"-"`
	SubTotal      float64       `json:"sub_total" db:"sub_total"`
	TaxAmount     float64       `json:"tax_amount" db:"tax_amount"`
	TotalAmount   float64       `json:"total_amount" db:"total_amount"`
	Currency      string        `json:"currency" db:"currency"`
	Status        InvoiceStatus `json:"status" db:"status"`
	GeneratedAt   time.Time     `json:"generated_at" db:"generated_at"`
	SentAt        *time.Time    `json:"sent_at,omitempty" db:"sent_at"`
}

type InvoiceItem struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Amount      float64 `json:"amount"`
}
