package menu

type pos struct {
	x, y int
}

var (
	menuPos = pos{2, 3}
	menuOps = []string{
		"Start",
		"Quit",
	}
	curOpt int
)

func handleMenuNav(mag int) {
	curOpt += mag
	optCount := len(menuOps) - 1
	if curOpt > optCount {
		curOpt = 0
	}
	if curOpt < 0 {
		curOpt = optCount
	}
}

func handleOptSelect() (quit bool) {
	switch curOpt {
	case 1:
		return true
	}
	return false
}

func StartMainMenu() {
	renderMenu()
	listenMenuKeyboard()
}
