package iojwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (ep *Endpoint) SigninHS(c *gin.Context) {
	defer c.Request.Body.Close()

	var request credentials //model รับ input จาก body

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, msg, err := EnCodeHS(request)
	if err != nil {
		//return err
		c.JSON(msg.Status, msg)
		return
	}
	if result.Value == "" {
		c.JSON(msg.Status, msg)
		return
	}
	log.Infof("HS256 : %+s", result.Value)

	c.JSON(http.StatusOK, result)
	return

}

func (ep *Endpoint) Register(c *gin.Context) {
	defer c.Request.Body.Close()

	var request inputRegister //model รับ input จาก body

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, msg, err := CreateUser(request)
	if err != nil {
		//return err
		c.JSON(msg.Status, msg)
		return
	}

	log.Info(result.Value)

	c.JSON(http.StatusOK, result.Value)
	return

}
