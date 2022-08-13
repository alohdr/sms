package utils

import "time"

const (
	SecretToken          string        = "IWANSECRET"
	ExpiredTokenDuration time.Duration = time.Minute * 300

	SuccessResponseMessege string = "Success Retriving Data"
	StatusCodeSuccess      string = "200"
)
