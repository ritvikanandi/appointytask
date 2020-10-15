package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
)
type article struct{
	id string `json:"id"`
	title string `json:"title"`
	subtitle string `json:"subtitle"`
	content string `json:"content"`
	timestamp time.Time `json:"timestamp"`

}

var Articles []article

func allArticles(w http.ResponseWriter, r *http.Request){
	switch r.Method {
    case "GET":
		fmt.Println("Endpoint Hit")
		json.NewEncoder(w).Encode(Articles)
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
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

}

func getArticle(w http.ResponseWriter, r *http.Request){
		key := r.URL.Path[len("/articles/"):]
		for i := 0; i < len(Articles); i++ {

			if key == Articles[i].id {
				fmt.Println("Article Hit")
				json.NewEncoder(w).Encode(Articles[i])
			}
		}



}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Hit")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/articles/", getArticle)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main(){
	handleRequests()
}