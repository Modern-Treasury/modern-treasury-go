package moderntreasury

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type WebhookService struct {
	Options []option.RequestOption
}

func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// To verify that a webhook was actually sent by Modern Treasury, every payload is
// signed with a signature that is passed through the `X-Signature` HTTP header.
//
// This method will generate a signature based off of your webhook key which can be
// found in the Developer Settings,
// https://app.moderntreasury.com/developers/webhooks, and the webhook payload.
//
// You can then compare the generated signature with the signature sent with the
// request, if they match then the webhook was sent by Modern Treasury.
func (r *WebhookService) GetSignature(payload []byte, key string) (res string, err error) {
	h := hmac.New(sha256.New, []byte(key))
	_, err = h.Write([]byte(payload))
	if err != nil {
		return "", err
	}

	signature := hex.EncodeToString(h.Sum(nil))
	return signature, nil

}

// Returns whether or not the webhook payload was sent by Modern Treasury.
func (r *WebhookService) ValidateSignature(payload []byte, key string, headers http.Header) (res bool, err error) {
	expectedSignature := headers.Get("X-Signature")
	if expectedSignature == "" {
		return false, errors.New("Could not find an X-Signature header")
	}

	signature, err := r.GetSignature(payload, key)
	if err != nil {
		return false, err
	}

	return signature == expectedSignature, nil

}
