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

// shfind row starts at 'start' looks ahead to find the first and last indices of a <row> tag. it return the first and
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
		for ; ix < end && runes[ix] != '<'; ix++ {
			first = ix
		}

		// TODO - remove this debugging
		window := shdebug(runes, ix, 5)
		use(window)

		// set index to the first rune inside of the tag marker
		ix++

		if done(ix, end) {
			return -1, -1
		}

		cur := runes[ix]

		// skip any namespace
		for ; ix < end && cur != ':'; ix++ {
			// TODO - remove this debugging
			cur = runes[ix]
			peekRune := shdebug(runes, ix, 1)
			use(peekRune)
		}

		if done(ix, end) {
			return -1, -1
		}

		//for ; ix < end && cur != ':' && cur != 'r' && cur != '>' && cur != ' ' && cur != '=' && cur != '"'; ix++ {
		//	// TODO - remove this debugging
		//	cur = runes[ix]
		//	peekRune := shdebug(runes, ix, 1)
		//	use(peekRune)
		//}
		//
		//if done(ix, end) {
		//	return -1, -1
		//}

		cur = runes[ix]
		if cur == ':' {
			ix++
		}
		//else if cur == '>' {
		//	continue startTagLoop
		//}

		if done(ix, end) {
			return -1, -1
		}

		// check for 'row '
		if cur != 'r' {
			continue startTagLoop
		}

		ix++
		// TODO - remove this debugging
		cur = runes[ix]
		peekRune := shdebug(runes, ix, 1)
		use(peekRune)

		if done(ix, end) {
			return -1, -1
		}

		cur = runes[ix]

		if cur != 'o' {
			continue startTagLoop
		}

		ix++

		if done(ix, end) {
			return -1, -1
		}

		cur = runes[ix]

		if cur != 'w' {
			continue startTagLoop
		}

		ix++

		if done(ix, end) {
			return -1, -1
		}

		cur = runes[ix]

		if cur != ' ' && cur != '>' {
			continue startTagLoop
		}

		// if we reach here then we have successfully identified the start of a <row> tag
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

		cur := runes[ix]

		if cur != '/' {
			// this is not a close tag
			continue closeTagLoop
		}

		// set index to the first rune after the close slash marker </
		ix++

		if done(ix, end) {
			return -1, -1
		}

		cur = runes[ix]

		for ; ix < end && cur != ':' && cur != 'r' && cur != '>' && cur != '=' && cur != '"'; ix++ {
			// TODO - remove this debugging
			cur = runes[ix]
			peekRune := shdebug(runes, ix, 1)
			use(peekRune)
		}

		// skip any namespace
		if done(ix, end) {
			return -1, -1
		}

		cur = runes[ix]
		if cur == ':' {
			ix++
		} else if cur == '>' {
			continue closeTagLoop
		}

		if done(ix, end) {
			return -1, -1
		}

		// check for 'row '
		if cur != 'r' {
			continue closeTagLoop
		}

		ix++

		if done(ix, end) {
			return -1, -1
		}

		cur = runes[ix]

		if cur != 'o' {
			continue closeTagLoop
		}

		ix++

		if done(ix, end) {
			return -1, -1
		}

		cur = runes[ix]

		if cur != 'w' {
			continue closeTagLoop
		}

		ix++

		for ; ix < end && cur != '>'; ix++ {
			cur = runes[ix]
		}

		if done(ix, end) {
			return -1, -1
		} else if cur != '>' {
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
