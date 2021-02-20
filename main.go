package main

import (
	// "encoding/json"
	"fmt"
	"log"

	// "os"

	"github.com/gocolly/colly"
)

type post struct {
	Title   string
	Upvotes string
	Link    string
}

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
		// handle selectors
		title := e.ChildText("a.title.may-blank")
		upvotes := e.ChildText("div.score.likes")
		link = e.ChildAttr("a.title.may-blank", "href")
		// call method
		findPosts(link, title, upvotes, e)
		// dataToJSON(posts, "output.json")
	})

	// visit out base URL
	c.Visit(link)

}

func findPosts(link, title, upvotes string, e *colly.HTMLElement) []post {
	// All Posts
	var allPosts []post

	// adds all posts to Post Struct
	posts := append(allPosts, post{title, upvotes, link})

	// print slice check
	// fmt.Println(posts)

	return posts
}

// func dataToJSON(posts []post, fileName string) {
// 	postJSON, err := json.MarshalIndent(posts, "", " ")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	f, err := os.Create(fileName)
// 	f.Write(postJSON)
// 	f.Close()
// }
