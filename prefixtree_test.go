package prefix

import (
	"fmt"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree("Test")

	//Check Name
	if tree.GetName() != "Test" {
		t.Errorf("Expected Name: Test: %v instead.", tree.GetName())
	}

	//Check size
	if tree.GetSize() != 0 {
		t.Errorf("Expected Size: 0, instead had size %v", tree.GetSize())
	}

	//Get Root rune
	if tree.GetRoot().GetChar() != '0' {
		t.Errorf("Expected Rune 0 instead had rune %v", tree.GetRoot().GetChar())
	}

}

func TestInsertKeyword(t *testing.T) {
	tree := NewTree("Test")

	tree.InsertKeyword("apple", "becky")

	//fmt.Println(tree.Contains("apple"))
}

func TestPrintTrie(t *testing.T) {
	tree := NewTree("Test")
	tree.InsertKeyword("pear", "becky")
	tree.InsertKeyword("able", "karen")
	tree.InsertKeyword("bear", "susan")
	tree.InsertKeyword("mamale", "getrude")

	var testCases = []struct {
		strings  []string
		expected []string
	}{
		{tree.PrintTrie(), []string{"pear", "able", "bear", "mamale"}},
	}

	for _, test := range testCases {

		if fmt.Sprintf("%v", test.strings) != fmt.Sprintf("%v", test.expected) {
			t.Errorf("Expected %s got %s instead", test.expected, test.strings)
		}
	}
}
