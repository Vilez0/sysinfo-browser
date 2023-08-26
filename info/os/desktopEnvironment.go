package osutils

import (
	"os"
)

func Desktop() string {
	desktop := os.Getenv(`DESKTOP_SESSION`)
	return desktop
}
