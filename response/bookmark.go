package response

type ArtistBookmark struct {
	ArtURL        string       `json:"artUrl"`
	ArtistName    string       `json:"artistName"`
	BookmarkToken string       `json:"bookmarkToken"`
	DateCreated   DateResponse `json:"dateCreated"`
	MusicToken    string       `json:"musicToken"`
}

type BookmarkAddArtistBookmark struct {
	ArtistBookmark
}

type SongBookmark struct {
	AlbumName     string       `json:"artistName"`
	ArtURL        string       `json:"artUrl"`
	ArtistName    string       `json:"artistName"`
	BookmarkToken string       `json:"bookmarkToken"`
	DateCreated   DateResponse `json:"dateCreated"`
	MusicToken    string       `json:"musicToken"`
	SampleGain    string       `json:"sampleGain"`
	SampleURL     string       `json:"sampleUrl"`
	SongName      string       `json:"songName"`
}

type BookmarkAddSongBookmark struct {
	SongBookmark
}
