package fetcher

type DataFetcher interface {
	Fetch(source string) (string, error)
}
