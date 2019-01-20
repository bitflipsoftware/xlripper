package xlsx

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// unzip is a reference function that I found on the Internet
// TODO - remove this function
func unzip(src string, dest string) ([]string, error) {

	var filenames []string
	b := []byte(src)
	brdr := bytes.NewReader(b)
	r, err := zip.NewReader(brdr, int64(len(b)))

	//r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	//defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}
