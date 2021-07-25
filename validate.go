package main

import (
	"fmt"

	"github.com/kubewarden/gjson"
	kubewarden "github.com/kubewarden/policy-sdk-go"
)

func validate(payload []byte) ([]byte, error) {
	if !gjson.ValidBytes(payload) {
		return kubewarden.RejectRequest(
			kubewarden.Message("Not a valid JSON document"),
			kubewarden.Code(400))
	}

	settings, err := NewSettingsFromValidationReq(payload)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(400))
	}

	data := gjson.GetBytes(
		payload,
		"request.object.metadata.labels")
	//logger.InfoWithFields("breaking news !", func(e onelog.Entry) {
	//	e.String("data", data.Str)
	//})

	data.ForEach(func(key, value gjson.Result) bool {
		label := key.String()

		logger.Info(label)

		if settings.DenyPalindromeKey == true {
			if isPalindrome(label) {
				logger.Info("Found palindrome")
				err = fmt.Errorf("Label has a palindrome key %s, hence denied", label)
				return false
			}
		}

		return true
	})

	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.NoCode)
	}

	return kubewarden.AcceptRequest()
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
