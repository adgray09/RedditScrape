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

	var link string
	link = "https://old.reddit.com/r/wow/"

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
		// grab title
		title := e.ChildText("a.title.may-blank")

		// grab upvotes
		upvotes := findUpvotes(e)

		// grab link
		link = e.ChildAttr("a.title.may-blank", "href")
		link = "https://old.reddit.com" + link

		// put data into struct
		posts := findPosts(link, title, upvotes, e)

		// jsonify data
		dataToJSON(posts, "output.json")
	})

	// c.OnHTML(".nav-buttons", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.ChildAttr("a", "href"))
	// })

	// visit our base URL
	c.Visit(link)

}

func findPosts(link, title, upvotes string, e *colly.HTMLElement) []Post {

	// adds all posts to Post Struct
	allPosts = append(allPosts, Post{title, upvotes, link})

	// print slice check
	// fmt.Println(posts)

	return allPosts
}

func findUpvotes(e *colly.HTMLElement) string {
	upvotes := e.ChildText("div.score.likes")

	// handling case of not having upvotes
	if upvotes == "•" {
		upvotes = "0"
	}

	return upvotes
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
