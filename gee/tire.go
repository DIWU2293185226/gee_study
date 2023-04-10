package gee

import "strings"

type Node struct {
	Pattern  string
	part     string
	Iswild   bool
	children []*Node
}

// 第一个匹配成功的节点，插入时使用，从这个节点往后插或者从这里向下找
func (n *Node) Matchchild(part string) *Node {
	for _, child := range n.children {
		if child.part == part || child.Iswild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，提供一个节点集合供后续查找
func (n *Node) Matchchildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range n.children {
		if child.part == part || child.Iswild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *Node) Insert(pattern string, parts []string, height int) {
	//到头了就不插了,待插入的pattern就给这个节点了
	if len(parts) == height {
		n.Pattern = pattern
		return
	}
	//拿出当前层的节点，查找有没有，没有就增加直到高度达到模式串长度，有就继续递归查找
	part := parts[height]
	child := n.Matchchild(part)

	if child == nil {
		child = &Node{
			part:   part,
			Iswild: part[0] == '*' || part[0] == ':',
		}
		n.children = append(n.children, child)
	}
	child.Insert(pattern, parts, height+1)
}

func (n *Node) Search(parts []string, hegint int) *Node {
	//匹配到头，查看pattern是否为空，不为空则成功
	if len(parts) == hegint || strings.HasPrefix(n.part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}
	part := parts[hegint]
	children := n.Matchchildren(part)
	for _, child := range children {
		result := child.Search(parts, hegint+1)
		if result != nil {
			return result
		}
	}

	return nil
}
