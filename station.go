package gopiano

import (
	"denniskupec.com/gopiano/request"
	"denniskupec.com/gopiano/response"
)

// StationAddFeedback adds feedback (thumbs up or down, or star or ban if you prefer) to a song.
// Argument trackToken is the token identifying a track. Obtained from Client.StationGetPlaylist
// Argument isPositive is a bool which if true is a "star" and if false is a "ban".
func (c *Client) StationAddFeedback(trackToken string, isPositive bool) (*response.StationAddFeedback, error) {
	requestData := request.AddFeedback{
		TrackToken: trackToken,
		IsPositive: isPositive,
		UserToken:  c.Token(),
	}

	var resp response.StationAddFeedback
	return &resp, c.Call(requestData, &resp)
}

// StationAddMusic adds an additional music seed to an existing station.
// Argument musicToken is obtained from Client.MusicSearch
// Argument stationToken is obtained from Client.UserGetStationList
func (c *Client) StationAddMusic(musicToken, stationToken string) (*response.StationAddMusic, error) {
	requestData := request.AddMusic{
		MusicToken:   musicToken,
		StationToken: stationToken,
		UserToken:    c.Token(),
	}

	var resp response.StationAddMusic
	return &resp, c.Call(requestData, &resp)
}

// StationCreateStationTrack creates a new station from a specified track.
// Argument trackToken is a token of a song or artist obtained from Client.StationGetPlaylist.
// Argument musicType is either "song" or "artist" specifying the type of track being used.
func (c *Client) StationCreateStationTrack(trackToken, musicType string) (*response.StationCreateStation, error) {
	requestData := request.CreateStation{
		TrackToken: trackToken,
		MusicType:  musicType,
		UserToken:  c.Token(),
	}

	var resp response.StationCreateStation
	return &resp, c.Call(requestData, &resp)
}

// StationCreateStationMusic creates a new station from a music search result.
// Argument musicToken is obtained from Client.MusicSearch.
func (c *Client) StationCreateStationMusic(musicToken string) (*response.StationCreateStation, error) {
	requestData := request.CreateStation{
		MusicToken: musicToken,
		UserToken:  c.Token(),
	}

	var resp response.StationCreateStation
	return &resp, c.Call(requestData, &resp)
}

// StationDeleteFeedback deletes feedback (thumbs up/down) on a particular tracks feedback ID.
func (c *Client) StationDeleteFeedback(feedbackID string) error {
	requestData := request.DeleteFeedback{
		FeedbackID: feedbackID,
		UserToken:  c.Token(),
	}

	var resp interface{}
	return c.Call(requestData, &resp)
}

// StationDeleteMusic removes seed music identified by a seedID from a station.
func (c *Client) StationDeleteMusic(seedID string) error {
	requestData := request.DeleteMusic{
		SeedID:    seedID,
		UserToken: c.Token(),
	}

	var resp interface{}
	return c.Call(requestData, &resp)
}

// StationDeleteStation removes a station identified by a stationToken.
func (c *Client) StationDeleteStation(stationToken string) error {
	requestData := request.DeleteStation{
		StationToken: stationToken,
		UserToken:    c.Token(),
	}

	var resp interface{}
	return c.Call(requestData, &resp)
}

// StationGetGenreStations retrieves a list of predefined "genre stations".
func (c *Client) StationGetGenreStations() (*response.StationGetGenreStations, error) {
	requestData := request.GetGenreStations(c.Token())

	var resp response.StationGetGenreStations
	return &resp, c.Call(requestData, &resp)
}

// StationGetPlaylist retrieves a playlist for a specified token.
// Argument stationToken is a obtained from User.GetStationList.
// Note: an error response with code 0 may mean you've called getPlaylist too much.
func (c *Client) StationGetPlaylist(stationToken string) (*response.StationGetPlaylist, error) {
	requestData := request.GetPlaylist{
		StationToken: stationToken,
		UserToken:    c.Token(),
	}

	var resp response.StationGetPlaylist
	return &resp, c.Call(requestData, &resp)
}

// StationGetStation retrieves station details.
// Argument stationToken is obtained from Client.UserGetStationList
// Argument includeExtendedAttributes will include music seed and feedback IDs in response.
func (c *Client) StationGetStation(stationToken string, includeExtendedAttributes bool) (*response.StationGetStation, error) {
	requestData := request.GetStation{
		StationToken:              stationToken,
		IncludeExtendedAttributes: includeExtendedAttributes,
		UserToken:                 c.Token(),
	}

	var resp response.StationGetStation
	return &resp, c.Call(requestData, &resp)
}

// StationShareStation shares a station with provided email addresses.
// Arguments stationID and stationToken obtained from Client.UserGetStationList
// Argument emails is a list of email addresses.
func (c *Client) StationShareStation(stationID, stationToken string, emails []string) error {
	requestData := request.ShareStation{
		StationToken: stationToken,
		StationID:    stationID,
		Emails:       emails,
		UserToken:    c.Token(),
	}

	var resp interface{}
	return c.Call(requestData, &resp)
}

// StationRenameStation sets a new name for a station.
func (c *Client) StationRenameStation(stationToken, stationName string) (*response.StationRenameStation, error) {
	requestData := request.RenameStation{
		StationToken: stationToken,
		StationName:  stationName,
		UserToken:    c.Token(),
	}

	var resp response.StationRenameStation
	return &resp, c.Call(requestData, &resp)
}

// StationTransformSharedStation copies a shared station and creates a user-editable station.
func (c *Client) StationTransformSharedStation(stationToken string) (*response.StationTransformSharedStation, error) {
	requestData := request.TransformSharedStation{
		StationToken: stationToken,
		UserToken:    c.Token(),
	}

	var resp response.StationTransformSharedStation
	return &resp, c.Call(requestData, &resp)
}
