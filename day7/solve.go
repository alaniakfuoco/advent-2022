package day7

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

type sysfile struct {
	name     string
	isDir    bool
	size     int
	files    []*sysfile
	location *sysfile
}

func Solve(path string) {
	data := util.GetFileLines(path)

	root := newDir("/", nil)
	currentFile := &root

	//var lsFlag bool
	for _, line := range data {
		if line != "" && line != "$ cd /" { // skip first line or empty lines
			parts := strings.Split(line, " ")
			switch parts[0] {
			case "$":
				// lsFlag = false
				switch parts[1] {
				case "cd":
					if parts[2] == ".." {
						currentFile = currentFile.location
					} else if parts[2] == "/" {
						currentFile = &root
					} else {
						for _, f := range currentFile.files {
							if f.isDir && f.name == parts[2] {
								currentFile = f
								break
							}
						}
					}
				case "ls":
					// lsFlag = true
				}
			case "dir":
				d := newDir(parts[1], currentFile)
				currentFile.files = append(currentFile.files, &d)
			default:
				size, _ := strconv.Atoi(parts[0])
				f := newFile(parts[1], currentFile, size)
				currentFile.files = append(currentFile.files, &f)
			}
		}
	}

	//printDirectory(root, 0)
	//fmt.Println()
	// smallDirs := []*sysfile{}
	// getAllSmallDirectories(root, 30000000, &smallDirs)
	//fmt.Println(len(smallDirs))
	totalSpace := 70000000
	usedSpace := getDirectorySize(root)
	requiredSpace := 30000000
	freeSpace := totalSpace - usedSpace
	threshold := requiredSpace - freeSpace

	bigDirs := []*sysfile{}
	getAllBigDirectories(root, threshold, &bigDirs)

	for _, f := range bigDirs {
		fmt.Println(f.name)
	}

	min := 4294967295
	for _, d := range bigDirs {
		size := getDirectorySize(*d)
		if size < min {
			min = size
		}
	}

	fmt.Println(min)
}

func Solve2(path string) {
	// data := util.GetFileLines(path)

	// for _, line := range data {

	// }
}

func newDir(name string, location *sysfile) sysfile {
	return sysfile{
		name:     name,
		isDir:    true,
		files:    []*sysfile{},
		location: location,
	}
}

func newFile(name string, location *sysfile, size int) sysfile {
	return sysfile{
		name:     name,
		size:     size,
		location: location,
	}
}

func getAllBigDirectories(d sysfile, threshold int, result *[]*sysfile) {
	for _, f := range d.files {
		if f.isDir {
			if getDirectorySize(*f) >= threshold {
				*result = append(*result, f)
				getAllBigDirectories(*f, threshold, result)
			}
		}
	}
}

func getAllSmallDirectories(d sysfile, threshold int, result *[]*sysfile) {
	//fmt.Println("checking dir:", d.name, "with size:", getDirectorySize(d))
	for _, f := range d.files {
		if f.isDir {
			if getDirectorySize(*f) <= threshold {
				*result = append(*result, f)
			}
			getAllSmallDirectories(*f, threshold, result)
		}
	}
}

func getDirectorySize(d sysfile) int {
	result := 0
	for _, f := range d.files {
		if f.isDir {
			result += getDirectorySize(*f)
		} else {
			result += f.size
		}
	}
	return result
}

func printDirectory(d sysfile, depth int) {
	fmt.Println(d.name, "size:", getDirectorySize(d))
	for _, f := range d.files {
		printTabs(depth + 1)
		if f.isDir {
			printDirectory(*f, depth+1)
		} else {
			fmt.Println(f.name)
		}
	}
}

func printTabs(amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print("  ")
	}
}
