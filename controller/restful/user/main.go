package user

import (
	"aura-test/env"
	"aura-test/pkg/config"
	"aura-test/pkg/log"
	"context"
	"net/http"

	biz "aura-test/service/user"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

// Login
// @Summary user login
// @Description user login
// @Tags User
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Security token
// @Param payload body biz.ReqOfLogin true "user object"
// @Success 200 {object} Response{}
// @Failure 400 {object} Response{}
// @Failure 500 {object} Response{}
// @Router /v1/login [post]
func Login(c *gin.Context) {
	req := new(biz.ReqOfLogin)
	if err := c.Bind(req); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, Response{
			Status:      "failed",
			Description: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, config.ConForge().GetDuration(env.APITimeout))
	defer cancel()

	err := biz.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Status:      "failed",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:      "success",
		Description: "ok",
	})
}
