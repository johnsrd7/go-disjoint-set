package disjointset

import (
	"fmt"
	"testing"
)

func TestMakeSet(t *testing.T) {
	// Int value
	dsInt := MakeSet(5)

	if dsInt.Value != 5 {
		t.Errorf("Expected: %d, Actual: %d\n", 5, dsInt.Value)
	}

	if dsInt.parent != nil {
		t.Error("MakeSet should create a set with no parent")
	}

	if dsInt.rank != 0 {
		t.Error("MakeSet should create a set with rank 0")
	}

	// String value
	dsStr := MakeSet("bad_wolf")

	if dsStr.Value != "bad_wolf" {
		t.Errorf("Expected: %s, Actual: %s\n", "bad_wolf", dsStr.Value)
	}

	if dsStr.parent != nil {
		t.Error("MakeSet should create a set with no parent")
	}

	if dsStr.rank != 0 {
		t.Error("MakeSet should create a set with rank 0")
	}

	// Disjoint set value
	dsDS := MakeSet(dsInt)

	if dsDS.Value != dsInt {
		t.Errorf("Expected: %v, Actual: %v\n", dsInt, dsDS.Value)
	}

	if dsDS.parent != nil {
		t.Error("MakeSet should create a set with no parent")
	}

	if dsDS.rank != 0 {
		t.Error("MakeSet should create a set with rank 0")
	}
}

func TestFind(t *testing.T) {
	dss := []*DisjointSet{}

	for i := 0; i < 10; i++ {
		dss = append(dss, MakeSet(i))
	}

	// Structure:
	// 0
	// 1->0
	// 2
	// 3->0
	// 4->3
	// 5->2
	// 6
	// 7
	// 8->5
	// 9->5
	dss[1].parent = dss[0]
	dss[3].parent = dss[0]
	dss[4].parent = dss[3]
	dss[5].parent = dss[2]
	dss[8].parent = dss[5]
	dss[9].parent = dss[5]

	parents := []*DisjointSet{dss[0], dss[0], dss[2], dss[0], dss[0], dss[2], dss[6], dss[7], dss[2], dss[2]}
	for idx, ds := range dss {
		parent, err := ds.Find()
		if err != nil {
			t.Errorf("Unexpected error in Find: %s\n", err.Error())
		}

		if parent != parents[idx] {
			t.Errorf("Expected parent: %v, Actual parent: %v\n", parents[idx], parent)
		}
	}
}

func TestUnion(t *testing.T) {
	dss := []*DisjointSet{}

	for i := 0; i < 10; i++ {
		dss = append(dss, MakeSet(i))
	}

	// Structure:
	// 0
	// 1->0
	// 2
	// 3->0
	// 4->3
	// 5->2
	// 6
	// 7
	// 8->5
	// 9->5

	// Test when the roots are the same
	err := Union(dss[0], dss[0])
	if err != nil {
		t.Error(err)
	}
	if dss[0].parent != nil {
		t.Error("Union of roots should do nothing")
	}

	// Test 1 level union
	err = Union(dss[0], dss[1])
	if err != nil {
		t.Error(err)
	}
	if dss[1].parent != dss[0] {
		t.Error(fmt.Sprintf("Expected parent: %v, Actual parent: %v", dss[0], dss[1].parent))
	}
}
