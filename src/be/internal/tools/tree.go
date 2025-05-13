package tools

import (
	"fmt"
	"strings"
)

func ParseTree(s string) (*Node, error) {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return nil, fmt.Errorf("empty string")
	}

	i := 0
	for i < len(s) && s[i] != '(' {
		i++
	}

	if i == len(s) {
		return &Node{Name: s}, nil
	}

	name := s[:i]

	node := &Node{
		Name:     name,
		Children: make([]*Node, 0, 2),
	}

	content := s[i:]
	level := 0
	start := 1

	for j := 1; j < len(content); j++ {
		if content[j] == '(' {
			level++
		} else if content[j] == ')' {
			level--
		}

		if (level == 0 && j < len(content)-1 && content[j+1] == '(') || (level == -1) {
			if j > start {
				child, err := ParseTree(content[start : j+1])
				if err != nil {
					return nil, err
				}
				node.Children = append(node.Children, child)
			}

			if level == -1 {
				break
			}

			start = j + 1
		}
	}

	return node, nil
}

func CountNodes(node *Node) int {
	if node == nil {
		return 0
	}
	
	count := 1
	for _, child := range node.Children {
		count += CountNodes(child)
	}
	
	return count
}

func SortTree(node *Node) *Node {
	if node == nil || len(node.Children) < 2 {
		return node
	}
	
	for i := range node.Children {
		node.Children[i] = SortTree(node.Children[i])
	}
	
	if len(node.Children) == 2 {
		leftCount := CountNodes(node.Children[0])
		rightCount := CountNodes(node.Children[1])
		
		if rightCount > leftCount {
			node.Children[0], node.Children[1] = node.Children[1], node.Children[0]
		}
	}
	
	return node
}