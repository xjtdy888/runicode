runicode
========

"Roon-eh-code" is a wrapper package to let you use `[]rune` like `string` in Go. Certain operations on native strings containing Unicode characters may return unexpected results. For example, character positioning (like substring indexes given by `strings.Index()`) is actually a byte—not character—offset. Runicode treats multibyte characters as single characters, true to `[]rune` form.



Installation
-------------

Just the usual:

    $ go get github.com/smartystreets/runicode




Usage
------

First:

```go
import "github.com/smartystreets/runicode"
```

Then:

```go
str := runicode.New("My string here")  // runicode.String, which is []rune under the hood
```

Or if you prefer to cast instead:

```go
str := runicode.String("My string here")  // runicode.New() does this for you anyway
```

Some more:

```go
native := "Hello, 世界!"
runi   := runicode.New(native)

nativeLen := len(native)   // 14 (bytes)
runiLen   := len(runi)     // 10 (characters)

idx1 := strings.Index(native, "界")   // 10 (byte offset)
idx2 := runi.Index("界")              // 8  (true character offset)

// Ascending lexicographical sort is available
sortable := []String {
	runicode.New("世界!"),
	runicode.New("Hello,"),
}
sorted := runicode.Sort(sortable)

// Most common functions from the strings package are available as methods
padded := runicode.New("  Woah too much whitespace	")
padded = padded.TrimSpace()   // Runicode strings are immutable
padded = padded.ToUpper()     // Each method/function call will return a different instance

// If you prefer to use string literals, call the package funtions instead
runicode.HasSuffix(runi, "界!")
runicode.Split(runi, ", ")
runicode.Replace(runi, "Hello,", "Goodbye,")
```

Please feel free to fork and contribute to this simple project.

