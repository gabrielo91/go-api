package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	Articles "./models/article"
	"github.com/gorilla/mux"
)

var articles = Articles.ArticlesList{
	Articles.Article{Id: 1, Title: "Hello", Desc: "Article Description", Content:"Article Content"},
	Articles.Article{Id: 2, Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Id: 3, Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Id: 4, Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Id: 5, Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Id: 6, Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Id: 7, Title: "World", Desc: "Article Description World", Content:"Article Content World"},
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", html.EscapeString(r.URL.Path))
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi")
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting Articles")
	json.NewEncoder(w).Encode(articles)
}

func getArticleById (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for _, article := range articles  {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}


func main(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",  handleMain)
	myRouter.HandleFunc("/hi",  handleHi)
	myRouter.HandleFunc("/articles",  getArticles)
	myRouter.HandleFunc("/articles/{id}",  getArticleById)
	http.Handle("/", myRouter)



	log.Fatal(http.ListenAndServe(":8081", nil))
}