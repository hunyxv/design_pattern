package mediator

// Mediator 中介接口
type Mediator interface {
	Enquiry(float64) []string
}

// Realtor 房屋中介
type Realtor struct {
	source []Landlorder
}

// Enquiry 询价
func (r *Realtor) Enquiry(price float64) (houes []string) {
	for _, i := range r.source {
		if i.Enquiry(price) {
			houes = append(houes, i.Addr())
		}
	}
	return
}

// Landlorder 房东接口
type Landlorder interface {
	Enquiry(float64) bool
	Addr() string
}

// Landlord .
type Landlord struct {
	addr  string
	price float64
}

// Enquiry 询价
func (l *Landlord) Enquiry(p float64) bool {
	return p >= l.price
}

// Addr 地址
func (l *Landlord) Addr() string {
	return l.addr
}
