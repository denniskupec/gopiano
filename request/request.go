/*
Structs for use with json.Marshal when sending requests to the Pandora API.
*/
package request

type AuthPartnerLogin struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DeviceModel string `json:"deviceModel"`
	Version     string `json:"version"`
	IncludeURLs bool   `json:"includeUrls,omitempty"`
}

type AuthUserLogin struct {
	PartnerAuthToken              string `json:"partnerAuthToken"`
	Username                      string `json:"username"`
	Password                      string `json:"password"`
	LoginType                     string `json:"loginType"` // Should always be "user"
	SyncTime                      int    `json:"syncTime"`
	IncludeAdAttributes           bool   `json:"includeAdAttributes,omitempty"`
	IncludeDemographics           bool   `json:"IncludeDemographics,omitempty"`
	IncludePandoraOneInfo         bool   `json:"includePandoraOneInfo,omitempty"` // Appears to do nothing.
	IncludeStationArtURL          bool   `json:"includeStationArtUrl,omitempty"`
	IncludeSubscriptionExpiration bool   `json:"includeSubscriptionExpiration,omitempty"`
	ReturnCapped                  bool   `json:"returnCapped,omitempty"`
	ReturnGenreStations           bool   `json:"returnGenreStations,omitempty"`
	ReturnStationList             bool   `json:"returnStationList,omitempty"`
}

type UserToken struct {
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type UserGetBookmarks UserToken
type UserGetStationListChecksum UserToken
type UserCanSubscribe UserToken

type UserCreateUser struct {
	AccountType      string `json:"accountType"`
	BirthYear        int    `json:"birthYear"`
	CountryCode      string `json:"countryCode"`
	EmailOptin       bool   `json:"emailOptin"`
	Gender           string `json:"gender"`
	PartnerAuthToken string `json:"partnerAuthToken"`
	Password         string `json:"password"`
	RegisteredType   string `json:"registeredType"`
	SyncTime         int    `json:"syncTime"`
	Username         string `json:"username"`
	ZipCode          int    `json:"zip"`
}

type UserEmailPassword struct {
	PartnerAuthToken string `json:"partnerAuthToken"`
	SyncTime         int    `json:"syncTime"`
	Username         string `json:"username"`
}

type UserGetStationList struct {
	IncludeStationArtURL bool `json:"includeStationArtUrl,omitempty"`
	UserToken
}

type UserSetQuickMix struct {
	QuickMixStationIDs []string `json:"quickMixStationIds"`
	UserToken
}

type trackAction struct {
	TrackToken string `json:"trackToken"`
	UserToken
}
type UserSleepSong trackAction
type BookmarkAddArtistBookmark trackAction
type BookmarkAddSongBookmark trackAction

type MusicSearch struct {
	SearchText string `json:"searchText"`
	UserToken
}

type StationCreateStation struct {
	MusicToken string `json:"musicToken,omitempty"`
	TrackToken string `json:"trackToken,omitempty"`
	MusicType  string `json:"musicType,omitempty"`
	UserToken
}

type StationDeleteStation struct {
	StationToken string `json:"stationToken"`
	UserToken
}

type StationAddFeedback struct {
	TrackToken string `json:"trackToken"`
	IsPositive bool   `json:"isPositive"`
	UserToken
}

type StationDeleteFeedback struct {
	FeedbackID string `json:"feedbackId"`
	UserToken
}

type StationAddMusic struct {
	MusicToken   string `json:"musicToken"`
	StationToken string `json:"stationToken"`
	UserToken
}

type StationDeleteMusic struct {
	SeedID string `json:"seedId"`
	UserToken
}

type StationGetGenreStations UserToken
type StationGetGenreStationsChecksum UserToken

type StationGetPlaylist struct {
	StationToken string `json:"stationToken"`
	UserToken
}

type StationGetStation struct {
	StationToken              string `json:"stationToken"`
	IncludeExtendedAttributes bool   `json:"includeExtendedAttributes,omitempty"`
	UserToken
}

type StationShareStation struct {
	StationID    string   `json:"stationId"`
	StationToken string   `json:"stationToken"`
	Emails       []string `json:"emails"`
	UserToken
}

type StationRenameStation struct {
	StationToken string `json:"stationToken"`
	StationName  string `json:"stationName"`
	UserToken
}

type StationTransformSharedStation struct {
	StationToken string `json:"stationToken"`
	UserToken
}

type ExplainTrack struct {
	TrackToken string `json:"trackToken"`
	UserToken
}
