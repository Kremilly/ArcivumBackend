package utils

import (
	"fmt"
    "strings"
    "path/filepath"
)

func FormatBytes(bytes uint) string {
    const unit = 1024
    if bytes < unit {
        return fmt.Sprintf("%d B", bytes)
    }

    div, exp := int64(unit), 0
    for n := bytes / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }

    return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func GetMime(filePath string) string {
	mimeTypes := map[string]string{
		".sql":  "application/sql",
		".json": "application/json",
		".xml":  "application/xml",
		".html": "text/html",
		".txt":  "text/plain",
		".pdf":  "application/pdf",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	if mimeType, exists := mimeTypes[ext]; exists {
		return mimeType
	}

	return "application/octet-stream"
}

func StringToUint(s string) uint {
	var n uint

	fmt.Sscanf(s, "%d", &n)
	return n
}
