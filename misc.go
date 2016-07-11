package gopiano

import (
	"denniskupec.com/gopiano/request"
	"denniskupec.com/gopiano/response"
)

// Client.ExplainTrack retrieves an incomplete list of attributes assigned specified son by the
// Music Genome Project
// Calls API method "track.explainTrack"
func (c *Client) ExplainTrack(trackToken string) (*response.ExplainTrack, error) {
	requestData := request.ExplainTrack{
		TrackToken: trackToken,
		UserToken:  c.Token(),
	}

	var resp response.ExplainTrack
	err := c.BlowfishJSONCall(c.formatURL("http://", "track.explainTrack"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.MusicSearch searches for music, which can be used to create a new or add seeds to a station.
// Calls API method "music.search"
func (c *Client) MusicSearch(searchText string) (*response.MusicSearch, error) {
	requestData := request.MusicSearch{
		SearchText: searchText,
		UserToken:  c.Token(),
	}

	var resp response.MusicSearch
	err := c.BlowfishJSONCall(c.formatURL("http://", "music.search"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.AddArtistBookmark bookmarks an artist.
// Argument trackToken is a token of a specific artist.
// Calls API method "bookmark.addArtistBookmark"
func (c *Client) BookmarkAddArtistBookmark(trackToken string) (*response.BookmarkAddArtistBookmark, error) {
	requestData := request.AddArtistBookmark{
		TrackToken: trackToken,
		UserToken:  c.Token(),
	}

	var resp response.BookmarkAddArtistBookmark
	err := c.BlowfishJSONCall(c.formatURL("http://", "bookmark.addArtistBookmark"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.BookmarkAddSongBookmark bookmarks a song.
// Argument trackToken is a token of a specific song.
// Calls API method "bookmark.addSongBookmark"
func (c *Client) BookmarkAddSongBookmark(trackToken string) (*response.BookmarkAddSongBookmark, error) {
	requestData := request.AddSongBookmark{
		TrackToken: trackToken,
		UserToken:  c.Token(),
	}

	var resp response.BookmarkAddSongBookmark
	err := c.BlowfishJSONCall(c.formatURL("http://", "bookmark.addSongBookmark"), requestData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
