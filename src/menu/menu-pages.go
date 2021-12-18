package menu

var (
	mainMenu = menu{
		curIdx: 0,
		opts: []opt{
			{
				label: "Start",
			},
			{
				label: "Quit",
				onSelect: func() bool {
					return true
				},
			},
		},
	}
)
