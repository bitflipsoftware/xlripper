package xlsx

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const strContentTypes = "[Content_Types].xml"

// zinfo represents info about how to find the xlsx parts inside of the zip package
type zinfo struct {
	contentTypesFound bool
	contentTypesIndex int
}

// zstruct represents the zip file reader and metadata about what was found in the xlsx package
type zstruct struct {
	r    *zip.Reader
	info zinfo
}

func zopen(zipData string) (z zstruct, err error) {
	b := []byte(zipData)
	brdr := bytes.NewReader(b)
	zr, err := zip.NewReader(brdr, int64(len(b)))

	if err != nil {
		return zstruct{}, err
	} else if zr == nil {
		return zstruct{}, errors.New("a nil zip.Reader was encountered")
	}

	z, err = zinit(zr)

	if err != nil {
		return zstruct{}, err
	}

	return z, nil
}

// zinit requires an open, error free *zip.Reader and returns a fully constructed zstruct
func zinit(zr *zip.Reader) (z zstruct, err error) {
	z.r = zr

	z.info, err = zparseContentTypes(zr, z.info)

	if err != nil {
		return zstruct{}, err
	}

	z.info, err = zparseRels(zr, z.info)

	if err != nil {
		return zstruct{}, err
	}

	return z, nil
}

func zparseContentTypes(zr *zip.Reader, zi zinfo) (zout zinfo, err error) {
	if zr == nil {
		return zi, errors.New("nil zip.Reader")
	}

	for ix, file := range zr.File {
		lcactual := strings.ToLower(file.FileHeader.Name)
		lcexpect := strings.ToLower(strContentTypes)
		lenactual := len(lcactual)
		lenexpect := len(lcexpect)

		if lenactual < lenexpect {
			continue
		}

		if lcactual[lenactual-lenexpect:] == lcexpect {
			zi.contentTypesFound = true
			zi.contentTypesIndex = ix
			break
		}
	}

	if zi.contentTypesIndex < 0 {
		return zi, err
	}

	file := zr.File[zi.contentTypesIndex]
	ctt := contentTypesTypes{}
	ctbuf := bytes.Buffer{}
	ctwri := bufio.NewWriter(&ctbuf)
	ofile, err := file.Open()

	if err != nil {
		return zi, err
	} else {
		defer ofile.Close()
	}

	io.Copy(ctwri, ofile)
	err = xml.Unmarshal(ctbuf.Bytes(), &ctt)

	if err != nil {
		return zi, err
	}

	xmlFile, err := os.Open("users.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	return zi, nil
}

func zparseRels(zr *zip.Reader, zi zinfo) (zout zinfo, err error) {
	return zi, nil
}

// unzip is a reference function that I found on the Internet
// TODO - remove this function
func unzip(src string, dest string) ([]string, error) {

	var filenames []string
	z, err := zopen(src)

	//z, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	//defer z.Close()

	for _, f := range z.r.File {

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

type ContentTypeItem struct {
	//Default     xml.Name `xml:"Default"`
	//Override    xml.Name `xml:"Override"`
	ContentType string `xml:"ContentType,attr""`
	Extension   string `xml:"Extension,attr"`
	PartName    string `xml:"PartName,attr"`
}

type ContentTypeDefault struct {
	XMLName xml.Name `xml:"Default"`
	ContentTypeItem
}

type ContentTypeOverride struct {
	XMLName xml.Name `xml:"Override"`
	ContentTypeItem
}

type ContentTypes struct {
	xmlns     string                `xml:"xmlns,attr"`
	defaults  []ContentTypeDefault  `xml:"Default"`
	overrides []ContentTypeOverride `xml:"Override"`
}

type contentTypesTypes struct {
	XMLName   xml.Name              `xml:"Types"`
	Defaults  []ContentTypeDefault  `xml:"Default"`
	Overrides []ContentTypeOverride `xml:"Override"`
}
