package roundrobin

type RoundRobin struct {
	vals []any
	curr int
}

func New() *RoundRobin {
	return &RoundRobin{
		curr: 0,
	}
}

func (r *RoundRobin) Add(v any) {
	r.vals = append(r.vals, v)
}

func (r *RoundRobin) Next() (interface{}, bool) {
	if len(r.vals) == 0 {
		return nil, false
	}

	idx := r.curr % len(r.vals)
	r.curr++

	return r.vals[idx], true
}
