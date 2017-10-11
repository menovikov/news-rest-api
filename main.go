package main

import (
	models "./models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

const (
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_HOST     = ""
	DB_NAME     = ""
)
const layout = "2010-01-01"

func postList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	v := r.URL.Query()

	start, _ := time.Parse(layout, v.Get("start"))
	end, _ := time.Parse(layout, v.Get("end"))
	link := v.Get("link")
	title := v.Get("title")
	content := v.Get("content")
	params := models.NewParams(title, content, link, start.Format(layout), end.Format(layout))

	posts, err := models.AllPosts(*params)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, posts)
}

func postListAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	v := r.URL.Query()

	start, _ := time.Parse(layout, v.Get("start"))
	end, _ := time.Parse(layout, v.Get("end"))
	link := v.Get("link")
	title := v.Get("title")
	content := v.Get("content")
	params := models.NewParams(title, content, link, start.Format(layout), end.Format(layout))

	posts, err := models.AllPosts(*params)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func main() {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	models.InitDB(dbinfo)
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/")))).
		Methods("GET")
	router.HandleFunc("/", postList).Methods("GET")
	router.HandleFunc("/api/", postListAPI)
	http.ListenAndServe(":8080", router)
}
