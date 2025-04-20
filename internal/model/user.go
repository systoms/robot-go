package model

import "time"

type User struct {
    ID        uint64    `json:"id" gorm:"primary_key"`
    CompanyID uint64    `json:"company_id"`
    Username  string    `json:"username"`
    Password  string    `json:"-"`
    Status    int8      `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}