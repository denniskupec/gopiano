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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/crypto/blowfish"

	"denniskupec.com/gopiano/coder"
	"denniskupec.com/gopiano/responses"
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
	req, err := http.NewRequest("POST", callUrl, body)
	if err != nil {
		return err
	}
	//req.Header.Add("User-Agent", "gopiano")
	req.Header.Add("Content-type", "text/plain")

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var errResp responses.ErrorResponse
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseBody, &errResp)
	if err != nil {
		return err
	}

	if errResp.Stat == "fail" {
		if message, ok := responses.ErrorCodeMap[errResp.Code]; ok {
			errResp.Message = message
		}
		return errResp
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) formatURL(protocol, method string) string {
	urlArgs := url.Values{
		"method": {method},
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

	return protocol + c.description.BaseURL + "?" + urlArgs.Encode()
}

// Client.BlowfishCall first encrypts the body before calling PandoraCall.
// Arguments are identical to PandoraCall.
func (c *Client) BlowfishCall(protocol string, method string, body io.Reader, data interface{}) error {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	enc := coder.New(c.encrypter)
	enc.Write(bodyBytes)

	return PandoraCall(c.formatURL(protocol, method), enc, data)
}

// Most calls require a SyncTime int argument (Unix epoch). We store our current time offset
// but must calculate the SyncTime for each call. This method does that.
func (c *Client) GetSyncTime() int {
	return int(time.Now().Add(c.timeOffset).Unix())
}
