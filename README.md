# kount
Go SDK for Kount's RSI API 

# Install
```
go get -u github.com/phpsquid/kount
```

# Examples
Look through this package's source code and read [Kount's documentation](https://kount.github.io/docs/introduction/) for more info.
```go
package main

import (
	"github.com/phpsquid/kount/data"
	"github.com/phpsquid/kount/request"
	"github.com/phpsquid/kount/settings"
	"log"
)

func main() {
        // initialize settings
	s := settings.New(
		"123456", // your merchant id
		"https://risk.test.kount.net", // ris url
		"your-api-key", // your api key
		"1b^jIFD)e1@<ZKuH\"A+?Ea`p+ATAo6@:Wee+EM+(FD5Z2/N<", // config key
	)
	i := request.NewInquiry(s)
	i.SetSessionID("someSessionID")
	i.SetPayment("TOKEN", "token123")
	i.SetUnique("123")
	i.SetTotal("1000")
	i.SetEmail("email@example.com")
	i.SetWebsite("DEFAULT")
	i.SetMack("Y")
	i.SetIpAddress("127.0.0.1")

	cart := []data.CartItem{
		data.CartItem{
			ProductType: "T-Shirt",
			ItemName:    "Blue T-Shirt",
			Description: "A really cool blue shirt with a Gopher on it",
			Quantity:    "1",
			Price:       "1000",
		},
	}
	i.SetCart(cart)

	res, err := i.GetResponse()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetAuto())
}
```
