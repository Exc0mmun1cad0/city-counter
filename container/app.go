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
	ctx := context.Background()

	countryCode := r.URL.Path

	// cityCnt, err := a.cache.Get(ctx, countryCode)
	// if err != nil {
	// 	log.Println("cannot get value from redis: %s", err)
	// 	return
	// }
	// if cityCnt == "0" {
	// 	cityCnt, err := a.storage.CountCitiesInCountry(countryCode)
	// 	if err != nil {
	// 		log.Println("cannot get value from postgres: %w", err)
	// 	}
	// 	fmt.Fprintf(w, "В %s %d городов (из постгреса)", countryCode, cityCnt)
	// 	a.cache.Insert(ctx, countryCode, strconv.Itoa(cityCnt))
	// } else {
	// 	fmt.Fprintf(w, "В %s %s городов (из кэша)", countryCode, cityCnt)
	// }

	cityCnt, err := a.storage.CountCitiesInCountry(countryCode)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "%d", cityCnt)

	_ = ctx
}
