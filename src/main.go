package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	Articles "./models/article"
	"github.com/gorilla/mux"
)

var articles = Articles.ArticlesList{
	Articles.Article{Title: "Hello", Desc: "Article Description", Content:"Article Content"},
	Articles.Article{Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Title: "World", Desc: "Article Description World", Content:"Article Content World"},
	Articles.Article{Title: "World", Desc: "Article Description World", Content:"Article Content World"},
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
	key := vars["id"]

	fmt.Fprintf(w, "Key: " + key)
}


func main(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",  handleMain)
	myRouter.HandleFunc("/hi",  handleHi)
	myRouter.HandleFunc("/articles",  getArticles)
	myRouter.HandleFunc("/articles/{id}",  getArticles)
	http.Handle("/", r)



	log.Fatal(http.ListenAndServe(":8081", nil))
}