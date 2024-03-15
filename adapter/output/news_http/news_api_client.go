package news_http

import (
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/adapter/output/model/response"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/domain"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/env"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/rest_err"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type newsClient struct {
}

func NewNewsClient() *newsClient {
	return &newsClient{}
}

func NewsClientInit() *newsClient {
	client = resty.New().SetHostURL("https://newsapi.org/v2/")
	return &newsClient{}
}

var (
	client *resty.Client
)

func (n *newsClient) GetNewsPort(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {

	newsResponse := response.NewsClientResponse{}

	_, err := client.R().
		SetQueryParams(map[string]string{
			"q":      newsDomain.Subject,
			"from":   newsDomain.From,
			"apiKey": env.GetNewsTokenAPI(),
		}).SetResult(&newsResponse).
		Get("/everything")

	if err != nil {
		return nil, rest_err.NewInternalServerError("error when trying to get news")

	}

	newsDomainResponse := &domain.NewsDomain{}
	copier.Copy(newsDomainResponse, newsResponse)

	return newsDomainResponse, nil
}
