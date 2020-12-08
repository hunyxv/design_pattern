package mediator

import "testing"

func TestMediator(t *testing.T) {
	landlord1 := &Landlord{price: 1600, addr: "八里桥xxx"}
	landlord2 := &Landlord{price: 1800, addr: "大兴xxx"}
	landlord3 := &Landlord{price: 2200, addr: "国贸xxx"}

	realtor := &Realtor{
		source: []Landlorder{landlord1,landlord2,landlord3},
	}

	result := realtor.Enquiry(1900)
	t.Logf("%+v", result)
}
