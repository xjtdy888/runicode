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

func (self String) Contains(sub string) bool {
	return strings.Contains(self.stringValue, sub)
}

func (self String) Count(sub string) int {
	return strings.Count(self.stringValue, sub)
}
