package coder

import (
	"io/ioutil"
	"log"
	"testing"

	"golang.org/x/crypto/blowfish"
)

func TestEncoderSanity(t *testing.T) {
	var tests [][]byte
	for i := 0; i < 16; i++ {
		tests = append(tests, make([]byte, i))
	}

	for _, data := range tests {
		c, err := blowfish.NewCipher([]byte{0xff})
		if err != nil {
			t.Fatal(err)
		}
		enc := encoder{blow: c}

		n, err := enc.Write(data)
		if err != nil {
			t.Fatal(err)
		}
		if len(data) != n {
			t.Fatal("not enough data written")
		}

		output, err := ioutil.ReadAll(&enc)
		if err != nil {
			t.Fatal(err)
		}

		log.Println(n, len(output))
		log.Printf("%s\n", output)
	}
}
