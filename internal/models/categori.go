package models

type Categoridbres struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CategoriesRes struct {
	Status      string          `json:"status"`
	Description string          `json:"description"`
	Categories  []Categoridbres `json:"categories"`
}
