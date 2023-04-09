package service

import (
	"github.com/harsha-aqfer/learn-go/methods-return-structs-or-interfaces/parking-lot/internal/db"
	"github.com/harsha-aqfer/learn-go/methods-return-structs-or-interfaces/parking-lot/pkg"
)

type Handler struct {
	db db.DB
}

func NewHandler(db db.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) AddFloor(floor pkg.Floor) error {
	return h.db.AddFloor(floor)
}

func (h *Handler) AddSpot(floorID string, spot pkg.Spot) error {
	return h.db.AddSpot(floorID, spot)
}
