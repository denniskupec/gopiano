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
	"strings"
	"time"

	"golang.org/x/crypto/blowfish"

	"denniskupec.com/gopiano/coder"
	"denniskupec.com/gopiano/request"
	"denniskupec.com/gopiano/response"
)

// Describes a particular type of client to emulate.
type ClientDescription struct {
	DeviceModel string
	Username    string
	Password    string
	BaseURL     string
	EncryptKey  string
	DecryptKey  string
	Version     string
}

// Class for a Client object.
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

// Create a new Client with specified ClientDescription
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
func (c *Client) decrypt(data string) (string, error) {
	chunks := make([]string, 0)
	for i := 0; i < len(data); i += 16 {
		var buf [16]byte
		var decoded, decrypted [8]byte
		copy(buf[:], data[i:])
		_, err := hex.Decode(decoded[:], buf[:])
		if err != nil {
			return "", err
		}
		c.decrypter.Decrypt(decrypted[:], decoded[:])
		chunks = append(chunks, strings.Trim(string(decrypted[:]), "\x00"))
	}
	return strings.Join(chunks, ""), nil
}

// Client.PandoraCall is the basic function to send an HTTP POST to pandora.com.
// Arguments: protocol is either "https://" or "http://", method is whatever must be in
// the "method" url argument and specifies the remote procedure to call, body is an io.Reader
// to be passed directly into http.Post, and data is to be passed to json.Unmarshal to parse
// the JSON response.
func PandoraCall(callUrl string, body io.Reader, data interface{}) error {
	resp, err := http.Post(callUrl, "text/plain", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var wrap response.Wrapper
	if err := json.NewDecoder(resp.Body).Decode(&wrap); err != nil {
		return err
	}

	if wrap.Stat == "fail" {
		if message, ok := response.ErrorCodeMap[wrap.Code]; ok {
			wrap.Message = message
		}
		return wrap.ErrorResponse
	}

	return json.Unmarshal(wrap.Result, &data)
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

func (c *Client) Call(req request.Type, data interface{}) error {
	enc := coder.New(c.encrypter)
	if err := json.NewEncoder(enc).Encode(req); err != nil {
		return err
	}

	return PandoraCall(c.formatURL(req), enc, data)
}

// Most calls require a SyncTime int argument (Unix epoch). We store our current time offset
// but must calculate the SyncTime for each call. This method does that.
func (c *Client) GetSyncTime() int {
	return int(time.Now().Add(c.timeOffset).Unix())
}

func (c *Client) Token() request.UserToken {
	return request.UserToken{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
}
