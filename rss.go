package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	/*
	res.Body.Close() is being deferred. It is a common practice to close the response body after making an HTTP request.
	The Close() method is used to release any resources associated with the response body, such as network connections or file handles.

	By deferring the Close() method call, we ensure that it will be executed at the end of the surrounding function,
	regardless of any early returns or panics that may occur before that point. This helps prevent resource leaks and ensures proper cleanup
	*/
	defer res.Body.Close()

	/* 
	reads data from an HTTP response body into a byte slice `dat`
	
	*/
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil

}