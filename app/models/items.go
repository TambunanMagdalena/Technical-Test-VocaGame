package models

import "time"

type MasterItem struct {
	ID        int        `json:"id"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Category  string     `json:"category,omitempty"`
	Stock     int        `json:"stock"`
	CreatedBy string     `json:"created_by,omitempty"`
	UpdatedBy string     `json:"updated_by,omitempty"`
	DeletedBy string     `json:"deleted_by,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Barcode   string     `json:"barcode,omitempty"`
}

// Untuk response client (opsional, jika ingin response custom)
type MasterItemResponse struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
	Barcode  string `json:"barcode,omitempty"`
}


// Untuk parsing file import (CSV/Excel)
type ImportItem struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
}

type FilterRequest struct {
	Code     string `json:"code,omitempty"`
	Name     string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
}

func (MasterItem) TableName() string {
	return "master_item" // pastikan sesuai dengan nama tabel di database
}
