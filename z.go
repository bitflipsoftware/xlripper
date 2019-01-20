package xlsx

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bitflip-software/xlsx/xmlprivate"
)

const (
	strContentTypes = "[Content_Types].xml"
	strRels         = "_rels/.rels"
	strWorkbookRels = "_rels/workbook.xml.rels"
)

// zinfo represents info about how to find the xlsx parts inside of the zip package
type zinfo struct {
	contentTypesIndex int
	contentTypes      xmlprivate.ContentTypes
	relsIndex         int
	rels              xmlprivate.Rels
	wkbkName          string
	wkbkIndex         int
	wkbk              *zip.File
}

// zstruct represents the zip file reader and metadata about what was found in the xlsx package
type zstruct struct {
	r    *zip.Reader
	info zinfo
}

// zopen parses all of the necessary information from the xlsx package into a usable data structure
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

	z.info, err = zparseWorkbookLocation(zr, z.info)

	if err != nil {
		return zstruct{}, err
	}

	z.info, err = zparseWorkbookRels(zr, z.info)

	if err != nil {
		return zstruct{}, err
	}

	return z, nil
}

func zparseContentTypes(zr *zip.Reader, zi zinfo) (zout zinfo, err error) {
	zi.contentTypesIndex = zfind(zr, strContentTypes)

	if zi.contentTypesIndex < 0 {
		return zi, err
	}

	file := zr.File[zi.contentTypesIndex]
	ctt := xmlprivate.ContentTypes{}
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

	if len(ctt.Defaults) == 0 && len(ctt.Overrides) == 0 {
		return zi, fmt.Errorf("the %s file has no contents", strContentTypes)
	}

	zi.contentTypes = ctt
	return zi, nil
}

// zparseWorkbookLocation must come after zparseRels
func zparseWorkbookLocation(zr *zip.Reader, zi zinfo) (zout zinfo, err error) {
	// examples seen so far
	// http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	// http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument

	wkbkRelsIndex := -1

	for ix, rel := range zi.rels.Rels {
		rx, _ := regexp.Compile(`.+officeDocument.+officeDocument$`)
		match := rx.Match([]byte(rel.Type))
		if match {
			wkbkRelsIndex = ix
		}
	}

	if wkbkRelsIndex < 0 {
		for ix, rel := range zi.rels.Rels {
			rx, _ := regexp.Compile(`workbook\.xml$`)
			match := rx.Match([]byte(rel.Type))
			if match {
				wkbkRelsIndex = ix
			}
		}
	}

	if wkbkRelsIndex < 0 {
		return zi, nil
	}

	wkb := zi.rels.Rels[wkbkRelsIndex].Target
	zi.wkbkName = removeLeadingSlash(wkb)
	zi.wkbkIndex = zfind(zr, wkb)

	if zi.wkbkIndex < 0 {
		return zi, errors.New("the workbook could not be found")
	}

	zi.wkbk = zr.File[zi.wkbkIndex]
	return zi, nil
}

func wkbkRelsPath(wkbkPath string) (wkbkRelsPath string) {
	dir := filepath.Dir(wkbkPath)
	path := path.Join(dir, strWorkbookRels)
	return path
}

func zfind(zr *zip.Reader, filename string) (index int) {
	filename = removeLeadingSlash(filename)

	for ix, file := range zr.File {
		lcActual := strings.ToLower(removeLeadingSlash(file.FileHeader.Name))
		lcToFind := strings.ToLower(filename)
		lenActual := len(lcActual)
		lenToFind := len(lcToFind)

		if lenActual < lenToFind {
			continue
		}

		if lcActual[lenActual-lenToFind:] == lcToFind {
			return ix
		}
	}

	return -1
}

func removeLeadingSlash(instr string) (outstr string) {
	if len(instr) == 0 {
		return instr
	} else if len(instr) == 1 && instr == "/" {
		return ""
	} else if len(instr) == 1 && instr != "/" {
		return instr
	}

	var first rune
	for _, r := range instr {
		first = r
		break
	}

	if first == '/' {
		return instr[1:]
	}

	return instr
}

func zparseRels(zr *zip.Reader, zi zinfo) (zout zinfo, err error) {
	zi.relsIndex = zfind(zr, strRels)

	if zi.contentTypesIndex < 0 {
		return zi, err
	}

	file := zr.File[zi.relsIndex]
	xstruct := xmlprivate.Rels{}
	fbuf := bytes.Buffer{}
	fwrite := bufio.NewWriter(&fbuf)
	ofile, err := file.Open()

	if err != nil {
		return zi, err
	} else {
		defer ofile.Close()
	}

	io.Copy(fwrite, ofile)
	err = xml.Unmarshal(fbuf.Bytes(), &xstruct)

	if err != nil {
		return zi, err
	}

	zi.rels = xstruct

	return zi, nil
}

// zparseWorkbookRels requires that the workbook has been found
func zparseWorkbookRels(zr *zip.Reader, zi zinfo) (zout zinfo, err error) {
	wrelsName := wkbkRelsPath(zi.wkbkName)
	ix := zfind(zr, wrelsName)

	if ix < 0 {
		return zi, fmt.Errorf("workbook rels '%s' could not be found", wrelsName)
	}

	return zi, nil
}
