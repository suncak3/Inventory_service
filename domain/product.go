package domain

type Product struct {
	ID          uint     `json:"id" db:"id"`
	Name        string   `json:"name" db:"name" binding:"required"`
	Description string   `json:"description" db:"description" binding:"required"`
	Price       float64  `json:"price" db:"price" binding:"required"`
	StockLevel  int      `json:"stock_level" db:"stock_level" binding:"required"`
	CategoryID  uint     `json:"category_id" db:"category_id" binding:"required"`
	Category    Category `json:"category" db:"category"`
}

type Category struct {
	ID   uint   `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
