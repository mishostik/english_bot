package secure

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

func CalcSignature(secret string, message string) string {
	fmt.Println(secret, message)
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

var mtx = sync.Mutex{}

func CalcInternalId(anyString string) (string, error) {

	mtx.Lock()
	defer mtx.Unlock()
	hasher := sha512.New()
	if _, err := hasher.Write([]byte(fmt.Sprintf("%s_%s", anyString, time.Now().Format(time.RFC3339)))); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
