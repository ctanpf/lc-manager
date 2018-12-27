package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

var length int

func main() {
	fmt.Println("Hello")
	argsWithoutProg := os.Args[1:]
	filename := argsWithoutProg[0]
	dir := argsWithoutProg[1]
	fmt.Println("***************\nProcessing todo\n***************")
	todo := getID(filename)
	fmt.Println("***************\nProcessing done\n***************")
	done := getIDfromDir(dir)
	fmt.Println("***************\nProcessing diff\n***************")
	diff := getDiff(todo, done)
	fmt.Println("***************\nGenerating file\n***************")
	generateFile(diff, 10)
}

func generateFile(diff []int, num int) {
	var chosen []string
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(len(diff))
	for _, r := range p[:10] {
		chosen = append(chosen, strconv.Itoa(diff[r]))
	}

	for _, elem := range chosen {
		_, err := exec.Command("leetcode", "show", elem, "-gxl", "java").Output()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("leetcode show %s -gxl java\n", elem)
	}

	fmt.Println("***************\nGenerating submission file\n***************")

	for _, elem := range chosen {
		_, err := exec.Command("leetcode", "submission", elem).Output()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("leetcode submission %s\n", elem)
	}
}

func getDiff(todo []int, done []int) []int {
	var diff []int
	m := make(map[int]bool)
	for _, elem := range todo {
		m[elem] = true
	}
	for _, elem := range done {
		delete(m, elem)
	}
	for key, value := range m {
		if value {
			diff = append(diff, key)
		}
	}
	fmt.Printf("%d files found in diff\n", len(diff))
	return diff
}

func getID(filename string) []int {
	var res []int
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile(`\[(...)\]`)
		firstslice := r.FindString(line)
		r, _ = regexp.Compile(`\d+`)
		idstr := r.FindString(firstslice)
		id, _ := strconv.Atoi(idstr)
		res = append(res, id)
	}
	fmt.Printf("%d files found in todo\n", len(res))
	return res
}

func getIDfromDir(dir string) []int {
	var res []int
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		file := f.Name()
		r, _ := regexp.Compile(`.java$`)
		if r.MatchString(file) {
			r, _ = regexp.Compile(`^\d+`)
			idstr := r.FindString(file)
			id, _ := strconv.Atoi(idstr)
			res = append(res, id)
		}
	}
	fmt.Printf("%d files found in done\n", len(res))
	return res
}
