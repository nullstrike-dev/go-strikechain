package storm

type Node struct {
	Address string
	Peers   map[string]bool
}

func NewNode(addr string) *Node {
	return &Node{
		Address: addr,
		Peers:   make(map[string]bool),
	}
}
