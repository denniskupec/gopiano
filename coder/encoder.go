package coder

import (
	"bytes"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/blowfish"
)

func New(cipher *blowfish.Cipher) io.ReadWriter {
	return &encoder{blow: cipher}
}

type encoder struct {
	in, out bytes.Buffer
	blow    *blowfish.Cipher
}

func (enc *encoder) Write(p []byte) (n int, err error) {
	n, err = enc.in.Write(p)

	for 8 <= enc.in.Len() {
		enc.encodeBlock()
	}

	return n, err
}

func (enc *encoder) Read(p []byte) (n int, err error) {
	n, err = enc.out.Read(p)
	if err != io.EOF {
		return n, err
	}

	enc.Flush()
	return enc.out.Read(p)
}

// Flush makes any remaining data available to the reader
// even if it does not fall on a block size.
func (enc *encoder) Flush() {
	for 0 < enc.in.Len() {
		enc.encodeBlock()
	}
}

func (enc *encoder) encodeBlock() {
	var buf1 [8]byte
	enc.in.Read(buf1[:])

	enc.blow.Encrypt(buf1[:], buf1[:])

	var buf2 [16]byte
	hex.Encode(buf2[:], buf1[:])

	enc.out.Write(buf2[:])
}
