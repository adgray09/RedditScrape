package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// "os"

	"github.com/gocolly/colly"
)

// post struct
type Post struct {
	Title   string `json:"title"`
	Upvotes string `json:"upvotes"`
	Link    string `json:"link"`
}

// All Posts
var allPosts []Post

func main() {

	link := "https://old.reddit.com/r/wow/"

	// call Colly method
	visitSite(link)

}

func visitSite(link string) {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("div.thing", func(e *colly.HTMLElement) {

		// jsonify data
		dataToJSON(findPosts(e), "output.json")
	})

	c.OnHTML(".nav-buttons", func(e *colly.HTMLElement) {
		fmt.Println("NEXT HIT")
		e.Request.Visit(e.ChildAttr("a", "href"))
	})

	// visit our base URL
	c.Visit(link)

	c.Wait()

}

func findPosts(e *colly.HTMLElement) []Post {

	link := e.ChildAttr("a.title.may-blank", "href")
	link = "https://old.reddit.com" + link
	title := e.ChildText("a.title.may-blank")
	upvotes := e.ChildText("div.score.likes")

	// handling case of not having upvotes
	if upvotes == "•" {
		upvotes = "0"
	}

	// add data to slice
	allPosts = append(allPosts, Post{title, upvotes, link})

	return allPosts
}

func dataToJSON(posts []Post, fileName string) {

	jsonData, err := json.MarshalIndent(posts, "", "    ")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(jsonData))

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	f.Write(jsonData)

	f.Close()

}
