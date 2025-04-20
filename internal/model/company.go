package model

import "time"

type Company struct {
    ID          uint64    `json:"id" gorm:"primary_key"`
    CompanyCode string    `json:"company_code"`
    CompanyName string    `json:"company_name"`
    Status      int8      `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}