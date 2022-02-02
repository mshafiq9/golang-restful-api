package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "log"
    "net/http"
)


type Article struct {
    Id string       `json:"id"`
    Title string    `json:"title"`
    Desc string     `json:"desc"`
    Content string  `json:"content"`
}

// declaring a global Articles array to simulate as Articles database
var Articles []Article


// I didn't like /favicon.io generating a requst to homePage so!
func doNothing(w http.ResponseWriter, r *http.Request) {}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnSingleArticle")

    vars := mux.Vars(r)
    id := vars["id"]

    // loop over Articles and return the matching article encoded as JSON
    for _, article := range Articles {
        if article.Id == id {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: createNewArticle")

    // get the POST request body and unmarshal into a new Article struct
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article
    json.Unmarshal(reqBody, &article)

    // append new article to Articles database array
    Articles = append(Articles, article)

    // return the new article
    json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: updateArticle")

    // get the PUT request body and unmarshal into a new Article struct
    reqBody, _ := ioutil.ReadAll(r.Body)
    var newArticle Article
    json.Unmarshal(reqBody, &newArticle)

    // loop over Articles and update the Article at the matching id index
    for index, article := range Articles {
        if article.Id == newArticle.Id {
            Articles[index] = newArticle
            json.NewEncoder(w).Encode(newArticle)
        }
    }
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: deleteArticle")

    // parse the path params and extract the id of article to delete
    vars := mux.Vars(r)
    id := vars["id"]

    // loop through all the articles and delete the one with the matching Id
    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
            break
        }
    }
}

func handleRequests() {
    // create a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/favicon.ico", doNothing)
    myRouter.HandleFunc("/", homePage)

    // get all articles
    myRouter.HandleFunc("/articles", returnAllArticles)

    // CRUD operations for Articles
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")

    http.Handle("/", myRouter)
    log.Fatal(http.ListenAndServe(":10000", nil))
}


func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")

    Articles = []Article {
        Article {Id: "1", Title: "Hello 1", Desc: "Article 1 Desc", Content: "Article 1 Content"},
        Article {Id: "2", Title: "Hello 2", Desc: "Article 2 Desc", Content: "Article 2 Content"},
    }

    handleRequests()
}
