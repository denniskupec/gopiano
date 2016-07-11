/*
Structs used with json.Unmarshal in processing responses from the Pandora API.
*/
package response

import (
	"encoding/json"
	"time"
)

// DateResponse is used repeatedly in places where Pandora returns a JSON object
// called dateCreated.
// Most of the data is rubish without a little processing but you can use GetDate()
// and also Time is just a nice UNIX epoch.
type DateResponse struct {
	Nanos          int `json:"nano"`
	Seconds        int `json:"seconds"`
	Year           int `json:"year"`
	Month          int `json:"month"`
	Hours          int `json:"hours"`
	Time           int `json:"time"`
	Date           int `json:"date"`
	Minutes        int `json:"minutes"`
	Day            int `json:"day"`
	TimezoneOffset int `json:"timezoneOffset"`
}

// Get this mess of ints as a time.Time object. Much nicer.
func (d DateResponse) GetDate() time.Time {
	return time.Date(1900+d.Year, time.Month(d.Month), d.Date, d.Hours, d.Minutes, d.Seconds,
		d.Nanos, time.FixedZone("Local Time", d.TimezoneOffset*60))
}

type MusicSearch struct {
	NearMatchesAvailable bool   `json:"nearMatchesAvailable"`
	Explanation          string `json:"explanation"`
	Songs                []struct {
		ArtistName string `json:"artistName"`
		MusicToken string `json:"musicToken"`
		SongName   string `json:"songName"`
		Score      int    `json:"score"`
	} `json:"songs"`
	Artists []struct {
		ArtistName  string `json:"artistName"`
		MusicToken  string `json:"musicToken"`
		LikelyMatch bool   `json:"likelyMatch"`
		Score       int    `json:"score"`
	} `json:"artists"`
}

type FeedbackResponse struct {
	ArtistName  string       `json:"artistName"`
	SongName    string       `json:"songName"`
	DateCreated DateResponse `json:"dateCreated"`
	FeedbackID  string       `json:"feedbackId"`
	IsPositive  bool         `json:"isPositive"`
}

type ExplainTrack struct {
	Explanations []struct {
		FocustTraitName string `json:"focusTraitName"`
		FocusTraitID    string `json:"focustTraitId"`
	} `json:"explanations"`
}

type Wrapper struct {
	ErrorResponse
	Result json.RawMessage `json:"result"`
}
