package menu

type pos struct {
	x, y int
}

type menuPagesType struct {
	curIdx     int
	menus      []*menu
	displayPos pos
}

func (m *menuPagesType) changeMenuPage(pageIdx int) {
	m.curIdx = pageIdx
	m.menus[m.curIdx].curIdx = 0
}

const (
	MAIN_MENU int = iota
	START_MENU
	OFFLINE_MENU
	ONLINE_MENU
	HOST_MENU
)

func InitMenu() {
	menuPages := menuPagesType{
		displayPos: pos{5, 5},
		curIdx:     0,
		menus:      []*menu{&mainMenu, &startMenu, &offlineMenu, &onlineMenu, &hostMenu},
	}
	menuPages.renderMenu()
	if callback := menuPages.keyboardLoop(); callback != nil {
		callback()
	}
}
