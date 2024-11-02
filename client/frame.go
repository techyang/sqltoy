package client

import (
	"database/sql"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"os"
	"strings"
	"time"
)

func NewTreeViewModel(rootNodes []*TreeNode) *TreeViewModel {
	tvm := new(TreeViewModel)
	tvm.rootNodes = rootNodes
	return tvm
}

// 创建一个示例数据结构
func createSampleData() []*TreeNode {
	// 第一级：数据库会话名称
	session1 := &TreeNode{Title: "Session 1"}
	session2 := &TreeNode{Title: "Session 2"}

	// 第二级：数据量列表
	dataList1 := &TreeNode{Title: "Data List 1"}
	dataList2 := &TreeNode{Title: "Data List 2"}

	// 第三级：表列表
	tableList1 := &TreeNode{Title: "Table 1"}
	tableList2 := &TreeNode{Title: "Table 2"}

	// 构建树结构
	session1.Children = []*TreeNode{dataList1, dataList2}
	session2.Children = []*TreeNode{&TreeNode{Title: "Data List 3"}}
	dataList1.Children = []*TreeNode{tableList1, tableList2}

	return []*TreeNode{session1, session2}
}

type MyTreeView struct {
	walk.TreeView
	//items []OrganizationTreeModel
}

type VehicleTypeAddForm struct {
	*walk.Composite

	parent  walk.Container
	mainWin *TabMainWindow

	scroll *walk.ScrollView

	treeView            *walk.TreeView
	orgNameEdit         *walk.LineEdit
	vehicleTypeNameEdit *walk.LineEdit
	//notIntelligentRB         *walk.RadioButton
	//intelligentRB            *walk.RadioButton
	notFilterMissingColumnRB *walk.RadioButton
	filterMissingColumnRB    *walk.RadioButton

	canGroup *walk.Composite

	chb []*walk.CheckBox
	le  []*walk.LineEdit
	ne  []*walk.NumberEdit

	addVhBtn *walk.PushButton
}

func NewOrganizationTreeModel2() *OrganizationTreeModel {
	model := new(OrganizationTreeModel)

	org1 := NewOrganization("1", "abc", nil)
	org11 := NewOrganization("11", "abc11", org1)
	org12 := NewOrganization("12", "abc22", org1)
	org11.Text()
	org12.Text()
	model.roots = append(model.roots, org1)
	model.roots = append(model.roots, NewOrganization("2", "dsfsdfsd", nil))
	/*	orgs: = new Organization()
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(orgs); i++ {
			if orgs[i] == nil {
				continue
			}
			model.roots = append(model.roots, NewOrganization(orgs[i].Id, orgs[i].OrgName, nil))
		}*/

	return model
}

func NewOrganizationTreeModel() (*OrganizationTreeModel, error) {
	model := new(OrganizationTreeModel)

	/*	orgs, err := (&OrganizationEntity{SearchKey: ""}).ListSubOrg()
		if err != nil {
			return nil, err
		}*/

	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := `SHOW DATABASES `

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows *sql.Rows
	rows, err = stmt.Query()

	if err != nil {
		return nil, err
	}

	var databases []string
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			log.Fatal(err)
		}
		databases = append(databases, dbName)
		model.roots = append(model.roots, NewOrganization(dbName, dbName, nil))
		fmt.Println("- " + dbName)
	}

	/*for i := 0; i < len(orgs); i++ {
		if orgs[i] == nil {
			continue
		}
		model.roots = append(model.roots, NewOrganization(orgs[i].Id, orgs[i].OrgName, nil))
	}*/

	return model, nil
}

// 假设 tv 是一个 *walk.TreeView 实例，model 是一个 walk.TreeModel 实例

func (tv *MyTreeView) expandAllNodes(model walk.TreeModel) {
	// 获取子节点的数量
	count := model.RootCount()

	for i := 0; i < count; i++ {
		// 获取子节点
		child := model.RootAt(i)

		// 展开该节点
		tv.SetExpanded(child, true)

		// 递归展开所有子节点
		//expandAllNodes( model)
	}
}

var lnTextEdit LineNumberTextEdit

// MySQL connection parameters
const (
	USER     = "root"          // MySQL username
	PASSWORD = "password"      // MySQL password
	HOST     = "localhost"     // MySQL server host
	DATABASE = "your_database" // Database name
)

// DataModel to be used in TableView
type DataModel struct {
	walk.TableModelBase
	items   []map[string]interface{}
	columns []string
}

func (m *DataModel) RowCount() int {
	return len(m.items)
}

