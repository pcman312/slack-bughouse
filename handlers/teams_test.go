package handlers

import (
	"math/rand"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestSelectTeams(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	input := []string{
		"one",
		"two",
		"three",
		"four",
	}
	teams, err := selectTeams(rng, input)
	Ok(t, err)

	actualPlayers := []string{}

	for _, team := range teams {
		actualPlayers = append(actualPlayers, team.players...)
	}

	if len(actualPlayers) != len(input) {
		t.Fatalf("expected actual to have length [%d] but was [%d]", len(input), len(actualPlayers))
	}

	for _, player := range actualPlayers {
		In(t, player, input)
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\n%s:%d: unexpected error: %s\n", filepath.Base(file), line, err)
	}
}

func In(t *testing.T, str string, arr []string) {
	for _, a := range arr {
		if str == a {
			return
		}
	}
	t.Fatalf("[%s] not found in slice", str)
}
