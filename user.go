package gopiano

import (
	"denniskupec.com/gopiano/request"
	"denniskupec.com/gopiano/response"
)

// Client.UserCanSubscribe returns whehter a user is subscribed or can subscribe
// to the premium Pandora One service.
// Calls API method "user.canSubscribe"
func (c *Client) UserCanSubscribe() (*response.UserCanSubscribe, error) {
	requestData := request.CanSubscribe{
		UserToken: c.Token(),
	}

	var resp response.UserCanSubscribe
	return &resp, c.Call(requestData, &resp)
}

// Client.UserCreateUser creates a new Pandora user.
// Argument username must be in the form of an email address. gender must be either "male" or "female".
// countryCode must be "US".
// Calls API method "user.createUser"
func (c *Client) UserCreateUser(username, password, gender, countryCode string, zipCode, birthYear int, emailOptin bool) (*response.UserCreateUser, error) {
	requestData := request.CreateUser{
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
	if err := c.Call(requestData, &resp); err != nil {
		return nil, err
	}

	// Set user data onto client for later use.
	c.userAuthToken = resp.UserAuthToken
	c.userID = resp.UserID

	return &resp, nil
}

// Client.UserEmailPassword resends registration email, maybe?
// Calls API method "user.emaillPassword"
func (c *Client) UserEmailPassword(username string) error {
	requestData := request.EmailPassword{
		Username:         username,
		PartnerAuthToken: c.partnerAuthToken,
		SyncTime:         c.GetSyncTime(),
	}

	var resp interface{}
	return c.Call(requestData, &resp)
}

// Client.UserGetBookmarks returns the users bookmarked artists and songs.
// Also see BookmarkAddArtistBookmark and BookmarkAddSongBookmark.
// Calls API method "user.getBookmarks"
func (c *Client) UserGetBookmarks() (*response.UserGetBookmarks, error) {
	requestData := request.GetBookmarks(c.Token())

	var resp response.UserGetBookmarks
	return &resp, c.Call(requestData, &resp)
}

// Client.UserGetStationList gets the list of a users stations.
// Call API method "user.getStationList"
func (c *Client) UserGetStationList(includeStationArtURL bool) (*response.UserGetStationList, error) {
	requestData := request.GetStationList{
		IncludeStationArtURL: includeStationArtURL,
		UserToken:            c.Token(),
	}

	var resp response.UserGetStationList
	return &resp, c.Call(requestData, &resp)
}

// Client.UserGetStationList returns the checksum of the user's station list.
// Call API method "user.getStationListChecksum"
func (c *Client) UserGetStationListChecksum() (*response.UserGetStationListChecksum, error) {
	requestData := request.GetStationListChecksum(c.Token())

	var resp response.UserGetStationListChecksum
	return &resp, c.Call(requestData, &resp)
}

// Client.UserSetQuickMix selects the stations that should be in the special QuickMix station.
// Call API method "user.setQuickMix"
func (c *Client) UserSetQuickMix(stationIDs []string) error {
	requestData := request.SetQuickMix{
		QuickMixStationIDs: stationIDs,
		UserToken:          c.Token(),
	}

	var resp interface{}
	return c.Call(requestData, &resp)
}

// Client.UserSleepSong marks a song to be not played again for 1 month.
// Calls API method "user.sleepSong"
func (c *Client) UserSleepSong(trackToken string) error {
	requestData := request.SleepSong{
		TrackToken: trackToken,
		UserToken:  c.Token(),
	}

	var resp interface{}
	return c.Call(requestData, &resp)
}
