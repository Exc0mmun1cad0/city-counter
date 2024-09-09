package storage

type Storage interface {
	CountCitiesInCountry(countryCode string) (int, error)
}
