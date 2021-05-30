package shopping

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"project_os/service/iojwt"
	"time"
)

func searchType(request inputProductType) (result resListProduct, msg messageResponse, err error) {

	var list []listProduct
	list, err = getProduct(request)
	if err != nil {
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "notFound"}
		return
	}

	msg = messageResponse{
		Status:             http.StatusOK,
		MessageDescription: "reset password success",
	}
	result = resListProduct{
		list,
	}

	return
}

func UpdateTransaction(request inputBuyProduct, claims iojwt.DecodeClaims) (result parseCode, msg messageResponse, err error) {
	var detail product
	detail, err = getProductById(request)
	if err != nil {
		log.Error("not found product in system")
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "not found product in system"}

		return
	}
	log.Info(detail)
	reduce := detail.Unit - request.Unit
	price := request.Unit * int64(detail.PricePerUnit)
	priceFloat := float64(price)

	if request.Unit > detail.Unit {
		log.Error("Enough stock")
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "Enough stock"}
		return
	}

	update, err := upDateUnitPlus(request, reduce)
	if err != nil {
		log.Error("not found product in system")
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "not found product in system"}
		return
	}
	log.Info(update)

	var order = orderModel{
		LoginUuid:   claims.LoginUuid,
		UserName:    claims.UserName,
		ProductName: detail.ProductName,
		ProductBrand: detail.ProductBrand,
		Unit:        request.Unit,
		TotalPrice:  priceFloat,
		DateTime:    time.Now(),
	}

	transaction, err := upDateTransaction(order)
	if err != nil {
		log.Error("not found product in system")
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "not found product in system"}
		return
	}
	log.Info(transaction)

	msg = messageResponse{
		Status:             http.StatusOK,
		MessageDescription: "update stock success",
	}
	return
}
