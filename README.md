runicode
========

"Roon-eh-code" lets you use `[]rune` like `string` in Go. Certain operations on native strings containing Unicode characters may return unexpected results. For example, character positioning (like substring indexes given by `strings.Index()`) is actually a byte offset. Runicode treats multibyte characters as single characters, true to `[]rune` form.