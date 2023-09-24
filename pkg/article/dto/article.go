package dto

type GetArticleRequest struct {
	Title string `json:"title"`
}

type GetArticleResponse struct {
	Title string `json:"title"`
}
