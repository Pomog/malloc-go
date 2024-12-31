package hash

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
)

type Node struct {
	Key   uint32
	Value string
}

func newNode(value string) (*Node, error) {
	if value == "" {
		return nil, errors.New("value cannot be empty")
	}

	h := sha256.Sum256([]byte(value))

	key := binary.BigEndian.Uint32(h[:32])

	node := &Node{
		Key:   key,
		Value: value,
	}

	return node, nil
}

func (n *Node) PrintValue() {
    fmt.Printf("\n**\nKey: %d\n, Value: %s\n", n.Key, n.Value)
}
