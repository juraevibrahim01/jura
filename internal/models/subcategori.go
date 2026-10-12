package models

type SubCategoriResdb struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SubCategoriesRes struct {
	Status        string             `json:"status"`
	Description   string             `json:"description"`
	SubCategories []SubCategoriResdb `json:"subcategories"`
}
