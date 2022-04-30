// https://dev.splunk.com/enterprise/docs/releaseapps/packageapps/packagingtoolkit/
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pmeyerson/cryptochecker/client"
)

func main() {
	fiatCurrency := flag.String(
		"fiat", "USD", "The name of the fiat currency you would like to know your crypto price in",
	)

	nameOfCrypto := flag.String(
		"crypto", "BTC", "Input the name of the crypto you wish to price",
	)

	flag.Parse()

	crypto, err := client.FetchCrypto(*fiatCurrency, *nameOfCrypto)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(crypto)
}
