package main

import (
	"flag"
	"fmt"
	"github.com/namsral/microdata"
	"log"
	"strconv"
	"time"
)

type options struct {
	Interval   time.Duration
	EventURL   string
	PriceLimit int
}

var opt = options{}

func init() {
	flag.DurationVar(&opt.Interval, "interval", time.Minute, "The interval at which to check for tickets")
	flag.StringVar(&opt.EventURL, "url", "", "The URL of the event on TicketSwap")
	flag.IntVar(&opt.PriceLimit, "price", 50, "The maximum price you're willing to pay")

	flag.Parse()

	if opt.EventURL == "" {
		log.Fatal("The event URL is required")
	}

	fmt.Println("Looking for tickets on", opt.EventURL)
	fmt.Println("* Scanning every", opt.Interval)
	fmt.Println("* Price limit is set to", opt.PriceLimit)
}

func main() {
	ticker := time.NewTicker(opt.Interval)
	defer ticker.Stop()

	for range ticker.C {
		go ping()
	}
}

func ping() {
	data, err := microdata.ParseURL(opt.EventURL)
	if err != nil {
		log.Fatal(err)
	}

	if len(data.Items) == 0 {
		log.Fatal("You are probably blocked for refreshing too often.")
	}

	tickets := data.Items[0].Properties["tickets"]
	for _, t := range tickets {
		var ticket *microdata.Item
		ticket = t.(*microdata.Item)

		var currency, url string
		var quantity, price int

		currency = ticket.Properties["currency"][0].(string)
		url = ticket.Properties["offerurl"][0].(string)
		quantity, err = strconv.Atoi(ticket.Properties["quantity"][0].(string))
		if err != nil {
			continue
		}

		price, err = strconv.Atoi(ticket.Properties["price"][0].(string))
		if err != nil {
			continue
		}

		if currency != "EUR" {
			continue
		}

		if price > opt.PriceLimit {
			log.Println("No more eligible tickets available")
			break
		}

		log.Println("Found", quantity, "ticket(s) for", price, url)
	}
}
