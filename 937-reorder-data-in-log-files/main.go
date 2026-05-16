package main

import (
	"regexp"
	"slices"
	"strings"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func reorderLogFiles(logs []string) []string {

	sortedLogs := make([]string, len(logs))

	rightIDX := len(logs) - 1
	leftIDX := 0
	for i := rightIDX; i >= 0; i-- {
		leftIDX, rightIDX = insert(sortedLogs, logs[i], leftIDX, rightIDX)
	}

	slices.SortFunc(sortedLogs[:leftIDX], compare)

	return sortedLogs

}

// Just a wrapper to reduce indentation
// returns new left index and right index
func insert(slice []string, entry string, leftIDX, rightIDX int) (int, int) {
	if isTextEntry(entry) {
		return insertText(slice, entry, leftIDX), rightIDX
	} else if isNumEntry(entry) {
		return leftIDX, insertNum(slice, entry, rightIDX)
	} else {
		return leftIDX, rightIDX
	}
}

// Insertion sort from left to right
func insertText(slice []string, entry string, idx int) int {
	slice[idx] = entry
	return idx + 1

}

// Insert into first space right to left
func insertNum(slice []string, entry string, idx int) int {
	slice[idx] = entry
	return idx - 1
}

// left < right : -1
// left > right : +1
// left = right : 0
func compare(left string, right string) int {

	leftid, leftlog, _ := strings.Cut(left, " ")
	rightid, rightlog, _ := strings.Cut(right, " ")

	if logCompare := strings.Compare(leftlog, rightlog); logCompare != 0 {
		return logCompare
	}

	if idCompare := strings.Compare(leftid, rightid); idCompare != 0 {
		return idCompare
	}

	return 0
}

func isTextEntry(entry string) bool {
	slice := strings.Split(entry, " ")

	// If the slice only has the first entry, the identifier, then we
	// will treat it as a text entry for now
	if len(slice) < 2 {
		return true
	}

	return !digitCheck.MatchString(slice[1])
}

func isNumEntry(entry string) bool {
	slice := strings.Split(entry, " ")

	// If the slice only has the first entry, the identifier, then we
	// will treat it as a text entry for now
	if len(slice) < 2 {
		return false
	}

	return digitCheck.MatchString(slice[1])
}
