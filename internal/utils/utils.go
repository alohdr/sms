package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CreateOtp() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Int()
	arrString := strings.Split(strconv.Itoa(v), "")
	otp := arrString[len(arrString)-6:]
	otpstr := strings.Join(otp, "")

	return otpstr
}

func Expired(now time.Time, v string) (bool, error) {

	exp, err := time.ParseInLocation(time.RFC3339, v, now.Location())
	if err != nil {
		return false, err
	}

	if now.Sub(exp).Seconds() <= 0 {
		return false, nil
	}

	return true, nil
}
