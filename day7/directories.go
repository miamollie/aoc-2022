package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	cd       = "cd"
	ls       = "ls"
	cmd      = "$"
	rootDir  = "/"
	upALevel = ".."
	dirname  = "dir"
)

const (
	TOTAL_DISK_SPACE       = 70000000
	REQUIRED_UPGRADE_SPACE = 30000000
)

func main() {
	stacksFilePath := os.Args[1]
	stacksFile, err := os.Open(stacksFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer stacksFile.Close()

	scanner := bufio.NewScanner(stacksFile)
	root := Node{
		name:     "/",
		size:     0,
		children: make(map[string]*Node, 0),
	}

	cur := &root
	for scanner.Scan() {
		line := scanner.Text()
		entries := strings.Split(line, " ")
		if entries[0] == cmd {
			switch entries[1] {
			case cd:
				newDir := entries[2]
				if newDir == upALevel {
					cur = cur.parent
				} else {
					cur = cur.findChild(newDir)
				}
			case ls:
			}
		} else {
			// it wasn't a command, so add to the tree
			if entries[0] == dirname {
				cur.insert(entries[1], 0, true)
			} else {
				size, err := strconv.Atoi(entries[0])
				if err != nil {
					panic(err)
				}
				cur.insert(entries[1], size, false)
			}
		}
	}

	// t := root.calcTotalCappedSize(100000)

	// fmt.Printf("Total: %#v \n", t)

	remainingSpace := TOTAL_DISK_SPACE - root.size
	fmt.Printf("Root size: %v\n", root.size)
	requiredSpace := REQUIRED_UPGRADE_SPACE - remainingSpace
	fmt.Printf("Required space: %v\n", requiredSpace)

	d := root.getAllDirSizes(requiredSpace)
	sort.Ints(d[:]) //sort the array

	var smallestToDelete int
	for _, v := range d {
		if v > requiredSpace {
			smallestToDelete = v
			break
		}
	}

	fmt.Println(len(d))

	fmt.Printf("Should delete directory: %#v \n", smallestToDelete)
}

// Node represents tree with integer size and a string name.
type Node struct {
	name     string
	size     int
	children map[string]*Node
	parent   *Node
}

func (n *Node) insert(name string, size int, isDirectory bool) *Node {
	// already a child
	if c := n.findChild(name); c != nil {
		return n
	}

	var children map[string]*Node = nil
	if isDirectory {
		children = make(map[string]*Node, 0)
	}
	newN := Node{name: name, size: size, parent: n, children: children}
	n.children[name] = &newN

	//update the parent's size to include the size of children (directory children start with 0 size)
	cur := n
	for cur != nil {
		cur.size += size
		cur = cur.parent
	}
	return n
}

func (n *Node) findChild(name string) *Node {
	if v, ok := n.children[name]; ok {
		return v
	}
	return nil
}

func (n *Node) getAllDirSizes(requiredSpace int) []int {
	dirSizes := make([]int, 0)

	for _, v := range n.children {
		if v.children != nil { //don't include files, just directories
			dirSizes = append(dirSizes, v.size)
			childSizes := v.getAllDirSizes(requiredSpace)
			dirSizes = append(dirSizes, childSizes...)
		}
	}
	return dirSizes
}

func (n *Node) calcTotalCappedSize(maxDirSize int) int {
	fmt.Printf("Node %#v | size  %#v \n", n.name, n.size)
	cappedTotalSize := 0
	if n.children == nil {
		fmt.Println("Leaf node, don't add to count")
		return cappedTotalSize // no need to count children again, their size has already been added to the parent
	}

	for _, v := range n.children {
		childDirTotalSize := v.calcTotalCappedSize(maxDirSize)
		cappedTotalSize += childDirTotalSize
	}

	if n.size < maxDirSize {
		cappedTotalSize += n.size
	}

	return cappedTotalSize
}
