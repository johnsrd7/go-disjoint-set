package disjointset

import "errors"

type DisjointSet struct {
	Value  interface{}
	parent *DisjointSet
	rank   int
}

func MakeSet(value interface{}) *DisjointSet {
	return &DisjointSet{value, nil, 0}
}

func Find(ds *DisjointSet) (*DisjointSet, error) {
	if ds == nil {
		return nil, errors.New("Invalid pointer for Find")
	}

	if ds.parent == nil {
		return ds, nil
	}

	return Find(ds.parent)
}

func Union(x *DisjointSet, y *DisjointSet) error {
	xRoot, xerr := Find(x)
	if xerr != nil {
		return xerr
	}
	yRoot, yerr := Find(y)
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
