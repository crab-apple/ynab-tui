package test

import (
	"ynabtui/app/app/ui"
)

type FakeScreenSupplier struct {
}

func (FakeScreenSupplier) FirstLoad() ui.Screen {
	return nil
}
