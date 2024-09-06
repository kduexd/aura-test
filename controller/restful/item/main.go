package item

import (
	"aura-test/env"
	"aura-test/pkg/config"
	"context"
	"net/http"

	biz "aura-test/service/item"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status string      `json:"status"`
	Item   interface{} `json:"item"`
}

// List
// @Summary get item list
// @Description get item list
// @Tags Item
// @Accept multipart/form-data
// @Produce application/json
// @Security token
// @Param username query string true "username"
// @Param item query string true "type of item"
// @Success 200 {object} Response{item=biz.ResOfList}
// @Failure 400 {object} Response{}
// @Failure 500 {object} Response{}
// @Router /v1/item [get]
func List(c *gin.Context) {
	req := new(biz.ReqOfList)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, Response{
			Status: "failed",
			Item:   make([]string, 0),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, config.ConForge().GetDuration(env.APITimeout))
	defer cancel()

	res, err := biz.List(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Status: "failed",
			Item:   make([]string, 0),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status: "success",
		Item:   res,
	})
}
