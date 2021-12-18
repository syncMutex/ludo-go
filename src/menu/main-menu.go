package menu

type pos struct {
	x, y int
}

type menuPagesType struct {
	curIdx     int
	menus      []menu
	displayPos pos
}

var (
	menuPages = menuPagesType{
		displayPos: pos{5, 5},
		curIdx:     0,
		menus:      []menu{mainMenu},
	}
)

func StartMainMenu() {
	menuPages.renderMenu()
	menuPages.keyboardLoop()
}
