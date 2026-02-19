package mime_manager

import (
	"fmt"
	"mime"
	"net/http"
	"strings"
	"sync"
)

var mutex sync.RWMutex

type ExtensionAllowed map[string]bool

var allowedExtensionsForFileType = map[string]ExtensionAllowed{}

type MimeAllowed map[string]bool

var allowedMimeForFileType = map[string]MimeAllowed{}

// The extensions have to be with a leading dot, as in ".html". Panics if no mime type assocoated with extension.
func AddAllowedExtensions(fileType string, extensions []string) {
	mutex.Lock()
	defer mutex.Unlock()

	if allowedExtensionsForFileType[fileType] == nil {
		allowedExtensionsForFileType[fileType] = ExtensionAllowed{}
	}
	if allowedMimeForFileType[fileType] == nil {
		allowedMimeForFileType[fileType] = MimeAllowed{}
	}

	for _, s := range extensions {
		ext := strings.ToLower(s)
		mime := mime.TypeByExtension(ext)
		if mime == "" {
			panic(fmt.Errorf("no mime type assocoated with %s", ext))
		}
		allowedExtensionsForFileType[fileType][ext] = true
		allowedMimeForFileType[fileType][mime] = true
	}

}

func GetFileExtension(data []byte, fileType string) (string, error) {
	mimeType := detectMimeType(data)
	err := IsMimeAllowed(mimeType, fileType)
	if err != nil {
		return "", err
	}

	exts, err := mime.ExtensionsByType(mimeType)
	if err != nil {
		return "", err
	}

	mutex.RLock()
	defer mutex.RUnlock()

	allowedExtensions, ok := allowedExtensionsForFileType[fileType]
	if !ok {
		return "", fmt.Errorf("file type '%s' not allowed", fileType)
	}

	for _, ext := range exts {
		if _, ok := allowedExtensions[ext]; ok {
			return ext, nil
		}
	}

	return "", fmt.Errorf("invalid file or extension not allowed for detected mime: '%s'", mimeType)
}

func IsFileAllowed(data []byte, fileType string) error {
	mimeType := detectMimeType(data)
	err := IsMimeAllowed(mimeType, fileType)
	if err != nil {
		return err
	}

	return nil
}

func IsMimeAllowed(mimeType, fileType string) error {
	mutex.RLock()
	defer mutex.RUnlock()

	allowedMime, ok := allowedMimeForFileType[fileType]
	if !ok {
		return fmt.Errorf("file type '%s' not allowed", fileType)
	}

	if _, ok := allowedMime[mimeType]; !ok {
		return fmt.Errorf("mime type '%s' is not allowed for file type '%s'", mimeType, fileType)
	}

	return nil
}

func detectMimeType(data []byte) string {
	return http.DetectContentType(data)
}
