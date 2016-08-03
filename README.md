# Ticketswap Alert

Check TicketSwap for event tickets below a certain price

## Installation

Run the following to install the `ticketswap-alert` binary to your `$GOPATH/bin`:
```
go get github.com/rjelierse/ticketswap-alert/...
```

## Usage

Usage of `ticketswap-alert`:

- **interval** *duration* The interval at which to check for tickets (default 1m0s)
- **price**    *int*      The maximum price you're willing to pay (default 50)
- **url**      *string*   The URL of the event on TicketSwap

## Caveats

* You'll be regularly marked as a bot for continuous requests to the same URL
  The interval may vary

## TODO

* [ ] Implement a notification service instead of logging to console
