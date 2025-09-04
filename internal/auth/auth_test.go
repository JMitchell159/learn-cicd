package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header1 := http.Header{}
	header1.Add("Authorization", "ApiKey Funny")
	header2 := http.Header{}
	header2.Add("Authorization", "ApiKey Funny Haha")
	header3 := http.Header{}
	header3.Add("Authorization", "ApiKey")
	header4 := http.Header{}
	header4.Add("Authorization", "Funny ApiKey")
	header5 := http.Header{}
	header5.Add("Authorization", "")
	cases := []struct {
		input          http.Header
		expectedString string
		expectedErr    error
	}{
		{
			input:          header1,
			expectedString: "Funny",
			expectedErr:    nil,
		},
		{
			input:          header2,
			expectedString: "Funny",
			expectedErr:    nil,
		},
		{
			input:          header3,
			expectedString: "",
			expectedErr:    fmt.Errorf("malformed authorization header"),
		},
		{
			input:          header4,
			expectedString: "",
			expectedErr:    fmt.Errorf("malformed authorization header"),
		},
		{
			input:          header5,
			expectedString: "",
			expectedErr:    fmt.Errorf("no authorization header included"),
		},
		{
			input:          http.Header{},
			expectedString: "",
			expectedErr:    fmt.Errorf("no authorization header included"),
		},
	}

	for _, c := range cases {
		actual, err := GetAPIKey(c.input)
		if err == nil {
			if c.expectedErr != nil {
				t.Errorf("expected %v error and was given none", c.expectedErr)
				t.Fail()
			}
		} else {
			if c.expectedErr == nil {
				t.Errorf("expected no error and was given %v", err)
				t.Fail()
			} else if err.Error() != c.expectedErr.Error() {
				t.Errorf("expected %v error and was given %v", c.expectedErr, err)
				t.Fail()
			}
		}
		if actual != c.expectedString {
			t.Errorf("expected %v, was given %v", c.expectedString, actual)
			t.Fail()
		}
	}
}
