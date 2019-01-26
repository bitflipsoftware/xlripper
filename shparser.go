package xlsx

import (
	"bytes"
	"fmt"
	"io"
	"unicode"
)

var badPair = indexPair{-1, -1}
var badTagLoc = tagLoc{badPair, badPair}

const (
	lChevron = '<'
	rChevron = '>'
)

func shparse(zs zstruct, sheetIndex int) (Sheet, error) {
	sh := NewSheet()

	if sheetIndex < 0 || sheetIndex >= len(zs.info.sheetMeta) {
		return sh, fmt.Errorf("bad sheet index '%d', the zstruct has only '%d' sheets", sheetIndex, len(zs.info.sheetMeta))
	}

	meta := zs.info.sheetMeta[sheetIndex]
	sh.Name = meta.sheetName
	sh.Index = sheetIndex
	data, err := shload(meta)

	if err != nil {
		return sh, err
	}

	next := 0
	first, last := 0, 0
	for first != -1 && last != -1 {
		tloc := shfindRow(data, next, len(data))
		if first != -1 && last != -1 {
			rowRunes := data[first : last+1]
			// TODO - send rune slice down a pipeline
			// TODO - remove this debugging
			str := string(rowRunes)
			fmt.Print(str)
			use(tloc)
		}
		next = last + 1
	}

	return sh, nil
}

// shload reads the worksheet file and returns the unzipped data therein as a slice of runes
func shload(meta sheetMeta) ([]rune, error) {
	if meta.file == nil {
		return make([]rune, 0), fmt.Errorf("the file is nil")
	}

	reader, err := meta.file.Open()

	if err != nil {
		return make([]rune, 0), err
	}

	defer reader.Close()
	buf := bytes.Buffer{}
	io.Copy(&buf, reader)
	data := string(buf.Bytes())
	return []rune(data), nil
}

// shadvance starts at 'first' and advances until it finds 'r' then returns the index of 'r'. returns -1 if 'r' is not
// found
func shadvance(runes []rune, start int, r rune) int {
	e := len(runes)
	ix := start

	if start < 0 {
		return -1
	}

	for ; ix < e; ix++ {
		if runes[ix] == r {
			return ix
		}
	}

	return -1
}

// shfind row starts at 'first' looks ahead to find the first and last indices of a <row> tag. it return the first and
// last indices of the row tag. that is, if you take data[first:last+1] you will get exactly the complete row tag.
// a return of -1, -1 indicates that there was no row found
func shfindRow(runes []rune, first, last int) tagLoc {
	ix := shSetFirst(runes, first)
	e := shSetLast(runes, last)
	s := shdebug(runes, 0, 10000)
	use(s)
	x := shTagFind(runes, ix, e, "row")
	return x
}

// shdebug is used in debugging to view a chunk of data as a string instead of a rune slice (i.e. so you can log it or
// view it in a debugger). index is your area of interest, window is the number of chars before and after to include.
func shdebug(runes []rune, index, window int) string {
	a := maxi(0, index-window)
	b := mini(len(runes), index+window+1)

	if a == b {
		return ""
	}

	return string(runes[a:b])
}

// shbad returns true if the index is out of range
func shbad(runes []rune, ix int) bool {
	if ix < 0 {
		return true
	}

	if ix >= len(runes) {
		return true
	}

	return false
}

func shFindNamespaceColon(runes []rune, first, last int) int {
	e := shSetLast(runes, last)
	ix := shSetFirst(runes, first)
	var r rune
	namespaceColonPos := -1

findNamespaceColon:
	for localIX := ix; localIX <= e; localIX++ {
		r = runes[localIX]
		if r == '<' {
			return -1
		} else if r == ':' {
			namespaceColonPos = localIX
			break findNamespaceColon
		} else if r == '>' {
			break findNamespaceColon
		}
	}

	return namespaceColonPos
}

// shIsTag returns true if the tag matches the desired element and false if it does not. specify whether it is an open
// tag or a close tag with isCloseTag. 'first' must be pointing to the first char 'inside' the tag, that is after '<'
// or '</'. returns the location of the closing '>' or -1 if the tag is not well formed or does not match elem
func shTagCompletion(runes []rune, first, last int, elem string) int {
	e := shSetLast(runes, last)
	ix := shSetFirst(runes, first)
	var r rune
	peek := shdebug(runes, ix, 3)
	use(peek)
	namespaceColonPos := shFindNamespaceColon(runes, ix, e)
	if namespaceColonPos > 0 {
		ix = namespaceColonPos + 1
	}

	r = runes[ix]
	shdebugchar(r)
	shdebugrange(runes, ix)

	// we should be pointing at the first rune of the element name now
	if ix > e || r == '<' || r == ':' || r == ' ' || r == '>' {
		return -1
	}

	elemRunes := []rune(elem)
	elemLen := len(elemRunes)
	for elemIx := 0; elemIx < elemLen; elemIx++ {
		if ix > e {
			return -1
		}
		r = runes[ix]
		if r == '>' {
			return -1
		} else if r == ':' {
			return -1
		} else if r != elemRunes[elemIx] {
			return -1
		}
		ix++
	}

	// proceed to close
	for ; ix <= e; ix++ {
		r = runes[ix]
		if r == '>' {
			return ix
		}
	}

	return -1
}

