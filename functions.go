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

func Join(pieces []String, sep string) String {
	return New(sep).Join(pieces)
}

func Repeat(s String, times int) String {
	return s.Repeat(times)
}

func Replace(s String, old, new string, limit int) String {
	return s.Replace(New(old), New(new), limit)
}

func Split(s String, sep string) []String {
	return s.Split(New(sep))
}

func SplitAfter(s String, delim string) []String {
	return s.SplitAfter(New(delim))
}

func Title(s string) String {
	return New(s).Title()
}

func ToTitle(s string) String {
	return New(s).ToTitle()
}

func ToLower(s string) String {
	return New(s).ToLower()
}

func ToUpper(s string) String {
	return New(s).ToUpper()
}

func Trim(s String, cutset string) String {
	return s.Trim(New(cutset))
}

func TrimPrefix(s String, prefix string) String {
	return s.TrimPrefix(New(prefix))
}

func TrimSuffix(s String, suffix string) String {
	return s.TrimSuffix(New(suffix))
}

func TrimLeft(s String, cutset string) String {
	return s.TrimLeft(New(cutset))
}

func TrimRight(s String, cutset string) String {
	return s.TrimRight(New(cutset))
}
