package disjointset

import (
	"fmt"
	"testing"
)

func TestMakeSet(t *testing.T) {
	// Int value
	dsInt := MakeSet(5)

	if dsInt.Value != 5 {
		t.Error(fmt.Sprintf("Expected: %d, Actual: %d", 5, dsInt.Value))
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
		t.Error(fmt.Sprintf("Expected: %s, Actual: %s", "bad_wolf", dsStr.Value))
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
		t.Error(fmt.Sprintf("Expected: %v, Actual: %v", dsInt, dsDS.Value))
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
	err := Union(dss[0], dss[1])
	if err != nil {
		t.Error(fmt.Sprintf("Unexpected error in building parent structure with Union: %s", err.Error()))
	}
	err = Union(dss[0], dss[3])
	if err != nil {
		t.Error(fmt.Sprintf("Unexpected error in building parent structure with Union: %s", err.Error()))
	}
	err = Union(dss[3], dss[4])
	if err != nil {
		t.Error(fmt.Sprintf("Unexpected error in building parent structure with Union: %s", err.Error()))
	}
	err = Union(dss[2], dss[5])
	if err != nil {
		t.Error(fmt.Sprintf("Unexpected error in building parent structure with Union: %s", err.Error()))
	}
	err = Union(dss[5], dss[8])
	if err != nil {
		t.Error(fmt.Sprintf("Unexpected error in building parent structure with Union: %s", err.Error()))
	}
	err = Union(dss[5], dss[9])
	if err != nil {
		t.Error(fmt.Sprintf("Unexpected error in building parent structure with Union: %s", err.Error()))
	}

	parents := []*DisjointSet{dss[0], dss[0], dss[2], dss[0], dss[0], dss[2], dss[6], dss[7], dss[2], dss[2]}
	for idx, ds := range dss {
		parent, err := Find(ds)
		if err != nil {
			t.Error(fmt.Sprintf("Unexpected error in Find: %s", err.Error()))
		}

		if parent != parents[idx] {
			t.Error(fmt.Sprintf("Expected parent: %v, Actual parent: %v", parents[idx], parent))
		}
	}
}
