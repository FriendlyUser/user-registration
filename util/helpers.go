package util
import (
	"os"
	"log"
)


func IsEmpty(data string) bool {
    if len(data) == 0 {
        return true
    } else {
        return false
    }
}

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	// TODO update to exist only on critical envs not being set
	if !exists {
		value = fallback
		log.Fatal(key + " is not set")
	}
	return value
}
