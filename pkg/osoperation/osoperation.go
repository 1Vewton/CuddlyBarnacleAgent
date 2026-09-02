package osoperation

import (
	"strings"
)

// PathToWindows converts path form to windows type
func PathToWindows(
	path string,
) string {
	return strings.ReplaceAll(
		path,
		`\`,
		"/",
	)
}
