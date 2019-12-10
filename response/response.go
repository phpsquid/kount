package response

import (
	"github.com/phpsquid/kount/data"
	"strconv"
	"strings"
)

type Response struct {
	Raw  string            // Raw response string
	Data map[string]string // a map containing extracted response body values
}

// Digest parses the raw response string and updates Data map[string]string
func (r *Response) Digest() {
	data := make(map[string]string)
	r.Data = data

	// split raw response on new line character
	splitRaw := strings.Split(r.Raw, "\n")

	// loop over the new []string and update Data map[string]string
	for _, v := range splitRaw {
		splitValue := strings.Split(v, "=")
		if len(splitValue) == 2 {
			r.Data[splitValue[0]] = splitValue[1]
		}
	}
}

// Get an explicit parameter from the response
func (r *Response) GetParam(key string) string {
	return r.Data[key]
}

// Get the version number
func (r *Response) GetVersion() string {
	return r.GetParam("VERS")
}

// Get the RIS mode
func (r *Response) GetMode() string {
	return r.GetParam("MODE")
}

// Get the transaction id
func (r *Response) GetTransactionId() string {
	return r.GetParam("TRAN")
}

// Get the merchant id
func (r *Response) GetMerchantId() string {
	return r.GetParam("MERC")
}

// Get the merchant gateway's customer id for Kount Central
func (r *Response) GetKCCustomerId() string {
	return r.GetParam("KC_CUSTOMER_ID")
}

// Get the session id
func (r *Response) GetSessionId() string {
	return r.GetParam("SESS")
}

// Get the site
func (r *Response) GetSite() string {
	return r.GetParam("SITE")
}

// Get the merchant order number
func (r *Response) GetOrderNumber() string {
	return r.GetParam("ORDR")
}

// Get the RIS auto response (A/R/D)
func (r *Response) GetAuto() string {
	return r.GetParam("AUTO")
}

// Get the merchant defined decision reason code.
func (r *Response) GetReasonCode() string {
	return r.GetParam("REASON_CODE")
}

// Get the Kount score
func (r *Response) GetScore() string {
	return r.GetParam("SCOR")
}

// Get the Kount OMNI score
func (r *Response) GetOmniscore() string {
	return r.GetParam("OMNISCORE")
}

// Get the geox
func (r *Response) GetGeox() string {
	return r.GetParam("GEOX")
}

// Get the card brand
func (r *Response) GetBrand() string {
	return r.GetParam("BRND")
}

// Get the 6 week velocity
func (r *Response) GetVelo() string {
	return r.GetParam("VELO")
}

// Get the 6 hour velocity
func (r *Response) GetVmax() string {
	return r.GetParam("VMAX")
}

// Get the network type
func (r *Response) GetNetwork() string {
	return r.GetParam("NETW")
}

// Get the 'know your customer' flash
func (r *Response) GetKnowYourCustomer() string {
	return r.GetParam("KYCF")
}

// Get the region
func (r *Response) GetRegion() string {
	return r.GetParam("REGN")
}

// Get the kaptcha flag, enabled upon request and for when a RIS record has
func (r *Response) GetKaptcha() string {
	return r.GetParam("KAPT")
}

// Get a string representing whether the remote device is using a proxy ("Y" or "N")
func (r *Response) GetProxy() string {
	return r.GetParam("PROXY")
}

// Get the number of transactions associated with the email
func (r *Response) GetEmails() string {
	return r.GetParam("EMAILS")
}

// Get the two character country code setting in the remote device's browser
func (r *Response) GetHttpCountry() string {
	return r.GetParam("HTTP_COUNTRY")
}

// Get a string representing the time zone of the customer as a 3 digit number
func (r *Response) GetTimeZone() string {
	return r.GetParam("TIMEZONE")
}

// Get the number of transactions associated with the credit card
func (r *Response) GetCards() string {
	return r.GetParam("CARDS")
}

/**
 * Get a string representing whether the end device is a remotely
 * controlled computer
 */
func (r *Response) GetPCRemote() string {
	return r.GetParam("PC_REMOTE")
}

// Get the number of transactions associated with the particular device
func (r *Response) GetDevices() string {
	return r.GetParam("DEVICES")
}

/**
 * Get a string representing the five layers (Operating System, SSL, HTTP,
 * Flash, JavaScript) of the remote device
 */
func (r *Response) GetDeviceLayers() string {
	return r.GetParam("DEVICE_LAYERS")
}

// Get the mobile device's wireless application protocol
func (r *Response) GetMobileForwarder() string {
	return r.GetParam("MOBILE_FORWARDER")
}

