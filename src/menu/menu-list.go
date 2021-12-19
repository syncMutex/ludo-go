package menu

var (
	mainMenu = menu{
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
					mpt.changeMenuPage(2)
					return false, nil
				},
			},
			{
				label: "Online",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					return false, nil
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(0)
					return false, nil
				},
			},
		},
	}

	offlineMenu = menu{
		opts: []opt{
			{
				label:   "No of Players: ",
				subOpts: []string{"2", "3", "4"},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(1)
					return false, nil
				},
			},
		},
	}
)
