package core

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

func getHighScore() (int, error) {
	u, err := user.Current()
	if err != nil {
		return 0, err
	}

	p := ""
	switch runtime.GOOS {
	case "windows":
		p = fmt.Sprintf("C:\\Users\\%s\\AppData", u.Username)
	case "linux":
		p = fmt.Sprintf("/users/%s", u.Username)
	case "darwin":
		p = fmt.Sprintf("/Users/%s/Library/Application Support/Goasteroids", u.Username)
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		if err := os.Mkdir(p, 0755); err != nil {
			return 0, err
		}

		if _, err := os.Stat(p + "/highscore.txt"); os.IsNotExist(err) {
			err := os.WriteFile(p+"/highscore.txt", []byte("0"), 0750)
			if err != nil {
				return 0, err
			}
		}
	}

	content, err := os.ReadFile(p + "/highscore.txt")
	if err != nil {
		return 0, err
	}

	score := string(content)
	score = strings.TrimSpace(score)
	s, err := strconv.Atoi(score)
	if err != nil {
		return 0, err
	}

	return s, nil
}

func updateHighScore(score int) error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	p := ""
	switch runtime.GOOS {
	case "windows":
		p = fmt.Sprintf("C:\\Users\\%s\\AppData\\highscore.txt", u.Username)
	case "linux":
		p = fmt.Sprintf("/users/%s/highscore.txt", u.Username)
	case "darwin":
		p = fmt.Sprintf("/Users/%s/Library/Application Support/Goasteroids/highscore.txt", u.Username)
	}

	s := fmt.Sprintf("%d", score)

	return os.WriteFile(p, []byte(s), 0750)
}
