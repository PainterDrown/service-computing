package ah

import "time"

type customer struct{}

type destination struct{}

type quote struct{}

type weather struct{}

type summary struct {
	destination destination
	quote       quote
	weather     weather
}

func getCustomerDetails() customer {
	time.Sleep(150 * time.Millisecond)
	return customer{}
}

func getDestination(c customer) [10]destination {
	time.Sleep(250 * time.Millisecond)
	return [10]destination{}
}

func getWeatherForcast(d destination) weather {
	time.Sleep(330 * time.Millisecond)
	return weather{}
}

func getQuote(d destination) quote {
	time.Sleep(170 * time.Millisecond)
	return quote{}
}
