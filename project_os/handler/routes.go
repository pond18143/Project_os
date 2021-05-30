package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"project_os/service/iojwt"
	"project_os/service/shopping"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute() http.Handler {


	iojwt := iojwt.NewEndpoint()
	shopping := shopping.NewEndpoint()



	txIoJwt := []route{
		{
			Name:        "login",
			Description: "login",
			Method:      http.MethodPost,
			Pattern:     "/login",
			Endpoint:    iojwt.SigninHS,
		},
		{
			Name:        "register",
			Description: "register",
			Method:      http.MethodPost,
			Pattern:     "/register",
			Endpoint:    iojwt.Register,
		},
	}

	txShopping := []route{
		{
			Name:        "searchType",
			Description: "searchType",
			Method:      http.MethodPost,
			Pattern:     "/searchType",
			Endpoint:    shopping.SearchByType,
		},
		{
			Name:        "buyProduct",
			Description: "buyProduct",
			Method:      http.MethodPost,
			Pattern:     "/buyProduct",
			Endpoint:    shopping.BuyProduct,
		},
	}

	ro := gin.New()



	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}


	store = ro.Group("/oauth")
	for _, e := range txIoJwt {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	store = ro.Group("/shopping")
	for _, e := range txShopping {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
