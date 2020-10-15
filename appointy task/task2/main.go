package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
	"net/url"
)
type article struct{
	id string `json:"id"`
	title string `json:"title"`
	subtitle string `json:"subtitle"`
	content string `json:"content"`
	timestamp time.Time `json:"timestamp"`

}

type Articles []article

func allArticles(w http.ResponseWriter, r *http.Request){
	switch r.Method {
    case "GET": 
    	articles := Articles
		fmt.Println("Endpoint Hit")
		json.NewEncoder(w).Encode(articles)
	case "POST":
		reqBody := ioutil.ReadAll(r.Body)
    		var tempArticle article
    		err := json.Unmarshal(reqBody, &tempArticle)
    		if err != nil {
    			fmt.Println(err)
    		}
    		tempArticle.timestamp = time.Now()
    		Articles = append(Articles, tempArticle)
    default:
    	fmt.Fprintf(w, "Invalid request generated!")

}

func getarticle(w http.ResponseWriter, r *http.Request){
		key := r.URL.Path[len("/articles/"):]
		articles := Articles
		for i := 0; i < len(articles); i++ {

			if key == articles[i].id {
				fmt.Println("Article Hit")
				json.NewEncoder(w).Encode(articles[i])
			}
		}



}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Hit")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/articles/", getarticle)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main(){
	handleRequests()
}