package response

type NewsClientResponse struct {
	Status       string
	TotalResults string
	Articles     []ArticlResponse
}

type ArticlResponse struct {
	Source      ArticleSourceResponse
	Id          string
	Name        string
	Author      string
	Title       string
	URl         string
	UrlToImage  string
	PublishedAt string
	Content     string
}

type ArticleSourceResponse struct {
	Id   *string
	Name string
}
