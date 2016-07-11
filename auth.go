package gopiano

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"

	"denniskupec.com/gopiano/request"
	"denniskupec.com/gopiano/response"
)

// Client.AuthPartnerLogin establishes a Partner session with provided
// API username and password and receives a PartnerAuthToken, PartnerID and SyncTime
// which are stored for later calls.
// Calls API method "auth.partnerLogin"
func (c *Client) AuthPartnerLogin() (*response.AuthPartnerLogin, error) {
	requestData := request.PartnerLogin{
		Username:    c.description.Username,
		Password:    c.description.Password,
		Version:     c.description.Version,
		DeviceModel: c.description.DeviceModel,
		IncludeURLs: true,
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp response.AuthPartnerLogin
	err = PandoraCall(c.formatURL("https://", "auth.partnerLogin"), requestDataReader, &resp)
	if err != nil {
		// TODO Handle error
		return nil, err
	}

	syncTime, err := c.decrypt(resp.SyncTime)
	if err != nil {
		return nil, err
	}
	resp.SyncTime = string(syncTime[4:14])
	i, err := strconv.ParseInt(resp.SyncTime, 10, 32)
	if err != nil {
		return nil, err
	}

	// Set partner data onto client for later use.
	c.timeOffset = time.Unix(i, 0).Sub(time.Now())
	c.partnerAuthToken = resp.PartnerAuthToken
	c.partnerID = resp.PartnerID

	return &resp, nil
}

// Client.AuthUserLogin logs in a username and password pair.
// Receives the UserAuthToken which is used in subsequent calls.
// You must call AuthPartnerLogin first, and then either this method or UserCreateUser
// before you proceed.
// Calls API method "auth.userLogin"
func (c *Client) AuthUserLogin(username, password string) (*response.AuthUserLogin, error) {
	requestData := request.UserLogin{
		PartnerAuthToken: c.partnerAuthToken,
		LoginType:        "user",
		Username:         username,
		Password:         password,
		SyncTime:         c.GetSyncTime(),
	}

	var resp response.AuthUserLogin
	err := c.BlowfishJSONCall(c.formatURL("https://", "auth.userLogin"), requestData, &resp)
	if err != nil {
		// TODO Handle error
		return nil, err
	}

	// Set user data onto client for later use.
	c.userAuthToken = resp.UserAuthToken
	c.userID = resp.UserID

	return &resp, nil
}
