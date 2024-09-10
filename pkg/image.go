package pkg

import (
	"path/filepath"
	"strings"
)

func isImageByExt(fileName string) bool {
	switch strings.ToLower(filepath.Ext(fileName)) {
	case ".jpg", ".jpeg", ".png", ".pdf":
		return true
	default:
		return false
	}
}

func isImage(bs []byte, fileName string) bool {
	split := strings.Split(fileName, ".")
	suffix := split[len(split)-1]
	suffix = strings.ToLower(suffix)
	switch suffix {
	case "jpg", "jpeg":
		if len(bs) < 2 {
			return false
		}
		if bs[0] == 0xFF && bs[1] == 0xD8 {
			return true
		}
	case "png":
		if len(bs) < 7 {
			return false
		}

		if bs[0] == 0x89 && bs[1] == 0x50 && bs[2] == 0x4E && bs[3] == 0x47 &&
			bs[4] == 0x0D && bs[5] == 0x0A && bs[6] == 0x1A {
			return true
		}
	}
	return false
}

func isPDF(name string) bool {
	return strings.HasSuffix(strings.ToLower(name), ".pdf")
}
