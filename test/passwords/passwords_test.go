package passwords

import (
	"log"
	"testing"

	"github.com/sanket143/lisa/src/system/passwords"
)

func init() {
	log.Println("Testing password")
}

// TestBasicLog Function
func TestBasicLog(t *testing.T) {
	pass := "TEST"
	hash, err := passwords.Encrypt(pass)

	if err != nil {
		t.Error(err.Error())
	}

	log.Println(hash)

	ok := passwords.IsValid(hash, pass)
	if !ok {
		t.Error("Password not matching... hashing is not working")
	}

	log.Println("Successfully tested hashing!")
}
