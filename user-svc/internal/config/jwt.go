package config

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	jwtSignature []byte
	expiredTime  = 1000
)

func GenerateJwtSignature() {
	rand.Seed(time.Now().UTC().UnixNano())
	jwtSignature = []byte(strconv.FormatUint(rand.Uint64(), 10))
}

func GetJwtExpiredTime() int {
	return expiredTime
}

var GetJwtSignature = func() []byte {
	return jwtSignature
}
