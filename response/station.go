package response

type Station struct {
	SuppressVideoAds bool         `json:"suppressVideoAds"`
	StationID        string       `json:"stationId"`
	AllowAddMusic    bool         `json:"allowAddMusic"`
	DateCreated      DateResponse `json:"dateCreated"`
	StationDetailURL string       `json:"stationDetailUrl"`
	ArtURL           string       `json:"artUrl"`
	RequiresCleanAds bool         `json:"requiresCleanAds"`
	StationToken     string       `json:"stationToken"`
	StationName      string       `json:"stationName"`
	Music            struct {
		Songs []struct {
			SeedID      string       `json:"seedId"`
			ArtistName  string       `json:"artistName"`
			SongName    string       `json:"songName"`
			DateCreated DateResponse `json:"dateCreated"`
		} `json:"songs"`
		Artists []struct {
			SeedID      string       `json:"seedId"`
			ArtistName  string       `json:"artistName"`
			DateCreated DateResponse `json:"dateCreated"`
		} `json:"artists"`
	} `json:"music"`
	IsShared           bool     `json:"isShared"`
	AllowDelete        bool     `json:"allowDelete"`
	Genre              []string `json:"genre"`
	IsQuickMix         bool     `json:"isQuickMix"`
	AllowRename        bool     `json:"allowRename"`
	StationSharingURL  string   `json:"stationSharingUrl"`
	QuickMixStationIDs []string `json:"quickMixStationIds"`
	Feedback           struct {
		ThumbsDown []FeedbackResponse `json:"thumbsDown"`
		ThumbsUp   []FeedbackResponse `json:"thumbsUp"`
	} `json:"feedback"`
}

type StationList []Station

// Make Station implement sort.Interface
func (s StationList) Len() int {
	return len(s)
}

func (s StationList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StationList) Less(i, j int) bool {
	return s[i].StationName < s[j].StationName
}

type StationAddFeedback struct {
	Result FeedbackResponse `json:"result"`
}

type StationAddMusic struct {
	ArtistName  string       `json:"artistName"`
	DateCreated DateResponse `json:"dateCreated"`
	SeedID      string       `json:"seedId"`
}

type StationResponse struct {
	Result Station `json:"result"`
}
type StationCreateStation StationResponse
type StationGetStation StationResponse
type StationRenameStation StationResponse
type StationTransformSharedStation StationResponse

type StationGetGenreStations struct {
	Categories []struct {
		CategoryName string `json:"categoryName"`
		Stations     []struct {
			StationToken string `json:"stationToken"`
			StationName  string `json:"stationName"`
			StationID    string `json:"stationId"`
		}
	} `json:"categories"`
}

type StationGetGenreStationsChecksum struct {
	Checksum string `json:"checksum"`
}

type StationGetPlaylist struct {
	Items []struct {
		TrackToken      string `json:"trackToken"`
		ArtistName      string `json:"artistName"`
		AlbumName       string `json:"albumName"`
		AmazonAlbumURL  string `json:"amazonAlbumUrl"`
		SongExplorerURL string `json:"songExplorerUrl"`
		AlbumArtURL     string `json:"albumArtUrl"`
		ArtistDetailURL string `json:"artistDetailUrl"`
		AudioURLMap     map[string]struct {
			Bitrate  string `json:"bitrate"`
			Encoding string `json:"encoding"`
			AudioURL string `json:"audioUrl"`
			Protocol string `json:"protocol"`
		} `json:"audioUrlMap"`
		ITunesSongURL          string `json:"itunesSongUrl"`
		AmazonAlbumAsin        string `json:"amazonAlbumAsin"`
		AmazonAlbumDigitalAsin string `json:"amazonAlbumDigitalAsin"`
		ArtistExplorerURL      string `json:"artistExplorerUrl"`
		SongName               string `json:"songName"`
		AlbumDetailURL         string `json:"albumDetailUrl"`
		SongDetailURL          string `json:"songDetailUrl"`
		StationID              string `json:"stationId"`
		SongRating             int    `json:"songRating"`
		TrackGain              string `json:"trackGain"`
		AlbumExplorerURL       string `json:"albumExplorerUrl"`
		AllowFeedback          bool   `json:"allowFeedback"`
		AmazonSongDigitalAsin  string `json:"amazonSongDigitalAsin"`
		NowPlayingStationAdURL string `json:"nowPlayingStationAdUrl"`
		AdToken                string `json:"adToken"`
	} `json:"items"`
}
