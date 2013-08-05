package main

import (
	"bytes"
	"container/list"
	"github.com/travissimon/formatting"
	"io"
	"strings"
)

type ViewWriter struct {
	context           *ParseContext
	rootNode          *Node
	writer            io.Writer
	destinationName   string
	writingCodeOutput bool
}

func NewViewWriter(wr io.Writer, context *ParseContext, rootNode *Node, writerName string) *ViewWriter {
	vw := &ViewWriter{
		context:         context,
		rootNode:        rootNode,
		writer:          wr,
		destinationName: writerName,
	}

	return vw
}

func (w *ViewWriter) WriteView() {
	src := formatting.NewIndentingWriter(w.writer)

	htmlArr, srcOut, patterns := w.processNodes()

	src.Printf("package %s\n", w.context.pkg)
	src.Println("")
	src.Println("// THIS IS A GENERATED FILE, EDITS WILL BE OVERWRITTEN")
	src.Println("// EDIT THE .haml FILE INSTEAD")
	src.Println("")
	src.Println("import (")
	src.IncrIndent()
	src.Println("\"fmt\"")
	src.Println("\"html/template\"")
	src.Println("\"net/http\"")
	for _, imp := range w.context.imports {
		src.Printf("%q\n", imp)
	}
	src.DecrIndent()
	src.Println(")")
	src.Println("")
	src.Printf("func New%sWriter() (*%sWriter) {\n", w.destinationName, w.destinationName)
	src.IncrIndent()
	src.Printf("wr := &%sWriter{}\n", w.destinationName)
	src.Println("")
	src.Printf("for idx, pattern := range %sTemplatePatterns {\n", w.destinationName)
	src.IncrIndent()
	src.Printf("tmpl, err := template.New(\"%sTemplates\" + string(idx)).Parse(pattern)\n", w.destinationName)
	src.Println("if err != nil {")
	src.IncrIndent()
	src.Println("fmt.Errorf(\"Could not parse template: %d\", idx)")
	src.Println("panic(err)")
	src.DecrIndent()
	src.Println("}")
	src.Printf("%sTemplates = append(%sTemplates, tmpl)\n", w.destinationName, w.destinationName)
	src.DecrIndent()
	src.Println("}")
	src.Println("return wr")
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("type %sWriter struct {\n", w.destinationName)
	src.IncrIndent()
	src.Printf("data %s\n", w.context.dataType)
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("func (wr *%sWriter) SetData(data interface{}) {\n", w.destinationName)
	src.IncrIndent()
	src.Printf("wr.data = data.(%s)\n", w.context.dataType)
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("var %sHtml = [...]string{\n", w.destinationName)
	src.Println(htmlArr)
	src.Println("}")
	src.Println("")
	src.Printf("var %sTemplatePatterns = []string{\n", w.destinationName)
	src.Print(patterns)
	src.Println("}")
	src.Println("")
	src.Printf("var %sTemplates = make([]*template.Template, 0, len(%sTemplatePatterns))\n", w.destinationName, w.destinationName)
	src.Println("")
	src.Printf("func (wr %sWriter) Execute(w http.ResponseWriter, r *http.Request) {\n", w.destinationName)
	src.IncrIndent()
	src.Println("wr.ExecuteData(w, r, wr.data)")
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("func (wr *%sWriter) ExecuteData(w http.ResponseWriter, r *http.Request, data %s) {\n", w.destinationName, w.context.dataType)
	src.Printf("	var err error = nil\n")
	// output from processNodes
	// This is calls to htmlArray and code generated Prints
	src.Print(srcOut)

	src.IncrIndent()
	src.Printf("if err != nil {")
	src.IncrIndent()
	src.Printf("err = nil")
	src.DecrIndent()
	src.Printf("}")
	src.DecrIndent()
	src.Println("}")
	src.Println("")
	src.Printf("func handle%sError(err error) {\n", w.destinationName)
	src.IncrIndent()
	src.Printf("if err != nil {")
	src.IncrIndent()
	src.Printf("fmt.Println(err)")
	src.DecrIndent()
	src.Printf("}")
	src.DecrIndent()
	src.Printf("}")
}

