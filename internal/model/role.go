package model

import "time"

type Role struct {
    ID        uint64    `json:"id" gorm:"primary_key"`
    CompanyID uint64    `json:"company_id"`
    RoleName  string    `json:"role_name"`
    RoleCode  string    `json:"role_code"`
    Status    int8      `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserRole struct {
    ID        uint64    `json:"id" gorm:"primary_key"`
    UserID    uint64    `json:"user_id"`
    RoleID    uint64    `json:"role_id"`
    CreatedAt time.Time `json:"created_at"`
}