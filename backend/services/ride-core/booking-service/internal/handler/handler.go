package handler

import (
	"net/http"

	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/dto"
	"github.com/enjoys-in/BhaiyaChalo/services/ride-core/booking-service/internal/ports"
)

type BookingHandler struct {
	svc ports.BookingService
}

func NewBookingHandler(svc ports.BookingService) *BookingHandler {
	return &BookingHandler{svc: svc}
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateBookingRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID == "" || req.CityID == "" || req.VehicleType == "" {
		errorJSON(w, http.StatusBadRequest, "missing required fields")
		return
	}

	booking, err := h.svc.CreateBooking(r.Context(), &req)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := dto.BookingResponse{
		ID:             booking.ID,
		UserID:         booking.UserID,
		CityID:         booking.CityID,
		PickupLat:      booking.PickupLat,
		PickupLng:      booking.PickupLng,
		PickupAddress:  booking.PickupAddress,
		DropLat:        booking.DropLat,
		DropLng:        booking.DropLng,
		DropAddress:    booking.DropAddress,
		VehicleType:    booking.VehicleType,
		EstimatedFare:  booking.EstimatedFare,
		FinalFare:      booking.FinalFare,
		PromoCode:      booking.PromoCode,
		DiscountAmount: booking.DiscountAmount,
		Status:         string(booking.Status),
		DriverID:       booking.DriverID,
		PaymentMethod:  booking.PaymentMethod,
		CreatedAt:      booking.CreatedAt,
		UpdatedAt:      booking.UpdatedAt,
	}

	successJSON(w, http.StatusCreated, "", resp)
}

func (h *BookingHandler) GetBooking(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		errorJSON(w, http.StatusBadRequest, "booking id is required")
		return
	}

	booking, err := h.svc.GetBooking(r.Context(), id)
	if err != nil {
		errorJSON(w, http.StatusNotFound, "booking not found")
		return
	}

	resp := dto.BookingResponse{
		ID:             booking.ID,
		UserID:         booking.UserID,
		CityID:         booking.CityID,
		PickupLat:      booking.PickupLat,
		PickupLng:      booking.PickupLng,
		PickupAddress:  booking.PickupAddress,
		DropLat:        booking.DropLat,
		DropLng:        booking.DropLng,
		DropAddress:    booking.DropAddress,
		VehicleType:    booking.VehicleType,
		EstimatedFare:  booking.EstimatedFare,
		FinalFare:      booking.FinalFare,
		PromoCode:      booking.PromoCode,
		DiscountAmount: booking.DiscountAmount,
		Status:         string(booking.Status),
		DriverID:       booking.DriverID,
		PaymentMethod:  booking.PaymentMethod,
		CreatedAt:      booking.CreatedAt,
		UpdatedAt:      booking.UpdatedAt,
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *BookingHandler) CancelBooking(w http.ResponseWriter, r *http.Request) {
	var req dto.CancelBookingRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" {
		errorJSON(w, http.StatusBadRequest, "booking_id is required")
		return
	}

	if err := h.svc.CancelBooking(r.Context(), &req); err != nil {
		errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "", map[string]string{"status": "cancelled"})
}

func (h *BookingHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateStatusRequest
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" || req.Status == "" {
		errorJSON(w, http.StatusBadRequest, "booking_id and status are required")
		return
	}

	if err := h.svc.UpdateStatus(r.Context(), &req); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := dto.BookingStatusResponse{
		BookingID: req.BookingID,
		Status:    req.Status,
	}

	successJSON(w, http.StatusOK, "", resp)
}

func (h *BookingHandler) AssignDriver(w http.ResponseWriter, r *http.Request) {
	type assignReq struct {
		BookingID string `json:"booking_id"`
		DriverID  string `json:"driver_id"`
	}

	var req assignReq
	if err := decodeJSON(r, &req); err != nil {
		errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.BookingID == "" || req.DriverID == "" {
		errorJSON(w, http.StatusBadRequest, "booking_id and driver_id are required")
		return
	}

	if err := h.svc.AssignDriver(r.Context(), req.BookingID, req.DriverID); err != nil {
		errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	successJSON(w, http.StatusOK, "driver assigned successfully", map[string]string{
		"booking_id": req.BookingID,
		"driver_id":  req.DriverID,
		"status":     "driver_assigned",
	})
}
