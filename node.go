package main

import (
	"container/list"
)

// Represents ... wait for it... name and value strings
type nameValueStr struct {
	name  string
	value string
}

// A parse node (roughly translates directly to an html tag)
type Node struct {
	name        string
	text        string
	id          *nameValueStr // treat id and class differently so that we can
	class       *nameValueStr // write them to the front of the attr list
	attributes  *list.List
	children    *list.List
	selfClosing bool
}

func NewNode(name string) *Node {
	return &Node{
		name:        name,
		text:        "",
		id:          nil,
		class:       nil,
		attributes:  new(list.List),
		children:    new(list.List),
		selfClosing: false,
	}
}

// appends a node to the children list
func (n *Node) appendNode(NodeToAdd *Node) {
	n.children.PushBack(NodeToAdd)
}

// sets the id of the node
func (n *Node) setId(id string) {
	n.id = &nameValueStr{
		name:  "id",
		value: id,
	}
}

// adds a class to the current node
func (n *Node) addClass(class string) {
	if n.class == nil {
		n.class = &nameValueStr{
			name:  "class",
			value: "",
		}
	}

	if len(n.class.value) > 0 {
		n.class.value += " "
	}

	n.class.value += class
}

// adds an attribute to this node
func (n *Node) addAttribute(attr *nameValueStr) {
	n.attributes.PushBack(attr)
}


// recursive function to dump a node's content
func (n *Node) dumpNode(str string, indent int) string {
	if len(str) > 0 {
		str += "\n"
	}

	for i := 0; i < indent; i++ {
		str += "\t"
	}

	str += n.name

	if n.id != nil {
		str += " " + getAttrStr(n.id)
	}

	if n.class != nil {
		str += " " + getAttrStr(n.class)
	}

	for attrEl := n.attributes.Front(); attrEl != nil; attrEl = attrEl.Next() {
		attr := attrEl.Value.(*nameValueStr)
		str += " " + getAttrStr(attr)
	}

	if len(n.text) > 0 {
		str += " (" + n.text + ")"
	}

	for n := n.children.Front(); n != nil; n = n.Next() {
		str = n.Value.(*Node).dumpNode(str, indent+1)
	}

	return str
}

// returns a string representation of an attribute
func getAttrStr(nv *nameValueStr) string {
	return nv.name + "='" + nv.value + "'"
}
