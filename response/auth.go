package response

type AuthPartnerLogin struct {
	SyncTime         string `json:"syncTime"`
	StationSkipLimit int    `json:"stationSkipLimit"`
	PartnerAuthToken string `json:"partnerAuthToken"`
	PartnerID        string `json:"partnerId"`
	StationSkipUnit  string `json:"stationSkipUnit"`
	DeviceProperties struct {
		VideoAdRefreshInterval int `json:"videoAdRefreshInterval"`
		VideoAdUniqueInterval  int `json:"videoAdUniqueInterval"`
		AdRefreshInterval      int `json:"adRefreshInterval"`
		VideoAdStartInterval   int `json:"videoAdStartInterval"`
	} `json:"deviceProperties"`
	Urls struct {
		AutoComplete string `json:"autoComplete"`
	} `json:"urls"`
}

type AuthUserLogin struct {
	CanListen                   bool   `json:"canListen"`
	HasAudioAds                 bool   `json:"hasAudioAds"`
	IsCapped                    bool   `json:"isCapped,omitempty"`
	ListeningTimeoutAlertMsgUri string `json:"listeningTimeoutAlertMsgUri"`
	ListeningTimeoutMinutes     string `json:"listeningTimeoutMinutes"`
	MaxStationsAllowed          int    `json:"maxStationsAllowed"`
	MinimumAdRefreshInterval    int    `json:"minimumAdRefreshInterval"`
	NowPlayingURL               string `json:"nowPlayingUrl"`
	SplashScreenAdURL           string `json:"splashScreenAdUrl"`
	StationCreationAdURL        string `json:"stationCreationAdUrl"`
	UserAuthToken               string `json:"userAuthToken"`
	UserID                      string `json:"userId"`
	UserProfileURL              string `json:"userProfileUrl"`
	Username                    string `json:"username"`
	VideoAdURL                  string `json:"videoAdUrl"`
}
