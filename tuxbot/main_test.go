// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package main

import (
	"strings"
	"testing"
)

func TestBuildLinux(t *testing.T) {
	bot := &tuxbot{}
	out, err := bot.Run("", []string{"build", "linux"})
	if err != nil {
		t.Fatal(err)
	}
	if out != "Dry Run: Doing that would run `prerelease.sh`" {
		t.Errorf("Unexpected output: %s", out)
	}
}

func TestInvalidUsage(t *testing.T) {
	bot := &tuxbot{}
	out, err := bot.Run("", []string{"build", "oops"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(out, "```\nI don't know what you mean by") {
		t.Errorf("Unexpected output: %s", out)
	}
}
