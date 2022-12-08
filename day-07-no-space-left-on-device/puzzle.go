package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const GOTOROOT = "goto-root"
const GOTOCHILDDIR = "goto-child-dir"
const GOTOPARENTDIR = "goto-parent-dir"
const FILEDETAILS = "file-details"
const VOID = "void"

type file struct {
	name string
	size int
}
type dir struct {
	name         string
	files        []file
	ptrParentDir *dir
	ptrChildDirs []*dir
	is_root      bool
}

func main() {
	fmt.Println("Result of part-1: ", solvePart1(100000, []string{}))
	fmt.Println("Result of part-2: ", solvePart2(70000000, 30000000, []string{}))
}

func solvePart1(threshold int, lines []string) (totalSize int) {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	var rootDir dir = parseOutput(lines)
	var dirs []dir = collectDirs(rootDir)

	for _, dir := range dirs {
		if getDirSize(dir) < threshold {
			totalSize += getDirSize(dir)
		}
	}

	return
}

func solvePart2(discSize int, requiredSpace int, lines []string) int {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	var rootDir dir = parseOutput(lines)

	return freeableSize(discSize, requiredSpace, rootDir)
}

func freeableSize(discSize int, requiredSpace int, rootDir dir) (size int) {
	var allDirs []dir = collectDirs(rootDir)
	var usedSpace int = getDirSize(rootDir)
	var unusedSpace int = discSize - usedSpace
	var minSize int = requiredSpace - unusedSpace

	for _, dir := range allDirs {
		dirSize := getDirSize(dir)

		if dirSize > minSize && (size == 0 || size > dirSize) {
			size = dirSize
		}
	}

	return
}

func collectDirs(dir dir) (dirs []dir) {
	for _, ptrChildDir := range dir.ptrChildDirs {
		dirs = append(dirs, collectDirs(*ptrChildDir)...)
	}

	return append(dirs, dir)
}

func getDirSize(dir dir) (size int) {
	for _, file := range dir.files {
		size += file.size
	}

	for _, ptrChildDir := range dir.ptrChildDirs {
		size += getDirSize(*ptrChildDir)
	}

	return
}

func parseOutput(lines []string) dir {
	var rootDir dir = buildRootDir()
	var ptrActiveDir *dir = &rootDir

	var command, details string
	for _, line := range lines {
		command, details = parseLine(line)

		switch command {
		case GOTOROOT:
			ptrActiveDir = &rootDir
		case GOTOCHILDDIR:
			ptrNewDir := goToChildDir(ptrActiveDir, details)
			ptrActiveDir = ptrNewDir
		case GOTOPARENTDIR:
			ptrActiveDir = goToParentDir(ptrActiveDir)
		case FILEDETAILS:
			addFile(ptrActiveDir, details)
		case VOID:

		default:
			fmt.Println("unknown command: '", command, "'")
			panic("unsupported command")
		}
	}

	return rootDir
}

func addFile(ptrCurrentDir *dir, details string) {
	var file file

	parts := strings.Split(details, " ")
	file.size = conv_str_to_i(parts[0])
	file.name = parts[1]

	ptrCurrentDir.files = append(ptrCurrentDir.files, file)

	return
}

func goToChildDir(ptrCurrentDir *dir, dirname string) *dir {
	var newDir dir
	newDir.name = dirname
	newDir.is_root = false
	newDir.ptrParentDir = ptrCurrentDir

	ptrCurrentDir.ptrChildDirs = append(ptrCurrentDir.ptrChildDirs, &newDir)

	return &newDir
}

func goToParentDir(ptrCurrentDir *dir) *dir {
	return ptrCurrentDir.ptrParentDir
}

func parseLine(line string) (command string, details string) {
	details = line

	switch {
	case line == "$ cd /":
		command = GOTOROOT
	case line == "$ cd ..":
		command = GOTOPARENTDIR
	case line == "$ ls":
		command = VOID
	case line == "dir ":
		command = VOID
	case strings.Index(line, "$ cd ") == 0:
		command = GOTOCHILDDIR
		details = string(line[5:])
	case strings.Index(line, "dir ") == 0:
		command = VOID
	case strings.Index(line, " ") > 0:
		// Hope this is always a line with file details
		command = FILEDETAILS
	default:
		fmt.Println("unknown output: '", line, "'")
		panic("unexpected terminal output")
	}

	return
}

func buildRootDir() (root dir) {
	root.name = "root"
	root.is_root = true

	return
}

func read_file_as_blocks() (blocks [][]string) {
	var block_inputs []string = strings.Split(read_file(), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
}

func read_file_as_lines() []string {
	return strings.Split(read_file(), "\n")
}

func read_file() (content string) {
	raw, err := os.ReadFile("input.txt")
	check(err)
	content = strings.TrimSuffix(string(raw), "\n")
	return
}

func conv_str_to_i(s string) (i int) {
	i, err := strconv.Atoi(s)
	check(err)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
