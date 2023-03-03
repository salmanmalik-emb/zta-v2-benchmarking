package main

import (
	"capture_results/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	resultHandler := http.NewResultHandler()

	hopResultHandler := http.NewHopResultHandler()
	g := router.Group("")

	g.POST("/results", resultHandler.SaveResult)
	g.GET("/results", resultHandler.GetResults)
	g.GET("/results/clear", resultHandler.ClearResults)

	g.POST("/hop-results", hopResultHandler.SaveResult)
	g.GET("/hop-results", hopResultHandler.GetResults)
	g.GET("/hop-results/clear", hopResultHandler.ClearResults)

	if err := router.Run("0.0.0.0:8888"); err != nil {
		panic(err.Error())
	}
}
