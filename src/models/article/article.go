package article

type Article struct {
	Id int `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type ArticlesList []Article