/**
 * Get a string representing whether or not the remote device is voice
 * controlled
 */
func (r *Response) GetVoiceDevice() string {
	return r.GetParam("VOICE_DEVICE")
}

// Get local time of the remote device in the YYYY-MM-DD format
func (r *Response) GetLocalTime() string {
	return r.GetParam("LOCALTIME")
}

// Get the mobile device type
func (r *Response) GetMobileType() string {
	return r.GetParam("MOBILE_TYPE")
}

// Get the device finger print
func (r *Response) GetFingerPrint() string {
	return r.GetParam("FINGERPRINT")
}

// Get a string representing whether or not the remote device allows flash
func (r *Response) GetFlash() string {
	return r.GetParam("FLASH")
}

// Get the language setting on the remote device
func (r *Response) GetLanguage() string {
	return r.GetParam("LANGUAGE")
}

// Get the remote device's country of origin as a two character code
func (r *Response) GetCountry() string {
	return r.GetParam("COUNTRY")
}

// Get a string representing whether the remote device allows JavaScript
func (r *Response) GetJavaScript() string {
	return r.GetParam("JAVASCRIPT")
}

// Get a string representing whether the remote device allows cookies
func (r *Response) GetCookies() string {
	return r.GetParam("COOKIES")
}

// Get a string representing whether the remote device is a mobile device
func (r *Response) GetMobileDevice() string {
	return r.GetParam("MOBILE_DEVICE")
}

/**
 * Get MasterCard Fraud Score associated with the RIS transaction. Please
 * contact your Kount representative to enable support for this feature in
 * your merchant account.
 */
func (r *Response) GetMasterCardFraudScore() string {
	return r.GetParam("MASTERCARD")
}

// Get pierced IP address
func (r *Response) GetPiercedIPAddress() string {
	return r.GetParam("PIP_IPAD")
}

// Get latitude of pierced IP address
func (r *Response) GetPiercedIPAddressLatitude() string {
	return r.GetParam("PIP_LAT")
}

// Get longitude of pierced IP address
func (r *Response) GetPiercedIPAddressLongitude() string {
	return r.GetParam("PIP_LON")
}

// Get country of pierced IP address
func (r *Response) GetPiercedIPAddressCountry() string {
	return r.GetParam("PIP_COUNTRY")
}

// Get region of pierced IP address
func (r *Response) GetPiercedIPAddressRegion() string {
	return r.GetParam("PIP_REGION")
}

// Get city of pierced IP address
func (r *Response) GetPiercedIPAddressCity() string {
	return r.GetParam("PIP_CITY")
}

// Get organization of pierced IP address
func (r *Response) GetPiercedIPAddressOrganization() string {
	return r.GetParam("PIP_ORG")
}

// Get proxy IP address
func (r *Response) GetIPAddress() string {
	return r.GetParam("IP_IPAD")
}

// Get latitude of proxy IP address
func (r *Response) GetIPAddressLatitude() string {
	return r.GetParam("IP_LAT")
}

// Get longitude of proxy IP address
func (r *Response) GetIPAddressLongitude() string {
	return r.GetParam("IP_LON")
}

// Get country of proxy IP address
func (r *Response) GetIPAddressCountry() string {
	return r.GetParam("IP_COUNTRY")
}

// Get region of proxy IP address
func (r *Response) GetIPAddressRegion() string {
	return r.GetParam("IP_REGION")
}

// Get city of proxy IP address
func (r *Response) GetIPAddressCity() string {
	return r.GetParam("IP_CITY")
}

// Get organization of proxy IP address
func (r *Response) GetIPAddressOrganization() string {
	return r.GetParam("IP_ORG")
}

// Get date device first seen
func (r *Response) GetDateDeviceFirstSeen() string {
	return r.GetParam("DDFS")
}

// Get user agent string
func (r *Response) GetUserAgentString() string {
	return r.GetParam("UAS")
}

// Get device screen resolution
func (r *Response) GetDeviceScreenResolution() string {
	return r.GetParam("DSR")
}

// Get operating system (derived from user agent string)
func (r *Response) GetOS() string {
	return r.GetParam("OS")
}

// Get browser (derived from user agent string)
func (r *Response) GetBrowser() string {
	return r.GetParam("BROWSER")
}

// Implement Stringer interface
func (r *Response) String() string {
	return r.Raw
}

// Get a possible error code
func (r *Response) GetErrorCode() string {
	return r.GetParam("ERRO")
}

