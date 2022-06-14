package controller

import (
	"community/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindAllDiscussPost(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	findDiss := dao.SelectDiscussPosts(page, limit)
	ctx.JSON(http.StatusOK, gin.H{
		"data": findDiss,
	})
}
