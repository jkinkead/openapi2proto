package openapi2proto

import (
	"testing"
)

func TestPathMethodToName(t *testing.T) {
	tests := []struct {
		tName, path, method, want string
	}{
		{ "multiple path separators", "/one/two-three.four", "get", "GetOneTwoThreeFour" },
		{ "removes braces", "/one.{two}/[three]/(four)", "get", "GetOneTwoThreeFour" },
		{ "removes trailing .json", "/one/two.json", "get", "GetOneTwo" },
		{ "removes query", "/one/two?bad={param}", "get", "GetOneTwo" },
		{ "removes query, braces, and .json", "/one/{two}.json?bad={param}", "get", "GetOneTwo" },
	}

	for _, test := range tests {
		t.Run(test.tName, func(t *testing.T) {
			got := pathMethodToName(test.path, test.method)
			if got != test.want {
				t.Errorf("[%s] expected %q got %q", test.tName, test.want, got)
			}
		})
	}
}
