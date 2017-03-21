package disjointset

// DisjointSet that implements the disjoint set ADT with API Find/Union.
type DisjointSet struct {
	Value  interface{}
	parent *DisjointSet
	rank   int
}

// MakeSet returns a pointer to a new DisjoinSet that stores the given
// value. The returned DisjointSet has no parent.
func MakeSet(value interface{}) *DisjointSet {
	return &DisjointSet{value, nil, 0}
}

// Find returns a pointer to the DisjointSet that is the parent of this
// DisjointSet. If this DisjointSet has no parent, then Find will return
// this DisjointSet.
func (ds *DisjointSet) Find() (*DisjointSet, error) {
	if ds.parent == nil {
		return ds, nil
	}

	return ds.parent.Find()
}

// Union merges the two DisjointSets into one DisjointSet.
func Union(x *DisjointSet, y *DisjointSet) error {
	xRoot, xerr := x.Find()
	if xerr != nil {
		return xerr
	}
	yRoot, yerr := y.Find()
	if yerr != nil {
		return yerr
	}

	if xRoot == yRoot {
		return nil
	}

	if xRoot.rank < yRoot.rank {
		xRoot.parent = yRoot
	} else if xRoot.rank > yRoot.rank {
		yRoot.parent = xRoot
	} else {
		yRoot.parent = xRoot
		xRoot.rank = xRoot.rank + 1
	}

	return nil
}
