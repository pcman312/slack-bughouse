package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func HandleTeams(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	req.ParseForm()

	names := strings.Split(req.Form["text"][0], " ")

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	teams, err := selectTeams(rng, names)
	if err != nil {
		io.WriteString(rw, "There are not enough players!")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// build json to send back to slack

	message := `{
        "response_type": "in_channel",
        "text": "*Would you like to play a game?*",
        "mrkdwn": true,
        "attachments": [
            {
                "text": "{{.}}",
                "mrkdwn_in": ["text"]
            }
            ]
        }`

	results := "*Team 1 White:* _" + teams[0].players[0] +
		"_\n*Team 1 Black*: _" + teams[0].players[1] +
		"_\n\n*Team 2 White:* _" + teams[1].players[0] +
		"_\n*Team 2 Black:* _" + teams[1].players[1] + "_"

	tmpl, err := template.New("message").Parse(message)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer

	err = tmpl.Execute(&tpl, results)
	if err != nil {
		panic(err)
	}

	io.WriteString(rw, tpl.String())
	rw.WriteHeader(http.StatusOK)
}

type team struct {
	players []string
}

func selectTeams(rng *rand.Rand, players []string) (teams []team, err error) {
	if len(players) < 4 {
		return teams, fmt.Errorf("must specify at least 4 players")
	}

	perm := rng.Perm(len(players))

	team1 := team{
		players: []string{
			players[perm[0]],
			players[perm[1]],
		},
	}
	teams = append(teams, team1)

	team2 := team{
		players: []string{
			players[perm[2]],
			players[perm[3]],
		},
	}

	teams = append(teams, team2)

	return teams, nil
}
