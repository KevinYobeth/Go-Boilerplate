package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Repo interface {
	Insert(ctx context.Context, order Order) error
	FindById(ctx context.Context, orderId uint64) (Order, error)
	DeleteById(ctx context.Context, orderId uint64) error
	Update(ctx context.Context, order Order) error
	FindAll(ctx context.Context, page FindAllPage) (FindResult, error)
}

type OrderInterface struct {
	Repo Repo
}

func (h *OrderInterface) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		CustomerID uuid.UUID  `json:"customer_id"`
		LineItems  []LineItem `json:"line_items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	order := Order{
		OrderId:    rand.Uint64(),
		CustomerId: body.CustomerID,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}

	err := h.Repo.Insert(r.Context(), order)
	if err != nil {
		fmt.Println("failed to insert: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println("failed to marshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}

func (h *OrderInterface) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const size = 50
	res, err := h.Repo.FindAll(r.Context(), FindAllPage{
		Offset: cursor,
		Size:   size,
	})
	if err != nil {
		fmt.Println("failed to find all:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		Items []Order `json:"items"`
		Next  uint64  `json:"next,omitempty"`
	}
	response.Items = res.Orders
	response.Next = res.Cursor

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("failed to marshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h *OrderInterface) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	o, err := h.Repo.FindById(r.Context(), orderID)
	if errors.Is(err, ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(o); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *OrderInterface) UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderId, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	theOrder, err := h.Repo.FindById(r.Context(), orderId)
	if errors.Is(err, ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	const completedStatus = "completed"
	const shippedStatus = "shipped"
	now := time.Now().UTC()

	switch body.Status {
	case shippedStatus:
		if theOrder.ShippedAt != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		theOrder.ShippedAt = &now
	case completedStatus:
		if theOrder.CompletedAt != nil || theOrder.ShippedAt == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		theOrder.CompletedAt = &now
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Repo.Update(r.Context(), theOrder)
	if err != nil {
		fmt.Println("failed to update:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(theOrder); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *OrderInterface) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")

	const base = 10
	const bitSize = 64

	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Repo.DeleteById(r.Context(), orderID)
	if errors.Is(err, ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to delete:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
