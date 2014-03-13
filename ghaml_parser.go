package main

import (
	"log"
)

type ParseContext struct {
	pkg     string
	data    []string
    types   map[string]string
	imports []string
}

type GhamlParser struct {
	name    string
	lineNo  int
	lexer   *lexer
	root    *Node
	context *ParseContext

	// used during construction of the tree
	tags *stack
}

// initialises and returns a new Ghaml parser
func NewParser(name, input string) *GhamlParser {
	g := &GhamlParser{
		name:   name,
		lexer:  lex(name, input),
		lineNo: 0,
		tags:   new(stack),
		context: &ParseContext{
			pkg:     "main",
			data:    make([]string, 0),
			types:   make(map[string]string),
			imports: make([]string, 0),
		},
	}

	g.root = NewNode("root")
	tagIndent := &tagIndentation{
		node:        g.root,
		indentLevel: 0,
	}
	g.tags.push(tagIndent)

	return g
}

// stack to track tag indentation for closing tags
type tagIndentation struct {
	node        *Node
	indentLevel int
}

// adds an import to the context
func (g *GhamlParser) addImport(imp string) {
	g.context.imports = append(g.context.imports, imp)
}

// sets the 'data' parameter's type
func (g *GhamlParser) setDataType(dataName string, dataType string) {
	g.context.data = append(g.context.data, dataName)
	g.context.types[dataName] = dataType
}

// setPackage sets the package for the generated code
func (g *GhamlParser) setPackage(pkgName string) {
	g.context.pkg = pkgName
}

// gets the current indent tag
func (g *GhamlParser) getCurrentTagIndent() *tagIndentation {
	return g.tags.peek().(*tagIndentation)
}

// gets the current parse node
func (g *GhamlParser) getCurrentNode() *Node {
	return g.getCurrentTagIndent().node
}

// gets the number of spaces in the current indent
func (g *GhamlParser) getCurrentIndentation() int {
	return g.getCurrentTagIndent().indentLevel
}

// initiates the parsing process
func (g *GhamlParser) Parse() {
	dataName := ""
	for {
		lexeme := g.lexer.nextItem()

		switch lexeme.typ {
		case itemAttributeName:
			g.handleAttribute(lexeme.val)
		case itemImport:
			g.addImport(lexeme.val)
		case itemDataName:
			dataName = lexeme.val
		case itemDataType:
			g.setDataType(dataName, lexeme.val)
		case itemPackage:
			g.setPackage(lexeme.val)
		case itemDoctype:
			g.handleDoctype(lexeme)
		case itemCodeOutputLiteral, itemCodeOutputValue, itemCodeOutputRaw, itemCodeExecution:
			g.handleCodeOutput(lexeme)
		case itemIndentation:
			indentation := lexeme.val
			nextLexeme := g.lexer.nextItem()
			g.parseLineStart(indentation, nextLexeme)
		case itemSelfClosingTagIdentifier:
			g.getCurrentNode().selfClosing = true
		case itemText:
			g.getCurrentNode().text += lexeme.val
		case itemId:
			g.getCurrentNode().setId(lexeme.val)
		case itemClass:
			g.getCurrentNode().addClass(lexeme.val)
		case itemEOF:
			return
		}
	}
}

// parses an attribute
func (g *GhamlParser) handleAttribute(attributeName string) {
	nextLexeme := g.lexer.nextItem()
	if nextLexeme.typ != itemAttributeValue {
		doError(g.lexer.lineNumber(), "expected attribute value, received "+nextLexeme.val)
	}
	value := nextLexeme.val
	attr := &nameValueStr{
		name:  attributeName,
		value: value,
	}
	g.getCurrentNode().addAttribute(attr)
}

// parses a doctype (!!!)
func (g *GhamlParser) handleDoctype(l lexeme) {
	n := NewNode("doctype")
	n.text = l.val
	n.selfClosing = true
	g.getCurrentNode().appendNode(n)
}

// parses a code output token (= ...)
func (g *GhamlParser) handleCodeOutput(l lexeme) {
	switch l.typ {
	case itemCodeOutputLiteral:
		g.buildCodeNode("code_output_literal", l)
	case itemCodeOutputValue:
		g.buildCodeNode("code_output_value", l)
	case itemCodeOutputRaw:
		g.buildCodeNode("code_output_raw", l)
	case itemCodeExecution:
		g.buildCodeNode("code_execution", l)
	}
}

func (g *GhamlParser) buildCodeNode(nodeName string, l lexeme) {
	n := NewNode(nodeName)
	n.text = l.val
	g.getCurrentNode().appendNode(n)
}

// parses significant whitespace, and then handles the first node on the line
func (g *GhamlParser) parseLineStart(indentation string, firstItem lexeme) {
	indentationAmnt := len(indentation)

	for g.tags.count() > 1 && g.getCurrentIndentation() >= indentationAmnt {
		g.tags.pop()
	}

	if firstItem.typ == itemText && g.getCurrentNode().children.Len() == 0 {
		// text continuing from a previous line
		g.getCurrentNode().text += " " + firstItem.val
		return
	}

	var firstNode *Node

	switch firstItem.typ {
	case itemTag:
		firstNode = NewNode(firstItem.val)
	case itemText:
		// text on a new line, but not a continuation
		firstNode = NewNode("")
		firstNode.text = firstItem.val
	case itemCodeOutputLiteral:
		firstNode = NewNode("code_output_literal")
		firstNode.text = firstItem.val
	case itemCodeOutputValue:
		firstNode = NewNode("code_output_value")
		firstNode.text = firstItem.val
	case itemCodeOutputRaw:
		firstNode = NewNode("code_output_raw")
		firstNode.text = firstItem.val
	case itemCodeExecution:
		firstNode = NewNode("code_execution")
		firstNode.text = firstItem.val
	default:
		firstNode = NewNode("div")
	}

	tagIndentation := buildTagIndentation(firstNode, indentation)
	g.getCurrentNode().appendNode(tagIndentation.node)
	g.tags.push(tagIndentation)

	// if this was an implicit div, we stll need to add the id or class data
	switch firstItem.typ {
	case itemId:
		g.getCurrentNode().setId(firstItem.val)
	case itemClass:
		g.getCurrentNode().addClass(firstItem.val)
	}
}

func doError(lineNo int, msg string) {
	log.Printf("error line %q: %q", lineNo, msg)
}

func buildTagIndentation(n *Node, indentation string) *tagIndentation {
	return &tagIndentation{
		node:        n,
		indentLevel: len(indentation),
	}
}

// utility function to create a string dump of the content
func (g *GhamlParser) dumpNodes() string {
	result := ""

	for n := g.root.children.Front(); n != nil; n = n.Next() {
		result = n.Value.(*Node).dumpNode(result, 0)
	}

	return result
}
