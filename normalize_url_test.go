package main
import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
		},
		{
			name:"remove suffix",
			inputURL:"www.boot.dev/blog/path/",
			expected:"www.boot.dev/blog/path",
		},
		{
			name:"remove Prefix suffix",
			inputURL:"http://www.boot.dev/blog/path/",
			expected:"www.boot.dev/blog/path",
		},
		{
			name:"Valid Url",
			inputURL:"www.boot.dev/blog/path",
			expected:"www.boot.dev/blog/path",
		},
        // add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}


func TestGetHeadingFromHTMLBasic(t *testing.T) {
	tests := []struct {
		name          string
		inputBody      string
		expected      string
	}{
		{
			name:"Valid h1",
			inputBody:"<html><body><h1>Test Title</h1></body></html>",
			expected:"Test Title",
		},
		{
			name:"Valid h2",
			inputBody:"<html><body><h2>Test Title</h2></body></html>",
			expected:"Test Title",
		},
		{
			name:"Valid h1,h2",
			inputBody:"<html><body><h1>CiaoPippo</h1><h2>Test Title</h2></body></html>",
			expected:"CiaoPippo",

		},
	
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual:= getHeadingFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name          string
		inputBody      string
		expected      string
	}{
		{
			name:"Valid p Main",
			inputBody: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
			expected:"Main paragraph.",
		},
		{
			name:"Multiple p Main",
			inputBody:`<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Ciao</p>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
			expected:"Ciao",
		},
		{
			name:"Outside p",
			inputBody:"<html><body><h1>CiaoPippo</h1><p>Prova</p></body></html>",
			expected:"Prova",

		},
		{
			name:"No Paragragh",
			inputBody:"<html><body><h1>CiaoPippo</h1></body></html>",
			expected:"",
		},
	
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual:= getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}



