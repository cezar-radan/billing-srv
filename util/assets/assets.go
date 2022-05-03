package assets

import (
	"embed"
	"os"
	"path/filepath"
	"strings"
)

var (
	prefixToEmbededPath string = "util/assets/"

	//go:embed images
	resi       embed.FS
	imageFiles = map[string]string{
		"ledgertech_icon1": "images/ledgertech_icon1.png",
		"ledgertech_icon2": "images/ledgertech_icon2.png",
		"ledgertech_logo1": "images/ledgertech_logo1.png",
	}
	//go:embed fonts
	resf      embed.FS
	fontFiles = map[string]string{
		"OpenSansHebrew-Bold":    "fonts/OpenSansHebrew-Bold.ttf",
		"OpenSansHebrew-Italic":  "fonts/OpenSansHebrew-Italic.ttf",
		"OpenSansHebrew-Regular": "fonts/OpenSansHebrew-Regular.ttf",
	}
)

type EmbededFile struct {
	FileName      string
	FileShortPath string
	FileLongPath  string
	FileContent   []byte
}

func LoadEmbeddedFonts() (map[string]*EmbededFile, error) {
	out := map[string]*EmbededFile{}

	for k, v := range fontFiles {
		b, err := resf.ReadFile(v)
		if err != nil {
			return map[string]*EmbededFile{}, err
		}

		ff := EmbededFile{
			FileName:      k,
			FileShortPath: v,
			FileLongPath:  filepath.Join(prefixToEmbededPath, v),
			//FileContent:   b,
		}
		_ = b

		out[k] = &ff
	}

	return out, nil
}

func UserFonts(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}
	ff := []string(nil)
	for _, f := range files {
		if isTrueType(f.Name()) {
			fn := filepath.Join(dir, f.Name())
			ff = append(ff, fn)
		}
	}
	return ff, nil
}

func isTrueType(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".ttf")
}

func LoadEmbeddedImages() (map[string]*EmbededFile, error) {
	out := map[string]*EmbededFile{}

	for k, v := range imageFiles {
		b, err := resi.ReadFile(v)
		if err != nil {
			return map[string]*EmbededFile{}, err
		}

		ifl := EmbededFile{
			FileName:      k,
			FileShortPath: v,
			FileLongPath:  filepath.Join(prefixToEmbededPath, v),
			//FileContent: b,
		}
		_ = b

		out[k] = &ifl
	}

	return out, nil
}
