package epi_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/goepijudge/epi"
)

func checkApplyPermutation() error {
	//TODO
	return nil
}

func TestApplyPermutation(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "apply_permutation.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Perm    []int
		A       []int
		Details string
	}

	parser, err := csv.NewParser(file, &csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Perm,
			&tc.A,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			ApplyPermutation(tc.Perm, tc.A)
			err := checkApplyPermutation()
			if err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Errorf("parsing error: %w", err)
	}
}
