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
	return &resp, c.Call(requestData, &resp)
}

// Client.MusicSearch searches for music, which can be used to create a new or add seeds to a station.
// Calls API method "music.search"
func (c *Client) MusicSearch(searchText string) (*response.MusicSearch, error) {
	requestData := request.MusicSearch{
		SearchText: searchText,
		UserToken:  c.Token(),
	}

	var resp response.MusicSearch
	return &resp, c.Call(requestData, &resp)
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
	return &resp, c.Call(requestData, &resp)
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
	return &resp, c.Call(requestData, &resp)
}
