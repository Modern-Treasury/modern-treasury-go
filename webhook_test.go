// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"net/http"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

func TestGetSignature(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)

	payload := []byte("foo")
	key := "foo"
	expected := "08ba357e274f528065766c770a639abf6809b39ccfd37c2a3157c7f51954da0a"
	signature, err := c.Webhooks.GetSignature(payload, key)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if expected != signature {
		t.Fatalf("expected signatures to match: %s, %s", expected, signature)
	}
}

func TestValidateSignature(t *testing.T) {
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)

	payload := []byte(`{"foo":"bar"}`)
	headers := http.Header{
		"X-Signature": []string{"57e14f6f354543f0101fb06ea24df7731d90087b76651e3497345e22a3622940"},
	}
	key := "foo"

	valid, err := c.Webhooks.ValidateSignature(payload, key, headers)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !valid {
		t.Errorf("Expected signature to be valid")
	}

	headers.Set("X-Signature", "hello")
	valid, err = c.Webhooks.ValidateSignature(payload, key, headers)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if valid {
		t.Errorf("Expected signature to be invalid")
	}

	headers.Del("X-Signature")
	_, err = c.Webhooks.ValidateSignature(payload, key, headers)
	if err == nil {
		t.Errorf("Expected error due to missing X-Signature header")
	}
}
