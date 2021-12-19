package menu

type opt struct {
	label    string
	onSelect func(*menuPagesType) (bool, callback)
	curIdx   int
	subOpts  []string
}

type menu struct {
	curIdx int
	opts   []opt
}

func (m *menu) handleOptNav(mag int) {
	maxOptIdx := len(m.opts) - 1

	m.curIdx += mag

	if m.curIdx < 0 {
		m.curIdx = maxOptIdx
	} else if m.curIdx == maxOptIdx+1 {
		m.curIdx = 0
	}
}

func (m *menuPagesType) handleOptSelect() (bool, callback) {
	curMenu := m.menus[m.curIdx]
	return curMenu.opts[curMenu.curIdx].onSelect(m)
}

func (m *menuPagesType) changeMenuPage(pageIdx int) {
	m.curIdx = pageIdx
}
