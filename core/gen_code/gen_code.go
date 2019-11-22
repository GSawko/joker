package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/candid82/joker/core"
	_ "github.com/candid82/joker/std/string"
)

type FileInfo struct {
	name     string
	filename string
}

/* The entries must be ordered such that a given namespace depends
/* only upon namespaces loaded above it. E.g. joker.template depends
/* on joker.walk, so is listed afterwards, not in alphabetical
/* order. */
var files []FileInfo = []FileInfo{
	{
		name:     "<joker.core>",
		filename: "core.joke",
	},
	// {
	// 	name:     "<joker.repl>",
	// 	filename: "repl.joke",
	// },
	// {
	// 	name:     "<joker.walk>",
	// 	filename: "walk.joke",
	// },
	// {
	// 	name:     "<joker.template>",
	// 	filename: "template.joke",
	// },
	// {
	// 	name:     "<joker.test>",
	// 	filename: "test.joke",
	// },
	// {
	// 	name:     "<joker.set>",
	// 	filename: "set.joke",
	// },
	// {
	// 	name:     "<joker.tools.cli>",
	// 	filename: "tools_cli.joke",
	// },
	// {
	// 	name:     "<joker.core>",
	// 	filename: "linter_all.joke",
	// },
	// {
	// 	name:     "<joker.core>",
	// 	filename: "linter_joker.joke",
	// },
	// {
	// 	name:     "<joker.core>",
	// 	filename: "linter_cljx.joke",
	// },
	// {
	// 	name:     "<joker.core>",
	// 	filename: "linter_clj.joke",
	// },
	// {
	// 	name:     "<joker.core>",
	// 	filename: "linter_cljs.joke",
	// },
}

const hextable = "0123456789abcdef"
const masterFile = "a_code.go"

func main() {
	codeWriterEnv := &CodeWriterEnv{
		NeedSyms:     map[*string]struct{}{},
		NeedStrs:     map[string]struct{}{},
		NeedBindings: map[*Binding]struct{}{},
		NeedKeywords: map[uint32]Keyword{},
	}

	GLOBAL_ENV.FindNamespace(MakeSymbol("user")).ReferAll(GLOBAL_ENV.CoreNamespace)
	for _, f := range files {
		fileTemplate := `// Generated by gen_code. Don't modify manually!

package core

func init() {
	{name}NamespaceInfo = internalNamespaceInfo{init: {name}Init, generated: {name}NamespaceInfo.generated, available: true}
}

{inits}
func {name}Init() {
{interns}
}
`

		GLOBAL_ENV.SetCurrentNamespace(GLOBAL_ENV.CoreNamespace)
		content, err := ioutil.ReadFile("data/" + f.filename)
		if err != nil {
			panic(err)
		}

		var inits, interns string
		inits, interns, err = CodeWriter(NewReader(bytes.NewReader(content), f.name), codeWriterEnv)
		PanicOnErr(err)

		name := f.filename[0 : len(f.filename)-5] // assumes .joke extension
		newFile := "a_" + name + "_code.go"
		if newFile <= masterFile {
			panic(fmt.Sprintf("I think Go initializes things alphabetically, so %s must come after %s due to dependencies; rename accordingly",
				newFile, masterFile))
		}
		fileContent := strings.Replace(strings.Replace(strings.ReplaceAll(fileTemplate, "{name}", name), "{inits}", inits, 1), "{interns}", interns, 1)
		ioutil.WriteFile(newFile, []byte(fileContent), 0666)
	}

	fileContent := `// Generated by gen_code. Don't modify manually!

package core

{strDefs}
{symDefs}
{kwDefs}
{bindingDefs}
func init() {
{strInterns}
{symInterns}
}
`

	bindingDefs := ""
	for b, _ := range codeWriterEnv.NeedBindings {
		bindingDefs += fmt.Sprintf(`
var binding_%p = Binding{
	name: sym_%s,
	index: %d,
	frame: %d,
	isUsed: %v,
}
`[1:],
			b, NameAsGo(*b.Name()), b.Index(), b.Frame(), b.IsUsed())

		codeWriterEnv.NeedSyms[b.Name()] = struct{}{}
	}

	symDefs := ""
	symInterns := ""
	for s, _ := range codeWriterEnv.NeedSyms {
		name := NameAsGo(*s)
		symDefs += fmt.Sprintf(`
var sym_%s = Symbol{}
`[1:],
			name)

		codeWriterEnv.NeedStrs[*s] = struct{}{}
		symInterns += fmt.Sprintf(`
	sym_%s.name = string_%s
`[1:],
			name, name)
	}

	kwDefs := ""
	for _, k := range codeWriterEnv.NeedKeywords {
		ns := "nil"
		if k.NsField() != nil {
			ns = "string_" + NameAsGo(*k.NsField())

		}
		name := "string_" + NameAsGo(*k.NameField())

		kwId := fmt.Sprintf("kw_%d", k.HashField())

		kwDefs += fmt.Sprintf(`
var %s = Keyword{
	ns: %s,
	name: %s,
}`,
			kwId, ns, name)
	}

	strDefs := ""
	strInterns := ""
	for s, _ := range codeWriterEnv.NeedStrs {
		name := NameAsGo(s)
		strDefs += fmt.Sprintf(`
var string_%s *string
`[1:],
			name)

		strInterns += fmt.Sprintf(`
	string_%s = STRINGS.Intern("%s")
`[1:],
			name, s)
	}

	var tr = [][2]string{
		{"{strDefs}", strDefs},
		{"{symDefs}", symDefs},
		{"{kwDefs}", kwDefs},
		{"{bindingDefs}", bindingDefs},
		{"{strInterns}", strInterns},
		{"{symInterns}", symInterns},
	}
	for _, t := range tr {
		fileContent = strings.Replace(fileContent, t[0], t[1], 1)
	}

	ioutil.WriteFile(masterFile, []byte(fileContent), 0666)
}
