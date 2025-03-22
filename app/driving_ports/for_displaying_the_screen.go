package driving_ports

import "ynabtui/app/app/ui"

type ForDisplayingTheScreen interface {
	FirstLoad() ui.Screen
}
