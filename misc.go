package gopiano

import (
	"denniskupec.com/gopiano/requests"
	"denniskupec.com/gopiano/responses"
)

// Client.ExplainTrack retrieves an incomplete list of attributes assigned specified son by the
// Music Genome Project
// Calls API method "track.explainTrack"
func (c *Client) ExplainTrack(trackToken string) (*responses.ExplainTrack, error) {
	requestData := requests.ExplainTrack{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	var resp responses.ExplainTrack
	err := c.BlowfishJSONCall(c.formatURL("http://", "track.explainTrack"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.MusicSearch searches for music, which can be used to create a new or add seeds to a station.
// Calls API method "music.search"
func (c *Client) MusicSearch(searchText string) (*responses.MusicSearch, error) {
	requestData := requests.MusicSearch{
		SearchText:    searchText,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	var resp responses.MusicSearch
	err := c.BlowfishJSONCall(c.formatURL("http://", "music.search"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.AddArtistBookmark bookmarks an artist.
// Argument trackToken is a token of a specific artist.
// Calls API method "bookmark.addArtistBookmark"
func (c *Client) BookmarkAddArtistBookmark(trackToken string) (*responses.BookmarkAddArtistBookmark, error) {
	requestData := requests.BookmarkAddArtistBookmark{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	var resp responses.BookmarkAddArtistBookmark
	err := c.BlowfishJSONCall(c.formatURL("http://", "bookmark.addArtistBookmark"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.BookmarkAddSongBookmark bookmarks a song.
// Argument trackToken is a token of a specific song.
// Calls API method "bookmark.addSongBookmark"
func (c *Client) BookmarkAddSongBookmark(trackToken string) (*responses.BookmarkAddSongBookmark, error) {
	requestData := requests.BookmarkAddSongBookmark{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	var resp responses.BookmarkAddSongBookmark
	err := c.BlowfishJSONCall(c.formatURL("http://", "bookmark.addSongBookmark"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
