package concurrency

type WebChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultCh := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultCh <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultCh
		results[r.string] = r.bool
	}

	return results
}
