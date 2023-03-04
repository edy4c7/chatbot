package main

import (
	"github.com/edy4c7/go-discord-bot/internal/sessions"
)

func main() {
	if err := sessions.NewSession(sessions.Discord); err != nil {
		panic(err)
	}
}
