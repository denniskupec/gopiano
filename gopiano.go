/*
Package gopiano provides a thin wrapper library around the Pandora.com client API.

This client API has been reverse engineered and documentation is available at
https://6xq.net/pandora-apidoc/

The package provides a Client struct with a myriad of methods which interact with the
Pandora JSON API's own methods. Each method returns a struct of the parsed JSON data and an error.
All of the responses that these methods return can be found in the responses subpackage. There
is also a requests subpackage but mostly you don't need to bother with those; they get instantiated
by these client methods.
*/
package gopiano

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/crypto/blowfish"

	"denniskupec.com/gopiano/coder"
	"denniskupec.com/gopiano/request"
	"denniskupec.com/gopiano/response"
)

// ClientDescription describes a particular type of client to emulate.
type ClientDescription struct {
	DeviceModel string
	Username    string
	Password    string
	BaseURL     string
	EncryptKey  string
	DecryptKey  string
	Version     string
}

// Client information needed to interface with pandora API.
type Client struct {
	description      ClientDescription
	encrypter        *blowfish.Cipher
	decrypter        *blowfish.Cipher
	timeOffset       time.Duration
	partnerAuthToken string
	partnerID        string
	userAuthToken    string
	userID           string
}

// NewClient creates a new Client with specified ClientDescription
func NewClient(d ClientDescription) (*Client, error) {
	encrypter, err := blowfish.NewCipher([]byte(d.EncryptKey))
	if err != nil {
		return nil, err
	}
	decrypter, err := blowfish.NewCipher([]byte(d.DecryptKey))
	if err != nil {
		return nil, err
	}
	return &Client{
		description: d,
		encrypter:   encrypter,
		decrypter:   decrypter,
	}, nil
}

// Blowfish decrypts a string in ECB mode.
// Some data returned from the Pandora API is encrypted. This decrypts it.
// The key for the decryption is provided by the ClientDescription.
//
// Decryption is done inplace and returns the number of bytes successfully decrypted.
func (c *Client) decrypt(data []byte) (n int, err error) {
	for in, out := data, data; 16 <= len(in); in, out = in[16:], out[8:] {
		if _, err := hex.Decode(out[:8], in[:16]); err != nil {
			return n, err
		}

		c.decrypter.Decrypt(out[:8], out[:8])

		n += 8
	}

	for 0 < n && data[n-1] == '\x00' {
		n--
	}

	return n, nil
}

// PandoraCall is the basic function to send an HTTP POST to pandora.com.
func PandoraCall(callURL string, body io.Reader) (json.RawMessage, error) {
	resp, err := http.Post(callURL, "text/plain", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var wrap response.Wrapper
	if err := json.NewDecoder(resp.Body).Decode(&wrap); err != nil {
		return nil, err
	}

	if wrap.Stat == "ok" {
		return wrap.Result, nil
	}

	if message, ok := response.ErrorCodeMap[wrap.Code]; ok {
		wrap.ErrorResponse.Message = message
	}

	return nil, wrap.ErrorResponse
}

func (c *Client) formatURL(req request.Type) string {
	urlArgs := url.Values{
		"method": {req.Method()},
	}

	if c.partnerID != "" {
		urlArgs.Add("partner_id", c.partnerID)
	}
	if c.userID != "" {
		urlArgs.Add("user_id", c.userID)
	}
	if c.partnerAuthToken != "" && c.userAuthToken == "" {
		urlArgs.Add("auth_token", c.partnerAuthToken)
	} else if c.userAuthToken != "" {
		urlArgs.Add("auth_token", c.userAuthToken)
	}

	return req.Protocol().URL() + c.description.BaseURL + "?" + urlArgs.Encode()
}

// Call makes the given request to pandora and unmarshals the result into
// the 'data' argument.
func (c *Client) Call(req request.Type, data interface{}) error {
	enc := coder.New(c.encrypter)
	if err := json.NewEncoder(enc).Encode(req); err != nil {
		return err
	}

	res, err := PandoraCall(c.formatURL(req), enc)
	if err != nil {
		return err
	}

	return json.Unmarshal(res, data)
}

// GetSyncTime returns a calculated SyncTime (Unix epoch) which is required
// for most calls.
func (c *Client) GetSyncTime() int {
	return int(time.Now().Add(c.timeOffset).Unix())
}

// Token returns UserToken needed for some requests
func (c *Client) Token() request.UserToken {
	return request.UserToken{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
}
