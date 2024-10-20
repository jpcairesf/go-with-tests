package concurrency

type WebChecker func(string) bool

func CheckWebsites(wc WebChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(url string) {
			results[url] = wc(url)
		}(url)
	}

	return results
}
