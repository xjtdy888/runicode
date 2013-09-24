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

		Convey("When a string contains a substring", func() {
			Convey("The 'Contains' method should return true", func() {
				So(value.Contains("界!"), ShouldBeTrue)
			})
		})

		Convey("When a string does NOT contain a substring", func() {
			Convey("The 'Contains' method should return false", func() {
				So(value.Contains("!界"), ShouldBeFalse)
			})
		})

		Convey("When a string contains occurrences of a substring", func() {
			Convey("The 'Count' method should return the number of non-overlapping occurrences", func() {
				So(value.Count("l"), ShouldEqual, 2)
				So(value.Count("世"), ShouldEqual, 1)
			})
		})

		Convey("When a string does NOT contain occurrences of a substring", func() {
			Convey("The 'Count' method should return 0", func() {
				So(value.Count("Blah"), ShouldEqual, 0)
				So(value.Count("ቷቪ"), ShouldEqual, 0)
			})
		})
	})
}
