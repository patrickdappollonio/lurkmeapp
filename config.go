package main

import (
	"errors"
	"os"
	"strings"
)

func getLoginInfo() (string, string, error) {
	var (
		lurkerUsername = os.Getenv("USERNAME")
		lurkerToken    = os.Getenv("TOKEN")
	)

	if lurkerUsername == "" {
		return "", "", errors.New("no username defined in $USERNAME")
	}

	if lurkerToken == "" {
		return "", "", errors.New("no token defined in $TOKEN")
	}

	return lurkerUsername, lurkerToken, nil
}

func env(key, defval string) string {
	if v, found := os.LookupEnv(key); found {
		if s := strings.TrimSpace(v); s != "" {
			return s
		}

		return defval
	}

	return defval
}
