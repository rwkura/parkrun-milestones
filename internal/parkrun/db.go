package parkrun

import (
	"fmt"
	"os"
	"path"
	"time"

	download "github.com/rwkura/parkrun-milestones/internal/download"
	file "github.com/rwkura/parkrun-milestones/internal/file"
)

var MaxFileAge time.Duration = 24 * time.Hour

func CachePath(format string, a ...any) (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(base, "parkrun-milestones", fmt.Sprintf(format, a...)), nil
}

func DownloadAndRead(url string, fileName string) (string, time.Time, error) {
	filePath, err := CachePath(fileName)
	if err != nil {
		return "", time.Time{}, err
	}

	if err := download.DownloadFile(url, filePath, MaxFileAge); err != nil {
		return "", time.Time{}, err
	}

	buf, err := file.ReadFile(filePath)
	if err != nil {
		return "", time.Time{}, err
	}

	t, err := file.GetMtime(filePath)
	if err != nil {
		return "", time.Time{}, err
	}

	return buf, t, err
}

func DownloadAndReadMaxMtime(url string, fileName string, maxMtime time.Time) (string, time.Time, error) {
	filePath, err := CachePath(fileName)
	if err != nil {
		return "", time.Time{}, err
	}

	if err := download.DownloadFileMaxMtime(url, filePath, maxMtime); err != nil {
		return "", time.Time{}, err
	}

	buf, err := file.ReadFile(filePath)
	if err != nil {
		return "", time.Time{}, err
	}

	t, err := file.GetMtime(filePath)
	if err != nil {
		return "", time.Time{}, err
	}

	return buf, t, err
}
