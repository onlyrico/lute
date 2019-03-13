// Lute - A structural markdown engine.
// Copyright (C) 2019, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package lute

import (
	"fmt"
)

// Node represents a node in ast. https://github.com/syntax-tree/mdast
type Node interface {
	Type() NodeType
	Position() Pos
	String() string
	HTML() string
}

// NodeType identifies the type of a parse tree node.
type NodeType int

func (t NodeType) Type() NodeType {
	return t
}

const (
	NodeParent NodeType = iota
	NodeRoot
	NodeParagraph
	NodeHeading
	NodeThematicBreak
	NodeBlockquote
	NodeList
	NodeListItem
	NodeTable
	NodeTableRow
	NodeTableCell
	NodeHTML
	NodeCode
	NodeYAML
	NodeDefinition
	NodeFootnoteDefinition
	NodeText
	NodeEmphasis
	NodeStrong
	NodeDelete
	NodeInlineCode
	NodeBreak
	NodeLink
	NodeImage
	NodeLinkReference
	NodeImageReference
	NodeFootnote
	NodeFootnoteReference
)

// Nodes.

type Root struct {
	NodeType
	Pos
	Children []Node
}

func (n *Root) String() string {
	return fmt.Sprintf("%s", n.Children)
}

func (n *Root) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("%s\n", content)
}

func (n *Root) append(c Node) {
	n.Children = append(n.Children, c)
}

type Paragraph struct {
	NodeType
	Pos

	Children []Node
}

func (n *Paragraph) String() string {
	return fmt.Sprintf("%s", n.Children)
}

func (n *Paragraph) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("<p>%s</p>\n", content)
}

func (n *Paragraph) append(c Node) {
	n.Children = append(n.Children, c)
}

func (n *Paragraph) trim() {
	size := len(n.Children)
	if 1 > size {
		return
	}

	initialNoneWhitespace := 0
	for i := initialNoneWhitespace; i < size/2; i++ {
		if NodeBreak == n.Children[i].Type() {
			initialNoneWhitespace++
		}
	}

	finalNoneWhitespace := size
	for i := finalNoneWhitespace - 1; size/2 <= i; i-- {
		if NodeBreak == n.Children[i].Type() {
			finalNoneWhitespace--
		}
	}

	n.Children = n.Children[initialNoneWhitespace:finalNoneWhitespace]
}

type Heading struct {
	NodeType
	Pos
	Children []Node

	Depth int
}

func (n Heading) String() string {
	return fmt.Sprintf("# %s", n.Children)
}

func (n *Heading) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("<h%d>%s</h%d>", n.Depth, content, n.Depth)
}

type ThematicBreak struct {
	NodeType
	Pos
}

func (n *ThematicBreak) String() string {
	return fmt.Sprintf("'***'")
}

func (n *ThematicBreak) HTML() string {
	return fmt.Sprintf("<hr>")
}

type Blockquote struct {
	NodeType
	Pos
	Children []Node
}

func (n *Blockquote) String() string {
	return fmt.Sprintf("%s", n.Children)
}

func (n *Blockquote) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("<blockquote>%s</blockquote>", content)
}

type List struct {
	NodeType
	Pos
	Children []Node

	Ordered bool
	Start   int
	Spread  bool
}

func (n *List) String() string {
	return fmt.Sprintf("%s", n.Children)
}

func (n *List) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("<ul>\n%s</ul>", content)
}

func (n *List) append(c Node) {
	n.Children = append(n.Children, c)
}

type ListItem struct {
	NodeType
	Pos
	Children []Node

	Checked  bool
	Spread   bool

	Tree *Tree
}

func (n *ListItem) String() string {
	return fmt.Sprintf("%s", n.Children)
}

func (n *ListItem) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("<li>\n%s</li>\n", content)
}

func (n *ListItem) append(c Node) {
	n.Children = append(n.Children, c)
	n.Tree.CurNode = c
}

type Table struct {
	NodeType
	Pos
	Children []Node

	Align    alignType
}

type TableRow struct {
	NodeType
	Pos
	Children []Node
}

type TableCell struct {
	NodeType
	Pos
	Children []Node
}

type HTML struct {
	NodeType
	Pos
	Value string
}

type Code struct {
	NodeType
	Pos
	Value string

	Lang string
	Meta string
}

func (n *Code) String() string {
	return fmt.Sprintf("```%s```", n.Value)
}

func (n *Code) HTML() string {
	return fmt.Sprintf("<pre><code>%s</code></pre>", n.Value)
}

type YAML struct {
	NodeType
	Pos
	Value string
}

type Definition struct {
	NodeType
	Pos
	Association
	Resource
}

type FootnoteDefinition struct {
	NodeType
	Pos
	Children []Node

	Association
}

type Text struct {
	NodeType
	Pos
	Value string
}

func (n *Text) String() string {
	return fmt.Sprintf("'%s'", n.Value)
}

func (n *Text) HTML() string {
	return fmt.Sprintf("%s", n.Value)
}

type Emphasis struct {
	NodeType
	Pos
	Children []Node
}

func (n *Emphasis) String() string {
	return fmt.Sprintf("*%v*", n.Children)
}

func (n *Emphasis) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("<em>%s</em>", content)
}

type Strong struct {
	NodeType
	Pos
	Children []Node
}

func (n *Strong) String() string {
	return fmt.Sprintf("**%v**", n.Children)
}

func (n *Strong) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("<strong>%s</strong>", content)
}

type Delete struct {
	NodeType
	Pos
	Children []Node
}

func (n *Delete) String() string {
	return fmt.Sprintf("~~%v~~", n.Children)
}

func (n *Delete) HTML() string {
	content := html(n.Children)

	return fmt.Sprintf("<del>%s</del>", content)
}

type InlineCode struct {
	NodeType
	Pos
	Value string
}

func (n InlineCode) String() string {
	return fmt.Sprintf("`%s`", n.Value)
}

func (n InlineCode) HTML() string {
	return fmt.Sprintf("<code>%s</code>", n.Value)
}

type Break struct {
	NodeType
	Pos
}

func (n *Break) String() string {
	return fmt.Sprint("\n")
}

func (n *Break) HTML() string {
	return fmt.Sprintf("\n")
}

type Link struct {
	NodeType
	Pos
	Children []Node

	Resource
}

type Image struct {
	NodeType
	Pos
	Resource
	Alternative
}

func (n Image) String() string {
	return fmt.Sprintf("%s", n.URL)
}

type LinkReference struct {
	NodeType
	Pos
	Children []Node

	Reference
}

type ImageReference struct {
	NodeType
	Pos
	Reference
	Alternative
}

type Footnote struct {
	NodeType
	Pos
	Children []Node
}

type FootnoteReference struct {
	NodeType
	Pos
	Association
}

// Mixins.

type Resource struct {
	URL   string
	Title string
}

type Association struct {
	Identifier string
	Label      string
}

type Reference struct {
	ReferenceType referenceType
	Association
}

type Alternative struct {
	Alt string
}

// Enumerations.

type alignType string
type referenceType string

func html(children []Node) string {
	var ret string
	for _, c := range children {
		ret += c.HTML()
	}

	return ret
}
