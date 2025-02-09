package models

type Category struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Checked bool `json:"checked"`
}