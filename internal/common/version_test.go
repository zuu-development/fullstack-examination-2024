package common

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestVersionString(t *testing.T) {
	v := Version{
		Version: "v1.2.3",
	}

	expected := "v1.2.3"
	actual := v.String()

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("Version.String() mismatch (-expected +actual):\n%s", diff)
	}
}
