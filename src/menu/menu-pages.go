package menu

var (
	mainMenu = menu{
		curIdx: 0,
		opts: []opt{
			{
				label: "Start",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(1)
					return false, nil
				},
			},
			{
				label: "Quit",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					return true, nil
				},
			},
		},
	}

	startMenu = menu{
		opts: []opt{
			{
				label: "Offline",
				onSelect: func(mpt *menuPagesType) (bool, callback) {

					return false, nil
				},
			},
			{
				label: "Online",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					return false, nil
				},
			},
		},
	}
)
