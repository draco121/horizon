package utils

import (
	"fmt"
	"os"
	"path"
)

func BaseDir() string {
	appDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	appDir = path.Join(appDir, "botman")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		// Directory does not exist, create it
		_ = os.MkdirAll(appDir, 0755) // 0755 is the default permission
		fmt.Printf("Directory '%s' created successfully.\n", appDir)
	}
	return appDir
}

func CreateBotSpace(botId, projectId string) (string, error) {
	appDir := BaseDir()
	botSpace := path.Join(appDir, projectId, botId)
	if _, err := os.Stat(botSpace); os.IsNotExist(err) {
		// Directory does not exist, create it
		_ = os.MkdirAll(botSpace, 0755) // 0755 is the default permission
		fmt.Printf("Directory '%s' created successfully.\n", botSpace)
	}
	return botSpace, nil
}
