package runicode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestString(t *testing.T) {
	var input string
	var value String

	Convey("Subject: String type", t, func() {
		input = "Hello, 世界!"
		value = New(input)

		Convey("String input with Unicode should yield same output", func() {
			So(value.String(), ShouldEqual, input)
		})

		Convey("The 'Len' method should return the correct length (character count)", func() {
			So(value.Len(), ShouldEqual, 10)
		})

		Convey("When a String contains a sub-String", func() {
			Convey("The 'Contains' method should return true", func() {
				So(value.Contains(New("界!")), ShouldBeTrue)
			})
		})

		Convey("When a String does NOT contain a sub-String", func() {
			Convey("The 'Contains' method should return false", func() {
				So(value.Contains(New("!界")), ShouldBeFalse)
			})
		})

		Convey("When a String contains a certain sub-String", func() {
			Convey("The 'Count' method should return the number of non-overlapping occurrences", func() {
				So(value.Count(New("l")), ShouldEqual, 2)
				So(value.Count(New("世")), ShouldEqual, 1)
			})

			Convey("The 'Index' method should return the first correct 0-based character position", func() {
				So(value.Index(New("llo,")), ShouldEqual, 2)
				So(value.Index(New(" ")), ShouldEqual, 6)
				So(value.Index(New("世")), ShouldEqual, 7)
				So(value.Index(New("界")), ShouldEqual, 8)
				So(value.Index(New("!")), ShouldEqual, 9)
			})
		})

		Convey("When a String does NOT contain occurrences of a sub-String", func() {
			Convey("The 'Count' method should return 0", func() {
				So(value.Count(New("Blah")), ShouldEqual, 0)
				So(value.Count(New("ቷቪ")), ShouldEqual, 0)
			})
		})

		Convey("When a String is split using the 'Fields' method", func() {
			Convey("It should return an array of Strings", func() {
				So(value.Fields(), ShouldResemble, []String{New("Hello,"), New("世界!")})
			})
		})

		Convey("When a String has a given prefix", func() {
			Convey("The 'HasPrefix' method should return true", func() {
				So(value.HasPrefix(New("He")), ShouldBeTrue)
				So(value.HasPrefix(New("Hello, 世")), ShouldBeTrue)
			})
		})

		Convey("When a String does NOT have a given prefix", func() {
			Convey("The 'HasPrefix' method should return false", func() {
				So(value.HasPrefix(New("lo,")), ShouldBeFalse)
				So(value.HasPrefix(New("世界!")), ShouldBeFalse)
			})
		})

		Convey("When a String has a given suffix", func() {
			Convey("The 'HasSuffix' method should return true", func() {
				So(value.HasSuffix(New("世界!")), ShouldBeTrue)
				So(value.HasSuffix(New("!")), ShouldBeTrue)
			})
		})

		Convey("When a String does NOT have a given suffix", func() {
			Convey("The 'HasSuffix' method should return false", func() {
				So(value.HasSuffix(New("世界")), ShouldBeFalse)
				So(value.HasSuffix(New("Hello")), ShouldBeFalse)
			})
		})
	})
}
