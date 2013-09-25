package runicode

import (
	"strings"
)

type String struct {
	runes       []rune
	stringValue string
}

func New(input string) String {
	return String{[]rune(input), input}
}

func (self String) String() string {
	return self.stringValue
}

func (self String) Len() int {
	return len(self.runes)
}

func (self String) Contains(sub String) bool {
	return strings.Contains(self.String(), sub.String())
}

func (self String) Count(sub String) int {
	return strings.Count(self.String(), sub.String())
}

func (self String) Fields() []String {
	rawFields := strings.Fields(self.String())
	fields := []String{}
	for _, field := range rawFields {
		fields = append(fields, New(field))
	}
	return fields
}

func (self String) HasPrefix(prefix String) bool {
	return strings.HasPrefix(self.String(), prefix.String())
}

func (self String) HasSuffix(prefix String) bool {
	return strings.HasSuffix(self.String(), prefix.String())
}

func (self String) Index(other String) int {

	/*
		TODO -- Look into using array splices like here: http://golang.org/src/pkg/strings/strings.go?s=3598:3631#L161


		switch {
		case other.Len() == 0:
			return 0
		case other.Len() == 1:
			for i := 0; i < self.Len(); i++ {
				if self.runes[i] == other.runes[0] {
					return i
				}
			}
			return -1
		case self.Len() == other.Len():
			if self.String() == other.String() {
				return 0
			}
			return -1
		case other.Len() > self.Len():
			return -1
		default:
			for i := 0; i < self.Len()-other.Len(); i++ {
				if self.runes[i] == other.runes[0] && self.runes[i:i+other.Len()] == other {
					return i
				}
			}
		}*/

	if other.Len() == 0 {
		return 0
	}

	for i := 0; i < self.Len(); i++ {
		if self.runes[i] == other.runes[0] {

			found := true
			for j := 0; j < other.Len(); j++ {
				if self.Len() <= i+j || other.runes[j] != self.runes[i+j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}

func (self String) Join(pieces []String) String {
	result := ""
	for i, piece := range pieces {
		if i > 0 {
			result += self.String()
		}
		result += piece.String()
	}
	return New(result)
}

// TODO: LastIndex()

func (self String) Repeat(times int) String {
	if times < 1 {
		return New("")
	}
	return New(strings.Repeat(self.String(), times))
}

func (self String) Replace(old, new String, limit int) String {
	return New(strings.Replace(self.String(), old.String(), new.String(), limit))
}

func (self String) Split(sep String) []String {
	rawResults := strings.Split(self.String(), sep.String())
	results := make([]String, len(rawResults))
	for i, str := range rawResults {
		results[i] = New(str)
	}
	return results
}
