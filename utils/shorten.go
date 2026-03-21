package utils
import (
	"encoding/base64"
	"fmt"
	"time"
)

func GetShortCode() string {
	fmt.Println("Shortening URL")
	ts := time.Now().UnixNano()
	fmt.Println("Timestamp: ", ts)

	ts_bytes := []byte(fmt.Sprintf("%d", ts))
	key := base64.StdEncoding.EncodeToString(ts_bytes)
	fmt.Println("Key: ", key)
	//removing the last two chars since they are usually equal sign
	key = key[:len(key)-2]
	// returining the last chars after 16 chars, these are almost alway different
	return key[16:]
}