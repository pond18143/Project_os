package shopping

import (
	mssql "project_os/database/mssql"
)

func getProduct(request inputProductType) (detail []listProduct, err error) {

	if err = mssql.DB.Select("pr.product_name ,pr.price_per_unit ,br.brand_name ").
		Table("shopping.dbo.product AS pr , shopping.dbo.brand AS br ").
		Where("pr.product_brand = br.id AND br.id = ? ", request.Id).
		Find(&detail).Error; err != nil {
		return
	}

	return
}

func getProductById(request inputBuyProduct) (detail product, err error) {

	if err = mssql.DB.
		Table("shopping.dbo.product").
		Where("id = ?", request.Id).
		Find(&detail).Error; err != nil {
		return
	}
	return
}

func upDateUnitPlus(request inputBuyProduct , newUnit int64) (detail parseCode, err error) {
	if err = mssql.DB.Table("shopping.dbo.product").
		Where("id = ?", request.Id).
		Update(map[string]interface{}{
			"unit": newUnit,
		}).
		Error; err != nil {
		return
	}
	detail = parseCode{
		Value: "update product success",
	}
	return

}

func upDateTransaction (request orderModel) (detail parseCode, err error) {
	if err = mssql.DB.Table("shopping.dbo.order_transaction").
		Save(&request).Error; err != nil {
		return
	}
	detail = parseCode{
		Value: "update order transaction success",
	}
	return
}

