package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ahmed-khlifi/go-rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreatFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
		URL string `json:"url"`

	}

	decoder := json.NewDecoder(r.Body)
	paramas := parameters{}
	err := decoder.Decode(&paramas)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
 		UserID: user.ID,
		FeedID: paramas.FeedID,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

 
func (apiCfg *apiConfig) handleGetFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {
 

	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, feeds)
}
 
func (apiCfg *apiConfig) handleDeleteFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, struct{}{})
}
