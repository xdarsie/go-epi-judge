package do_lists_overlap_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stefantds/csvdecoder"

	. "github.com/stefantds/go-epi-judge/epi/do_lists_overlap"
	"github.com/stefantds/go-epi-judge/list"
)

func TestOverlappingLists(t *testing.T) {
	testFileName := testConfig.TestDataFolder + "/" + "do_lists_overlap.tsv"
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatalf("could not open file %s: %v", testFileName, err)
	}
	defer file.Close()

	type TestCase struct {
		L0      list.ListNodeDecoder
		L1      list.ListNodeDecoder
		Common  list.ListNodeDecoder
		Cycle0  int
		Cycle1  int
		Details string
	}

	parser, err := csvdecoder.NewWithConfig(file, csvdecoder.Config{Comma: '\t', IgnoreHeaders: true})
	if err != nil {
		t.Fatalf("could not parse file %s: %s", testFileName, err)
	}

	for i := 0; parser.Next(); i++ {
		tc := TestCase{}
		if err := parser.Scan(
			&tc.L0,
			&tc.L1,
			&tc.Common,
			&tc.Cycle0,
			&tc.Cycle1,
			&tc.Details,
		); err != nil {
			t.Fatal(err)
		}

		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			if err := overlappingListsWrapper(tc.L0.Value, tc.L1.Value, tc.Common.Value, tc.Cycle0, tc.Cycle1); err != nil {
				t.Error(err)
			}
		})
	}
	if err = parser.Err(); err != nil {
		t.Fatalf("parsing error: %s", err)
	}
}

func overlappingListsWrapper(l0 *list.ListNode, l1 *list.ListNode, common *list.ListNode, cycle0 int, cycle1 int) error {
	if common != nil {
		if l0 == nil {
			l0 = common
		} else {
			it := l0
			for it.Next != nil {
				it = it.Next
			}
			it.Next = common
		}

		if l1 == nil {
			l1 = common
		} else {
			it := l1
			for it.Next != nil {
				it = it.Next
			}
			it.Next = common
		}
	}

	if cycle0 != -1 && l0 != nil {
		last := l0
		for last.Next != nil {
			last = last.Next
		}

		it := l0
		for ; cycle0 > 0; cycle0-- {
			if it == nil {
				panic("invalid input data")
			}
			it = it.Next
		}
		last.Next = it
	}

	if cycle1 != -1 && l1 != nil {
		last := l1
		for last.Next != nil {
			last = last.Next
		}

		it := l1
		for ; cycle1 > 0; cycle1-- {
			if it == nil {
				panic("invalid input data")
			}
			it = it.Next
		}
		last.Next = it
	}

	commonNodes := make(map[int]bool)
	for it := common; it != nil; it = it.Next {
		if _, ok := commonNodes[it.Data.(int)]; ok {
			break
		}

		commonNodes[it.Data.(int)] = true
	}

	result := OverlappingLists(l0, l1)

	if len(commonNodes) == 0 {
		if result != nil {
			return errors.New("invalid result")
		}
	} else {
		if result == nil {
			return errors.New("invalid result")
		}

		_, ok := commonNodes[result.Data.(int)]
		if !ok {
			return errors.New("invalid result")
		}
	}

	return nil
}
