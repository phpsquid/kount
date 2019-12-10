// The request package defines a struct for building and sending a request to Kount
package request

import (
	"github.com/phpsquid/kount/response"
	"github.com/phpsquid/kount/settings"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Version           = "0700"
	PyplType          = "PYPL"
	GoogType          = "GOOG"
	GiftCardType      = "GIFT"
	GDMPType          = "GDMP"
	NoneType          = "NONE"
	CardType          = "CARD"
	CheckType         = "CHEK"
	BLMLType          = "BLML"
	APAYType          = "APAY"
	BPAYType          = "BPAY"
	CarteBleueType    = "CARTE_BLEUE"
	ELVType           = "ELV"
	GiroPayType       = "GIROPAY"
	InteracType       = "INTERAC"
	MercadePagoType   = "MERCADE_PAGO"
	NetellerType      = "NETELLER"
	POLIType          = "POLI"
	SEPAType          = "SEPA"
	SkrillType        = "SKRILL"
	SofortType        = "SOFORT"
	TokenType         = "TOKEN"
	ConnectionTimeout = 30
)

type Request struct {
	*settings.Settings
	data              map[string]string
	connectionTimeout int
	url               string
	apiKey            string
}

func newRequest(settings *settings.Settings) *Request {
	data := make(map[string]string)
	req := &Request{
		Settings: settings,
		data:     data,
	}
	// set merchant id
	req.SetMerchantID(settings.GetMerchantID())
	return req
}

// Set a parameter for the request.
func (r *Request) SetParm(key, value string) {
	r.data[key] = value
}

// Set the merchant id assigned by Kount.
func (r *Request) SetMerchantID(id string) {
	r.data["MERC"] = id
}

// Set the merchant gateway's customer id for Kount Central.
func (r *Request) SetKCCustomerID(id string) {
	r.data["CUSTOMER_ID"] = id
}

// Set the maximum number of seconds for RIS connection function to timeout.
func (r *Request) SetConnectionTimeout(timeout int) {
	r.connectionTimeout = timeout
}

// Set the version number
func (r *Request) SetVersion(version string) {
	r.data["VERS"] = version
}

// Set the session id. Must be unique over a 30-day span.
func (r *Request) SetSessionID(id string) {
	r.data["SESS"] = id
}

// Set the order number
func (r *Request) SetOrderNumber(orderNumber string) {
	r.data["ORDR"] = orderNumber
}

// Set the mack
func (r *Request) SetMack(mack string) {
	r.data["MACK"] = mack
}

// Set the Authorization status.
func (r *Request) SetAUTH(auth string) {
	r.data["AUTH"] = auth
}

/**
 * Set the Bankcard AVS zip code reply.
 * Address Verification System Zip Code verification response returned to merchant from
 * processor. Acceptable values are ‘M’ for match, ’N’ for no match, or ‘X’ for unsupported
 * or unavailable.
 */
func (r *Request) SetAVSZ(avsz string) {
	r.data["AVSZ"] = avsz
}

/**
 * Set the Bankcard AVS street address reply.
 * Address Verification System Street verification response returned to merchant from
 * processor. Acceptable values are ’M’ for match, ’N’ for no-match, or ’X’ for
 * unsupported or unavailable.
 */
func (r *Request) SetAVST(avst string) {
	r.data["AVST"] = avst
}

/**
 * Set the Bankcard CVV/CVC/CVV2 reply.
 * Card Verification Value response returned to merchant from processor. Acceptable
 * values are ’M’ for match, ’N’ for no-match, or ’X’ unsupported or unavailable.
 */
func (r *Request) SetCVVR(cvvr string) {
	r.data["CVVR"] = cvvr
}

// Set the RIS target server URL
func (r *Request) SetURL(url string) {
	r.url = url
}

// Set the API key for authentication.
func (r *Request) SetAPIKey(key string) {
	r.apiKey = key
}

// Set the payment token.
func (r *Request) setPaymentToken(token string) {
	r.data["PTOK"] = token
}

// Set a Green Dot MoneyPak payment.
func (r *Request) SetGreenDotMoneyPakPayment(paymentID string) {
	r.data["PTYP"] = GDMPType
	r.setPaymentToken(paymentID)
}

// Set no payment.
func (r *Request) SetNoPayment() {
	r.data["PTYP"] = NoneType
	delete(r.data, "PTOK")
}

// Set a PayPal payment.
func (r *Request) SetPayPalPayment(payPalID string) {
	r.data["PTYP"] = PyplType
	r.setPaymentToken(payPalID)
}

// Set a Google payment.
func (r *Request) SetGooglePayment(googleID string) {
	r.data["PTYP"] = GoogType
	r.setPaymentToken(googleID)
}

// Set a gift card payment.
func (r *Request) SetGiftCardPayment(giftCardNumber string) {
	r.data["PTYP"] = GiftCardType
	r.setPaymentToken(giftCardNumber)
}

// Set a card payment.
func (r *Request) SetCardPayment(cardNumber string) {
	r.data["PTYP"] = CardType
	r.setPaymentToken(cardNumber)
}

// Set a check payment.
func (r *Request) SetCheckPayment(micr string) {
	r.data["PTYP"] = CheckType
	r.setPaymentToken(micr)
}