func (m *DataModel) Value(row, col int) interface{} {
	value := m.items[row][m.columns[col]]
	if byteValue, ok := value.([]byte); ok {
		return safeString(byteValue) // Convert byte slice to a safe string
	}
	return value
}

// safeString converts []byte to string and removes any NUL character if present
func safeString(value []byte) string {
	return strings.ReplaceAll(string(value), "\x00", "")
}
func (m *DataModel) ColumnCount() int {
	return len(m.columns)
}

func (m *DataModel) ResetData(newItems []map[string]interface{}, newColumns []string) {
	m.items = newItems
	m.columns = newColumns
	m.PublishRowsReset()
	m.PublishRowsChanged(0, len(newItems)-1)
}

// Fetch data from MySQL using the provided query
func fetchDataFromMySQL(query string) ([]map[string]interface{}, []string, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", USER, PASSWORD, HOST, DATABASE)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	var results []map[string]interface{}

	// Iterate over the rows
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))

		for i := range values {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, nil, err
		}

		row := make(map[string]interface{})
		for i, colName := range columns {
			row[colName] = values[i]
		}

		results = append(results, row)
	}

	return results, columns, nil
}

// InitWin 初始化界面
func InitWin() {

	tvm := new(TreeViewModel)
	tvm.rootNodes = createSampleData()

	tmw := new(TabMainWindow)
	var dataModel = new(DataModel)
	/*	var model = new(OrganizationTreeModel)
		p := NewOrganization("222", "orgs[i].OrgName2222", nil)
		NewOrganization("222-1", "orgs[i].OrgName2222", p)
		NewOrganization("222-2", "orgs[i].OrgName2222", p)
		model.roots = append(model.roots, NewOrganization("111", "orgs[i].OrgName111", nil))
		model.roots = append(model.roots, p)*/
	//treeView.SetModel(model)
	//vtaf := &VehicleTypeAddForm{}
	treeModel, err := NewOrganizationTreeModel()
	if err != nil {
		log.Fatal(err)
	}
	//vtaf.treeView.SetModel(treeModel)
	var treeView = new(MyTreeView)
	treeView.SetModel(treeModel)
	treeView.expandAllNodes(treeModel)
	var tableView *walk.TableView
	var textEdit *walk.TextEdit
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
						OnTriggered: func() {
							// 创建一个新窗口的实例并打开它
							newWin := NewCustomWindow(tmw)
							newWin.Run()

							/*	tp := tmw.newTab(fmt.Sprintf("【%s】从大数据导入", "配置窗口"))
								_, err := ImportCanFromBigdataPanel(tp, tmw)
								if err != nil {
									//mynotify.Error("窗口初始化失败," + err.Error())
								}*/

							//ImportCanFromBigdataPanel(tmw)
							//	tmw.DBSessionList(tmw)
						},
					},
					Action{
						Text: "&连接到",
						//	OnTriggered: tmw.openRelatedInstructionsPanel,
					},
					Action{
						Text:     "&新建窗口",
						Shortcut: Shortcut{walk.ModControl, walk.KeyN},
					},
					Action{
						Text:        "&新建查询标签页",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyT},
						OnTriggered: tmw.AddVehicleType,
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
						Text:        "&保存111",
						Shortcut:    Shortcut{walk.ModControl, walk.KeyS},
						OnTriggered: tmw.save,
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
						OnTriggered: tmw.close,
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
					Text:  "会话管理器",
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
					Text:  "运行",
					Image: "../img/document-properties.png",
					Items: []MenuItem{
						Action{
							Text:     "运行(V)",
							Shortcut: Shortcut{Key: walk.KeyF9},
							//OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:     "执行选择的部分(W)",
							Shortcut: Shortcut{Key: walk.KeyF1},
							//	OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text: "执行当前查询(X)",
							//	OnTriggered: mw.changeViewAction_Triggered,
							Shortcut: Shortcut{Modifiers: walk.ModControl | walk.ModShift, Key: walk.KeyF9},
						},
						Separator{},
						Action{
							Text:    "逐个发上查询(Y)",
							Checked: true,
							//	OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text: "一次性发送批处理(Z)",
							//	OnTriggered: mw.changeViewAction_Triggered,
						},
					},
				},
				Separator{},
				Action{
					Text:    "Special",
					Image:   "../img/system-shutdown.png",
					Enabled: Bind("isSpecialMode && enabledCB.Checked"),
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
						Layout: Grid{
							Columns: 2,
							/*	Margins:     Margins{Top: 3},
								MarginsZero: true,
								SpacingZero: true,
								Alignment:   AlignHVDefault,*/
							//	Alignment: AlignHCenterVCenter,
						},
						Children: []Widget{
							Composite{
								Layout: Grid{
									Columns: 2,
								},
								Children: []Widget{

									TextEdit{Text: "数据库过滤器", MaxSize: Size{Width: 20, Height: 10}, RowSpan: 1},
									TextEdit{Text: "表过滤器"},
									TreeView{
										//AssignTo: &treeView,
										Model: treeView.Model(),
										OnCurrentItemChanged: func() {
											//org := treeView.CurrentItem().(*Organization)
											//log.Print(org)
										},
										ColumnSpan: 2,
										MinSize:    Size{Height: 500},
									},

									HSpacer{RowSpan: 500},
									HSpacer{},
								},
							},
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
												Date:        time.Now(),
												Format:      "yyyy/MM/dd",
												ToolTipText: "请选择查询开始日期",
											},
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
													TabWidget{
														//Background: TransparentBrush{},
														AssignTo: &tmw.TabWidget,
													},
													TextEdit{
														Text:     "select * from infra_job;",
														AssignTo: &textEdit,
													},
													PushButton{
														Text: "Execute Query",
														OnClicked: func() {
															query := textEdit.Text()
															data, columns, err := fetchDataFromMySQL(query)
															if err != nil {
																walk.MsgBox(tmw, "Error", fmt.Sprintf("Query failed: %s", err), walk.MsgBoxIconError)
																return
															}

															// Reset the data model
															dataModel.ResetData(data, columns)

															// Clear existing columns
															tableView.Columns().Clear()

															// Dynamically create columns based on query result
															// Dynamically create columns based on query result
															for _, colName := range columns {
																column := walk.NewTableViewColumn()
																column.SetTitle(colName)
																tableView.Columns().Add(column)
															}
														},
													},
													TableView{
														AssignTo:         &tableView,
														AlternatingRowBG: true,
														//AlternatingRowBGColor: walk.RGB(239, 239, 239),
														ColumnsOrderable: true,
														Model:            dataModel,
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

	//InitVehicleTypePage(tmw)
	//tmw.newTab("设置")
	//tmw.newTab("我的")
	tmw.AddVehicleType()
	tmw.Run()

}

type TabMainWindow struct {
	*walk.MainWindow
	TabWidget      *walk.TabWidget
	targetPlatform *walk.Menu
	configMenu     *walk.Menu
	fileMenu       *walk.Menu
	helpMenu       *walk.Menu
}

func (mw *TabMainWindow) AddVehicleType() {
	tp := mw.newTab("查询  X")
	_, err := NewVehicleTypeAddPage(tp, mw)
	if err != nil {
		//mynotify.Error("窗口初始化失败," + err.Error())
	}
}

func NewVehicleTypeAddPage(parent walk.Container, mw *TabMainWindow) (*VehicleTypeAddForm, error) {
	vtaf := &VehicleTypeAddForm{}

	if err := (Composite{
		AssignTo: &vtaf.Composite,
		Layout:   HBox{},
		Children: []Widget{
			TextEdit{
				Text: "select * from tableName",
			},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	vtaf.mainWin = mw
	vtaf.parent = parent

	go func() {
		//treeModel, err := NewOrganizationTreeModel()
		//if err != nil {
		//log.Fatal(err)
		//}
		//vtaf.treeView.SetModel(treeModel)
	}()

	return vtaf, nil
}

func (mw *TabMainWindow) openCanConfig(url string) func() {
	return func() {
		//RunBuiltinWebView(url)

	}
}

// 创建新tab页面
func (mw *TabMainWindow) newTab(tabTitle string) *walk.TabPage {
	tp, err := walk.NewTabPage()
	// 错误处理
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

// 创建新页面
func (mw *TabMainWindow) newTab2(tabTitle string) *walk.TabPage {
	tabPage, _ := walk.NewTabPage()
	tabPage.SetTitle(tabTitle)

	// 创建一个关闭按钮
	closeButton, _ := walk.NewPushButton(tabPage)
	closeButton.SetText("x")
	closeButton.Clicked().Attach(func() {
		index := mw.TabWidget.Pages().Index(tabPage)
		if index != -1 {
			mw.TabWidget.Pages().Remove(tabPage) // 移除选项卡
		}
	})

	// 创建一个水平布局
	composite, _ := walk.NewComposite(tabPage)
	layout := walk.NewHBoxLayout()
	composite.SetLayout(layout)

	// 将标签和关闭按钮添加到 Composite 中
	titleLabel, _ := walk.NewLabel(composite)
	titleLabel.SetText(tabTitle)

	layout.StretchFactor(titleLabel)
	layout.StretchFactor(closeButton)
	//layout.AddWidget(closeButton)

	return tabPage
}

func (mw *TabMainWindow) about() {
	walk.MsgBox(mw, "", "SQLTOY\r\n\r\n新一代数据库客户端\r\n", walk.MsgBoxIconInformation)
}

// close the app
func (mw *TabMainWindow) close() {
	mw.Dispose()

}

func (mw *TabMainWindow) initFileMenu() {
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

func (mw *TabMainWindow) save() {
	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "SQL文件 (*.sql)|*.sql|文本文件 (*.txt)|*.txt|所有文件 (*.*)|*.*"

	//mw.edit.SetText("") //通过重定向变量设置TextEdit的Text
	if ok, err := dlg.ShowSave(mw); err != nil {

		//walk.MsgBox(mw, "Open", dlg.FilePath, walk.MsgBoxIconInformation)
		//mw.edit.AppendText("Error : File Open\r\n")
		return
	} else if !ok {
		//	mw.edit.AppendText("Cancel\r\n")
		return
	}
}

type EnvItem struct {
	name  string
	value string
}

type EnvModel struct {
	walk.ListModelBase
	items []EnvItem
}

func NewEnvModel() *EnvModel {
	env := os.Environ()

	m := &EnvModel{items: make([]EnvItem, len(env))}

	for i, e := range env {
		j := strings.Index(e, "=")
		if j == 0 {
			continue
		}

		name := e[0:j]
		value := strings.Replace(e[j+1:], ";", "\r\n", -1)

		m.items[i] = EnvItem{name, value}
	}

	return m
}

func (m *EnvModel) ItemCount() int {
	return len(m.items)
}

func (m *EnvModel) Value(index int) interface{} {
	return m.items[index].name
}

// TreeNode 表示树的节点
type TreeNode struct {
	Title    string
	Children []*TreeNode
}

// TreeViewModel 用于将 TreeNode 数据绑定到 TreeView
type TreeViewModel struct {
	walk.TreeModel
	rootNodes []*TreeNode
}

func (tvm *TreeViewModel) Count() int {
	return len(tvm.rootNodes)
}

func (tvm *TreeViewModel) Value(index int) interface{} {
	return tvm.rootNodes[index]
}

func (tvm *TreeViewModel) Parent(index int) int {
	node := tvm.rootNodes[index]
	if node == nil {
		return -1
	}
	return -1
}

func (tvm *TreeViewModel) Index(parent int, childIndex int, childCount int) int {
	if parent == -1 {
		return childIndex
	}
	parentNode := tvm.rootNodes[parent]
	if childIndex >= childCount {
		return -1
	}
	for i, child := range parentNode.Children {
		log.Print(child.Title)
		if i == childIndex {
			return len(tvm.rootNodes) + i
		}
	}
	return -1
}

func (tvm *TreeViewModel) ChildCount(index int) int {
	node := tvm.rootNodes[index]
	if node == nil {
		return 0
	}
	return len(node.Children)
}

func (tvm *TreeViewModel) Data() interface{} {
	return nil
}

type LineNumberTextEdit struct {
	*walk.Composite
	textEdit *walk.TextEdit
	listBox  *walk.ListBox
}

func (ln *LineNumberTextEdit) Create(builder *Builder) error {
	//TODO implement me
	//panic("implement me")
	return nil
}

func NewLineNumberTextEdit(parent walk.Container) (*LineNumberTextEdit, error) {
	lnTextEdit := new(LineNumberTextEdit)
	var err error

	if lnTextEdit.Composite, err = walk.NewComposite(parent); err != nil {
		return nil, err
	}

	layout := walk.NewHBoxLayout()
	layout.SetMargins(walk.Margins{5, 5, 5, 5})
	lnTextEdit.SetLayout(layout)

	if lnTextEdit.listBox, err = walk.NewListBox(lnTextEdit); err != nil {
		return nil, err
	}

	//lnTextEdit.listBox.SetFixedWidth(30)

	if lnTextEdit.textEdit, err = walk.NewTextEdit(lnTextEdit); err != nil {
		return nil, err
	}

	//lnTextEdit.textEdit.SetMultiline(true)
	lnTextEdit.textEdit.TextChanged().Attach(func() {
		lnTextEdit.updateLineNumbers()
	})
	lnTextEdit.textEdit.SetText("ddddddddddddddddd")
	return lnTextEdit, nil
}

func (ln *LineNumberTextEdit) updateLineNumbers() {
	text := ln.textEdit.Text()
	lines := strings.Count(text, "\n") + 1

	var lineNumbers []string
	for i := 1; i <= lines; i++ {
		lineNumbers = append(lineNumbers, "1")
	}

	ln.listBox.SetModel(lineNumbers)
}
