package menu

import (
	"ludo/src/common"
	"ludo/src/game"
	"ludo/src/network"

	"github.com/nsf/termbox-go"
)

type callback func() int

var (
	mainMenu = menu{
		opts: []opt{
			{
				label: "Start",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(START_MENU)
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
					mpt.changeMenuPage(OFFLINE_MENU)
					return false, nil
				},
			},
			{
				label: "Online",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(ONLINE_MENU)
					return false, nil
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(MAIN_MENU)
					return false, nil
				},
			},
		},
	}

	offlineSubOpts = []string{"Player", "Bot", "-"}
	onlineSubOpts  = []string{"Player", "-"}

	onlineMenu = menu{
		opts: []opt{
			{
				label: "Host",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(HOST_MENU)
					return false, nil
				},
			},
			{
				label: "Join",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					return true, func() int {
						return game.StartGameOnline("")
					}
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(START_MENU)
					return false, nil
				},
			},
		},
	}

	hostMenu = menu{
		opts: []opt{
			{
				label:   termbox.ColorBlue,
				subOpts: onlineSubOpts,
			},
			{
				label:   termbox.ColorRed,
				subOpts: onlineSubOpts,
			},
			{
				label:   termbox.ColorYellow,
				subOpts: onlineSubOpts,
			},
			{
				label:   termbox.ColorGreen,
				subOpts: onlineSubOpts,
			},
			{
				label: "Done",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					players := []common.PlayerData{}
					curMenuOpts := mpt.menus[mpt.curIdx].opts
					curMenuOpts = curMenuOpts[:len(curMenuOpts)-2]

					for _, opt := range curMenuOpts {
						players = append(players, common.PlayerData{Color: opt.label.(termbox.Attribute), Type: opt.subOpts[opt.curIdx]})
					}

					return true, func() int {
						termbox.Close()
						termbox.Init()
						go network.Host(players)
						return game.StartGameOnline("127.0.0.1:8080")
					}
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(ONLINE_MENU)
					return false, nil
				},
			},
		},
	}

	offlineMenu = menu{
		opts: []opt{
			{
				label:   termbox.ColorYellow,
				subOpts: offlineSubOpts,
			},
			{
				label:   termbox.ColorBlue,
				subOpts: offlineSubOpts,
			},
			{
				label:   termbox.ColorRed,
				subOpts: offlineSubOpts,
			},
			{
				label:   termbox.ColorGreen,
				subOpts: offlineSubOpts,
			},
			{
				label: "Done",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					players := []common.PlayerData{}
					curMenuOpts := mpt.menus[mpt.curIdx].opts
					curMenuOpts = curMenuOpts[:len(curMenuOpts)-2]

					for _, opt := range curMenuOpts {
						players = append(players, common.PlayerData{Color: opt.label.(termbox.Attribute), Type: opt.subOpts[opt.curIdx]})
					}
					return true, func() int {
						termbox.Close()
						termbox.Init()
						game.StartGameOffline(players)
						return 0
					}
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(START_MENU)
					return false, nil
				},
			},
		},
	}
)
