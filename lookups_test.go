package twilio

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestGetPhoneLookup(t *testing.T) {
	t.Parallel()
	client, s := getServer(phoneLookupResponse)
	defer s.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	lookup, err := client.Lookups.Get(ctx, "4157012312")
	if err != nil {
		t.Fatal(err)
	}
	if lookup.PhoneNumber != "+14157012312" {
		t.Errorf("expected PhoneNumber to be %s, got %s", "4157012312", lookup.PhoneNumber)
	}
	if lookup.CallerName.CallerName != "CCSF" {
		t.Errorf("expected CallerName to be %s, got %s", "CCSF", lookup.CallerName.CallerName)
	}
	if lookup.Carrier.Type != "landline" {
		t.Errorf("expected Carrier.Type to be %s, got %s", "landline", lookup.Carrier.Type)
	}
}

func TestRealLookup(t *testing.T) {
	cli := NewLookupClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"), nil)

	l, err := cli.Lookups.Get(context.TODO(), "4157012311")
	if err != nil {
		t.Error(err.Error())
	}
}
