package usecases

import (
	"bytes"
	"context"
	"encoding/base64"
	"image/png"
	"template-go/app/models"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type itemUsecase usecase

type ItemInterface interface {
	CreateMasterItem(ctx context.Context, reqBody models.ImportItem) (*models.MasterItemResponse, error)
	GetMasterItemByFilter(ctx context.Context, filter models.FilterRequest) ([]models.MasterItemResponse, error)
	UpdateMasterItem(ctx context.Context, code, name string, reqBody models.ImportItem) error
	DeleteMasterItem(ctx context.Context, code, name string) error
	ExportData(ctx context.Context, filter models.FilterRequest) ([]models.MasterItem, error)
}

// Implementasi usecase
func (u *itemUsecase) CreateMasterItem(ctx context.Context, reqBody models.ImportItem) (*models.MasterItemResponse, error) {
	now := time.Now()
	masterItem := models.MasterItem{
		Code:      reqBody.Code,
		Name:      reqBody.Name,
		Category:  reqBody.Category,
		Stock:     reqBody.Stock,
		CreatedAt: now,
		CreatedBy: "user", // bisa diganti dari context jika ada user login
	}

	// Generate QR code dari code+name (atau sesuai kebutuhan)
	qrStr, err := generateQRCodeBase64(reqBody.Code + "-" + reqBody.Name)
	if err != nil {
		return nil, err
	}
	masterItem.Barcode = qrStr

	item, err := u.Options.Repository.Item.CreateMasterItem(ctx, masterItem)
	if err != nil {
		return nil, err
	}

	resp := &models.MasterItemResponse{
		ID:       item.ID,
		Code:     item.Code,
		Name:     item.Name,
		Category: item.Category,
		Stock:    item.Stock,
		Barcode:  item.Barcode,
	}

	return resp, nil
}

func (u *itemUsecase) GetMasterItemByFilter(ctx context.Context, filter models.FilterRequest) ([]models.MasterItemResponse, error) {
	items, err := u.Options.Repository.Item.GetMasterItemByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	var resp []models.MasterItemResponse
	for _, item := range items {
		resp = append(resp, models.MasterItemResponse{
			ID:       item.ID,
			Code:     item.Code,
			Name:     item.Name,
			Category: item.Category,
			Stock:    item.Stock,
			Barcode:  item.Barcode,
		})
	}
	return resp, nil
}

func (u *itemUsecase) UpdateMasterItem(ctx context.Context, code, name string, reqBody models.ImportItem) error {
	update := models.MasterItem{
		Code:     reqBody.Code,
		Name:     reqBody.Name,
		Category: reqBody.Category,
		Stock:    reqBody.Stock,
		UpdatedAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
		UpdatedBy: "user", // replace with actual user if available
	}
	return u.Options.Repository.Item.UpdateMasterItem(ctx, code, name, update)
}

func (u *itemUsecase) DeleteMasterItem(ctx context.Context, code, name string) error {
	return u.Options.Repository.Item.DeleteMasterItem(ctx, code, name)
}

func (u *itemUsecase) ExportData(ctx context.Context, filter models.FilterRequest) ([]models.MasterItem, error) {
	return u.Options.Repository.Item.ExportData(ctx, filter)
}

func generateQRCodeBase64(data string) (string, error) {
	qrCode, err := qr.Encode(data, qr.M, qr.Auto)
	if err != nil {
		return "", err
	}
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, qrCode); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
