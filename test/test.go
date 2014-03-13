package main

// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN
// EDIT THE .haml FILE INSTEAD

import (
	"fmt"
	"html/template"
	"io"
	"os"
)

func NewTestWriter() (*TestWriter) {
	wr := &TestWriter{
		templates: make([]*template.Template, 0, len(testTemplatePatterns)),
	}
	
	for idx, pattern := range testTemplatePatterns {
		tmpl, err := template.New("testTemplates" + string(idx)).Parse(pattern)
		if err != nil {
			fmt.Errorf("Could not parse template: %d", idx)
			panic(err)
		}
		wr.templates = append(wr.templates, tmpl)
	}
	return wr
}

type TestWriter struct {
	templates []*template.Template
}

var testHtml = [...]string{
`<html>
	<head>
		<title>
			`,
			`
		</title>
	</head>
	<body>
		<h1>
			`,
			`
			<div></div>
		</h1>
		<div>
			 This is child content for the div above. Note that HAML is space-sensitive, so all text indented at this
			level is encased in the div.
		</div>
		<div id="id_div"> You can use the # operator as a shortcut to create a div with the given id.</div>
		<div class="implicit_class">
			 The .operator (think of the '.' css selector') lets you create a div with the given class. For example
			this text will be wrapped in a div that looks like
		</div>
		`,
		`
		<ul type="disc">
			`,
			`
			<li>
				`,
				`
			</li>
			`,
			`
		</ul>
	</body>
</html>
`,
}

var testTemplatePatterns = []string{
	`Hello, {{.}}`,
	`Hello, {{.}}`,
}

func (wr *TestWriter) Execute(w io.Writer,
			data string) {
	var err error = nil
		fmt.Fprint(w, testHtml[0])
	err = wr.templates[0].Execute(w, data)
	handleTestError(err)
	fmt.Fprint(w, testHtml[1])
	err = wr.templates[1].Execute(w, data)
	handleTestError(err)
	fmt.Fprint(w, testHtml[2])
	fmt.Fprint(w, "Unescaped (and dangerous) output: <i>", data, "</i>")
	fmt.Fprint(w, testHtml[3])
	for i := 0; i < 10; i++ {
		fmt.Fprint(w, testHtml[4])
		fmt.Fprint(w, "Item: ", i)
		fmt.Fprint(w, testHtml[5])
	}
	fmt.Fprint(w, testHtml[6])

	if err != nil {
		err = nil
	}
}

func handleTestError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
