package dto

type CreateWalletRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type CreditRequest struct {
	WalletID    string  `json:"wallet_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Reference   string  `json:"reference" validate:"required"`
	Description string  `json:"description"`
}

type DebitRequest struct {
	WalletID    string  `json:"wallet_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Reference   string  `json:"reference" validate:"required"`
	Description string  `json:"description"`
}

type GetBalanceRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
