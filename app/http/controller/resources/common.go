package resources

type ID struct {
	ID int `json:"id"`
}

type SkuGoods struct {
	GoodsID   int     `json:"goods_id"`
	GoodsName string  `json:"goods_name"`
	Image     string  `json:"image"`
	SkuID     int     `json:"sku_id"`
	SkuName   string  `json:"skuName"`
	SalePrice float64 `json:"sale_price"`
}

type Recharge struct {
	ID      int    `json:"id"`
	Mid     int    `json:"mid"`
	Mobile  string `json:"mobile"`
	Balance string `json:"balance"`
	CardNo  string `json:"card_no"`
	Status  int    `json:"status"`
}
