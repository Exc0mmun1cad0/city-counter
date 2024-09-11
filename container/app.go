package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"test-app/cache"
	"test-app/storage"
)

type App struct {
	storage storage.Storage
	cache   cache.Cache
}

func NewApp(storage storage.Storage, cache cache.Cache) *App {
	return &App{
		storage: storage,
		cache:   cache,
	}
}

func (a *App) GetStats(w http.ResponseWriter, r *http.Request) {
	// Background context for cache queries
	ctx := context.Background()

	// Getting and validating country code.
	// If there are any symbols except english capital letteres,
	// user will get 404.
	countryCode := r.URL.Path[1:]
	if !validateCountryCode(countryCode) {
		w.WriteHeader(404)
		return
	}

	// Getting number of cities in country by country code from cache
	cnt, err := a.cache.Get(ctx, countryCode)
	if err != nil {
		log.Printf("cannot get value from redis: %s", err)
	}

	// If value is equal to 0, there is no data in cache
	// So we have to do query to the database and cache that value for the future
	if cnt == 0 {
		cnt, err := a.storage.CountCitiesInCountry(countryCode)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		if cnt == 0 {
			w.WriteHeader(404)
			return
		}
		a.cache.Insert(ctx, countryCode, cnt)
		fmt.Fprintf(w, "%d (from db)", cnt)
		return
	}

	fmt.Fprintf(w, "%d (from cache)", cnt)
}

func validateCountryCode(countryCode string) bool {
	for _, sym := range countryCode {
		if !(sym >= 65 && sym <= 90) {
			return false
		}
	}

	return true
}
