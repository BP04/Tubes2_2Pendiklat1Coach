package tools

import (
	"strings"
)

func ParseTree(expr string) (*Node, error) {
	expr = strings.TrimSpace(expr)
	if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
		expr = expr[1 : len(expr)-1]
	}

	openIdx := strings.Index(expr, "(")
	if openIdx == -1 {
		return &Node{Name: expr}, nil
	}

	name := expr[:openIdx]

	root := &Node{
		Name:     name,
		Children: make([]*Node, 0, 2),
	}

	content := expr[openIdx:]

	var leftExpr, rightExpr string

	leftStart := 0
	parenCount := 0
	leftEnd := 0

	for i := 0; i < len(content); i++ {
		if content[i] == '(' {
			if parenCount == 0 {
				leftStart = i
			}
			parenCount++
		} else if content[i] == ')' {
			parenCount--
			if parenCount == 0 {
				leftEnd = i
				break
			}
		}
	}

	if leftStart < leftEnd {
		leftExpr = content[leftStart : leftEnd+1]
	}

	rightStart := leftEnd + 1
	parenCount = 0
	rightEnd := 0

	for i := rightStart; i < len(content); i++ {
		if content[i] == '(' {
			if parenCount == 0 {
				rightStart = i
			}
			parenCount++
		} else if content[i] == ')' {
			parenCount--
			if parenCount == 0 {
				rightEnd = i
				break
			}
		}
	}

	if rightStart < rightEnd {
		rightExpr = content[rightStart : rightEnd+1]
	}

	if leftExpr != "" {
		leftNode, err := ParseTree(leftExpr)
		if err != nil {
			return nil, err
		}
		root.Children = append(root.Children, leftNode)
	}

	if rightExpr != "" {
		rightNode, err := ParseTree(rightExpr)
		if err != nil {
			return nil, err
		}
		root.Children = append(root.Children, rightNode)
	}

	return root, nil
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
