package controller

import (
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/adapter/input/model/request"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/domain"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/port/input"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/logger"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/validation"
	"github.com/gin-gonic/gin"
)

type NewsController struct {
	NewUseCase input.NewsUseCase
}

func NewNewsController(NewUseCase input.NewsUseCase) *NewsController {
	return &NewsController{}
}

func (nc *NewsController) GetNews(c *gin.Context) {
	logger.Info("initiating get news request...")
	newRequest := new(request.NewRequest)

	if err := c.ShouldBindQuery(&newRequest); err != nil {
		logger.Error("Error binding request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	newsDomain := domain.NewsReqDomain{
		Subject: newRequest.Subject,
		From:    newRequest.From.Format("2006-01-02"),
	}

	newsResponseDomain, err := nc.NewUseCase.GetNewsService(newsDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(200, newsResponseDomain)
}
