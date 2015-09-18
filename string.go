package runicode

import (
	"sort"
	"strings"
	"strconv"
)

// Initially, we defined String to be of type struct { []rune, string }
// so the conversion to a string only had to happen once and could be
// stored. That was more efficient since nearly every method we implement
// uses the strings package. However, we found that it was valuable to
// preserve functionality with Go's builtin features for slices.
// If you need highly-optimized performance for *very* large strings,
// this package may not be for you. (Or you could fork it and use the struct.)

type String []rune

func New(input string) String {
	return String(input)
}

func (this String) String() string {
	return string(this)
}

func (this String) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(this))), nil
}

func (this String) Len() int {
	return len(this)
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
	switch {
	case other.Len() == 0:
		return 0
	case other.Len() == 1:
		for i := 0; i < self.Len(); i++ {
			if self[i] == other[0] {
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
		for i := 0; i <= self.Len()-other.Len(); i++ {
			found := true
			for j := 0; j < other.Len(); j++ {
				if self[i+j] != other[j] {
					found = false
				}
			}
			if found {
				return i
			}
		}
		return -1
	}
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
	return toRunicodeSlice(strings.Split(self.String(), sep.String()))
}

func (self String) SplitAfter(delim String) []String {
	return toRunicodeSlice(strings.SplitAfter(self.String(), delim.String()))
}

func (self String) Title() String {
	// From strings.Split() in the Go documentation:
	// "BUG: The rule Title uses for word boundaries does not handle Unicode punctuation properly."
	return New(strings.Title(self.String()))
}

func (self String) ToTitle() String {
	return New(strings.ToTitle(self.String()))
}

func (self String) ToLower() String {
	return New(strings.ToLower(self.String()))
}

func (self String) ToUpper() String {
	return New(strings.ToUpper(self.String()))
}

func (self String) Trim(cutset String) String {
	return New(strings.Trim(self.String(), cutset.String()))
}

func (self String) TrimPrefix(prefix String) String {
	return New(strings.TrimPrefix(self.String(), prefix.String()))
}

func (self String) TrimSuffix(suffix String) String {
	return New(strings.TrimSuffix(self.String(), suffix.String()))
}

func (self String) TrimLeft(cutset String) String {
	return New(strings.TrimLeft(self.String(), cutset.String()))
}

func (self String) TrimRight(cutset String) String {
	return New(strings.TrimRight(self.String(), cutset.String()))
}

func (self String) TrimSpace() String {
	return New(strings.TrimSpace(self.String()))
}

func toRunicodeSlice(s []string) []String {
	result := make([]String, len(s))
	for i, str := range s {
		result[i] = New(str)
	}
	return result
}

type stringSlice []String

func (self stringSlice) Less(i, j int) bool { return self[i].String() < self[j].String() }
func (self stringSlice) Len() int           { return len(self) }
func (self stringSlice) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func Sort(strings []String) []String {
	result := strings
	sort.Sort(stringSlice(result))
	return result
}
