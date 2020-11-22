package chainofresponsibility

import "testing"

func TestChineOfResponsibility(t *testing.T) {
	handler := &HRDepartment{}
	handler.SetNextHandler(&FinanceDepartment{})

	event := AskLeave
	handler.Handle(event)

	event = Reimburse
	handler.Handle(event)

	event = Other
	handler.Handle(event)
}