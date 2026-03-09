package models

import "time"

type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required,min=3,max=50"`
	Grade     int       `json:"grade" binding:"required,gte=0,lte=100"`
	Email     string    `json:"email" binding:"omitempty,email"`
	CreatedAt time.Time `json:"created_at"`
}
