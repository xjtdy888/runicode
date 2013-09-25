package runicode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRunicodePackageFunctions(t *testing.T) {
	var input string
	var value String

	Convey("Subject: Runicode package functions", t, func() {
		input = "Hello, 世界!"
		value = New(input)

		Convey("When a String contains a certain sub-String", func() {
			Convey("The 'Contains' function should return true", func() {
				So(Contains(value, "界!"), ShouldBeTrue)
			})

			Convey("The 'Count' function should return the number of non-overlapping occurrences", func() {
				So(Count(value, "l"), ShouldEqual, 2)
				So(Count(value, "世"), ShouldEqual, 1)
			})

			Convey("The 'Index' function should return the first correct 0-based character position", func() {
				So(Index(value, "llo,"), ShouldEqual, 2)
				So(Index(value, " "), ShouldEqual, 6)
				So(Index(value, "世"), ShouldEqual, 7)
				So(Index(value, "界"), ShouldEqual, 8)
				So(Index(value, "!"), ShouldEqual, 9)
			})
		})

		Convey("When a String does NOT contain occurrences of a sub-String", func() {
			Convey("The 'Contains' function should return false", func() {
				So(Contains(value, "!界"), ShouldBeFalse)
			})

			Convey("The 'Count' function should return 0", func() {
				So(Count(value, "Blah"), ShouldEqual, 0)
				So(Count(value, "ቷቪ"), ShouldEqual, 0)
			})
		})

		Convey("When a String is split using the 'Fields' function", func() {
			Convey("It should return an array of Strings", func() {
				So(Fields(value), ShouldResemble, []String{New("Hello,"), New("世界!")})
			})
		})

		Convey("When a String has a given prefix", func() {
			Convey("The 'HasPrefix' function should return true", func() {
				So(HasPrefix(value, "He"), ShouldBeTrue)
				So(HasPrefix(value, "Hello, 世"), ShouldBeTrue)
			})
		})

		Convey("When a String does NOT have a given prefix", func() {
			Convey("The 'HasPrefix' function should return false", func() {
				So(HasPrefix(value, "lo,"), ShouldBeFalse)
				So(HasPrefix(value, "世界!"), ShouldBeFalse)
			})
		})

		Convey("When a String has a given suffix", func() {
			Convey("The 'HasSuffix' function should return true", func() {
				So(HasSuffix(value, "世界!"), ShouldBeTrue)
				So(HasSuffix(value, "!"), ShouldBeTrue)
			})
		})

		Convey("When a String does NOT have a given suffix", func() {
			Convey("The 'HasSuffix' function should return false", func() {
				So(HasSuffix(value, "世界"), ShouldBeFalse)
				So(HasSuffix(value, "Hello"), ShouldBeFalse)
			})
		})

		Convey("When the 'Index' function is called with an empty sub-String", func() {
			Convey("It should return an index of 0", func() {
				So(Index(value, ""), ShouldEqual, 0)
			})
		})

		Convey("When a []String is joined by a concatenator String", func() {
			pieces := []String{New("Hello"), New("世界!")}
			sep := ", "

			Convey("Each element in the []String is joined by the concatenator", func() {
				So(Join(pieces, sep), shouldEqual, value)
			})
		})

		Convey("When the 'Repeat' function is called with a count less than 1", func() {
			Convey("It should return an empty String", func() {
				So(Repeat(value, 0), shouldEqual, New(""))
				So(Repeat(value, -1), shouldEqual, New(""))
			})
		})

		Convey("When the 'Repeat' function is called with a count of at least 1", func() {
			Convey("It should return that string repeated exactly that many times", func() {
				So(Repeat(value, 3), shouldEqual, New("Hello, 世界!Hello, 世界!Hello, 世界!"))
			})
		})

		Convey("When the 'Replace' function is called", func() {
			Convey("Instances of the sub-String should be replaced", func() {
				So(Replace(value, "Hello", "Goodbye", -1), shouldEqual, New("Goodbye, 世界!"))
				So(Replace(value, "Hello", "Goodbye", 0), shouldEqual, New("Hello, 世界!"))
				So(Replace(value, "l", "y", 1), shouldEqual, New("Heylo, 世界!"))
			})
		})

		Convey("When the 'Split' function is called", func() {
			Convey("The String should be split into []String by a separator", func() {
				So(Split(value, " "), ShouldResemble, []String{New("Hello,"), New("世界!")})
				So(Split(value, ", "), ShouldResemble, []String{New("Hello"), New("世界!")})
				So(Split(value, "  "), ShouldResemble, []String{New("Hello, 世界!")})
			})
		})

		Convey("When the 'SplitAfter' function is called", func() {
			Convey("The String should be split into []String after each separator", func() {
				So(SplitAfter(value, ","), ShouldResemble, []String{New("Hello,"), New(" 世界!")})
				So(SplitAfter(value, "界"), ShouldResemble, []String{New("Hello, 世界"), New("!")})
			})
		})

		Convey("When the 'Title' function is called", func() {
			Convey("The returned String should be title-cased (first letter of words capitalized)", func() {
				So(Title("hello, 世界!"), shouldEqual, value)
				So(Title("テストtitle"), shouldEqual, New("テストtitle"))
				So(Title("he llO, 世界!"), shouldEqual, New("He LlO, 世界!"))
			})
		})

		Convey("When the 'ToTitle' function is called", func() {
			Convey("The resulting String should be all title-cased", func() {
				So(ToTitle("hello, 世界!"), shouldEqual, New("HELLO, 世界!"))
			})
		})

		Convey("When the 'ToLower' function is called", func() {
			Convey("The resulting String should be all lower-cased", func() {
				So(ToLower("HELLO, 世界!"), shouldEqual, New("hello, 世界!"))
			})
		})

		Convey("When the 'ToUpper' function is called", func() {
			Convey("The resulting String should be all upper-cased", func() {
				So(ToUpper("hello, 世界!"), shouldEqual, New("HELLO, 世界!"))
			})
		})

		Convey("When trimming a String of certain characters", func() {
			Convey("Those characters should be removed from both ends of the string", func() {
				So(Trim(value, "H界!"), shouldEqual, New("ello, 世"))
			})
			Convey("And trimming only from the left, trimming should only happen on the left", func() {
				So(TrimLeft(value, "H界!"), shouldEqual, New("ello, 世界!"))
			})
			Convey("And trimming only from the right, trimming should only happen on the right", func() {
				So(TrimRight(value, "H界!"), shouldEqual, New("Hello, 世"))
			})
		})

		Convey("When trimming a prefix from a String", func() {
			Convey("If the string has the prefix, it should be stripped", func() {
				So(TrimPrefix(value, "Hello, "), shouldEqual, New("世界!"))
			})
			Convey("If the string does NOT has the prefix, it should not be changed", func() {
				So(TrimPrefix(value, "世界!"), shouldEqual, value)
			})
		})

		Convey("When trimming a suffix from a String", func() {
			Convey("If the string has the suffix, it should be stripped", func() {
				So(TrimSuffix(value, "世界!"), shouldEqual, New("Hello, "))
			})
			Convey("If the string does NOT has the suffix, it should not be changed", func() {
				So(TrimSuffix(value, "Hello,"), shouldEqual, value)
			})
		})
	})
}
