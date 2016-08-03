# Ticketswap Alert

> Check TicketSwap for event tickets below a certain price

## Usage

```
Usage of ticketswap-alert:
  -interval duration
    	The interval at which to check for tickets (default 1m0s)
  -price int
    	The maximum price you're willing to pay (default 50)
  -url string
    	The URL of the event on TicketSwap
```

## Caveats

* You'll be regularly marked as a bot by CloudFlare for continuous requests to the same URL

## TODO

* Implement a notification service instead of logging to console
