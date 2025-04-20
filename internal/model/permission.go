package model

import "time"

type Permission struct {
    ID             uint64    `json:"id" gorm:"primary_key"`
    PermissionName string    `json:"permission_name"`
    PermissionCode string    `json:"permission_code"`
    MenuID         uint64    `json:"menu_id"`
    Status         int8      `json:"status"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

type RolePermission struct {
    ID           uint64    `json:"id" gorm:"primary_key"`
    RoleID       uint64    `json:"role_id"`
    PermissionID uint64    `json:"permission_id"`
    CreatedAt    time.Time `json:"created_at"`
}