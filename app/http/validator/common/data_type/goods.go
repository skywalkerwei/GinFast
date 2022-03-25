package data_type

type Goods struct {
	ShopId     float64 `form:"shop_id" json:"shop_id"  binding:"required,min=1"`
	CategoryId float64 `form:"category_id" json:"category_id"  binding:"required,min=1"`
	Name       string  `form:"name" json:"name" binding:"required,min=1"`
	Img        string  `form:"img" json:"img" binding:"required,min=1"`
	Skus       []Sku   `form:"skus" json:"skus" binding:"required"`
}

type Sku struct {
	Name  string  `form:"name" json:"name" binding:"required,min=1"`
	Price float64 `form:"price" json:"price" binding:"required,min=0.01"`
}
