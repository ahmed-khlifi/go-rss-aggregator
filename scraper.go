package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ahmed-khlifi/go-rss-aggregator/internal/database"
)

func startScraping(db *database.Queries,
	 concurency int,
	 timeBetweenRequest time.Duration){
		fmt.Printf("Starting scraping on %v GoRoutines every %s duration", concurency, timeBetweenRequest)
		ticker := time.NewTicker(timeBetweenRequest)
		for ; ; <-ticker.C {
			feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurency))
			if err != nil {
				fmt.Printf("Error getting feeds to fetch: %v", err)
				continue
			}

			wg := &sync.WaitGroup{}
			for _, feed := range feeds {
				wg.Add(1)
				go scrapeFeed(db, wg, feed)
			}
			wg.Wait() // Wait for all goroutine to finish before moving to the next loop iteration
		}
	}


func scrapeFeed(db *database.Queries, 
	wg *sync.WaitGroup, 
	feed database.Feed,
	) {
	defer wg.Done()

	_,err:= db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Println("Error making feed as fetched")
	}

	resFedd, err := urlToFeed(feed.Url)
	if err != nil {
		fmt.Printf("Error fetching feed %s: %v", feed.Url, err)
		return
	}

	for _, item := range resFedd.Channel.Item {
		fmt.Println("Item: ", item.Title)
	}
	fmt.Printf("Feed %s collected, %v posts found", feed.Url, len(resFedd.Channel.Item))
}