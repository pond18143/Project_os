package shopping

import "time"

type listProduct struct {
	ProductName  string  `json:"product_name"`
	PricePerUnit float64 `json:"price_per_unit"`
	BrandName    string  `json:"brand_name"`
}

type inputProductType struct {
	Id int64 `json:"id"`
}

type messageResponse struct {
	Status             int    `json:"status"`
	MessageDescription string `json:"message_description"`
}

type resListProduct struct {
	detail []listProduct `json:"detail"`
}

type inputBuyProduct struct {
	Id   int64 `json:"id"`
	Unit int64 `json:"unit"`
}

type parseCode struct {
	Value string `json:"value"`
}

type product struct {
	ProductName  string  `json:"product_name"`
	Unit         int64   `json:"unit"`
	PricePerUnit float64 `json:"price_per_unit"`
	ProductBrand int     `json:"product_brand"`
	BrandName    string  `json:"brand_name"`
}

type orderModel struct {
	LoginUuid   string    `json:"login_uuid"`
	UserName    string    `json:"user_name"`
	ProductName string    `json:"product_name"`
	ProductBrand int     `json:"product_brand"`
	Unit        int64     `json:"unit"`
	TotalPrice  float64   `json:"total_price"`
	DateTime    time.Time `json:"date_time"`
}
