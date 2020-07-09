package anagrams_test

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	csv "github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/anagrams"
	"github.com/stefantds/go-epi-judge/utils"
)

func TestFindAnagrams(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "anagrams.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		Dictionary     []string
		ExpectedResult [][]string
		Details        string
	}

	parser, err := csv.NewParserWithConfig(file, csv.ParserConfig{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.Dictionary,
			&tc.ExpectedResult,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			result := FindAnagrams(tc.Dictionary)
			if !equal(result, tc.ExpectedResult) {
				t.Errorf("expected %v, got %v", tc.ExpectedResult, result)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func equal(result, expected [][]string) bool {
	for _, l := range expected {
		sort.Strings(l)
	}

	sort.Slice(expected, func(i, j int) bool {
		return utils.LexStringsCompare(expected[i], expected[j])
	})

	for _, l := range result {
		sort.Strings(l)
	}

	sort.Slice(result, func(i, j int) bool {
		return utils.LexStringsCompare(result[i], result[j])
	})

	return reflect.DeepEqual(expected, result)
}