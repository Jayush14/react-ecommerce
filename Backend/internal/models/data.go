package models

type Data struct {
	Products   []Product  `json:"products"`
	Categories []Category `json:"categories"`
	Brands     []Brand    `json:"brands"`
}
