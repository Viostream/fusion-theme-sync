package theme

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/FusionAuth/go-client/pkg/fusionauth"
)

var testTheme fusionauth.Theme

func init() {
	data, err := ioutil.ReadFile("../fixtures/fusionauth_theme.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &testTheme)
	if err != nil {
		log.Fatal(err)
	}
}

func TestLoadFromDisk(t *testing.T) {
	theme, err := LoadFromDisk("../fixtures/disk", map[string]string{})
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(theme.Templates.Samlv2Logout, "[#ftl/]\n[") {
		t.Errorf("Expected %v, got %v", "[#ftl/]\n[", theme.Templates.Samlv2Logout)
	}
}

func TestWriteToDisk(t *testing.T) {
	err := WriteToDisk(&testTheme, "test", map[string]string{})
	if err != nil {
		t.Error(err)
	}
}

func TestLoadSubstitutionsFromDisk(t *testing.T) {
	path := "../fixtures/substitutions.txt"
	subs, err := LoadSubstitutionsFromDisk(path)
	if err != nil {
		t.Fatal(err)
	}

	if e := "env1"; e != subs["complete-registration"] {
		t.Errorf("Expected %v, got %v", e, subs["complete-registration"])
	}

	if e := "env2"; e != subs["configure"] {
		t.Errorf("Expected %v, got %v", e, subs["configure"])
	}

	if e := "env3"; e != subs["configured"] {
		t.Errorf("Expected %v, got %v", e, subs["configured"])
	}
}

func TestInsertSubstitutions(t *testing.T) {
	raw := "blah=keep this\notherblah=\notherblahbut=keep this\nhahahhaa="
	out := insertSubstitutions(raw, map[string]string{
		"otherblah": "this is new",
		"hahahhaa":  "this is also new",
	})

	if e := "blah=keep this\notherblah=this is new\notherblahbut=keep this\nhahahhaa=this is also new"; e != out {
		t.Errorf("Expected %v, got %v", e, out)
	}
}

func TestRemoveSubstitutions(t *testing.T) {
	raw := "blah=keep this\notherblah=remove this\notherblahbut=keep this\nhahahhaa=remove this"
	out := removeSubstitutions(raw, map[string]string{"otherblah": "asd", "hahahhaa": "asd"})

	if e := "blah=keep this\notherblah=\notherblahbut=keep this\nhahahhaa="; e != out {
		t.Errorf("Expected %v, got %v", e, out)
	}
}
