package logs

import (
	"GoEncrypt/pkg/utils"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// writing logs to logs file.
func WriteLogs(content string) error {
	rootPath, err := utils.GetRootPath("go.mod")
	if err != nil {
		return err
	}

	logsFilePath := filepath.Join(rootPath, "data/logs", "logs.txt")

	file, err := os.OpenFile(logsFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[!] ERROR trying to write logs:", err)
		return err
	}
	defer file.Close()

	t := time.Date(2024, time.October, 17, 10, 30, 0, 0, time.UTC)
	dateStr := t.Format("2006-01-02 15:04:05")

	_, err = file.WriteString(dateStr + content + "\n")
	if err != nil {
		fmt.Println("[!] ERROR trying to write logs:", err)
		return err
	}

	return nil
}