// Set a bill-me-later payment.
func (r *Request) SetBillMeLaterPayment(blmlID string) {
	r.data["PTYP"] = BLMLType
	r.setPaymentToken(blmlID)
}

// Set a apple pay payment type.
func (r *Request) SetApplePayment(appleID string) {
	r.data["PTYP"] = APAYType
	r.setPaymentToken(appleID)
}

// Set a BPAY payment type
func (r *Request) SetBPAYPayment(bppID string) {
	r.data["PTYP"] = BPAYType
	r.setPaymentToken(bppID)
}

// Set a Carte Bleue payment type
func (r *Request) SetCarteBleuePayment(cbpID string) {
	r.data["PTYP"] = CarteBleueType
	r.setPaymentToken(cbpID)
}

// Set a ELV payment type
func (r *Request) SetELVPayment(elvpID string) {
	r.data["PTYP"] = ELVType
	r.setPaymentToken(elvpID)
}

// Set a GiroPay payment type
func (r *Request) SetGiroPayPayment(giroPayID string) {
	r.data["PTYP"] = GiroPayType
	r.setPaymentToken(giroPayID)
}

// Set a Interac payment type
func (r *Request) SetInteracPayment(interacID string) {
	r.data["PTYP"] = InteracType
	r.setPaymentToken(interacID)
}

// Set a Mercado Pago payment type
func (r *Request) SetMercadoPagoPayment(mercadoPagoID string) {
	r.data["PTYP"] = MercadePagoType
	r.setPaymentToken(mercadoPagoID)
}

// Set a Netellerpayment type
func (r *Request) SetNetellerPayment(netellerId string) {
	r.data["PTYP"] = NetellerType
	r.setPaymentToken(netellerId)
}

// Set a POLI type
func (r *Request) SetPoliPayment(popID string) {
	r.data["PTYP"] = POLIType
	r.setPaymentToken(popID)
}

// Set a Single Euro Payments Area payment type
func (r *Request) SetSepaPayment(sepaID string) {
	r.data["PTYP"] = SEPAType
	r.setPaymentToken(sepaID)
}

// Set a Skrill/Mooneybookers payment type
func (r *Request) SetSkrillPayment(skrillID string) {
	r.data["PTYP"] = SkrillType
	r.setPaymentToken(skrillID)
}

// Set a Sofort payment type
func (r *Request) SetSofortPayment(sofortID string) {
	r.data["PTYP"] = SofortType
	r.setPaymentToken(sofortID)
}

// Set a token payment type
func (r *Request) SetTokenPayment(tokenID string) {
	r.data["PTYP"] = TokenType
	r.setPaymentToken(tokenID)
}

// Set payment encoding with either KHASH or MASK values.
func (r *Request) SetPaymentEncoding() {
	r.data["PENC"] = "MASK"
}

// Sets a card payment type and masks the card number using the maskPaymentToken method.
func (r *Request) SetPaymentMasked(cardNumber string) {
	result := r.maskPaymentToken(cardNumber)
	r.data["PTOK"] = result
	r.data["PTYP"] = CardType
	r.SetPaymentEncoding()
}

/**
 * Set the masked payment token the following way.
 * First 6 characters remain the same, the following set of characters up to the last 4 are
 * replaced with 'X' character and the last 4 remain the same also.
 * Example: "0007380568572514" -> "000738XXXXXX2514"
 */
func (r *Request) maskPaymentToken(token string) string {
	// get first 6 characters
	result := token[0:6]

	for i := 6; i < (len(token) - 4); i++ {
		result += "X"
	}

	// get last 4 characters
	result += token[len(token)-4:]

	return result
}

// Set the last 4 characters on the payment token.
func (r *Request) SetPaymentTokenLast4(last4 string) {
	r.data["LAST4"] = last4
}

// Set the payment type and payment token. Payment token must be raw, i.e. NOT Khashed.
func (r *Request) SetPayment(paymentType, paymentToken string) {
	// check if last 4 is already set. if not then set it
	if val, ok := r.data["LAST4"]; !ok {
		if len(val) >= 4 {
			r.data["LAST4"] = paymentToken[len(paymentToken)-4:]
		} else {
			r.data["LAST4"] = paymentToken
		}
	}
	r.data["PTYP"] = paymentType
	r.data["PTOK"] = paymentToken
}

func (r *Request) GetResponse() (*response.Response, error) {
	myResp := &response.Response{}
	client := &http.Client{}

	// set timeout
	client.Timeout = time.Second * time.Duration(ConnectionTimeout)

	r.SetVersion(Version)

	form := url.Values{}
	for key, value := range r.data {
		form.Add(key, value)
	}

	req, err := http.NewRequest("POST", r.Settings.GetRISURL(), strings.NewReader(form.Encode()))
	if err != nil {
		return myResp, err
	}

	// set request headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Kount-Api-Key", r.Settings.GetAPIKey())
	req.Header.Add("X-Kount-Merc-Id", r.data["MERC"])

	resp, err := client.Do(req)
	if err != nil {
		return myResp, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return myResp, err
	}

	myResp.Raw = string(body)
	myResp.Digest()
	return myResp, nil
}
