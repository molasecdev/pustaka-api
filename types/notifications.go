package types

import "github.com/google/uuid"

type InputNotification struct {
	User_id uuid.UUID `json:"user_id" binding:"required"`
	Message string    `json:"message" binding:"required"`
}

type UpdateNotification struct {
	Read bool `json:"read" binding:"required"`
}
