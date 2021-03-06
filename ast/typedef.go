package ast

type Typedef struct {
	Address  string
	Type     string
	Children []Node
}

func parseTypedef(line string) *Typedef {
	groups := groupsFromRegex(
		"'(?P<type>.*)'",
		line,
	)

	return &Typedef{
		Address:  groups["address"],
		Type:     groups["type"],
		Children: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *Typedef) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
