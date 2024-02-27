package mst

type edge struct {
	a point
	b point
	w int
}

type edegeList []edge

func (es edegeList) Len() int { return len(es) }

func (es edegeList) Swap(i, j int) { es[i], es[j] = es[j], es[i] }

func (es edegeList) Less(i, j int) bool { return es[i].w < es[j].w }
