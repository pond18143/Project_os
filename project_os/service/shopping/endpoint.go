package shopping

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"net/http"
	"project_os/service/iojwt"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}


func (ep *Endpoint) SearchByType(c *gin.Context) {
	defer c.Request.Body.Close()

	token := c.Request.Header["Authorization"]
	if token == nil || token[0] == "" {
		msg := messageResponse{
			Status:             http.StatusUnauthorized,
			MessageDescription: "unauthorized"}
		//return err
		c.JSON(msg.Status, msg)
		return
	}

	decode, msgDecode, err := iojwt.DeCodeHS(token)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			msg := messageResponse{
				Status:             http.StatusUnauthorized,
				MessageDescription: "unauthorized"}
			//return err
			c.JSON(msg.Status, msg)
			return
		}
		//return err
		c.JSON(msgDecode.Status, msgDecode)
		return
	}

	log.Infof("loginId : %+v", decode.LoginUuid)
	log.Infof("userName : %+s", decode.UserName)
	log.Infof("companyId : %d", decode.Name)
	log.Infof("roleId : %d", decode.Email)
	log.Infof("-------------------------------")

	var request inputProductType //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, msg, err := searchType(request)
	if err != nil {
		//return err
		c.JSON(msg.Status, msg)
		return
	}
	log.Info(result.detail)
	c.JSON(msg.Status, result.detail)
	return

}


func (ep *Endpoint) BuyProduct(c *gin.Context) {
	defer c.Request.Body.Close()


	token := c.Request.Header["Authorization"]
	if token == nil || token[0] == "" {
		msg := messageResponse{
			Status:             http.StatusUnauthorized,
			MessageDescription: "unauthorized"}
		//return err
		c.JSON(msg.Status, msg)
		return
	}

	decode, msgDecode, err := iojwt.DeCodeHS(token)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			msg := messageResponse{
				Status:             http.StatusUnauthorized,
				MessageDescription: "unauthorized"}
			//return err
			c.JSON(msg.Status, msg)
			return
		}
		//return err
		c.JSON(msgDecode.Status, msgDecode)
		return
	}

	log.Infof("loginId : %+v", decode.LoginUuid)
	log.Infof("userName : %+s", decode.UserName)
	log.Infof("companyId : %d", decode.Name)
	log.Infof("roleId : %d", decode.Email)
	log.Infof("-------------------------------")

	var request inputBuyProduct //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	_, msg, err := UpdateTransaction(request ,decode)
	if err != nil {
		//return err
		c.JSON(msg.Status, msg)
		return
	}
	c.JSON(msg.Status, msg)
	return

}