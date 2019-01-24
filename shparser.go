package xlsx

import (
	"bytes"
	"fmt"
	"io"
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
		first, last = shfindRow(data, next)
		if first != -1 && last != -1 {
			rowRunes := data[first : last+1]
			// TODO - send rune slice down a pipeline
			// TODO - remove this debugging
			str := string(rowRunes)
			fmt.Print(str)
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
func shfindRow(runes []rune, start int) (int, int) {
	ix := start
	end := len(runes)
	done := func(current, theEnd int) bool { return current >= end }

	first := ix
	last := -1

startTagLoop:
	for {
		ix = shadvance(runes, start, '<')

		if shbad(runes, ix) {
			return -1, -1
		}

		// TODO - remove this debugging
		window := shdebug(runes, ix, 5)
		use(window)

		// set index to the first rune inside of the tag marker
		ix++

		if done(ix, end) {
			return -1, -1
		}

		// skip any namespace
		for ; ix < end && runes[ix] != ':' && runes[ix] != 'r' && runes[ix] != '>'; ix++ {
			str := string(runes[ix])
			use(str)
			// TODO - remove this debugging
			peekRune := shdebug(runes, ix, 1)
			use(peekRune)
		}

		if done(ix, end) {
			return -1, -1
		}

		//for ; ix < end && runes[ix] != ':' && runes[ix] != 'r' && runes[ix] != '>' && runes[ix] != ' ' && runes[ix] != '=' && runes[ix] != '"'; ix++ {
		//	// TODO - remove this debugging
		//
		//	peekRune := shdebug(runes, ix, 1)
		//	use(peekRune)
		//}
		//
		//if done(ix, end) {
		//	return -1, -1
		//}

		if runes[ix] == ':' {
			ix++
		}
		//else if runes[ix] == '>' {
		//	continue startTagLoop
		//}

		if done(ix, end) {
			return -1, -1
		}

		// check for 'row '
		if runes[ix] != 'r' {
			continue startTagLoop
		}

		if (ix-1 < 0) || ((runes[ix-1] != ':') && (runes[ix-1] != '<')) {
			continue startTagLoop
		}

		ix++
		// TODO - remove this debugging

		peekRune := shdebug(runes, ix, 1)
		use(peekRune)

		if done(ix, end) {
			return -1, -1
		}

		if runes[ix] != 'o' {
			continue startTagLoop
		}

		ix++

		if done(ix, end) {
			return -1, -1
		}

		if runes[ix] != 'w' {
			continue startTagLoop
		}

		ix++

		if done(ix, end) {
			return -1, -1
		}

		if runes[ix] != ' ' && runes[ix] != '>' {
			continue startTagLoop
		}

		// if we reach here then we have successfully identified the first of a <row> tag
		break startTagLoop
	}

	if first == -1 {
		panic("bug")
	}

closeTagLoop:
	for {
		for ; ix < end && runes[ix] != '<'; ix++ {
			// just advance ix
		}

		// TODO - remove this debugging
		window := shdebug(runes, ix, 5)
		use(window)

		// set index to the first rune inside of the tag marker
		ix++

		if done(ix, end) {
			return -1, -1
		}

		if runes[ix] != '/' {
			// this is not a close tag
			continue closeTagLoop
		}

		// set index to the first rune after the close slash marker </
		ix++

		if done(ix, end) {
			return -1, -1
		}

		for ; ix < end && runes[ix] != ':' && runes[ix] != 'r' && runes[ix] != '>' && runes[ix] != '=' && runes[ix] != '"'; ix++ {
			// TODO - remove this debugging

			peekRune := shdebug(runes, ix, 1)
			use(peekRune)
		}

		// skip any namespace
		if done(ix, end) {
			return -1, -1
		}

		if runes[ix] == ':' {
			ix++
		} else if runes[ix] == '>' {
			continue closeTagLoop
		}

		if done(ix, end) {
			return -1, -1
		}

		// check for 'row '
		if runes[ix] != 'r' {
			continue closeTagLoop
		}

		ix++

		if done(ix, end) {
			return -1, -1
		}

		if runes[ix] != 'o' {
			continue closeTagLoop
		}

		ix++

		if done(ix, end) {
			return -1, -1
		}

		if runes[ix] != 'w' {
			continue closeTagLoop
		}

		ix++

		for ; ix < end && runes[ix] != '>'; ix++ {

		}

		if done(ix, end) {
			return -1, -1
		} else if runes[ix] != '>' {
			return -1, -1
		}

		// if we reach here then we have successfully identified the end of a </row> tag
		last = ix - 1 // minus one because we obtained this ix from a for loop with increment which is one past cur
		break closeTagLoop
	}

	if first == -1 {
		panic("bug")
	} else if last == -1 {
		return -1, -1
	}

	return first, last
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

// shTagOpenFind returns the first and last indices of an element open tag with the name 'elem' (ignoring namespace).
// {-1, -1} indicates that no matching open tag was found. 'last' is the last rune that you want inspected for a closing
// tag. this is unlike slice indexing and more like traditional range indexing. enter -1 to go to the end of the runes.
func shTagOpenFind(runes []rune, first, last int, elem string) indexPair {
	e := shSetEnd(runes, last)
	return indexPair{-e, -e}
}

// shTagCloseFind returns the first and last indices of an element close tag with the name 'elem' (ignoring namespace).
// {-1, -1} indicates that no matching open tag was found. If elements of the same name are nested, the nested close
// tags are skipped. 'first' must be the first rune index that is inside of the element you want to find the close for.
// 'last' is the last rune that you want inspected for a closing tag. this is unlike slice indexing and more like
// traditional range indexing. enter -1 to go to the end of the runes
func shTagCloseFind(runes []rune, first, last int, elem string) indexPair {
	e := shSetEnd(runes, last)
	return indexPair{-e, -e}
}

// shTagFind returns the open and close locations for the desired tag 'elem' returns -1 (somewhere) if not found.
// 'last' is the last rune that you want inspected for a closing tag. this is unlike slice indexing and more like
// traditional range indexing. enter -1 to go to the end of the runes
func shTagFind(runes []rune, first, last int, elem string) tagLoc {
	return tagLoc{indexPair{-1, -1}, indexPair{-1, -1}}
}

// shSetEnd returns a safe 'last' value for loops on 'runes'
func shSetEnd(runes []rune, requestedLast int) int {
	l := len(runes)
	if requestedLast < 0 {
		return l
	} else if requestedLast > l-1 {
		return l - 1
	}
	return requestedLast
}
