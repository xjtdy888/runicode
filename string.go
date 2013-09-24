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
	return strings.Contains(self.stringValue, sub.String())
}

func (self String) Count(sub String) int {
	return strings.Count(self.stringValue, sub.String())
}

func (self String) Fields() []String {
	rawFields := strings.Fields(self.stringValue)
	fields := []String{}
	for _, field := range rawFields {
		fields = append(fields, New(field))
	}
	return fields
}

func (self String) HasPrefix(prefix String) bool {
	return strings.HasPrefix(self.stringValue, prefix.String())
}

func (self String) HasSuffix(prefix String) bool {
	return strings.HasSuffix(self.stringValue, prefix.String())
}

func (self String) Index(other String) int {
	for i := 0; i < self.Len(); i++ {
		startOfSubstring := self.runes[i] == sub.runes[0]
		if startOfSubstring {
			continue
		}

		found := true
		for j := 0; j < other.Len(); j++ {
			substringInterrupted := self.Len() <= i+j || other.runes[j] != self.runes[i+j]
			if substringInterrupted {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
