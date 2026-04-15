package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/wallets", h.CreateWallet)
	mux.HandleFunc("GET /api/v1/users/{userId}/wallet/balance", h.GetBalance)
	mux.HandleFunc("POST /api/v1/wallets/{walletId}/credit", h.CreditWallet)
	mux.HandleFunc("POST /api/v1/wallets/{walletId}/debit", h.DebitWallet)
	mux.HandleFunc("GET /api/v1/wallets/{walletId}/transactions", h.GetTransactions)
}
