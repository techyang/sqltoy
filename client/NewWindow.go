package client

import (
	"bufio"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"os"
	"sort"
	"strings"
)

// CustomWindow 结构体定义了新窗口
type CustomWindow struct {
	dialog      *walk.Dialog   // 窗口的实际对象
	lineEdit    *walk.LineEdit // 输入框组件
	dbLinkTable *walk.TableView
	owner       walk.Form // 父窗口，用于管理新窗口的拥有关系
}

// NewCustomWindow 创建并返回一个 CustomWindow 的实例
func NewCustomWindow(owner walk.Form) *CustomWindow {
	return &CustomWindow{
		owner: owner,
	}
}

// Run 显示并运行新窗口
func (cw *CustomWindow) Run() {
	var comboBox *walk.ComboBox
	var dependenceComboBox *walk.ComboBox
	//var settings Settings
	// 选项数据

	options := []Option{
		{Key: "01", Value: "MariaDB or MySQL (TCP/IP)"},
		{Key: "02", Value: "MariaDB or MySQL (named pipe)"},
		{Key: "03", Value: "MariaDB or MySQL (SSH tunnel)"},
		{Key: "04", Value: "ProxySQL Admin (Experimental)"},
		{Key: "05", Value: "MySQL on RDS"},
		{Key: "06", Value: "Microsoft SQL Server (named pipe) "},
		{Key: "07", Value: "Microsoft SQL Server (TCP/P)"},
		{Key: "08", Value: "Microsoft SQL Server (SPX/IPX)"},
		{Key: "09", Value: "Microsoft SQL Server (Banyan VINES)"},
		{Key: "10", Value: "Microsoft SQL Server (Windows RPC)"},
		{Key: "11", Value: "PostgreSQL (TCP/IP)"},
		{Key: "12", Value: "PostgreSQL (SSH tunnel)"},
	}

	// 提取显示值以供 ComboBox 使用
	displayValues := setOptions(options)

	dependenceLibsOptions := []Option{
		{Key: "01", Value: "libmariadb.dll"},
		{Key: "02", Value: "libmysql-6.1.dll"},
		{Key: "03", Value: "libmysql.dll"},
	}

	// 提取显示值以供 ComboBox 使用
	dependenceValues := setOptions(dependenceLibsOptions)

	var checkBox1, checkBox2 *walk.CheckBox
	//var numberEdit *walk.NumberEdit

	sessions, err := readDataFromFile("data.txt")
	if err != nil {
		log.Fatalf("读取数据文件失败: %v", err)
	}
	model := NewSessionModel(sessions)
	// 创建新窗口的对话框
	Dialog{
		AssignTo: &cw.dialog,
		Title:    "会话管理器",
		MinSize:  Size{Width: 800, Height: 600},
		Layout:   VBox{MarginsZero: true, SpacingZero: true},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Composite{
						Layout: VBox{

							/*	Margins:     Margins{Top: 3},
								MarginsZero: true,
								SpacingZero: true,
								Alignment:   AlignHVDefault,*/
							//	Alignment: AlignHCenterVCenter,
						},
						Children: []Widget{
							Composite{
								Layout: Grid{
									Columns: 1,
									//Margins: Margins{Top: 3},
									/*	MarginsZero: true,
										SpacingZero: true,*/
								},
								Children: []Widget{
									/*Label{
										//AssignTo: &vtc.statLbl,
										Text: "ddd",
										Font: Font{PointSize: 10},
									},*/
									//HSpacer{},
									TextEdit{Text: "数据库过滤器", MaxSize: Size{Width: 20, Height: 20}, RowSpan: 1},

									TableView{
										AssignTo:         &cw.dbLinkTable,
										AlternatingRowBG: true,
										//AlternatingRowBGColor: walk.RGB(239, 239, 239),
										ColumnsOrderable: true, // 启用列的排序功能
										Model:            model,

										Columns: []TableViewColumn{
											{Name: "SessionName", Title: "会话名称", Frozen: true, Width: 60, Alignment: AlignCenter},
											{Name: "Host", Title: "主机", Width: 120, Alignment: AlignCenter},
											{Name: "LastConnectTime", Title: "上次连接", Alignment: AlignCenter, Width: 60},
											{Name: "Remark", Title: "注释", Alignment: AlignCenter, Width: 60},
										},
									},
								},
							},
							Composite{
								Layout: HBox{
									//Columns: 1,
									//Margins: Margins{Top: 3},
									/*	MarginsZero: true,
										SpacingZero: true,*/
								},
								Children: []Widget{
									/*Label{
										//AssignTo: &vtc.statLbl,
										Text: "ddd",
										Font: Font{PointSize: 10},
									},*/
									//HSpacer{},
									/*PushButton{Text: "数据库过滤器", MaxSize: Size{Width: 50, Height: 20}, RowSpan: 1},
									PushButton{Text: "数据库过滤器", MaxSize: Size{Width: 50, Height: 20}, RowSpan: 1},
									PushButton{Text: "数据库过滤器", MaxSize: Size{Width: 50, Height: 20}, RowSpan: 1},*/
									ToolBar{
										ButtonStyle: ToolBarButtonImageBeforeText,
										Items: []MenuItem{
											//ActionRef{&openAction},
											Menu{
												Text:  "新建",
												Image: "/icons/disconnect.png",
												Items: []MenuItem{
													Action{
														Text: "在根文件夹创建会话（W）",
														//	OnTriggered: mw.newAction_Triggered,
													},
													Action{
														Text: "在选定的文件夹创建会话（X）",
														//OnTriggered: mw.newAction_Triggered,
													},
													Action{
														Text: "在根文件夹下创建子文件夹",
														//OnTriggered: mw.newAction_Triggered,
													},
													Action{
														Text: "在选定的文件夹下创建子文件夹",
														//OnTriggered: mw.newAction_Triggered,
													},
												},
												//OnTriggered: mw.newAction_Triggered,
											},

											Action{
												Text:    "   保存    ",
												Image:   "../img/system-shutdown.png",
												Enabled: Bind("isSpecialMode && enabledCB.Checked"),
												//OnTriggered: mw.specialAction_Triggered,

											},

											Action{
												Text:    "    删除  ",
												Image:   "../img/system-shutdown.png",
												Enabled: Bind("isSpecialMode && enabledCB.Checked"),
												//OnTriggered: mw.specialAction_Triggered,
											},
										},
									},
								},
							},

							//HSpacer{},
							//TextEdit{Text: "abc"},

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
											// 创建 TabWidget
											TabWidget{
												Pages: []TabPage{
													// 第一个标签页
													TabPage{
														Title:  "设置",
														Layout: VBox{},
														Children: []Widget{
															Composite{
																Layout: Grid{
																	Margins:     Margins{Top: 3},
																	MarginsZero: true,
																	SpacingZero: true,
																	//Alignment: AlignHVDefault,
																	Alignment: AlignHCenterVCenter,
																	Columns:   2,
																},
																Children: []Widget{

																	Label{
																		Text: "网络类型：",
																		//	ToolTip: "Enter text for Tab 2",
																		MinSize: Size{Width: 100},
																	},
																	ComboBox{
																		AssignTo:     &comboBox,
																		Model:        displayValues, // 设置下拉列表的选项
																		CurrentIndex: 0,
																		OnCurrentIndexChanged: func() {
																			// 获取当前选择的选项
																			currentIndex := comboBox.CurrentIndex()
																			if currentIndex >= 0 {
																				fmt.Printf("Selected: %s\n", options[currentIndex])
																			}
																		},
																	},
																	Label{
																		Text: "依赖库：",
																		//	ToolTip: "Enter text for Tab 2",
																	},
																	ComboBox{
																		AssignTo:     &dependenceComboBox,
																		Model:        dependenceValues, // 设置下拉列表的选项
																		CurrentIndex: 0,
																		OnCurrentIndexChanged: func() {
																			// 获取当前选择的选项
																			currentIndex := comboBox.CurrentIndex()
																			if currentIndex >= 0 {
																				fmt.Printf("Selected: %s\n", options[currentIndex])
																			}
																		},
																	},
																	Label{
																		Text: "主机名/IP：",
																		//	ToolTip: "Enter text for Tab 2",
																	},
																	Composite{
																		Layout: VBox{
																			Margins:     Margins{Top: 3},
																			MarginsZero: true,
																			SpacingZero: true,
																			//Alignment: AlignHVDefault,
																			Alignment: AlignHVDefault,
																		},
																		Children: []Widget{
																			LineEdit{
																				Name: "TextBox3",
																				//	ToolTip: "Enter text for Tab 3",
																			},
																			CheckBox{
																				AssignTo:  &checkBox1,
																				Alignment: AlignHVDefault,
																				Text:      "提提示身份认证", // 第一个复选框的文本
																				Checked:   false,     // 设置第一个复选框的默认状态（未选中）
																				OnCheckedChanged: func() {
																					// 当第一个复选框状态改变时触发
																					if checkBox1.Checked() {
																						fmt.Println("Option 1 is checked.")
																					} else {
																						fmt.Println("Option 1 is unchecked.")
																					}
																				},
																			},
																			// 第二个复选框
																			CheckBox{
																				AssignTo: &checkBox2,
																				Text:     "使用Windows 认证", // 第二个复选框的文本
																				Checked:  false,          // 设置第二个复选框的默认状态（未选中）
																				OnCheckedChanged: func() {
																					// 当第二个复选框状态改变时触发
																					if checkBox2.Checked() {
																						fmt.Println("Option 2 is checked.")
																					} else {
																						fmt.Println("Option 2 is unchecked.")
																					}
																				},
																			},
																		},
																	},

																	Label{

																		Text: "用户：",
																		//	ToolTip: "Enter text for Tab 2",
																	},
																	LineEdit{
																		Name: "TextBox3",
																		//	ToolTip: "Enter text for Tab 3",
																	},
																	Label{
																		Text: "密码：",
																		//	ToolTip: "Enter text for Tab 2",
																	},
																	LineEdit{
																		Name:         "TextBox3",
																		PasswordMode: true, // 设置为密码模式，输入字符会显示为星号
																		//	ToolTip: "Enter text for Tab 3",
																	},

																	Label{
																		Text: "端口：",
																		//	ToolTip: "Enter text for Tab 2",
																	},
																	// 数字选择框

																	LineEdit{
																		Name: "TextBox3",
																		//PasswordMode: true, // 设置为密码模式，输入字符会显示为星号
																		//	ToolTip: "Enter text for Tab 3",
																	},
																	/*NumberEdit{
																		AssignTo:  &numberEdit,
																		Value:     3306,  // 初始值
																		MinValue:  1000,  // 最小值
																		MaxValue:  10000, // 最大值
																		Increment: 1,     // 每次增加的值

																	},*/
																	Label{
																		Text: "数据库：",
																		//	ToolTip: "Enter text for Tab 2",
																	},
																	LineEdit{
																		Name: "TextBox3",
																		//	ToolTip: "Enter text for Tab 3",
																	},
																	Label{
																		Text: "注释：",
																		//RowSpan: 3,
																		//	ToolTip: "Enter text for Tab 2",
																	},
																	TextEdit{
																		Name: "TextBox3",
																		//Text: "TextBox3",
																		//Row:     5,    // 显示5行高度
																		VScroll: true, // 垂直滚动条
																		MaxSize: Size{Height: 100},
																	},
																},
															},
															Composite{
																Layout: HBox{
																	Margins:     Margins{Top: 3},
																	MarginsZero: true,
																	SpacingZero: true,
																	//Alignment: AlignHVDefault,
																	Alignment: AlignHCenterVCenter,
																	//Columns:   5,
																},
																Children: []Widget{
																	PushButton{
																		Text:    "打开",
																		MinSize: Size{Width: 100},
																	},
																	PushButton{
																		Text:    "取消",
																		MinSize: Size{Width: 100},
																	},
																	Composite{
																		Layout: HBox{
																			Margins:     Margins{Top: 3},
																			MarginsZero: true,
																			SpacingZero: true,
																			//Alignment: AlignHVDefault,
																			Alignment: AlignHCenterVCenter,
																			//Columns:   5,

																		},
																		Children: []Widget{
																			ToolBar{
																				ButtonStyle: ToolBarButtonImageBeforeText,
																				Items: []MenuItem{
																					//ActionRef{&openAction},
																					Menu{
																						Text: "    更多    ",
																						//Image: "/icons/disconnect.png",
																						Items: []MenuItem{
																							Action{
																								Text: "首选项(V)",
																								//	OnTriggered: mw.newAction_Triggered,
																							},
																							Action{
																								Text: "检查更新(W)..",
																								//OnTriggered: mw.newAction_Triggered,
																							},
																							Action{
																								Text: "导入配置文件 (X)..",
																								//OnTriggered: mw.newAction_Triggered,
																							},
																							Action{
																								Text: "导出配置文件(Y)...",
																								//OnTriggered: mw.newAction_Triggered,
																							},
																							Action{
																								Text: "通用帮助(Z)",
																								//OnTriggered: mw.newAction_Triggered,
																							},
																							Action{
																								Text: "关于 SQLTOY 12.5.0.6677(H)",
																								//OnTriggered: mw.newAction_Triggered,
																							},
																						},
																						//OnTriggered: mw.newAction_Triggered,

																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
													// 第二个标签页
													TabPage{
														Title:  "SSH隧道",
														Layout: VBox{},
														Children: []Widget{
															LineEdit{
																Name: "TextBox2",
																//	ToolTip: "Enter text for Tab 2",
															},
															PushButton{
																Text: "Button 2",
																OnClicked: func() {
																	fmt.Println("Button 2 clicked")
																},
															},
														},
													},
													// 第三个标签页
													TabPage{
														Title:  "高级",
														Layout: VBox{},
														Children: []Widget{
															LineEdit{
																Name: "TextBox3",
																//	ToolTip: "Enter text for Tab 3",
															},
															PushButton{
																Text: "Button 3",
																OnClicked: func() {
																	fmt.Println("Button 3 clicked")
																},
															},
														},
													},
													// 第四个标签页
													TabPage{
														Title:  "SSL",
														Layout: VBox{},
														Children: []Widget{
															LineEdit{
																Name: "TextBox4",
																//ToolTip: "Enter text for Tab 4",
															},
															PushButton{
																Text: "Button 4",
																OnClicked: func() {
																	fmt.Println("Button 4 clicked")
																},
															},
														},
													},
													// 第四个标签页
													TabPage{
														Title:  "统计信息",
														Layout: VBox{},
														Children: []Widget{
															LineEdit{
																Name: "TextBox4",
																//ToolTip: "Enter text for Tab 4",
															},
															PushButton{
																Text: "Button 4",
																OnClicked: func() {
																	fmt.Println("Button 4 clicked")
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}.Run(cw.owner)
}

func setOptions(options []Option) []string {
	displayValues := make([]string, len(options))
	for i, option := range options {
		displayValues[i] = option.Value
	}
	return displayValues
}

// 定义数据模型结构体
type Session struct {
	SessionName     string // 会话名称
	Host            string // 主机
	LastConnectTime string // 上次连接
	Remark          string // 注释
}

// 定义 Session 模型
type SessionModel struct {
	walk.TableModelBase
	walk.SorterBase
	items      []Session
	sortColumn int            // 当前排序列的索引
	sortOrder  walk.SortOrder // 排序顺序（升序或降序）
}

// NewSessionModel 创建一个新的 SessionModel
func NewSessionModel(data []Session) *SessionModel {
	m := &SessionModel{items: data}
	m.Sort(0, walk.SortAscending) // 默认按第一列升序排序
	return m
}

// RowCount 返回数据行数
func (m *SessionModel) RowCount() int {
	return len(m.items)
}

// Value 返回指定行和列的值
func (m *SessionModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.SessionName
	case 1:
		return item.Host
	case 2:
		return item.LastConnectTime
	case 3:
		return item.Remark
	}
	return nil
}

// ColumnSortable 返回指定列是否可排序
func (m *SessionModel) ColumnSortable(col int) bool {
	return true
}

// Sort 按指定列和排序顺序对数据排序
func (m *SessionModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {
		var less bool
		switch m.sortColumn {
		case 0:
			less = m.items[i].SessionName < m.items[j].SessionName
		case 1:
			less = m.items[i].Host < m.items[j].Host
		case 2:
			less = m.items[i].LastConnectTime < m.items[j].LastConnectTime
		case 3:
			less = m.items[i].Remark < m.items[j].Remark
		}

		if m.sortOrder == walk.SortAscending {
			return less
		}
		return !less
	})

	m.SorterBase.Sort(col, order) // 更新排序状态
	m.PublishRowsReset()          // 通知视图更新
	return nil
}

// 读取文件并解析数据
func readDataFromFile(filename string) ([]Session, error) {
	var sessions []Session

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) < 4 {
			continue // 跳过格式不正确的行
		}

		session := Session{
			SessionName:     parts[0],
			Host:            parts[1],
			LastConnectTime: parts[2],
			Remark:          parts[3],
		}

		sessions = append(sessions, session)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

// 定义一个结构体来表示显示和实际取值
type Option struct {
	Key   string // 实际取值
	Value string // 显示值
}

type Settings struct {
	netType       string // 网络类型
	dependenceLib string // 依赖库
	hosts         string // 主机/IP
	user          string //用户
	port          string // 端口

}
