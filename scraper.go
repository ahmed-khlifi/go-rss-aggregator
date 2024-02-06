package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ahmed-khlifi/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries,
	 concurency int,
	 timeBetweenRequest time.Duration){
		fmt.Printf("Starting scraping on %v GoRoutines every %s duration\n", concurency, timeBetweenRequest)
		ticker := time.NewTicker(timeBetweenRequest)
		for ; ; <-ticker.C {
			feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurency))
			if err != nil {
				fmt.Printf("Error getting feeds to fetch: %v\n", err)
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
		fmt.Printf("Error fetching feed %s: %v\n", feed.Url, err)
		return
	}

	for _, item := range resFedd.Channel.Item {
		
		description :=  sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		publishTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("Error parsing time %s: %v\n", item.PubDate, err)
			continue
		}

		url := sql.NullString{
			String: item.Link,
			Valid:  true,
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:         uuid.New(),
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
			Title:      item.Title,
			Description: description,
			PublishedAt: publishTime,
			Url:        url,
			FeedID:     feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				continue
			}
			fmt.Printf("Error creating post: %v\n", err)
		}
}
	fmt.Printf("Feed %s collected, %v posts found\n", feed.Url, len(resFedd.Channel.Item))
}