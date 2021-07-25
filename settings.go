package main

import (
	
	"github.com/kubewarden/gjson"
	kubewarden "github.com/kubewarden/policy-sdk-go"

	"fmt"
)

type Settings struct {
	DenyPalindromeKey bool `json:"deny_palindrome_key"`
}

// Builds a new Settings instance starting from a validation
// request payload:
// {
//    "request": ...,
//    "settings": {
//       "deny_palindrome_key": true
//    }
// }
func NewSettingsFromValidationReq(payload []byte) (Settings, error) {
	return newSettings(
		payload,
		"settings.deny_palindrome_key")
}

// Builds a new Settings instance starting from a Settings
// payload:
// {
//    "deny_palindrome_key": ...
// }
func NewSettingsFromValidateSettingsPayload(payload []byte) (Settings, error) {
	return newSettings(
		payload,
		"deny_palindrome_key")
}

func newSettings(payload []byte, paths ...string) (Settings, error) {
	if len(paths) != 1 {
		return Settings{}, fmt.Errorf("wrong number of json paths")
	}

	data := gjson.GetManyBytes(payload, paths...)

	deny_palindrome_key := true
	if data[0].Exists() {
		deny_palindrome_key = data[0].Bool()
	}

	return Settings{
		DenyPalindromeKey: deny_palindrome_key,
	}, nil
}

// No special check has to be done
func (s *Settings) Valid() bool {
	return true
}

func validateSettings(payload []byte) ([]byte, error) {
	logger.Info("validating settings")

	settings, err := NewSettingsFromValidateSettingsPayload(payload)
	if err != nil {
		return []byte{}, err
	}

	if settings.Valid() {
		logger.Info("accepting settings")
		return kubewarden.AcceptSettings()
	}

	logger.Warn("rejecting settings")
	return kubewarden.RejectSettings(kubewarden.Message("Provided settings are not valid"))
}
