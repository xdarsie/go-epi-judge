package epi_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi"
	"github.com/stefantds/go-epi-judge/tree"
)

func TestFindSuccessor(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "successor_in_tree.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Tree           tree.BinaryTreeDecoder
		NodeIdx        int
		ExpectedResult int
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Tree,
			&tc.NodeIdx,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := findSuccessorWrapper(tc.Tree.Value, tc.NodeIdx)
			if !reflect.DeepEqual(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}

func findSuccessorWrapper(inputTree *tree.BinaryTree, nodeIdx int) int {
	n := tree.MustFindNode(inputTree, nodeIdx).(*tree.BinaryTree)

	result := FindSuccessor(n)

	if result == nil {
		return -1
	}

	return result.Data.(int)
}
