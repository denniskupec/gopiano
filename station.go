package gopiano

import (
	"denniskupec.com/gopiano/request"
	"denniskupec.com/gopiano/response"
)

// Client.StationAddFeedback adds feedback (thumbs up or down, or star or ban if you prefer) to a song.
// Argument trackToken is the token identifying a track. Obtained from Client.StationGetPlaylist
// Argument isPositive is a bool which if true is a "star" and if false is a "ban".
// Calls API method "station.addFeedback"
func (c *Client) StationAddFeedback(trackToken string, isPositive bool) (*response.StationAddFeedback, error) {
	requestData := request.AddFeedback{
		TrackToken: trackToken,
		IsPositive: isPositive,
		UserToken:  c.Token(),
	}

	var resp response.StationAddFeedback
	err := c.BlowfishJSONCall(c.formatURL("http://", "station.addFeedback"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationAddMusic adds an additional music seed to an existing station.
// Argument musicToken is obtained from Client.MusicSearch
// Argument stationToken is obtained from Client.UserGetStationList
// Calls API method "station.addMusic"
func (c *Client) StationAddMusic(musicToken, stationToken string) (*response.StationAddMusic, error) {
	requestData := request.AddMusic{
		MusicToken:   musicToken,
		StationToken: stationToken,
		UserToken:    c.Token(),
	}

	var resp response.StationAddMusic
	err := c.BlowfishJSONCall(c.formatURL("http://", "station.addMusic"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationCreateStationTrack creates a new station from a specified track.
// Argument trackToken is a token of a song or artist obtained from Client.StationGetPlaylist.
// Argument musicType is either "song" or "artist" specifying the type of track being used.
// Calls API method "station.createStation"
func (c *Client) StationCreateStationTrack(trackToken, musicType string) (*response.StationCreateStation, error) {
	requestData := request.CreateStation{
		TrackToken: trackToken,
		MusicType:  musicType,
		UserToken:  c.Token(),
	}

	var resp response.StationCreateStation
	err := c.BlowfishJSONCall(c.formatURL("http://", "station.createStation"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationCreateStationMusic creates a new station from a music search result.
// Argument musicToken is obtained from Client.MusicSearch.
// Calls API method "station.createStation"
func (c *Client) StationCreateStationMusic(musicToken string) (*response.StationCreateStation, error) {
	requestData := request.CreateStation{
		MusicToken: musicToken,
		UserToken:  c.Token(),
	}

	var resp response.StationCreateStation
	err := c.BlowfishJSONCall(c.formatURL("http://", "station.createStation"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationDeleteFeedback deletes feedback (thumbs up/down) on a particular tracks feedback ID.
// Calls API method "station.deleteFeedback"
func (c *Client) StationDeleteFeedback(feedbackID string) error {
	requestData := request.DeleteFeedback{
		FeedbackID: feedbackID,
		UserToken:  c.Token(),
	}

	var resp interface{}
	return c.BlowfishJSONCall(c.formatURL("http://", "station.deleteFeedback"), requestData, &resp)
}

// Client.StationDeleteMusic removes seed music identified by a seedID from a station.
// Calls API method "station.deleteMusic"
func (c *Client) StationDeleteMusic(seedID string) error {
	requestData := request.DeleteMusic{
		SeedID:    seedID,
		UserToken: c.Token(),
	}

	var resp interface{}
	return c.BlowfishJSONCall(c.formatURL("http://", "station.deleteMusic"), requestData, &resp)
}

// Client.StationDeleteStation removes a station identified by a stationToken.
// Calls API method "station.deleteStation"
func (c *Client) StationDeleteStation(stationToken string) error {
	requestData := request.DeleteStation{
		StationToken: stationToken,
		UserToken:    c.Token(),
	}

	var resp interface{}
	return c.BlowfishJSONCall(c.formatURL("http://", "station.deleteStation"), requestData, &resp)
}

// Client.StationGetGenreStations retrieves a list of predefined "genre stations".
// Calls API method "station.getGenreStations"
func (c *Client) StationGetGenreStations() (*response.StationGetGenreStations, error) {
	requestData := request.GetGenreStations(c.Token())

	var resp response.StationGetGenreStations
	err := c.BlowfishJSONCall(c.formatURL("http://", "station.addGetGenreStations"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationGetPlaylist retrieves a playlist for a specified token.
// Argument stationToken is a obtained from User.GetStationList.
// Note: an error response with code 0 may mean you've called getPlaylist too much.
// Calls API method "station.getPlaylist"
func (c *Client) StationGetPlaylist(stationToken string) (*response.StationGetPlaylist, error) {
	requestData := request.GetPlaylist{
		StationToken: stationToken,
		UserToken:    c.Token(),
	}

	var resp response.StationGetPlaylist
	err := c.BlowfishJSONCall(c.formatURL("https://", "station.getPlaylist"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationGetStation retrieves station details.
// Argument stationToken is obtained from Client.UserGetStationList
// Argument includeExtendedAttributes will include music seed and feedback IDs in response.
// Calls API method "station.getStation"
func (c *Client) StationGetStation(stationToken string, includeExtendedAttributes bool) (*response.StationGetStation, error) {
	requestData := request.GetStation{
		StationToken:              stationToken,
		IncludeExtendedAttributes: includeExtendedAttributes,
		UserToken:                 c.Token(),
	}

	var resp response.StationGetStation
	err := c.BlowfishJSONCall(c.formatURL("http://", "station.getStation"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationShareStation shares a station with provided email addresses.
// Arguments stationID and stationToken obtained from Client.UserGetStationList
// Argument emails is a list of email addresses.
// Calls API method "station.shareStation"
func (c *Client) StationShareStation(stationID, stationToken string, emails []string) error {
	requestData := request.ShareStation{
		StationToken: stationToken,
		StationID:    stationID,
		Emails:       emails,
		UserToken:    c.Token(),
	}

	var resp interface{}
	return c.BlowfishJSONCall(c.formatURL("http://", "station.shareStation"), requestData, &resp)
}

// Client.StationRenameStation sets a new name for a station.
// Calls API method "station.renameStation"
func (c *Client) StationRenameStation(stationToken, stationName string) (*response.StationRenameStation, error) {
	requestData := request.RenameStation{
		StationToken: stationToken,
		StationName:  stationName,
		UserToken:    c.Token(),
	}

	var resp response.StationRenameStation
	err := c.BlowfishJSONCall(c.formatURL("http://", "station.renameStation"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationTransformSharedStation copies a shared station and creates a user-editable station.
// Calls API method "station.transformSharedStation"
func (c *Client) StationTransformSharedStation(stationToken string) (*response.StationTransformSharedStation, error) {
	requestData := request.TransformSharedStation{
		StationToken: stationToken,
		UserToken:    c.Token(),
	}

	var resp response.StationTransformSharedStation
	err := c.BlowfishJSONCall(c.formatURL("http://", "station.transformSharedStation"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
