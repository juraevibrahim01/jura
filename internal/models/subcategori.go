package models

type SubCategoriResdb struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SubCategoriesRes struct {
	Status        string             `json:"status"`
	Description   string             `json:"description"`
	SubCategories []SubCategoriResdb `json:"subcategories"`
}
