package service

import (
	"fmt"

	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/domain"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/port/output"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/logger"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/rest_err"
)

type NewsService struct {
	NewsPort output.NewsPort
}

func NewNewsService(NewsPort output.NewsPort) *NewsService {
	return &NewsService{}
}

func (ns *NewsService) GetNewsService(
	newsDomain domain.NewsReqDomain,
) (
	*domain.NewsDomain, *rest_err.RestErr,
) {
	logger.Info(fmt.Sprintf("initiating get news service subject: %s, from: %s", newsDomain.Subject, newsDomain.From))

	newsDomainResponse, err := ns.NewsPort.GetNewsPort(newsDomain)
	return newsDomainResponse, err

}