// Get rules triggered
func (r *Response) GetRulesTriggered() map[string]string {
	ruleCount := r.GetNumberRulesTriggered()
	rules := make(map[string]string)

	for i := 0; i < ruleCount; i++ {
		ruleID := r.GetParam("RULE_ID_" + strconv.Itoa(i))
		rules[ruleID] = r.GetParam("RULE_DESCRIPTION_" + strconv.Itoa(i))
	}

	return rules
}

// Get the number of rules triggered with the response.
func (r *Response) GetNumberRulesTriggered() int {
	i, err := strconv.Atoi(r.GetParam("RULES_TRIGGERED"))
	if err != nil {
		return 0
	}
	return i
}

// Get a string slice of the warnings associated with this response.
func (r *Response) GetWarnings() []string {
	warningCount := r.GetWarningCount()
	var warnings []string

	for i := 0; i < warningCount; i++ {
		warning := r.GetParam("WARNING_" + strconv.Itoa(i))
		warnings = append(warnings, warning)
	}

	return warnings
}

// Get the number of warnings triggered with the response.
func (r *Response) GetWarningCount() int {
	i, err := strconv.Atoi(r.GetParam("WARNING_COUNT"))
	if err != nil {
		return 0
	}
	return i
}

// Get a string slice of the errors associated with this response.
func (r *Response) GetErrors() []string {
	errorCount := r.GetErrorCount()
	var errors []string

	for i := 0; i < errorCount; i++ {
		errMessage := r.GetParam("ERROR_" + strconv.Itoa(i))
		errors = append(errors, errMessage)
	}

	return errors
}

// Get the number of errors triggered with the response.
func (r *Response) GetErrorCount() int {
	i, err := strconv.Atoi(r.GetParam("ERROR_COUNT"))
	if err != nil {
		return 0
	}
	return i
}

// Get a map[string]string of the rules counters associated with the response.
func (r *Response) GetCountersTriggered() map[string]string {
	numCounters := r.GetNumberCountersTriggered()
	counters := make(map[string]string)

	for i := 0; i < numCounters; i++ {
		counterName := r.GetParam("COUNTER_NAME_" + strconv.Itoa(i))
		counters[counterName] = r.GetParam("COUNTER_VALUE_" + strconv.Itoa(i))
	}

	return counters
}

// Get the number of rules counters triggered in the response.
func (r *Response) GetNumberCountersTriggered() int {
	i, err := strconv.Atoi(r.GetParam("COUNTERS_TRIGGERED"))
	if err != nil {
		return 0
	}
	return i
}

// Get a string slice of the Kount Central warnings associated with the response.
func (r *Response) GetKCWarnings() []string {
	warningCount := r.GetKCWarningCount()
	var warnings []string

	for i := 0; i < warningCount; i++ {
		warning := r.GetParam("KC_WARNING_" + strconv.Itoa(i))
		warnings = append(warnings, warning)
	}

	return warnings
}

// Get the number of Kount Central warnings associated with the response.
func (r *Response) GetKCWarningCount() int {
	i, err := strconv.Atoi(r.GetParam("KC_WARNING_COUNT"))
	if err != nil {
		return 0
	}
	return i
}

// Get a string slice of the Kount Central warnings associated with the response.
func (r *Response) GetKCErrors() []string {
	errorCount := r.GetKCErrorCount()
	var errors []string

	for i := 0; i < errorCount; i++ {
		warning := r.GetParam("KC_ERROR_" + strconv.Itoa(i))
		errors = append(errors, warning)
	}

	return errors
}

// Get the number of Kount Central errors associated with the response.
func (r *Response) GetKCErrorCount() int {
	i, err := strconv.Atoi(r.GetParam("KC_ERROR_COUNT"))
	if err != nil {
		return 0
	}
	return i
}

// Get the Kount Central threshold events associated with the decision
func (r *Response) GetKCEvents() []data.KCEvent {
	eventCount := r.GetKCEventCount()
	var events []data.KCEvent

	for i := 0; i < eventCount; i++ {
		decision := r.GetParam("KC_EVENT_" + strconv.Itoa(i) + "_DECISION")
		expression := r.GetParam("KC_EVENT_" + strconv.Itoa(i) + "_EXPRESSION")
		code := r.GetParam("KC_EVENT_" + strconv.Itoa(i) + "_CODE")
		kcEvent := data.KCEvent{
			Decision:   decision,
			Expression: expression,
			Code:       code,
		}
		events = append(events, kcEvent)
	}

	return events
}

// Get the number of Kount Central events associated with the response.
func (r *Response) GetKCEventCount() int {
	i, err := strconv.Atoi(r.GetParam("KC_TRIGGERED_COUNT"))
	if err != nil {
		return 0
	}
	return i
}