// processNodes generates code from the parsed haml nodes
// htmlArray is a string array of the raw html parts.
// src is the go source that calls html array, and then user defined code.
// usually looks like this:
//	fmt.Fprint(w, HtmlArray[0])
//	fmt.Fprint(w, "Custom string: ", data)
//  fmt.Fprint(w, HtmlArray[1]) ...
func (w *ViewWriter) processNodes() (htmlArray string, src string, patterns string) {
	htmlBuffer := bytes.NewBuffer(make([]byte, 0))
	htmlWriter := formatting.NewIndentingWriter(htmlBuffer)
	srcBuffer := bytes.NewBuffer(make([]byte, 0))
	srcWriter := formatting.NewIndentingWriter(srcBuffer)
	patternBuffer := bytes.NewBuffer(make([]byte, 0))
	patternWriter := formatting.NewIndentingWriter(patternBuffer)

	// initialise opening quote for htlmArray
	htmlWriter.Print("`")

	// initialise starting indent for src
	srcWriter.IncrIndent()

	// initialise starting indent and openning quote
	patternWriter.IncrIndent()

	htmlIndex := 0
	patternIndex := 0
	for n := w.rootNode.children.Front(); n != nil; n = n.Next() {
		htmlIndex, _ = w.writeNode(n, htmlWriter, srcWriter, patternWriter, htmlIndex, patternIndex)
	}

	// close quote for html Array
	htmlWriter.Print("`,")

	// Ensure final html is written
	srcWriter.Printf("fmt.Fprint(w, %sHtml[%d])\n", w.destinationName, htmlIndex)

	// if our last op was writing code, we need to close pattern string
	if w.writingCodeOutput {
		patternWriter.Println("`,")
	}

	htmlArray = htmlBuffer.String()
	src = srcBuffer.String()
	patterns = patternBuffer.String()
	return
}

type CodeOutputType int

const (
	Static CodeOutputType = iota
	Dynamic
	Raw
	Execution
)

// Recursive function to write parsed HAML Nodes
// We have to return a bool indicating if we have escaped any HTML (XSS protection)
// so that we know if we need to include the templating library for that function
func (w *ViewWriter) writeNode(el *list.Element, haml *formatting.IndentingWriter, src *formatting.IndentingWriter, pattern *formatting.IndentingWriter, currentHtmlIndex int, currentPatternIndex int) (htmlIndex, patternIndex int) {
	nd := el.Value.(*Node)

	htmlIndex = currentHtmlIndex
	patternIndex = currentPatternIndex

	if nd.name == "code_output_static" {
		return w.writeCodeOutput(el, haml, src, pattern, htmlIndex, patternIndex, Static)
	} else if nd.name == "code_output_dynamic" {
		return w.writeCodeOutput(el, haml, src, pattern, htmlIndex, patternIndex, Dynamic)
	} else if nd.name == "code_output_raw" {
		return w.writeCodeOutput(el, haml, src, pattern, htmlIndex, patternIndex, Raw)
	} else if nd.name == "code_execution" {
		return w.writeCodeOutput(el, haml, src, pattern, htmlIndex, patternIndex, Execution)
	}

	if w.writingCodeOutput {
		// we've finished writing code output and we're back to haml
		// so close off our pattern string
		pattern.Println("`,")
	}
	w.writingCodeOutput = false

	if nd.name != "" {
		haml.Printf("<%s", nd.name)
		if nd.id != nil {
			w.writeAttribute(nd.id, haml)
		}
		if nd.class != nil {
			w.writeAttribute(nd.class, haml)
		}

		for attrEl := nd.attributes.Front(); attrEl != nil; attrEl = attrEl.Next() {
			attr := attrEl.Value.(*nameValueStr)
			w.writeAttribute(attr, haml)
		}

		if nd.selfClosing {
			haml.Print(" />")
			if nd.text != "" {
				haml.Printf(" %s", nd.text)
			}
			haml.Println("")
			return
		} else {
			haml.Print(">")
		}
	}

	// Outputting text.

	// If tag only contains short text, add it on same line
	if w.canChildContentFitOnOneLine(nd) {
		if nd.name == "" {
			haml.Printf("%s\n", nd.text)
		} else {
			haml.Printf("%s</%s>\n", nd.text, nd.name)
		}
		return
	}

	// We either have long text, child tags or both
	// so we add it as indented child content
	haml.Println("")
	haml.IncrIndent()

	if len(nd.text) > 0 {
		w.writeLongText(nd.text, haml)
	}

	for n := nd.children.Front(); n != nil; n = n.Next() {
		htmlIndex, patternIndex = w.writeNode(n, haml, src, pattern, htmlIndex, patternIndex)
	}

	haml.DecrIndent()
	if nd.name != "" {
		haml.Printf("</%s>\n", nd.name)
	}

	return
}

var TEXT_BREAK_LENGTH = 100

func (vw *ViewWriter) writeLongText(text string, w *formatting.IndentingWriter) {
	// create index of spaces in string
	spaces := vw.getWhitespaceIndicies(text)

	// split string on space less than MAX_STRING_LENGTH
	start := 0

	for _, idx := range spaces {
		distance := idx - start
		if distance > TEXT_BREAK_LENGTH {
			w.Println(text[start:idx])
			start = idx + 1
		}
	}
	w.Println(text[start:])
}

