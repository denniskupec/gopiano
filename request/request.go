/*
Structs for use with json.Marshal when sending requests to the Pandora API.
*/
package request

import "strings"

type protocol int

func (p protocol) URL() string {
	switch p {
	case HTTP:
		return "http://"
	case HTTPS:
		return "https://"
	default:
		panic("can't get here")
	}
}

const (
	HTTP protocol = iota
	HTTPS
)

type Type interface {
	Method() string
	Protocol() protocol
}

// PartnerLogin -  auth.partnerLogin
//
// This request additionally serves as API version validation,
// time synchronization and endpoint detection and must be sent
// over a TLS-encrypted link. The POST body however is not encrypted.
type PartnerLogin struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DeviceModel string `json:"deviceModel"`
	Version     string `json:"version"`

	IncludeURLs                bool `json:"includeUrls,omitempty"`
	ReturnDeviceType           bool `json:"returnDeviceType,omitempty"`
	ReturnUpdatePromptVersions bool `json:"returnUpdatePromptVersions,omitempty"`
}

func (PartnerLogin) Method() string {
	return "auth.partnerLogin"
}

func (PartnerLogin) Protocol() protocol {
	return HTTPS
}

// UserLogin - auth.userLogin
//
// This request must be sent over a TLS-encrypted link. It authenticates
// the Pandora user by sending his username, usually his email address,
// and password as well as the partnerAuthToken obtained by Partner login.
//
// Additional response data can be requested by setting flags listed below.
type UserLogin struct {
	SyncTime         int    `json:"syncTime"`
	PartnerAuthToken string `json:"partnerAuthToken"`
	//

	LoginType string `json:"loginType"` // Should always be "user"
	Username  string `json:"username"`
	Password  string `json:"password"`

	StationArtSize string `json:"stationArtSize,omitempty"` // W130H130

	ReturnGenreStations   bool `json:"returnGenreStations,omitempty"`
	ReturnCapped          bool `json:"returnCapped,omitempty"`
	IncludePandoraOneInfo bool `json:"includePandoraOneInfo,omitempty"`
	IncludeAdAttributes   bool `json:"includeAdAttributes,omitempty"`
	ReturnStationList     bool `json:"returnStationList,omitempty"`
	IncludeStationArtURL  bool `json:"includeStationArtUrl,omitempty"`
	IncludeStationSeeds   bool `json:"includeStationSeeds,omitempty"`

	IncludeShuffleInsteadOfQuickMix bool `json:"includeShuffleInsteadOfQuickMix,omitempty"`
	ReturnCollectTrackLifetimeStats bool `json:"returnCollectTrackLifetimeStats,omitempty"`

	ReturnIsSubscriber bool `json:"returnIsSubscriber,omitempty"`
	XPlatformAdCapable bool `json:"xplatformAdCapable,omitempty"`

	ComplimentarySponsorSupported bool `json:"complimentarySponsorSupported,omitempty"`
	IncludeSubscriptionExpiration bool `json:"includeSubscriptionExpiration,omitempty"`

	ReturnHasUsedTrial    bool `json:"returnHasUsedTrial,omitempty"`
	ReturnUserstate       bool `json:"returnUserstate,omitempty"`
	IncludeAccountMessage bool `json:"includeAccountMessage,omitempty"`
	IncludeUserWebname    bool `json:"includeUserWebname,omitempty"`
	IncludeListeningHours bool `json:"includeListeningHours,omitempty"`

	IncludeFacebook   bool `json:"includeFacebook,omitempty"`
	IncludeTwitter    bool `json:"includeTwitter,omitempty"`
	IncludeGoogleplay bool `json:"includeGoogleplay,omitempty"`

	IncludeDailySkipLimit bool `json:"includeDailySkipLimit,omitempty"`
	IncludeSkipDelay      bool `json:"includeSkipDelay,omitempty"`

	IncludeShowUserRecommendations bool `json:"includeShowUserRecommendations,omitempty"`
	IncludeAdvertiserAttributes    bool `json:"includeAdvertiserAttributes,omitempty"`
}

func (UserLogin) Method() string {
	return "auth.userLogin"
}

func (UserLogin) Protocol() protocol {
	return HTTPS
}

