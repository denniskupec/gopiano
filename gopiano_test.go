package gopiano

import (
	"bytes"
	"flag"
	"testing"
)

var client *Client
var pUsername, pPassword string

func init() {
	flag.StringVar(&pUsername, "username", "", "Pandora login username")
	flag.StringVar(&pPassword, "password", "", "Pandora login password")
	flag.Parse()

	client, _ = NewClient(AndroidClient)
}

func Test_Decrypt_1(t *testing.T) {
	expected := []byte("foobar")
	testString := []byte("95b6027f2d427dc0")
	decrypted, err := client.decrypt(testString)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(decrypted, expected) != 0 {
		t.Error("decrypt failed.")
	} else {
		t.Log("decrypt passed")
	}
}

func Test_AuthPartnerLogin_1(t *testing.T) {
	response, err := client.AuthPartnerLogin()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_AuthUserLogin_1(t *testing.T) {
	response, err := client.AuthUserLogin(pUsername, pPassword)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserCanSubscribe_1(t *testing.T) {
	response, err := client.UserCanSubscribe()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserBetBookmarks_1(t *testing.T) {
	response, err := client.UserGetBookmarks()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserGetStationList_1(t *testing.T) {
	response, err := client.UserGetStationList(true)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserGetStationListChecksum_1(t *testing.T) {
	response, err := client.UserGetStationListChecksum()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}
