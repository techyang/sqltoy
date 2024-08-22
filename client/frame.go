package client

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"os"
	"strings"
)

// InitWin 初始化界面
func InitWin() {

	tmw := new(TabMainWindow)

	if err := (MainWindow{
		Title:      "SQLTOY",
		AssignTo:   &tmw.MainWindow,
		Background: SolidColorBrush{Color: walk.RGB(0xED, 0xED, 0xED)},
		//Background: GradientBrush{
		//	Vertexes: []walk.GradientVertex{
		//		{X: 0, Y: 0, Color: walk.RGB(255, 255, 127)},
		//		{X: 1, Y: 0, Color: walk.RGB(127, 191, 255)},
		//		{X: 0.5, Y: 0.5, Color: walk.RGB(255, 255, 255)},
		//		{X: 1, Y: 1, Color: walk.RGB(127, 255, 127)},
		//		{X: 0, Y: 1, Color: walk.RGB(255, 127, 127)},
		//	},
		//	Triangles: []walk.GradientTriangle{
		//		{0, 1, 2},
		//		{1, 3, 2},
		//		{3, 4, 2},
		//		{4, 0, 2},
		//	},
		//},
		MenuItems: []MenuItem{
			Menu{
				Text: "&文件",
				Items: []MenuItem{
					Action{
						Text:  "&会话管理器",
						Image: "/icons/disconnect.png",
						//OnTriggered: tmw.openCommandManagePanel,
					},
					Action{
						Text: "&连接到",
						//	OnTriggered: tmw.openRelatedInstructionsPanel,
					},
					Action{
						Text:        "&新建窗口",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyN},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&新建查询标签页",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyT},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&关闭查询标签页",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyF4},
						OnTriggered: tmw.about,
					},
					Action{
						Text: "&关闭所有查询标签页",
						//Shortcut:    Shortcut{walk.ModControl, walk.KeyN},
						OnTriggered: tmw.about,
					},
					Separator{},
					Action{
						Text:        "&加载SQL文件...",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: tmw.open,
					},
					Action{
						Text:        "&运行SQL文件...",
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&保存",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyS},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&保存为SQL片段...",
						OnTriggered: tmw.about,
					},
					Separator{},
					Action{
						Text:        "&导出配置文件...",
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&导入配置文件...",
						OnTriggered: tmw.about,
					},
					Separator{},
					Action{
						Text:        "&退出(X)",
						Shortcut:    Shortcut{walk.ModAlt, walk.KeyF4},
						OnTriggered: tmw.about,
					},
				},
			},
			Menu{
				Text:     "&编辑",
				AssignTo: &tmw.configMenu,
				Items: []MenuItem{
					Action{
						Text:        "&撤销(U)",
						Shortcut:    Shortcut{walk.ModAlt, walk.KeyBack},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&复制(C)",
						Shortcut:    Shortcut{walk.ModAlt, walk.KeyC},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&用制表符复制到空间(C)",
						Shortcut:    Shortcut{walk.ModControl | walk.ModAlt, walk.KeyC},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&粘贴(P)",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyV},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&下移当前行",
						Shortcut:    Shortcut{walk.ModAlt, walk.KeyDown},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&上移当前行",
						Shortcut:    Shortcut{walk.ModAlt, walk.KeyUp},
						OnTriggered: tmw.about,
					},
					Separator{},
					Action{
						Text:        "&选择所有",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyA},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&反向选择",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyI},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&过滤器面板",
						Shortcut:    Shortcut{walk.ModControl | walk.ModAlt, walk.KeyF},
						OnTriggered: tmw.about,
					},
				},
			},
			Menu{
				AssignTo: &tmw.helpMenu,
				Text:     "&搜索",
				Items: []MenuItem{
					Action{
						Text: "&帮助文档",
						OnTriggered: func() {
							dir, err := os.Getwd()
							if err == nil {
								dir = strings.ReplaceAll(dir, "\\", "/")
								//RunBuiltinWebView(fmt.Sprintf("%s/help.html", dir))
							}
						},
					},
					Action{
						Text:        "&关于",
						OnTriggered: tmw.about,
					},
				},
			},
			Menu{
				AssignTo: &tmw.helpMenu,
				Text:     "&工具",
				Items: []MenuItem{
					Menu{
						AssignTo: &tmw.helpMenu,
						Text:     "&刷新",
						Items: []MenuItem{
							Action{
								Text:        "&主机",
								OnTriggered: tmw.about,
							},
							Action{
								Text:        "&日志",
								OnTriggered: tmw.about,
							},
							Action{
								Text:        "&权限",
								OnTriggered: tmw.about,
							},
							Action{
								Text:        "&表",
								OnTriggered: tmw.about,
							},
							Action{
								Text:        "&带只读锁的表",
								OnTriggered: tmw.about,
							},
							Action{
								Text:        "&状态",
								OnTriggered: tmw.about,
							},
						},
					},
					Separator{},
					Action{
						Text:        "&用户管理",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&维护",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&批量表编辑",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&运行命令行",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&Sequal Suggest",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Separator{},
					Action{
						Text:        "&导出数据库为SQL脚本",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&导出表格的行",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Separator{},
					Action{
						Text:        "&导入CSV文件",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&插入文件到TEXT/BLOB字段...",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
					Separator{},
					Action{
						Text:        "&首选项",
						Image:       "/icons/disconnect.png",
						OnTriggered: tmw.about,
					},
				},
			},
			Menu{
				AssignTo: &tmw.helpMenu,
				Text:     "&转到",
				Items: []MenuItem{
					Action{ // Shortcut{walk.ModControl | walk.ModShift , walk.KeyTab},
						Text:        "&上一标签页(P)",
						Image:       "/icons/disconnect.png",
						Shortcut:    Shortcut{Modifiers: walk.ModControl | walk.ModShift, Key: walk.KeyTab},
						OnTriggered: tmw.about,
					},
					Action{ // Shortcut{walk.ModControl | walk.ModShift , walk.KeyTab},
						Text:        "&下一标签页(P)",
						Image:       "/icons/disconnect.png",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.KeyTab},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&上一结果标签页(P)",
						Image:       "/icons/disconnect.png",
						Shortcut:    Shortcut{Modifiers: walk.ModAlt, Key: walk.KeyLeft},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&下一结果标签页(P)",
						Image:       "/icons/disconnect.png",
						Shortcut:    Shortcut{Modifiers: walk.ModAlt, Key: walk.KeyRight},
						OnTriggered: tmw.about,
					},
					Separator{},
					Action{
						Text:        "&表过滤器",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.KeyE},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&数据库树",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.KeyD},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&切换到查询/结果",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.KeyE},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&多列的过滤器",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.KeyE},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&标签页1",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.Key1},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&标签页2",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.Key2},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&标签页3",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.Key3},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&标签页4",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.Key4},
						OnTriggered: tmw.about,
					},
					Action{
						Text:        "&标签页5",
						Shortcut:    Shortcut{Modifiers: walk.ModControl, Key: walk.Key5},
						OnTriggered: tmw.about,
					},
				},
			},
			Menu{
				AssignTo: &tmw.helpMenu,
				Text:     "&帮助",
				Items: []MenuItem{
					Action{
						Text: "&帮助文档",
						OnTriggered: func() {
							dir, err := os.Getwd()
							if err == nil {
								dir = strings.ReplaceAll(dir, "\\", "/")
								//RunBuiltinWebView(fmt.Sprintf("%s/help.html", dir))
							}
						},
					},
					Action{
						Text:        "&关于",
						OnTriggered: tmw.about,
					},
				},
			},
		},
		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				//ActionRef{&openAction},
				Menu{
					Text:  "New A",
					Image: "/icons/disconnect.png",
					Items: []MenuItem{
						Action{
							Text: "A",
							//	OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text: "B",
							//OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text: "C",
							//OnTriggered: mw.newAction_Triggered,
						},
					},
					//OnTriggered: mw.newAction_Triggered,
				},
				Separator{},
				Menu{
					Text:  "View",
					Image: "../img/document-properties.png",
					Items: []MenuItem{
						Action{
							Text: "X",
							//OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text: "Y",
							//	OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text: "Z",
							//	OnTriggered: mw.changeViewAction_Triggered,
						},
					},
				},
				Separator{},
				Action{
					Text:    "Special",
					Image:   "../img/system-shutdown.png",
					Enabled: Bind("isSpecialMode && enabledCB.Checked"),
					//OnTriggered: mw.specialAction_Triggered,
				},
			},
		},
		StatusBarItems: []StatusBarItem{
			StatusBarItem{
				//AssignTo: &sbi,
				//Icon:     icon1,
				Text:  "已连接：15:30",
				Width: 80,
				OnClicked: func() {
					/*if sbi.Text() == "click" {
						sbi.SetText("again")
						sbi.SetIcon(icon2)
					} else {
						sbi.SetText("click")
						sbi.SetIcon(icon1)
					}*/
				},
			},
			StatusBarItem{
				Text:        "left",
				ToolTipText: "no tooltip for me",
			},
			StatusBarItem{
				Text: "\tcenter",
			},
			StatusBarItem{
				Text: "\t\tright",
			},
			StatusBarItem{
				//Icon:        icon1,
				ToolTipText: "An icon with a tooltip",
			},
		},
		Size: Size{1200, 742},
		Layout: VBox{
			MarginsZero: true, SpacingZero: true,
			//
		},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Composite{
						Layout: VBox{
							Margins:     Margins{Top: 3},
							MarginsZero: true,
							SpacingZero: true,
							Alignment:   AlignHVDefault,
							//	Alignment: AlignHCenterVCenter,
						},
						Children: []Widget{
							Composite{
								Layout: HBox{
									Margins:     Margins{Top: 3},
									MarginsZero: true,
									SpacingZero: true,
								},
								Children: []Widget{
									Label{
										//AssignTo: &vtc.statLbl,
										Text: "ddd",
										Font: Font{PointSize: 10},
									},
									HSpacer{},
									TextEdit{Text: "数据库过滤器"},
									TextEdit{Text: "表过滤器"},
								},
							},
							HSpacer{},
							TextEdit{Text: "abc"},
						},
					},
					Composite{
						Layout: VBox{
							Margins: Margins{Top: 3},
							//	MarginsZero: true,
							//	SpacingZero: true,
							//Alignment: AlignHVDefault,
							Alignment: AlignHCenterVCenter,
						},
						Children: []Widget{
							VSplitter{
								Children: []Widget{
									Composite{
										Layout: Grid{
											Margins:     Margins{Top: 3},
											MarginsZero: true,
											SpacingZero: true,
											//Alignment: AlignHVDefault,
											Alignment: AlignHCenterVCenter,
										},
										Children: []Widget{
											Composite{
												Layout: Grid{
													Margins:     Margins{Top: 3},
													MarginsZero: true,
													SpacingZero: true,
												},
												Children: []Widget{

													TextEdit{Text: "数据库过滤器"},
													TextEdit{Text: "表过滤器"},
												},
											},
											Label{
												Text: "名称:",
											},
											DateEdit{
												MaxSize:     Size{Width: 78},
												Date:        "2024/08/26",
												Format:      "yyyy/MM/dd",
												ToolTipText: "请选择查询开始日期",
											},
											//TextEdit{Text: "数据库过滤器aaaaaaaaaaaaaabbbbb"},
										},
									},
									Composite{
										Layout: HBox{
											Margins: Margins{Top: 3},
											//	MarginsZero: true,
											//	SpacingZero: true,
											//Alignment: AlignHVDefault,
											Alignment: AlignHCenterVCenter,
										},
										Children: []Widget{
											Composite{
												Layout: VBox{
													Margins:     Margins{Top: 3},
													MarginsZero: true,
													SpacingZero: true,
												},
												Children: []Widget{

													TableView{
														//AssignTo:         &cfc.previewCanTbl,
														AlternatingRowBG: true,
														//AlternatingRowBGColor: walk.RGB(239, 239, 239),
														ColumnsOrderable: true,
														Columns: []TableViewColumn{
															{Name: "Index", Title: "#", Frozen: true, Width: 60, Alignment: AlignCenter},
															{Name: "GroupName", Title: "分组名", Width: 120, Alignment: AlignCenter},
															{Name: "Sort", Title: "分组排序", Alignment: AlignCenter, Width: 60},
															{Name: "Remark", Title: "分组标识", Alignment: AlignCenter, Width: 60},
															{Name: "Chinesename", Title: "别名", Alignment: AlignCenter, Width: 140},
															{Name: "FieldName", Title: "字段名", Alignment: AlignCenter, Width: 140},
															{Name: "OutfieldId", Title: "CAN编码", Width: 60, Alignment: AlignFar},
															{Name: "Unit", Title: "单位", Alignment: AlignCenter, Width: 60},
															{Name: "DataType", Title: "数据类型", Alignment: AlignCenter, Width: 100, FormatFunc: func(value interface{}) string {
																switch value {
																case "1":
																	return "日期时间"
																case "2":
																	return "数字枚举"
																case "3":
																	return "数据"
																case "4":
																	return "其他"
																case "5":
																	return "文本枚举"
																case "6":
																	return "文本多枚举值"
																case "7":
																	return "多字段组合多枚举值"
																default:
																	return ""
																}
															}},
															{Name: "Formula", Title: "转换公式", Alignment: AlignCenter, Width: 100},
															{Name: "Decimals", Title: "小数位", Alignment: AlignCenter, Width: 50},
															{Name: "DataMap", Title: "值域", Alignment: AlignCenter, Width: 160},
															{Name: "IsAlarm", Title: "是否软报警", Alignment: AlignCenter, Width: 75, FormatFunc: func(value interface{}) string {
																switch value {
																case "0":
																	return ""
																case "1":
																	return "√"
																default:
																	return ""
																}
															}},
															{Name: "IsAnalysable", Title: "是否可分析", Alignment: AlignCenter, Width: 75, FormatFunc: func(value interface{}) string {
																switch value {
																case 0:
																	return ""
																case 1:
																	return "√"
																case 2:
																	return "√√"
																default:
																	return ""
																}
															}},
															{Name: "IsDelete", Title: "是否删除", Alignment: AlignCenter, Width: 75, FormatFunc: func(value interface{}) string {
																switch value {
																case 0:
																	return ""
																case 1:
																	return "√"
																default:
																	return ""
																}
															}},
															{Name: "OutfieldSn", Title: "CAN排序", Alignment: AlignCenter, Width: 75},
														},
													},
												},
											},
										},
									}},
							},
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	//tmw.SetX((utils.GetSystemMetrics(0) - tmw.WidthPixels()) / 2)
	//tmw.SetY((utils.GetSystemMetrics(1) - tmw.HeightPixels()) / 2)

	//InitVehicleTypePage(tmw)
	//tmw.newTab("设置")
	//tmw.newTab("我的")
	tmw.Run()

}

type TabMainWindow struct {
	*walk.MainWindow
	TabWidget      *walk.TabWidget
	targetPlatform *walk.Menu
	configMenu     *walk.Menu

	helpMenu *walk.Menu
}

func InitVehicleTypePage(tmw *TabMainWindow) {
	tp, err := walk.NewTabPage()
	if err != nil {
		log.Fatal(err)
	}
	tp.SetWidth(50)
	if (tp.SetTitle(" 主页 ")); err != nil {
		log.Fatal(err)
	}

	tp.SetLayout(walk.NewHBoxLayout())

	if err := tmw.TabWidget.Pages().Add(tp); err != nil {
		log.Fatal(err)
	}

	if err := tmw.TabWidget.SetCurrentIndex(tmw.TabWidget.Pages().Len() - 1); err != nil {
		log.Fatal(err)
	}

}

func (mw *TabMainWindow) openCanConfig(url string) func() {
	return func() {
		//RunBuiltinWebView(url)
	}
}

// 创建新页面
func (mw *TabMainWindow) newTab(tabTitle string) *walk.TabPage {
	tp, err := walk.NewTabPage()
	if err != nil {
		log.Fatal(err)
	}

	if (tp.SetTitle(tabTitle)); err != nil {
		log.Fatal(err)
	}

	menu, err := walk.NewMenu()
	closeAction := walk.NewAction()
	closeAction.SetText("关闭")
	closeAction.Triggered().Attach(func() {
		mw.TabWidget.Pages().Remove(tp)
		tp.Dispose()
	})
	menu.Actions().Add(closeAction)

	tp.SetContextMenu(menu)
	tp.SetLayout(walk.NewHBoxLayout())

	if err := mw.TabWidget.Pages().Add(tp); err != nil {
		log.Fatal(err)
	}

	if err := mw.TabWidget.SetCurrentIndex(mw.TabWidget.Pages().Len() - 1); err != nil {
		log.Fatal(err)
	}

	return tp
}
func (mw *TabMainWindow) about() {
	walk.MsgBox(mw, "", "SQLTOY\r\n\r\n新一代数据库客户端\r\n", walk.MsgBoxIconInformation)
}

func (mw *TabMainWindow) open() {
	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "SQL文件 (*.sql)|*.sql|文本文件 (*.txt)|*.txt|所有文件 (*.*)|*.*"

	//mw.edit.SetText("") //通过重定向变量设置TextEdit的Text
	if ok, err := dlg.ShowOpen(mw); err != nil {

		walk.MsgBox(mw, "Open", dlg.FilePath, walk.MsgBoxIconInformation)
		//mw.edit.AppendText("Error : File Open\r\n")
		return
	} else if !ok {
		//	mw.edit.AppendText("Cancel\r\n")
		return
	}
}
