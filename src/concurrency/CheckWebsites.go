package concurrency

type WebsiteChecker func(string) bool
type result struct {
	// anonymous fields
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// a channel of result type
	resultChannel := make(chan result)

	for _, url := range urls {
		// use anonymous function here
		// 1. anonymous function is executed at the same time that they're declared => () at the end doing that
		// 2. it maintains the lexical scope they are defined in
		// (have access to the variables that are available at the point when declaring the anonymous)

		// // now each loop will start a new goroutine
		// go func() {
		// 	// problem: each goroutine refers to the same url variable, so they're writing the value
		// 	// 	that url has at the end of the iteration
		// 	results[url] = wc(url)
		// }()

		// /********************************************
		// Give each goroutine its own copy of "url"
		// ********************************************/
		// // u is a copy of url
		// go func(u string) {
		// 	// problem: race condition (concurrent map writes)
		// 	results[u] = wc(u)
		// }(url)

		/**************************
		use channels
		1. send result struct for each loop to resultChannel (In goroutine)
		**************************/
		go func(u string) {
			// <- : send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	/****************
	use channels
	2. assign values received from a channel (After goroutine)
	****************/
	for i := 0; i < len(urls); i++ {
		// <- : assign a value received from a channel to a variable
		r := <-resultChannel
		results[r.string] = r.bool
	}

	// time.Sleep(2 * time.Second)

	return results
}
