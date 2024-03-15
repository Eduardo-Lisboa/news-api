package routes

import (
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/adapter/input/controller"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/adapter/output/news_http"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newClient := news_http.NewNewsClient()
	newsService := service.NewNewsService(newClient)
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}
