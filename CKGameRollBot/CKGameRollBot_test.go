package ckgamerollbot_test

import (
	"testing"

	ckgamerollbot "github.com/TauTheProgrammer/CKGameRollBot/CKGameRollBot"
)

func TestCKGameRollBot(t *testing.T) {
	if ckgamerollbot.ABC() != "asdf" {
		t.Fatal("Unexpected output")
	}
}
