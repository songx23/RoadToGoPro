package stringset

import strutils "github.com/torden/go-strutil"

var strutil *strutils.StringProc

// UniqueStringSet is a special string array. It can only store unique values.
type UniqueStringSet []string

func init() {
	strutil = strutils.NewStringProc()
}

// NewUniqueStringSet creates a new UniqueStringSet. This function takes optional string arguments and append them
//to the new UniqueStringSet
func NewUniqueStringSet(str ...string) UniqueStringSet {
	set := UniqueStringSet{}
	for _, s := range str {
		set.Add(s)
	}
	return set
}

// Add appends the string argument to the UniqueStringSet if its value doesn't exist in the set.
func (s *UniqueStringSet) Add(newStr string) {
	if s.indexOf(newStr) == -1 {
		*s = append(*s, newStr)
	}
}

// Remove deletes the string argument from the UniqueStringSet if it exists in the set.
func (s *UniqueStringSet) Remove(rmvStr string) {
	rmvIdx := s.indexOf(rmvStr)
	set := *s
	if rmvIdx > -1 {
		*s = append(set[:rmvIdx], set[rmvIdx+1:]...)
	}
}

// Format takes each string in the UniqueStringSet and formatting the string so the first character of words changes
//to upper case
func (s *UniqueStringSet) Format() {
	formattedSet := UniqueStringSet{}
	for _, str := range *s {
		formattedSet = append(formattedSet, strutil.UpperCaseFirstWords(str))
	}
	*s = formattedSet
}

// indexOf returns the index of testStr in the UniqueStringSet. If testStr doesn't exist then index returned is -1
func (s *UniqueStringSet) indexOf(testStr string) int {
	for i, str := range *s {
		if str == testStr {
			return i
		}
	}
	return -1
}
