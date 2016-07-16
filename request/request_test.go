package request

import "testing"

func TestStreamTypeString(t *testing.T) {
	data := []struct {
		ST  streamType
		Str string
	}{
		{0, ""},
		{HTTP_40_AAC_MONO, "HTTP_40_AAC_MONO"},
		{HTTP_24_AACPLUS_ADTS, "HTTP_24_AACPLUS_ADTS"},
		{HTTP_32_WMA, "HTTP_32_WMA"},
		{HTTP_64_AAC | HTTP_64_AACPLUS | HTTP_128_MP3, "HTTP_64_AAC,HTTP_64_AACPLUS,HTTP_128_MP3"},
		{HTTP_64_AAC | HTTP_128_MP3 | HTTP_64_AACPLUS, "HTTP_64_AAC,HTTP_64_AACPLUS,HTTP_128_MP3"},
	}

	for _, d := range data {
		if out := d.ST.String(); out != d.Str {
			t.Errorf("\nexpected:\n\t%q\ngot:\n\t%q", d.Str, out)
		}
	}
}
