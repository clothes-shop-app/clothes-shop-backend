package models

type UploadProduct struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price       int      `json:"price" binding:"required,min=0"`
	CategoryID  string   `json:"categoryId"`
	Images      []string `json:"images"`
}
