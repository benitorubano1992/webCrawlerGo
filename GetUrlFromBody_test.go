package main
import (
	"testing"
	"reflect"
	"net/url"
)

func TestUrlsFromHtml(t *testing.T) {
	tests := []struct {
		name          string
		inputURL     string
		inputBody       string
		expected      []string
	}{
		{
			name:"singleLinkAbs",
			inputURL:"https://crawler-test.com",
			inputBody:`<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:[]string{"https://crawler-test.com"},
		},
		{
			name:"singleLinkRelative",
			inputURL:"https://crawler-test.com",
			inputBody:`<html><body><a href="/xyxz.html"><span>Boot.dev</span></a></body></html>`,
			expected:[]string{"https://crawler-test.com/xyxz.html"},
		},
		{
			name:"no body",
			inputURL:"https://crawler-test.com",
			inputBody:`<html><h1>pippo</h1></html>`,
			expected:nil,

		},
	{name:"multipleUrl",
	inputURL:"https://crawler-test.com",
	inputBody:`<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a><a href="/xyxz.html"><span>Pippo.dev</span></a></body></html>`,
	expected:[]string{"https://crawler-test.com","https://crawler-test.com/xyxz.html"},
	},
 }
 for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			 baseURL, err := url.Parse(tc.inputURL)
    if err != nil {
        t.Errorf("couldn't parse input URL: %v", err)
        return
    }

	actual, err := getURLsFromHTML(tc.inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, tc.expected) {
		t.Errorf("expected %v, got %v", tc.expected, actual)
	}
		})
	}

}