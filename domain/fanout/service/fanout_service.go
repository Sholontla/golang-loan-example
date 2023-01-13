package service

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/fanout"
	"github.com/projects/loans/utils/date_utils"
)

func DummyDataGeneratorA(a fanout.Worker1) fanout.Worker1 {
	return fanout.Worker1{
		ID1:          uuid.New(),
		Name1:        a.Name1,
		Description1: a.Description1,
		FloatNum1:    a.FloatNum1,
		CreatedAt1:   date_utils.GetNowString(),
	}
}

func DummyDataGeneratorB(b fanout.Worker2) fanout.Worker2 {
	return fanout.Worker2{
		ID2:          uuid.New(),
		Name2:        b.Name2,
		Description2: b.Description2,
		FloatNum2:    b.FloatNum2,
		CreatedAt2:   date_utils.GetNowString(),
	}
}

func DummyDataGeneratorC(c fanout.Worker3) fanout.Worker3 {
	return fanout.Worker3{
		ID3:          uuid.New(),
		Name3:        c.Name3,
		Description3: c.Description3,
		FloatNum3:    c.FloatNum3,
		CreatedAt3:   date_utils.GetNowString(),
	}
}

func DummyDataGeneratorD(d fanout.Worker4) fanout.Worker4 {
	return fanout.Worker4{
		ID4:          uuid.New(),
		Name4:        d.Name4,
		Description4: d.Description4,
		FloatNum4:    d.FloatNum4,
		CreatedAt4:   date_utils.GetNowString(),
	}
}

func FanOutService(a fanout.Worker1, b fanout.Worker2, c fanout.Worker3, d fanout.Worker4) (*fanout.Worker1, *fanout.Worker2, *fanout.Worker3, *fanout.Worker4, error) {

	a = DummyDataGeneratorA(a)
	b = DummyDataGeneratorB(b)
	c = DummyDataGeneratorC(c)
	d = DummyDataGeneratorD(d)

	fanout.GetProcessor(a, b, c, d)
	return &a, &b, &c, &d, nil
}
