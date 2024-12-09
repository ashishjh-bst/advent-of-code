package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type Disk struct {
	nextFreeBlock int
	storage       []interface{}
}

func Part1(input *string) string {
	disk := processInput(input)
	disk.compress()
	sum := disk.checksum()
	return fmt.Sprintf("%d", sum)
}

func Part2(input *string) string {
	disk := processInput(input)
	disk.defragmentedCompress()
	sum := disk.checksum()
	return fmt.Sprintf("%d", sum)
}

func processInput(input *string) *Disk {
	d := &Disk{}
	d.storage = make([]interface{}, 0)
	disk := strings.Split(*input, "")
	fileId := 0
	for i, j := range disk {
		intJ, _ := strconv.Atoi(j)
		isFile := i%2 == 0
		if isFile {
			for b := 0; b < intJ; b++ {
				d.storage = append(d.storage, fileId)
			}
			fileId++
		} else {
			for b := 0; b < intJ; b++ {
				d.storage = append(d.storage, ".")
			}
		}
	}
	return d
}

func (d *Disk) compress() {
	for i := len(d.storage) - 1; i >= 0; i-- {
		block := d.storage[i]
		switch block.(type) {
		case string:
			continue
		case int:
			d.getNextFreeBlock()
			if d.nextFreeBlock > i {
				return
			}
			d.storage[d.nextFreeBlock] = block
			d.storage[i] = "."
		}
	}
}

func (d *Disk) getFileStart(end int) (start int) {
	fileID := d.storage[end]
	for i := end; i > 0; i-- {
		if d.storage[i-1] != fileID {
			start = i
			return
		}
	}
	start = 0
	return
}

func (d *Disk) defragmentedCompress() {
	for i := len(d.storage) - 1; i >= 0; i-- {
		block := d.storage[i]
		switch block.(type) {
		case string:
			continue
		case int:
			fs := d.getFileStart(i)
			size := (i - fs) + 1
			bs := d.getNextFreeBlockBySize(size)
			if bs != -1 && bs < fs {
				d.moveFile(fs, bs, size)
			}
			i = fs
		}
	}
}

func (d *Disk) getNextFreeBlockBySize(size int) int {
	var currentBlockSize int
	for i := 0; i < len(d.storage); i++ {
		block := d.storage[i]
		switch block.(type) {
		case string:
			currentBlockSize++
			if currentBlockSize == size {
				return (i - currentBlockSize) + 1
			}
		case int:
			currentBlockSize = 0
		}
	}
	return -1
}

func (d *Disk) moveFile(fs, bs, size int) {
	for i := 0; i < size; i++ {
		d.storage[bs+i] = d.storage[fs+i]
		d.storage[fs+i] = "."
	}
}

func (d *Disk) getNextFreeBlock() {
	for i := d.nextFreeBlock; i < len(d.storage); i++ {
		switch d.storage[i].(type) {
		case string:
			d.nextFreeBlock = i
			return
		case int:
			continue
		}
	}
	d.nextFreeBlock = -1
}

func (d *Disk) checksum() int {
	var sum int
	for i := 0; i < len(d.storage); i++ {
		block := d.storage[i]
		switch t := block.(type) {
		case int:
			sum += i * t
		case string:
			continue
		}
	}
	return sum
}
