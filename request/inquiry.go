package request

import (
	"fmt"
	"github.com/phpsquid/kount/data"
	"github.com/phpsquid/kount/settings"
)

type Inquiry struct {
	*Request
}

func NewInquiry(settings *settings.Settings) *Inquiry {
	req := newRequest(settings)
	i := &Inquiry{req}
	// defaults
	i.SetMode("Q")
	i.SetCurrency("USD")
	i.SetParm("SDK", "CUST")
	return i
}

/**
 * Set the inquiry mode
 * Acceptable values are: "Q", "P", "W", "J"
 */
func (i *Inquiry) SetMode(mode string) {
	i.Request.SetParm("MODE", mode)
}

// Set the date of birth in the format YYYY-MM-DD.
func (i *Inquiry) SetDateOfBirth(dob string) {
	i.Request.SetParm("DOB", dob)
}

// Set the customer's gender. Either M(male) or F(female).
func (i *Inquiry) SetGender(gender string) {
	i.Request.SetParm("GENDER", gender)
}

// Set the value of a named user defined field.
func (i *Inquiry) SetUserDefinedField(label, value string) {
	key := "UDF[" + label + "]"
	i.Request.SetParm(key, value)
}

// Set the three character ISO-4217 currency code.
func (i *Inquiry) SetCurrency(currency string) {
	i.Request.SetParm("CURR", currency)
}

// Set the three character ISO-4217 currency code.
func (i *Inquiry) SetTotal(total string) {
	i.Request.SetParm("TOTL", total)
}

// Set the three character ISO-4217 currency code.
func (i *Inquiry) SetIpAddress(ipAddress string) {
	i.Request.SetParm("IPAD", ipAddress)
}

// Set the email address of the client.
func (i *Inquiry) SetEmail(email string) {
	i.Request.SetParm("EMAL", email)
}

// Set the ANI (Automatic Identification Number) received for the phone transaction.
func (i *Inquiry) SetANID(anid string) {
	i.Request.SetParm("ANID", anid)
}

// Set the name of the client.
func (i *Inquiry) SetName(name string) {
	i.Request.SetParm("NAME", name)
}

// Set the customer unique id or cookie
func (i *Inquiry) SetUnique(unique string) {
	i.Request.SetParm("UNIQ", unique)
}

// Set the unix epoc date when unique was first set.
func (i *Inquiry) SetEpoch(epoch string) {
	i.Request.SetParm("EPOC", epoch)
}

// Set cash amount of any feasible goods in the order
func (i *Inquiry) SetCash(cash string) {
	i.Request.SetParm("CASH", cash)
}

// Set shipment type
func (i *Inquiry) SetShipType(shipType string) {
	i.Request.SetParm("SHTP", shipType)
}

// Set the billing address
func (i *Inquiry) SetBillingAddress(address1, address2, city, state,
	postalCode, country, premise, street string) {

	i.Request.SetParm("B2A1", address1)
	i.Request.SetParm("B2A2", address2)
	i.Request.SetParm("B2CI", city)
	i.Request.SetParm("B2ST", state)
	i.Request.SetParm("B2PC", postalCode)
	i.Request.SetParm("B2CC", country)
	if premise != "" {
		i.Request.SetParm("BPREMISE", premise)
	}
	if street != "" {
		i.Request.SetParm("BSTREET", street)
	}
}

// Set the billing phone number
func (i *Inquiry) SetBillingPhoneNumber(phoneNumber string) {
	i.Request.SetParm("B2PN", phoneNumber)
}

// Set the shipping address
func (i *Inquiry) SetShippingAddress(address1, address2, city, state,
	postalCode, country, premise, street string) {

	i.Request.SetParm("S2A1", address1)
	i.Request.SetParm("S2A2", address2)
	i.Request.SetParm("S2CI", city)
	i.Request.SetParm("S2ST", state)
	i.Request.SetParm("S2PC", postalCode)
	i.Request.SetParm("S2CC", country)
	if premise != "" {
		i.Request.SetParm("SPREMISE", premise)
	}
	if street != "" {
		i.Request.SetParm("SSTREET", street)
	}
}

// Set the shipping phone number
func (i *Inquiry) SetShippingPhoneNumber(phoneNumber string) {
	i.Request.SetParm("S2PN", phoneNumber)
}

// Set the shipping name
func (i *Inquiry) SetShippingName(name string) {
	i.Request.SetParm("S2NM", name)
}

// Set the shipping email
func (i *Inquiry) SetShippingEmail(emailAddress string) {
	i.Request.SetParm("S2EM", emailAddress)
}

// Set the user agent string
func (i *Inquiry) SetUserAgent(userAgent string) {
	i.Request.SetParm("UAGT", userAgent)
}

// Set the website id (shortname) associated with this transaction
func (i *Inquiry) SetWebsite(site string) {
	i.Request.SetParm("SITE", site)
}

func (i *Inquiry) addItemToCart(index int, item data.CartItem) {
	productType := fmt.Sprintf("PROD_TYPE[%d]", index)
	itemName := fmt.Sprintf("PROD_ITEM[%d]", index)
	description := fmt.Sprintf("PROD_DESC[%d]", index)
	quantity := fmt.Sprintf("PROD_QUANT[%d]", index)
	price := fmt.Sprintf("PROD_PRICE[%d]", index)

	i.Request.SetParm(productType, item.ProductType)
	i.Request.SetParm(itemName, item.ItemName)
	i.Request.SetParm(description, item.Description)
	i.Request.SetParm(quantity, item.Quantity)
	i.Request.SetParm(price, item.Price)
}

// Set the shopping cart
func (i *Inquiry) SetCart(cart []data.CartItem) {
	for index, item := range cart {
		i.addItemToCart(index, item)
	}
}
