package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func isIp(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func isValidInput(arg string) {
	if arg == "" {
		fmt.Println("Invalid input: no args")
	} else if m, _ := regexp.MatchString("^[0-9]+$", arg); m {
		fmt.Println("Number")
	} else {
		fmt.Println("Not Number")
	}
}

func RegexExamples() {
	fmt.Println("IP check:")
	fmt.Println("127.0.0.1 isIp : ", isIp("127.0.0.1"))
	fmt.Println("12.0.0.1 isIp : ", isIp("12.0.0.1"))
	fmt.Println("1.0.0.1 isIp : ", isIp("1.0.0.1"))
	fmt.Println("1234.0.0.1 isIp : ", isIp("1234.0.0.1"))
	fmt.Println("\nInput validation:")
	fmt.Print("Input: 123 ")
	isValidInput("123")
	fmt.Print("Input: hello ")
	isValidInput("hello")
	fmt.Print("Input: h33lo ")
	isValidInput("h33lo")
}

func RegexWebParser() {
	resp, err := http.Get("https://osac.org.np")
	if err != nil {
		fmt.Println("http get error:")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	// Convert html tags to lowercase
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// Remove STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// Remove Script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// Remove all html code in < > and replace with newline
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	// Remove continuous newline
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}

func RegexFindExample() {
	str := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// Find the first match
	one := re.Find([]byte(str))
	fmt.Println("Find: ", string(one))

	// Find all matches and save to a slice, n less than 0 means return all matches
	// indicates length of slcie if its greater than 0
	all := re.FindAll([]byte(str), -1)
	fmt.Print("FindAll: ")
	for _, m := range all {
		fmt.Printf("%s, ", string(m))
	}

	// Find index of First match
	oneI := re.FindIndex([]byte(str))
	fmt.Println("\nFindIndex: ", oneI)

	// Find index of all matches, then n does the same job as above.
	allindex := re.FindAllIndex([]byte(str), -1)
	fmt.Print("FindAllIndex: ")
	for _, m := range allindex {
		fmt.Print(m, ", ")
	}

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// Find first submatch & return array, the 1st elem -> contains all elems
	// 2nd elem -> contains result of 1st
	// The first elem : "am learning Go language"
	// Second elem: " learning Go ", {NOTE: spaces too will be there}
	// Third elem: "uage"
	submatch := re2.FindSubmatch([]byte(str))
	fmt.Print("\nFindSubmatch: ")
	for _, v := range submatch {
		fmt.Print(string(v), ", ")
	}

	// same as FindIndex
	submatchindex := re2.FindSubmatchIndex([]byte(str))
	fmt.Println("\nFindSubmatchIndex: ", submatchindex)

	// FindAllSubmatches
	allsubmatches := re2.FindAllSubmatch([]byte(str), -1)
	fmt.Print("FindAllSubmatches: ")
	for _, v := range allsubmatches {
		for _, m := range v {
			fmt.Print(string(m), ", ")
		}
	}

	// FindAllSubmatchIndex, find index of all submatches.
	allSubmatchIndex := re2.FindAllSubmatchIndex([]byte(str), -1)
	fmt.Println("\nFindAllSubmatchIndex: ", allSubmatchIndex)
}

func RegexpExpandExample() {
	src := []byte(`
            call hello alice
            hello bob
            call hello eve
       `)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
