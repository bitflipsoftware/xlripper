package xlsx

import (
	"fmt"
	"path"
	"path/filepath"
	"time"
)

func mini(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxi(a, b int) int {
	if a > b {
		return a
	}

	return b
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

func wkbkRelsPath(wkbkPath string) (wkbkRelsPath string) {
	dir := filepath.Dir(wkbkPath)
	path := path.Join(dir, strWorkbookRels)
	return path
}

func joinWithWkbkPath(wkbkPath string, relPath string) string {
	dir := filepath.Dir(wkbkPath)
	path := path.Join(dir, removeLeadingSlash(relPath))
	return path
}

// silence the compiler complaint of an unused variable when you are trying to write programs
func use(anything interface{}) {
	if time.Now().Unix() < 0 {
		fmt.Print(anything)
	}
}
