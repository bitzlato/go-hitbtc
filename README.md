go-hitbtc
==========

go-hitbtc is an implementation of the HitBTC API (public and private) in Golang.

This version implement V2 HitBTC API.

### Differences from `github.com/saniales/go-hitbtc`

* Extended API error information. Error code, message, description added.
* Fixes bug with missing `Content-Type` header for `POST`, `PUT` requests
* Add `tradesReport` field to `Order` model

## Import
	import "github.com/bitzlato/go-hitbtc"
	
## Usage

In order to use the client with go's default http client settings you can do:

~~~ go
package main

import (
	"fmt"
	"github.com/bitzlato/go-hitbtc"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// hitbtc client
	hitbtc := hitbtc.New(API_KEY, API_SECRET)

	// Get balances
	balances, err := hitbtc.GetBalances()
	fmt.Println(err, balances)
}
~~~

In order to use custom settings for the http client do:

~~~ go
package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/bitzlato/go-hitbtc"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	// hitbtc client
	bc := hitbtc.NewWithCustomHttpClient(conf.hitbtc.ApiKey, conf.hitbtc.ApiSecret, httpClient)

	// Get balances
	balances, err := hitbtc.GetBalances()
	fmt.Println(err, balances)

	// Initialize websocket connection
	client, err := hitbtc.NewWSClient()
	if err != nil {
		handleError(err) // do something
	}
	defer client.Close()

	// Subscribe and handle
	tickerFeed, err := client.SubscribeTicker("ETHBTC")
	for {
		ticker := <-tickerFeed
		fmt.Println(ticker)
	}


}
~~~

See ["Examples" folder for more... examples](https://github.com/bitzlato/go-hitbtc/blob/master/examples/hitbtc.go)

# Projects using this library

- Golang Crypto Trading Bot: a framework to create trading bots easily and seamlessly (https://github.com/saniales/golang-crypto-trading-bot)
