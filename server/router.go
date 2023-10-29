package server

import (
	"api/server/common"
	"api/server/handles"
	"api/server/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/", convert(handles.Index))
	r.GET("/captcha/:type", convert(handles.Captcha))

	r.Use(middlewares.CCInterceptor)

	api := r.Group("/api")
	{
		api.GET("/getToken", convert(handles.GetToken))

		api.POST("/login", convert(handles.Login))
		api.POST("/reg", convert(handles.Register))
		userAuth := api.Group("/user", convert(middlewares.UserAuth))
		{
			userAuth.GET("/info", convert(handles.UserInfo))
		}
	}
}

func array(arr ...string) []string {
	return arr
}

func convert(c common.ApiHandleFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c(&common.ApiContext{
			Context: ctx,
		})
	}
}

func convertArray(c []func(*common.ApiContext)) []gin.HandlerFunc {
	var result []gin.HandlerFunc
	for _, val := range c {
		result = append(result, convert(val))
	}
	return result
}
