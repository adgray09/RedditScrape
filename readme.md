# RedditScrape

[![Go Report Card](https://goreportcard.com/badge/github.com/adgray09/RedditScrape)](https://goreportcard.com/report/github.com/adgray09/RedditScrape)

Web scraper built in Go using the Colly library to scrape all posts of the r/wow Subreddit
### 📚 Table of Contents
1. [Project Structure](#project-structure)
2. [Future Deliverables](#future-deliverables)
3. [Resources](#resources)

## Project Structure
```bash
📂 makescraper
├── README.md
└── scrape.go
└── .gitignore
```
### Flags
* -depth : Sets how far into the subreddit you scrape. Defaults to 5
```
$./web_scraper -depth=5
``` 

## Future deliverables
* Pulling text from individual posts
* Find tendencies from pulled data
* Adding flags for using different Reddit flairs

## Resources
- [**Colly** - Docs](http://go-colly.org/docs/)
- [**Go** - Docs](https://golang.org/)
