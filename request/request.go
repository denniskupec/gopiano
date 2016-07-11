/*
Structs for use with json.Marshal when sending requests to the Pandora API.
*/
package request

import "strings"

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

type UserToken struct {
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// GetBookmarks - user.getBookmarks
type GetBookmarks UserToken

// GetStationListChecksum - user.getStationListChecksum
//
// To check if the station list was modified by another client the checksum
// can be fetched. The response contains the new checksum.
type GetStationListChecksum UserToken

// CanSubscribe - user.canSubscribe
//
// Returns whether a user is subscribed or if they can subscribe to Pandora One.
// Can be useful to determine which Partner password to use.
type CanSubscribe struct {
	UserToken
	//

	IapVendor string `json:"iapVendor,omitempty"`
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

// EmailPassword - user.emailPassword
type EmailPassword struct {
	SyncTime         int    `json:"syncTime"`
	PartnerAuthToken string `json:"partnerAuthToken"`
	//

	Username string `json:"username"`
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

// SetQuickMix - user.setQuickMix
type SetQuickMix struct {
	UserToken
	//

	QuickMixStationIDs []string `json:"quickMixStationIds"`
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

// ExplainTrack - track.explainTrack
//
// A song can be banned from all stations temporarily (one month).
type ExplainTrack trackAction

// AddArtistBookmark - bookmark.addArtistBookmark
type AddArtistBookmark trackAction

// AddSongBookmark - bookmark.addSongBookmark
type AddSongBookmark trackAction

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

// AddMusic - station.addMusic
//
// Search results can be used to add new seeds to an existing station.
type AddMusic struct {
	UserToken
	//

	StationToken string `json:"stationToken"`
	MusicToken   string `json:"musicToken"`
}

// DeleteMusic - station.deleteMusic
//
// Seeds can be removed from a station, except for the last one.
type DeleteMusic struct {
	UserToken
	//

	SeedID string `json:"seedId"`
}

// RenameStation - station.renameStation
type RenameStation struct {
	UserToken
	//
	StationToken string `json:"stationToken"`
	StationName  string `json:"stationName"`
}

// DeleteStation - station.deleteStation
type DeleteStation struct {
	UserToken
	//

	StationToken string `json:"stationToken"`
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

// DeleteFeedback - station.deleteFeedback
//
// Feedback added by Rate track can be removed from the station.
type DeleteFeedback struct {
	UserToken
	//

	FeedbackID string `json:"feedbackId"`
}

// GetGenreStations - station.getGenreStations
//
// Pandora provides a list of predefined stations (“genre stations”).
type GetGenreStations UserToken

// GetGenreStationsChecksum - station.getGenreStationsChecksum
type GetGenreStationsChecksum struct {
	UserToken
	//

	IncludeGenreCategoryAdURL bool `json:"includeGenreCategoryAdUrl,omitempty"`
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

// TransformSharedStation - station.transformSharedStation
type TransformSharedStation struct {
	UserToken
	//

	StationToken string `json:"stationToken"`
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
