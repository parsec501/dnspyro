package cli

type Resolved struct {
	a   []string
	mx  []string
	ns  []string
	txt []string
}

func NewResolved() *Resolved {
	return &Resolved{}
}

func (r *Resolved) SetA(a string) {
	r.a = append(r.a, a)
}

func (r *Resolved) SetMX(mx string) {
	r.mx = append(r.mx, mx)
}

func (r *Resolved) SetNS(ns string) {
	r.ns = append(r.ns, ns)
}

func (r *Resolved) SetTXT(txt string) {
	r.txt = append(r.txt, txt)
}
