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
			result += self.stringValue
		}
		result += piece.String()
	}
	return New(result)
}