// shTagOpenFind returns the first and last indices of an element open tag with the name 'elem' (ignoring namespace).
// {-1, -1} indicates that no matching open tag was found. 'last' is the last rune that you want inspected for a closing
// tag. this is unlike slice indexing and more like traditional range indexing. enter -1 to go to the end of the runes.
func shTagOpenFind(runes []rune, first, last int, elem string) (found indexPair, lastCheckedIndex int) {
	e := shSetLast(runes, last)
	ix := shSetFirst(runes, first)
	var r rune
	foundFirst := -1

findOpenTag:
	for ; ix <= e; ix++ {
		r = runes[ix]
		if r == '<' {
			foundFirst = ix
			break findOpenTag
		}
	}

	ix++

	if ix > e {
		return badPair, ix
	}

	peek := shdebug(runes, ix, 3)
	use(peek)
	foundLast := shTagCompletion(runes, ix, e, elem)

	if foundLast <= ix {
		return badPair, ix
	}

	return indexPair{foundFirst, foundLast}, ix
}

// shTagCloseFind returns the first and last indices of an element close tag with the name 'elem' (ignoring namespace).
// {-1, -1} indicates that no matching open tag was found. If elements of the same name are nested, the nested close
// tags are skipped. 'first' must be the first rune index that is inside of the element you want to find the close for.
// 'last' is the last rune that you want inspected for a closing tag. this is unlike slice indexing and more like
// traditional range indexing. enter -1 to go to the end of the runes
func shTagCloseFind(runes []rune, first, last int, elem string) indexPair {
	e := shSetLast(runes, last)
	ix := shSetFirst(runes, first)
	var r rune
	foundFirst := -1

findLeftChevron:
	for ; ix <= e; ix++ {
		r = runes[ix]
		if r == '<' {
			foundFirst = ix
			break findLeftChevron
		}
	}

	ix++

	if ix > e {
		return badPair
	}

	r = runes[ix]

	if r != '/' {
		localElemName, localElemLast := shTagNameFind(runes, ix, e)

		if len(localElemName) == 0 || localElemLast < 0 {
			return badPair
		}

		ix = localElemLast

		if runes[ix] != rChevron {
			return badPair
		}

		ix++

		if ix > e {
			return badPair
		}

		peek := shdebug(runes, ix, 3)
		use(peek)
		// now we are inside of a nested element
		nestedCloseLoc := shTagCloseFind(runes, ix, e, localElemName)

		if nestedCloseLoc == badPair {
			return badPair
		}

		ix = nestedCloseLoc.last
		ix++

		if ix > e {
			return badPair
		}

		// now we have advanced beyond the nested element
		// we need to call ourself again to find the closing tag
		localFoundPair := shTagCloseFind(runes, ix, e, elem)
		return localFoundPair
	}

	ix++

	if ix > e {
		return badPair
	}

	for ; ix <= e && unicode.IsSpace(runes[ix]); ix++ {
		// advance past white space
	}

	if ix > e {
		return badPair
	}

	foundLast := shTagCompletion(runes, ix, last, elem)

	if foundLast <= ix {
		return badPair
	}

	return indexPair{foundFirst, foundLast}
}

// shTagFind returns the open and close locations for the desired tag 'elem' returns -1 (somewhere) if not found.
// 'last' is the last rune that you want inspected for a closing tag. this is unlike slice indexing and more like
// traditional range indexing. enter -1 to go to the end of the runes
func shTagFind(runes []rune, first, last int, elem string) tagLoc {
	ix := shSetFirst(runes, first)
	e := shSetLast(runes, last)
	open := badPair

	for ; ix <= e && open == badPair; ix++ {
		s := shdebug(runes, ix, 20)
		use(s)
		open, ix = shTagOpenFind(runes, ix, e, elem)
	}

	if open == badPair {
		return badTagLoc
	}

	close := shTagCloseFind(runes, open.last+1, last, elem)

	if close == badPair {
		return badTagLoc
	}

	return tagLoc{open, close}
}

// shSetLast returns a safe 'last' value for loops on 'runes'
func shSetLast(runes []rune, requestedLast int) int {
	l := len(runes)
	if requestedLast < 0 {
		return l
	} else if requestedLast > l-1 {
		return l - 1
	}
	return requestedLast
}

func shSetFirst(runes []rune, first int) int {
	if first < 0 {
		return 0
	} else if first > (len(runes) - 1) {
		return len(runes) - 1
	}
	return first
}

func shdebugchar(r rune) {
	s := string(r)
	fmt.Print(s)
}

func shdebugrange(runes []rune, currentIX int) {
	first := shSetFirst(runes, currentIX-5)
	last := shSetLast(runes, currentIX+5)
	s := string(runes[first : last+1])
	fmt.Print(s)
}

// shTagNameFind returns the name of an element and the position of the close '>' for that element. 'first' should be
// pointing at the first rune after '<' or '</'. if the element cannot be parsed, -1 is returned for 'lastPos'
func shTagNameFind(runes []rune, first, last int) (elem string, lastPos int) {
	e := shSetLast(runes, last)
	ix := shSetFirst(runes, first)

	for ; ix <= e && runes[ix] == ' '; ix++ {
		// advance the index
	}

	if ix > e {
		return "", -1
	} else if runes[ix] == ' ' {
		ix++
	}

	if ix > e {
		return "", -1
	}

	namespaceColonPos := shFindNamespaceColon(runes, ix, e)
	if namespaceColonPos >= 0 {
		ix = namespaceColonPos + 1
	}

	strbuf := bytes.Buffer{}

	for ; ix <= e && runes[ix] != ' ' && runes[ix] != '>' && runes[ix] != '=' && runes[ix] != '"'; ix++ {
		strbuf.WriteRune(runes[ix])
	}

	elem = strbuf.String()

	if len(elem) == 0 {
		return "", -1
	}

	if ix > e {
		return "", -1
	}

	for ; ix <= e && runes[ix] != '>'; ix++ {
		// advance ix to find the closing of the tag
	}

	if runes[ix] == '>' {
		return elem, ix
	}

	return "", -1
}
