package dto

import "time"

type TodoBodyRequest struct {
	TodoText string `json:"todoText"`
}

type TodoQueryRow struct {
	ID           uint      `gorm:"id"`
	CreatedAt    time.Time `gorm:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at"`
	DeletedAt    string    `gorm:"deleted_at"`
	UserIpID     string    `gorm:"user_ip_id"`
	TodoStatus   string    `gorm:"todo_status"`
	TodoText     string    `gorm:"todo_text"`
	ActiveStatus string    `gorm:"active_status"`
}

type TodoRespone struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    string    `json:"deleted_at"`
	UserIpID     string    `json:"user_ip_id"`
	TodoStatus   string    `json:"todo_status"`
	TodoText     string    `json:"todo_text"`
	ActiveStatus string    `json:"active_status"`
}
