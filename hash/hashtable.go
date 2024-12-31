package hash

type HashTable struct {
	buckets   [][]*Node
	tableSize uint32
}

func NewHashTable(tableSize uint32) *HashTable {
	return &HashTable{
		buckets:   make([][]*Node, tableSize),
		tableSize: tableSize,
	}
}

func (ht *HashTable) Add(node *Node) {
	if node == nil {
		return
	}

	index := node.Key % ht.tableSize
	currentBucket := ht.buckets[index]

	for i, existingNode := range currentBucket {
		if existingNode.Key == node.Key {
			ht.buckets[index][i].Value = node.Value
			return
		}
	}

	ht.buckets[index] = append(ht.buckets[index], node)
}
