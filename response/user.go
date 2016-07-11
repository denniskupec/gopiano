package response

type UserCanSubscribe struct {
	CanSubscribe bool `json:"canSubscribe"`
	IsSubscriber bool `json:"isSubscriber"`
}

type UserCreateUser AuthUserLogin

type UserGetBookmarks struct {
	Artists []ArtistBookmark `json:"artists"`
	Songs   []SongBookmark   `json:"songs"`
}

type UserGetStationList struct {
	Stations StationList `json:"stations"`
	Checksum string      `json:"checksum"`
}

type UserGetStationListChecksum struct {
	Checksum string `json:"checksum"`
}