func (vw *ViewWriter) getWhitespaceIndicies(text string) []int {
	spaces := make([]int, 0, 255)

	for i, c := range text {
		switch c {
		case ' ', '\t', '\n':
			spaces = append(spaces, i)
		}
	}

	return spaces
}

// Child content can only fit on one line when there is short
// text and no child nodes
func (w *ViewWriter) canChildContentFitOnOneLine(nd *Node) bool {
	return len(nd.text) < TEXT_BREAK_LENGTH && nd.children.Len() == 0
}

func (w *ViewWriter) writeAttribute(attribute *nameValueStr, haml *formatting.IndentingWriter) {
	haml.Printf(" %s=\"%s\"", attribute.name, attribute.value)
}

func (w *ViewWriter) writeCodeOutput(el *list.Element, haml *formatting.IndentingWriter, src *formatting.IndentingWriter, pattern *formatting.IndentingWriter, currentHtmlIndex int, currentPatternIndex int, nodeType CodeOutputType) (htmlIndex, patternIndex int) {
	nd := el.Value.(*Node)

	htmlIndex = currentHtmlIndex
	patternIndex = currentPatternIndex

	if !w.writingCodeOutput {
		// First code output node - close off haml node output:

		// end most recent haml output
		haml.Println("`,")
		// start next haml output (which will follow this code output
		haml.Println("`")

		// Add call to write html from array
		src.Printf("fmt.Fprint(w, %sHtml[%d])\n", w.destinationName, currentHtmlIndex)

		if nodeType == Static || nodeType == Dynamic {
			// Add calls to execute template. However, we need to know which
			// object to inject into template. Usually this will be 'data', but
			// can be something else inside a loop, for example, so we take the
			// object from the first dynamic element
			objectToInject := "data"
			if nodeType == Dynamic {
				objectToInject = getObjectName(nd.text)
			} else {
			LookaheadLoop:
				for n := el; n != nil; n = n.Next() {
					node := n.Value.(*Node)
					switch node.name {
					case "code_output_dynamic":
						objectToInject = getObjectName(node.text)
						break LookaheadLoop
					}
				}
			}
			src.Printf("err = %sTemplates[%d].Execute(w, %s)\n", w.destinationName, currentPatternIndex, objectToInject)
			src.Printf("handle%sError(err)\n", w.destinationName)
			// start a new pattern string
			pattern.Print("`")

			w.writingCodeOutput = true
			patternIndex++
		}

		htmlIndex++
	}

	// These stop writing patterns, so we need to close off pattern strings
	if w.writingCodeOutput && (nodeType == Raw || nodeType == Execution) {
		pattern.Println("`,")
		w.writingCodeOutput = false
	}

	// add call to print output
	switch nodeType {
	case Static:
		pattern.Print(nd.text)
	case Dynamic:
		// print the path to the desired object for the template pattern
		patternStr := "."
		index := strings.Index(nd.text, ".")
		if index > 0 {
			patternStr = nd.text[index:]
		}
		pattern.Printf("{{%s}}", patternStr)
	case Raw:
		src.Printf("fmt.Fprint(w, %s)\n", nd.text)
	case Execution:
		// attempt to keep formatting across user code.
		// Here we're checking to see if this is the end of a block statement
		// if so, we need to decrease indent
		first := getFirstChar(nd.text)
		if first == '}' {
			src.DecrIndent()
		}

		// add user's code
		src.Printf("%s\n", nd.text)

		// If user code ends in {, incr indent as they started a block statement
		last := getLastChar(nd.text)
		if last == '{' {
			src.IncrIndent()
		}
	}

	for n := nd.children.Front(); n != nil; n = n.Next() {
		htmlIndex, patternIndex = w.writeNode(n, haml, src, pattern, htmlIndex, patternIndex)
	}

	return
}

func getFirstChar(s string) byte {
	trimmed := strings.TrimLeft(s, "\t ")
	if len(trimmed) == 0 {
		return byte(0)
	}
	return trimmed[0]
}

func getLastChar(s string) byte {
	trimmed := strings.TrimRight(s, "\t ")
	if len(trimmed) == 0 {
		return byte(0)
	}
	return trimmed[len(trimmed)-1]
}

// returns the name of the object to inject
// into a template. Usually this is just 'data'
func getObjectName(s string) string {
	index := strings.Index(s, ".")
	if index < 0 {
		return s
	}
	return s[0:index]
}

// returns the path to the property to inject
// into a template.
func getObjectPath(s string) string {
	index := strings.Index(s, ".")
	if index < 0 {
		return "."
	}
	return s[index:]
}
