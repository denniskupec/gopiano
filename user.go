package gopiano

import (
	"denniskupec.com/gopiano/request"
	"denniskupec.com/gopiano/response"
)

// Client.UserCanSubscribe returns whehter a user is subscribed or can subscribe
// to the premium Pandora One service.
// Calls API method "user.canSubscribe"
func (c *Client) UserCanSubscribe() (*response.UserCanSubscribe, error) {
	requestData := request.UserCanSubscribe{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	var resp response.UserCanSubscribe
	err := c.BlowfishJSONCall(c.formatURL("http://", "user.canSubscribe"), requestData, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Client.UserCreateUser creates a new Pandora user.
// Argument username must be in the form of an email address. gender must be either "male" or "female".
// countryCode must be "US".
// Calls API method "user.createUser"
func (c *Client) UserCreateUser(username, password, gender, countryCode string, zipCode, birthYear int, emailOptin bool) (*response.UserCreateUser, error) {
	requestData := request.UserCreateUser{
		PartnerAuthToken: c.partnerAuthToken,
		AccountType:      "registered",
		RegisteredType:   "user",
		Username:         username,
		Password:         password,
		Gender:           gender,
		ZipCode:          zipCode,
		CountryCode:      countryCode,
		BirthYear:        birthYear,
		EmailOptin:       emailOptin,
		SyncTime:         c.GetSyncTime(),
	}

	var resp response.UserCreateUser
	err := c.BlowfishJSONCall(c.formatURL("https://", "user.createUser"), requestData, &resp)
	if err != nil {
		return nil, err
	}

	// Set user data onto client for later use.
	c.userAuthToken = resp.Result.UserAuthToken
	c.userID = resp.Result.UserID

	return &resp, nil
}

// Client.UserEmailPassword resends registration email, maybe?
// Calls API method "user.emaillPassword"
func (c *Client) UserEmailPassword(username string) error {
	requestData := request.UserEmailPassword{
		Username:         username,
		PartnerAuthToken: c.partnerAuthToken,
		SyncTime:         c.GetSyncTime(),
	}

	var resp interface{}
	return c.BlowfishJSONCall(c.formatURL("https://", "user.emailPassword"), requestData, &resp)
}

// Client.UserGetBookmarks returns the users bookmarked artists and songs.
// Also see BookmarkAddArtistBookmark and BookmarkAddSongBookmark.
// Calls API method "user.getBookmarks"
func (c *Client) UserGetBookmarks() (*response.UserGetBookmarks, error) {
	requestData := request.UserGetBookmarks{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	var resp response.UserGetBookmarks
	err := c.BlowfishJSONCall(c.formatURL("http://", "user.getBookmarks"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.UserGetStationList gets the list of a users stations.
// Call API method "user.getStationList"
func (c *Client) UserGetStationList(includeStationArtURL bool) (*response.UserGetStationList, error) {
	requestData := request.UserGetStationList{
		UserAuthToken:        c.userAuthToken,
		SyncTime:             c.GetSyncTime(),
		IncludeStationArtURL: includeStationArtURL,
	}

	var resp response.UserGetStationList
	err := c.BlowfishJSONCall(c.formatURL("http://", "user.getStationList"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.UserGetStationList returns the checksum of the user's station list.
// Call API method "user.getStationListChecksum"
func (c *Client) UserGetStationListChecksum() (*response.UserGetStationListChecksum, error) {
	requestData := request.UserGetStationListChecksum{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	var resp response.UserGetStationListChecksum
	err := c.BlowfishJSONCall(c.formatURL("http://", "user.getStationListChecksum"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.UserSetQuickMix selects the stations that should be in the special QuickMix station.
// Call API method "user.setQuickMix"
func (c *Client) UserSetQuickMix(stationIDs []string) error {
	requestData := request.UserSetQuickMix{
		QuickMixStationIDs: stationIDs,
		UserAuthToken:      c.userAuthToken,
		SyncTime:           c.GetSyncTime(),
	}

	var resp interface{}
	return c.BlowfishJSONCall(c.formatURL("https://", "user.setQuickMix"), requestData, &resp)
}

// Client.UserSleepSong marks a song to be not played again for 1 month.
// Calls API method "user.sleepSong"
func (c *Client) UserSleepSong(trackToken string) error {
	requestData := request.UserSleepSong{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	var resp interface{}
	return c.BlowfishJSONCall(c.formatURL("https://", "user.sleepSong"), requestData, &resp)
}
