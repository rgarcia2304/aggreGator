package main 

import(
	"fmt"
	"errors"
	"context"
	"net/http"
	"io"
	"encoding/xml"
	"html"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	rssResp := RSSFeed{}
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil{
		return &rssResp, errors.New("Issue with the creation of the new request")
	}

	//add header to the request of gator 
	req.Header.Set("User-Agent", "gator")
	resp, err := client.Do(req)
	if err != nil{
		return &rssResp, errors.New("client had issue with making the request")
	}

	//read the response of the request 
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return &rssResp, errors.New("There has been an issue reading the response body")
	}

	//create the object to marshal into 

	if err := xml.Unmarshal(body, &rssResp); err != nil{
		return &rssResp, errors.New("There has been an issue unmarshalling the data") 
	}

	//decode the escaped HTML entries
	rssResp.Channel.Title = html.UnescapeString(rssResp.Channel.Title)
	rssResp.Channel.Description = html.UnescapeString(rssResp.Channel.Description)
	
	return &rssResp, nil

}

func handlerAgg(s *state, cmd command) error{
	
	if len(cmd.args) != 3{
		return errors.New("Incorrect number of parameters passed in, must pass in a time")
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[2])
	if err != nil{
		return errors.New("A valid value must be passed in for duration to scrape")
	}
	consoleMsg := fmt.Sprintf("Collecting feeds every %v", time_between_reqs)
	fmt.Println(consoleMsg)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		fmt.Println("New Request Being Made")
		scrapeFeeds(s)
		}	
}
