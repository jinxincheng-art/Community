package routes

import (
	"community/config"
	diss "community/controller"
	"github.com/gin-gonic/gin"
)


func  InitRouter(r *gin.Engine) error{

	r.GET("/find",diss.FindAllDiscussPost)

	err := r.Run(config.Conf.Port)
	if err != nil {
		return err
	}
	return nil
}
