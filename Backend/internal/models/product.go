package models

import (
	"database/sql"
	// "github.com/lib/pq"
)
type Dimensions struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Depth  float64 `json:"depth"`
}

type Review struct {
	Rating        int    `json:"rating"`
	Comment       string `json:"comment"`
	Date          string `json:"date"`
	ReviewerName  string `json:"reviewerName"`
	ReviewerEmail string `json:"reviewerEmail"`
}

type Meta struct {
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Barcode   string `json:"barcode"`
	QRCode    string `json:"qrCode"`
}

type Product struct {
	ID                   int               `json:"id"`
	Title                sql.NullString    `json:"title"`
	Description          sql.NullString    `json:"description"`
	Category             sql.NullString    `json:"category"`
	Price                sql.NullFloat64   `json:"price"`
	DiscountPercentage   sql.NullFloat64   `json:"discountPercentage"`
	Rating               sql.NullFloat64   `json:"rating"`
	Stock                sql.NullInt32     `json:"stock"`
	Tags                 sql.NullString  `json:"tags"`
	Brand                sql.NullString    `json:"brand"`
	SKU                  sql.NullString    `json:"sku"`
	Weight               sql.NullInt32     `json:"weight"`
	Width                sql.NullFloat64   `json:"width"`
	Height               sql.NullFloat64   `json:"height"`
	Depth                sql.NullFloat64   `json:"depth"`
	WarrantyInformation  sql.NullString    `json:"warrantyInformation"`
	ShippingInformation  sql.NullString    `json:"shippingInformation"`
	AvailabilityStatus   sql.NullString    `json:"availabilityStatus"`
	ReturnPolicy         sql.NullString    `json:"returnPolicy"`
	MinimumOrderQuantity sql.NullInt32     `json:"minimumOrderQuantity"`
	CreatedAt            sql.NullString    `json:"createdAt"`
	UpdatedAt            sql.NullString    `json:"updatedAt"`
	Barcode              sql.NullString    `json:"barcode"`
	QRCode               sql.NullString    `json:"qrCode"`
	Images               sql.NullString  `json:"images"`
	Thumbnail            sql.NullString    `json:"thumbnail"`
}
