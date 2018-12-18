// This file is generated by generate-std.joke script. Do not edit manually!

package html

import (
	. "github.com/candid82/joker/core"
)

var htmlNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.html"))

var escape_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := html.EscapeString(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var unescape_ Proc = func(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := html.UnescapeString(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func init() {

	htmlNamespace.ResetMeta(MakeMeta(nil, "Provides functions for escaping and unescaping HTML text.", "1.0"))

	htmlNamespace.InternVar("escape", escape_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Escapes special characters like < to become &lt;. It escapes only five such characters: <, >, &, ' and ".`, "1.0"))

	htmlNamespace.InternVar("unescape", unescape_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Unescapes entities like &lt; to become <.`, "1.0"))

}
