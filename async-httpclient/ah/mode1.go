package ah

// Start1 .
func Start1() [10]summary {
	c := getCustomerDetails()
	ds := getDestination(c)
	var summaries [10]summary
	for i, d := range ds {
		q := getQuote(d)
		w := getWeatherForcast(d)
		summaries[i] = summary{destination: d, quote: q, weather: w}
	}
	return summaries
}