type UserToken struct {
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// GetBookmarks - user.getBookmarks
type GetBookmarks UserToken

func (GetBookmarks) Method() string {
	return "user.getBookmarks"
}

func (GetBookmarks) Protocol() protocol {
	return HTTP
}

// GetStationListChecksum - user.getStationListChecksum
//
// To check if the station list was modified by another client the checksum
// can be fetched. The response contains the new checksum.
type GetStationListChecksum UserToken

func (GetStationListChecksum) Method() string {
	return "user.getStationListChecksum"
}

func (GetStationListChecksum) Protocol() protocol {
	return HTTP
}

// CanSubscribe - user.canSubscribe
//
// Returns whether a user is subscribed or if they can subscribe to Pandora One.
// Can be useful to determine which Partner password to use.
type CanSubscribe struct {
	UserToken
	//

	IapVendor string `json:"iapVendor,omitempty"`
}

func (CanSubscribe) Method() string {
	return "user.canSubscribe"
}

func (CanSubscribe) Protocol() protocol {
	return HTTP
}

// CreateUser - user.createUser
type CreateUser struct {
	SyncTime         int    `json:"syncTime"`
	PartnerAuthToken string `json:"partnerAuthToken"`
	//

	Username       string `json:"username"`
	Password       string `json:"password"`
	Gender         string `json:"gender"`
	BirthYear      int    `json:"birthYear"`
	ZipCode        int    `json:"zip"`
	EmailOptin     bool   `json:"emailOptin"`
	CountryCode    string `json:"countryCode"`
	AccountType    string `json:"accountType"`    // registered
	RegisteredType string `json:"registeredType"` // user

	IncludePandoraOneInfo           bool `json:"includePandoraOneInfo,omitempty"`
	IncludeAccountMessage           bool `json:"includeAccountMessage,omitempty"`
	ReturnCollectTrackLifetimeStats bool `json:"returnCollectTrackLifetimeStats,omitempty"`
	ReturnIsSubscriber              bool `json:"returnIsSubscriber,omitempty"`
	XplatformAdCapable              bool `json:"xplatformAdCapable,omitempty"`
	IncludeFacebook                 bool `json:"includeFacebook,omitempty"`
	IncludeGoogleplay               bool `json:"includeGoogleplay,omitempty"`
	IncludeShowUserRecommendations  bool `json:"includeShowUserRecommendations,omitempty"`
	IncludeAdvertiserAttributes     bool `json:"includeAdvertiserAttributes,omitempty"`
}

func (CreateUser) Method() string {
	return "user.createUser"
}

func (CreateUser) Protocol() protocol {
	return HTTPS
}

// EmailPassword - user.emailPassword
type EmailPassword struct {
	SyncTime         int    `json:"syncTime"`
	PartnerAuthToken string `json:"partnerAuthToken"`
	//

	Username string `json:"username"`
}

func (EmailPassword) Method() string {
	return "user.emailPassword"
}

func (EmailPassword) Protocol() protocol {
	return HTTPS
}

// GetStationList - user.getStationList
type GetStationList struct {
	UserToken
	//

	StationArtSize string `json:"stationArtSize,omitempty"` // W130H130

	IncludeStationArtURL            bool `json:"includeStationArtUrl,omitempty"`
	IncludeAdAttributes             bool `json:"includeAdAttributes,omitempty"`
	IncludeStationSeeds             bool `json:"includeStationSeeds,omitempty"`
	IncludeShuffleInsteadOfQuickMix bool `json:"includeShuffleInsteadOfQuickMix,omitempty"`
	IncludeRecommendations          bool `json:"includeRecommendations,omitempty"`
	IncludeExplanations             bool `json:"includeExplanations,omitempty"`
}

func (GetStationList) Method() string {
	return "user.getStationList"
}

func (GetStationList) Protocol() protocol {
	return HTTP
}

// SetQuickMix - user.setQuickMix
type SetQuickMix struct {
	UserToken
	//

	QuickMixStationIDs []string `json:"quickMixStationIds"`
}

func (SetQuickMix) Method() string {
	return "user.setQuickMix"
}

func (SetQuickMix) Protocol() protocol {
	return HTTP
}

type trackAction struct {
	UserToken
	//

	TrackToken string `json:"trackToken"`
}

// SleepSong - user.sleepSong
//
// A song can be banned from all stations temporarily (one month).
type SleepSong trackAction

func (SleepSong) Method() string {
	return "user.sleepSong"
}

func (SleepSong) Protocol() protocol {
	return HTTP
}

// ExplainTrack - track.explainTrack
//
// A song can be banned from all stations temporarily (one month).
type ExplainTrack trackAction

func (ExplainTrack) Method() string {
	return "track.explainTrack"
}

func (ExplainTrack) Protocol() protocol {
	return HTTP
}

// AddArtistBookmark - bookmark.addArtistBookmark
type AddArtistBookmark trackAction

func (AddArtistBookmark) Method() string {
	return "bookmark.addArtistBookmark"
}

func (AddArtistBookmark) Protocol() protocol {
	return HTTP
}

// AddSongBookmark - bookmark.addSongBookmark
type AddSongBookmark trackAction

func (AddSongBookmark) Method() string {
	return "bookmark.addSongBookmark"
}

func (AddSongBookmark) Protocol() protocol {
	return HTTP
}

// MusicSearch - music.search
//
// This is a free text search that matches artist and track names.
type MusicSearch struct {
	UserToken
	//

	SearchText           string `json:"searchText"`
	IncludeNearMatches   bool   `json:"includeNearMatches,omitempty"`
	IncludeGenreStations bool   `json:"includeGenreStations,omitempty"`
}

func (MusicSearch) Method() string {
	return "music.search"
}

func (MusicSearch) Protocol() protocol {
	return HTTP
}

// CreateStation - station.createStation
//
// Stations can either be created with a musicToken obtained by Search
// or trackToken from playlists. The latter needs a musicType to specify
// whether the track itself or its artist should be used as seed.
type CreateStation struct {
	UserToken
	//

	TrackToken string `json:"trackToken,omitempty"`
	MusicType  string `json:"musicType,omitempty"` // 'song' or 'artist' ('song' for genre stations)
	MusicToken string `json:"musicToken,omitempty"`
}

func (CreateStation) Method() string {
	return "station.createStation"
}

func (CreateStation) Protocol() protocol {
	return HTTP
}

// AddMusic - station.addMusic
//
// Search results can be used to add new seeds to an existing station.
type AddMusic struct {
	UserToken
	//

	StationToken string `json:"stationToken"`
	MusicToken   string `json:"musicToken"`
}

func (AddMusic) Method() string {
	return "station.addMusic"
}

func (AddMusic) Protocol() protocol {
	return HTTP
}

// DeleteMusic - station.deleteMusic
//
// Seeds can be removed from a station, except for the last one.
type DeleteMusic struct {
	UserToken
	//

	SeedID string `json:"seedId"`
}

func (DeleteMusic) Method() string {
	return "station.deleteMusic"
}

func (DeleteMusic) Protocol() protocol {
	return HTTP
}

// RenameStation - station.renameStation
type RenameStation struct {
	UserToken
	//
	StationToken string `json:"stationToken"`
	StationName  string `json:"stationName"`
}

func (RenameStation) Method() string {
	return "station.renameStation"
}

func (RenameStation) Protocol() protocol {
	return HTTP
}

// DeleteStation - station.deleteStation
type DeleteStation struct {
	UserToken
	//

	StationToken string `json:"stationToken"`
}

func (DeleteStation) Method() string {
	return "station.deleteStation"
}

func (DeleteStation) Protocol() protocol {
	return HTTP
}

// GetStation - station.getStation
//
// Extended station information includes seeds and feedback.
type GetStation struct {
	UserToken
	//

	StationToken              string `json:"stationToken"`
	IncludeExtendedAttributes bool   `json:"includeExtendedAttributes,omitempty"`
}

func (GetStation) Method() string {
	return "station.getStation"
}

func (GetStation) Protocol() protocol {
	return HTTP
}

// AddFeedback - station.addFeedback
//
// Songs can be “loved” or “banned”. Both influence the music played
// on the station. Banned songs are never played again on this particular station.
type AddFeedback struct {
	UserToken
	//

	StationToken string `json:"stationToken"`
	TrackToken   string `json:"trackToken"`
	IsPositive   bool   `json:"isPositive"` // 'false' bans track
}

func (AddFeedback) Method() string {
	return "station.addFeedback"
}

func (AddFeedback) Protocol() protocol {
	return HTTP
}

// DeleteFeedback - station.deleteFeedback
//
// Feedback added by Rate track can be removed from the station.
type DeleteFeedback struct {
	UserToken
	//

	FeedbackID string `json:"feedbackId"`
}

func (DeleteFeedback) Method() string {
	return "station.deleteFeedback"
}

func (DeleteFeedback) Protocol() protocol {
	return HTTP
}

// GetGenreStations - station.getGenreStations
//
// Pandora provides a list of predefined stations (“genre stations”).
type GetGenreStations UserToken

func (GetGenreStations) Method() string {
	return "station.getGenreStations"
}

func (GetGenreStations) Protocol() protocol {
	return HTTP
}

// GetGenreStationsChecksum - station.getGenreStationsChecksum
type GetGenreStationsChecksum struct {
	UserToken
	//

	IncludeGenreCategoryAdURL bool `json:"includeGenreCategoryAdUrl,omitempty"`
}

func (GetGenreStationsChecksum) Method() string {
	return "station.getGenreStationsChecksum"
}

func (GetGenreStationsChecksum) Protocol() protocol {
	return HTTP
}

// ShareStation - station.shareStation
//
// Shares a station with the specified email addresses.
type ShareStation struct {
	UserToken
	//

	StationID    string   `json:"stationId"`
	StationToken string   `json:"stationToken"`
	Emails       []string `json:"emails"`
}

func (ShareStation) Method() string {
	return "station.shareStation"
}

func (ShareStation) Protocol() protocol {
	return HTTP
}

// TransformSharedStation - station.transformSharedStation
type TransformSharedStation struct {
	UserToken
	//

	StationToken string `json:"stationToken"`
}

func (TransformSharedStation) Method() string {
	return "station.transformSharedStation"
}

func (TransformSharedStation) Protocol() protocol {
	return HTTP
}

// GetPlaylist - station.getPlaylist
//
// This method must be sent over a TLS-encrypted connection.
type GetPlaylist struct {
	UserToken
	//

	StationToken string `json:"stationToken"`

	AdditionalAudioURL string `json:"additionalAudioUrl,omitempty"`

	StationIsStarting      bool `json:"stationIsStarting,omitempty"`
	IncludeTrackLength     bool `json:"includeTrackLength,omitempty"`
	IncludeAudioToken      bool `json:"includeAudioToken,omitempty"`
	XPlatformAdCapable     bool `json:"xplatformAdCapable,omitempty"`
	IncludeAudioReceiptURL bool `json:"includeAudioReceiptUrl,omitempty"`
	IncludeBackstageAdURL  bool `json:"includeBackstageAdUrl,omitempty"`
	IncludeSharingAdURL    bool `json:"includeSharingAdUrl,omitempty"`
	IncludeSocialAdURL     bool `json:"includeSocialAdUrl,omitempty"`

	IncludeCompetitiveSepIndicator bool `json:"includeCompetitiveSepIndicator,omitempty"`

	IncludeCompletePlaylist bool `json:"includeCompletePlaylist,omitempty"`
	IncludeTrackOptions     bool `json:"includeTrackOptions,omitempty"`
	AudioAdPodCapable       bool `json:"audioAdPodCapable,omitempty"`
}

func (GetPlaylist) Method() string {
	return "station.getPlaylist"
}

func (GetPlaylist) Protocol() protocol {
	return HTTPS
}

type streamType int

const (
	HTTP_40_AAC_MONO streamType = 1 << iota
	HTTP_64_AAC
	HTTP_32_AACPLUS
	HTTP_64_AACPLUS
	HTTP_24_AACPLUS_ADTS
	HTTP_32_AACPLUS_ADTS
	HTTP_64_AACPLUS_ADTS
	HTTP_128_MP3
	HTTP_32_WMA
)

var streamTypeNames = [...]string{
	"HTTP_40_AAC_MONO",
	"HTTP_64_AAC",
	"HTTP_32_AACPLUS",
	"HTTP_64_AACPLUS",
	"HTTP_24_AACPLUS_ADTS",
	"HTTP_32_AACPLUS_ADTS",
	"HTTP_64_AACPLUS_ADTS",
	"HTTP_128_MP3",
	"HTTP_32_WMA",
}

func (st streamType) String() string {
	stn := streamTypeNames
	if st&HTTP_40_AAC_MONO == 0 {
		stn[0] = ""
	}
	if st&HTTP_64_AAC == 0 {
		stn[1] = ""
	}
	if st&HTTP_32_AACPLUS == 0 {
		stn[2] = ""
	}
	if st&HTTP_64_AACPLUS == 0 {
		stn[3] = ""
	}
	if st&HTTP_24_AACPLUS_ADTS == 0 {
		stn[4] = ""
	}
	if st&HTTP_32_AACPLUS_ADTS == 0 {
		stn[5] = ""
	}
	if st&HTTP_64_AACPLUS_ADTS == 0 {
		stn[6] = ""
	}
	if st&HTTP_128_MP3 == 0 {
		stn[7] = ""
	}
	if st&HTTP_32_WMA == 0 {
		stn[8] = ""
	}
	return strings.Join(stn[:], ",")
}
