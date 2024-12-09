package day09

import (
	"aoc/utils"
	"slices"
	"strings"
)

type File struct {
	ID   int
	Size int
}

var filesystem1, filesystem2 []File

func init() {
	listOfNumbers := utils.ReadLines("res/day09.in")[0]
	// Extra 0 on diskmap helps for part 2 data preparation
	diskmap := strings.TrimSpace(listOfNumbers) + "0"

	for id := 0; id*2 < len(diskmap); id++ {
		size, free := int(diskmap[id*2]-'0'), int(diskmap[id*2+1]-'0')
		// For file system 1, keep track of every block (size is always 1),
		// assigning the id of free spaces as -1
		filesystem1 = append(filesystem1, slices.Repeat([]File{{id, 1}}, size)...)
		filesystem1 = append(filesystem1, slices.Repeat([]File{{-1, 1}}, free)...)
		// For file system 2, compact the representation
		// by keeping track of size of the block, still
		// assigning the id of free spaces as -1
		filesystem2 = append(filesystem2, File{id, size}, File{-1, free})
	}
}

func run(filesystem []File) (checksum int) {
	for file := len(filesystem) - 1; file >= 0; file-- {
		for free := 0; free < file; free++ {
			if filesystem[file].ID != -1 && filesystem[free].ID == -1 && filesystem[free].Size >= filesystem[file].Size {
				filesystem = defragment(filesystem, free, file)
			}
		}
	}
	return calculateChecksum(filesystem)
}

func defragment(filesystem []File, free int, file int) []File {
	// Insert file into free space
	filesystem = slices.Insert(filesystem, free, filesystem[file])
	// Recalibrate amount of free space
	// For part 1, it will be blocks of size 0 that will NOT affect checksum
	// For part 2, it might be blocks of small size that will affect checksum
	filesystem[file+1].ID, filesystem[free+1].ID = -1, -1
	filesystem[free+1].Size = filesystem[free+1].Size - filesystem[file+1].Size
	return filesystem
}

func calculateChecksum(filesystem []File) (checksum int) {
	index := 0
	for _, file := range filesystem {
		// For the size of the file, update index and possibly the checksum
		for range file.Size {
			// If it is not a free space, update checksum
			if file.ID != -1 {
				checksum += index * file.ID
			}
			index++
		}
	}
	return checksum
}

func Part1() int {
	return run(filesystem1)
}

func Part2() int {
	return run(filesystem2)
}
