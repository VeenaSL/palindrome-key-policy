#!/usr/bin/env bats

@test "reject because label key is palindrome" {
  run kwctl run policy.wasm -r test_data/deny.json --settings-json '{"deny_palindrome_key": true}'

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request rejected
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*false') -ne 0 ]
  [ $(expr "$output" : ".*Label has a palindrome key level, hence denied.*") -ne 0 ]
}

@test "accept because label is not palindrome" {
  run kwctl run policy.wasm -r test_data/allow.json --settings-json '{"deny_palindrome_key": true}'
  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}
