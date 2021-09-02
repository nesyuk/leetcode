package bstree

import "testing"

func TestCodec(t *testing.T) {
	c := NewCodec()
	encodedTrees := []string{"1", "1,2,", "1,2,3", "1,,3"}
	for _, encoded := range encodedTrees {
		tree := c.deserialize(encoded)
		encodedTree := c.serialize(tree)
		if encodedTree != encoded {
			Print(tree)
			t.Fatalf("expected: %s, got: %s", encoded, encodedTree)
		}
	}
}
