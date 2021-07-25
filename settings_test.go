package main

import (
	"testing"
)

func TestParsingSettingsWithValueProvidedFromValidationReq(t *testing.T) {
	request := `
	{
		"request": "doesn't matter here",
		"settings": {
			"deny_palindrome_key": true
		}
	}
	`
	rawRequest := []byte(request)

	settings, err := NewSettingsFromValidationReq(rawRequest)
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	if settings.DenyPalindromeKey != true {
		t.Errorf("Wrong value for DenyPalindromeKey")
	}

}

func TestSettingsAreValid(t *testing.T) {
	request := `
	{
		"deny_palindrome_key": true
	}
	`
	rawRequest := []byte(request)

	settings, err := NewSettingsFromValidateSettingsPayload(rawRequest)
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	if !settings.Valid() {
		t.Errorf("Settings are reported as not valid")
	}
}
