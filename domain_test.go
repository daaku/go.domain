package domain_test

import (
	"fmt"
	"github.com/daaku/go.domain"
	"testing"
)

var validRegistered = map[string]string{
	".fbrell.com":                    "fbrell.com",
	"local.fbrell.com":               "fbrell.com",
	"www.fbrell.com":                 "fbrell.com",
	"fbrell.com":                     "fbrell.com",
	"fbrell.edu":                     "fbrell.edu",
	"www.fbrell.edu":                 "fbrell.edu",
	"registered.com":                 "registered.com",
	"sub.registered.com":             "registered.com",
	"parliament.uk":                  "parliament.uk",
	"sub.registered.valid.uk":        "registered.valid.uk",
	"registered.somedom.kyoto.jp":    "registered.somedom.kyoto.jp",
	"sub.academy.museum":             "sub.academy.museum",
	"subsub.sub.academy.museum":      "sub.academy.museum",
	"sub.nic.pa":                     "nic.pa",
	"registered.sb":                  "registered.sb",
	"sub.registered.sb":              "registered.sb",
	"subsub.registered.something.zw": "registered.something.zw",
	"subsub.registered.9.bg":         "registered.9.bg",
	"registered.co.bi":               "registered.co.bi",
	"sub.registered.bi":              "registered.bi",
	"subsub.registered.ee":           "registered.ee",
}

var validTLD = map[string]string{
	".fbrell.com":                    "com",
	"local.fbrell.com":               "com",
	"www.fbrell.com":                 "com",
	"fbrell.com":                     "com",
	"fbrell.edu":                     "edu",
	"www.fbrell.edu":                 "edu",
	"registered.com":                 "com",
	"sub.registered.com":             "com",
	"parliament.uk":                  "uk",
	"sub.registered.valid.uk":        "valid.uk",
	"registered.somedom.kyoto.jp":    "somedom.kyoto.jp",
	"sub.academy.museum":             "academy.museum",
	"subsub.sub.academy.museum":      "academy.museum",
	"sub.nic.pa":                     "pa",
	"registered.sb":                  "sb",
	"sub.registered.sb":              "sb",
	"subsub.registered.something.zw": "something.zw",
	"subsub.registered.9.bg":         "9.bg",
	"registered.co.bi":               "co.bi",
	"sub.registered.bi":              "bi",
	"subsub.registered.ee":           "ee",
}

var invalidDomains = []string{
	"www.fbrell.foo",
	"local.fbrell.com:43600",
	"invalid-fqdn",
	"org",
	"academy.museum",
	"tokyo.jp",
}

func TestValidRegistered(t *testing.T) {
	for d, expected := range validRegistered {
		actual, err := domain.Registered(d)
		if err != nil {
			t.Errorf(`Error in Registered for domain "%s": %s`, d, err)
			continue
		}
		if expected != actual {
			t.Errorf(
				`Failed Registered for domain "%s". Was expecting "%s" but got "%s"`,
				d, expected, actual)
		}
	}
}

func TestValidTLD(t *testing.T) {
	for d, expected := range validTLD {
		actual, err := domain.TLD(d)
		if err != nil {
			t.Errorf(`Error in TLD for domain "%s": %s`, d, err)
			continue
		}
		if expected != actual {
			t.Errorf(
				`Failed TLD for domain "%s". Was expecting "%s" but got "%s"`,
				d, expected, actual)
		}
	}
}

func TestInvalid(t *testing.T) {
	for _, d := range invalidDomains {
		actual, err := domain.Registered(d)
		if err == nil {
			t.Errorf(`Was expecting error for invalid domain "%s" but got "%s"`,
				d, actual)
		}
	}
}

func ExampleRegistered() {
	var registered string

	registered, _ = domain.Registered("www.facebook.com")
	fmt.Println(registered)

	registered, _ = domain.Registered("apps.facebook.com")
	fmt.Println(registered)

	registered, err := domain.Registered("com")
	fmt.Println(err)
	// Output:
	// facebook.com
	// facebook.com
	// Invalid domain structure: com
}

func ExampleTLD() {
	var tld string

	tld, _ = domain.TLD("www.facebook.com")
	fmt.Println(tld)

	tld, _ = domain.TLD("apps.facebook.com")
	fmt.Println(tld)

	tld, err := domain.TLD("com")
	fmt.Println(err)
	// Output:
	// com
	// com
	// Invalid domain structure: com
}
