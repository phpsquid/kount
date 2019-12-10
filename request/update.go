package request

import "github.com/phpsquid/kount/settings"

type Update struct {
	*Request
}

func NewUpdate(settings *settings.Settings) *Update {
	req := newRequest(settings)
	u := &Update{req}
	// defaults
	u.SetMode("U")
	u.SetParm("SDK", "CUST")
	return u
}

/**
 * Set the update mode
 * Acceptable values are: "U" or "X"
 */
func (i *Update) SetMode(mode string) {
	i.Request.SetParm("MODE", mode)
}

// Set the transaction id received from the initial inquiry.
func (i *Update) SetTransactionId(transactionID string) {
	i.Request.SetParm("TRAN", transactionID)
}

// Set the refund/chargeback status of this transaction: R = Refund, C = Chargeback.
func (i *Update) SetRefundChargeback(rfcb string) {
	i.Request.SetParm("RFCB", rfcb)
}
