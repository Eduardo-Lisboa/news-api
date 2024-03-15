package domain

type NewsReqDomain struct {
	Subject string
	From    string
}

type NewsDomain struct {
	Status       string
	TotalResults string
	Articles     []Article
}

type Article struct {
	Source      string
	Id          string
	Name        string
	Author      string
	Title       string
	URl         string
	UrlToImage  string
	PublishedAt string
	Content     string
}
