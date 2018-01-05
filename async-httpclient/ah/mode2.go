package ah

// Start2 .
func Start2() [10]summary {
	c := getCustomerDetails()
	ds := getDestination(c)

	quoteDone := [10]chan struct{}{}
	weatherDone := [10]chan struct{}{}

	for i := range quoteDone {
		quoteDone[i] = make(chan struct{})
	}
	for i := range weatherDone {
		weatherDone[i] = make(chan struct{})
	}

	var summaries [10]summary

	for i, d := range ds {
		idx := i
		dest := d
		go func() {
			summaries[idx].quote = getQuote(dest)
			quoteDone[idx] <- struct{}{}
		}()
		go func() {
			summaries[idx].weather = getWeatherForcast(dest)
			weatherDone[idx] <- struct{}{}
		}()
	}

	for _, done := range quoteDone {
		<-done
	}
	for _, done := range weatherDone {
		<-done
	}
	return summaries
}
