package mime_manager

import (
	"shopapi/internal/supports"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	imageFileType = "image"
)

func TestMimeNabager(t *testing.T) {
	t.Parallel()

	t.Run("AddAllowedExtensions", func(t *testing.T) {
		require.NotPanics(t, func() {
			AddAllowedExtensions("image", []string{".JPEG", ".PNG"})
		})
	})

	t.Run("GetFileExtension", func(t *testing.T) {
		t.Parallel()

		require.NotPanics(t, func() {

			AddAllowedExtensions(imageFileType, []string{".JPEG", ".JPG", ".PNG"})
			ext, err := GetFileExtension(supports.TestImage, imageFileType)
			require.Nil(t, err)
			require.Equal(t, ext, ".jpeg")
		})
	})

	t.Run("IsFileAllowed", func(t *testing.T) {
		t.Parallel()

		require.NotPanics(t, func() {
			AddAllowedExtensions(imageFileType, []string{".JPEG", ".JPG", ".PNG"})
			err := IsFileAllowed(supports.TestImage, imageFileType)
			require.Nil(t, err)

			err = IsFileAllowed([]byte("Some text file not allowed"), imageFileType)
			require.NotNil(t, err)

			AddAllowedExtensions(imageFileType, []string{".txt"})
			err = IsFileAllowed([]byte("Some text file allowed now"), imageFileType)
			require.Nil(t, err)
		})
	})
}
