package runicode

func Contains(s String, sub string) bool {
	return s.Contains(New(sub))
}

func Count(s String, sub string) int {
	return s.Count(New(sub))
}

func Fields(s String) []String {
	return s.Fields()
}

func HasPrefix(s String, prefix string) bool {
	return s.HasPrefix(New(prefix))
}

func HasSuffix(s String, suffix string) bool {
	return s.HasSuffix(New(suffix))
}

func Index(s String, sub string) int {
	return s.Index(New(sub))
}

/*
func Join(pieces []String) String {
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

func Repeat(times int) String {
	if times < 1 {
		return New("")
	}
	return New(strings.Repeat(self.String(), times))
}

func Replace(old, new String, limit int) String {
	return New(strings.Replace(self.String(), old.String(), new.String(), limit))
}

func Split(sep String) []String {
	return toRunicodeSlice(strings.Split(self.String(), sep.String()))
}

func SplitAfter(delim String) []String {
	return toRunicodeSlice(strings.SplitAfter(self.String(), delim.String()))
}

func Title() String {
	// TODO ... from strings.Split() in the Go documentation:
	// "BUG: The rule Title uses for word boundaries does not handle Unicode punctuation properly."
	return New(strings.Title(self.String()))
}

func ToTitle() String {
	return New(strings.ToTitle(self.String()))
}

func ToLower() String {
	return New(strings.ToLower(self.String()))
}

func ToUpper() String {
	return New(strings.ToUpper(self.String()))
}

func Trim(cutset String) String {
	return New(strings.Trim(self.String(), cutset.String()))
}

func TrimPrefix(prefix String) String {
	return New(strings.TrimPrefix(self.String(), prefix.String()))
}

func TrimSuffix(suffix String) String {
	return New(strings.TrimSuffix(self.String(), suffix.String()))
}

func TrimLeft(cutset String) String {
	return New(strings.TrimLeft(self.String(), cutset.String()))
}

func TrimRight(cutset String) String {
	return New(strings.TrimRight(self.String(), cutset.String()))
}

func TrimSpace() String {
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

func (self stringSlice) Less(i, j int) bool { return self[i].stringValue < self[j].stringValue }
func (self stringSlice) Len() int           { return len(self) }
func (self stringSlice) Swap(i, j int)      { self[i], self[j] = self[j], self[i] }

func Sort(strings []String) []String {
	result := strings
	sort.Sort(stringSlice(result))
	return result
}*/
