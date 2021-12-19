package menu

type pos struct {
	x, y int
}

type menuPagesType struct {
	curIdx     int
	menus      []menu
	displayPos pos
}

func InitMenu() {
	menuPages := menuPagesType{
		displayPos: pos{5, 5},
		curIdx:     0,
		menus:      []menu{mainMenu, startMenu},
	}
	menuPages.renderMenu()
	if callback := menuPages.keyboardLoop(); callback != nil {
		callback()
	}
}
