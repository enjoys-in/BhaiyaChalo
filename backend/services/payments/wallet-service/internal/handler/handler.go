package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/payments/wallet-service/internal/ports"
)

type Handler struct {
	svc ports.WalletService
}

func NewHandler(svc ports.WalletService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" || req.Currency == "" {
		writeError(w, http.StatusBadRequest, "user_id and currency are required")
		return
	}

	wallet, err := h.svc.CreateWallet(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toWalletResponse(wallet))
}

func (h *Handler) CreditWallet(w http.ResponseWriter, r *http.Request) {
	walletID := extractPathParam(r, "walletId")

	var req dto.CreditRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.WalletID = walletID

	if req.Amount <= 0 || req.Reference == "" {
		writeError(w, http.StatusBadRequest, "positive amount and reference are required")
		return
	}

	txn, err := h.svc.Credit(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toTransactionResponse(txn))
}

func (h *Handler) DebitWallet(w http.ResponseWriter, r *http.Request) {
	walletID := extractPathParam(r, "walletId")

	var req dto.DebitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.WalletID = walletID

	if req.Amount <= 0 || req.Reference == "" {
		writeError(w, http.StatusBadRequest, "positive amount and reference are required")
		return
	}

	txn, err := h.svc.Debit(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, toTransactionResponse(txn))
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	userID := extractPathParam(r, "userId")
	if userID == "" {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	wallet, err := h.svc.GetBalance(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusNotFound, "wallet not found")
		return
	}

	writeJSON(w, http.StatusOK, toBalanceResponse(wallet))
}

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	walletID := extractPathParam(r, "walletId")
	limit, offset := parsePagination(r)

	txns, err := h.svc.GetTransactions(r.Context(), walletID, limit, offset)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var resp []interface{}
	for _, t := range txns {
		resp = append(resp, toTransactionResponse(&t))
	}
	writeJSON(w, http.StatusOK, resp)
}

func parsePagination(r *http.Request) (int, int) {
	limit := 20
	offset := 0
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}
	if v := r.URL.Query().Get("offset"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			offset = n
		}
	}
	return limit, offset
}
