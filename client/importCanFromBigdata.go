package client

import (
	"math/rand"

	"time"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

// ImportCanFromBigdataPanel 从大数据XML导入CAN
func ImportCanFromBigdataPanel(parent walk.Container, mainWin *TabMainWindow) (*CanAddPage, error) {
	rand.Seed(time.Now().UnixNano())

	var bigdataCan *SearchCanTable
	//var url string

	vtec := &CanAddPage{
		//searchModel:    &model.SearchModel{Items: []model.SearchItem{}},
		//targetCanModel: &TargetCanTable{items: []*model.CanDetail{}},
		searchCanModel: bigdataCan,
		//vehicleType:    vt,
		mainWin: mainWin,
	}

	//ef := &editForm{}

	//	fieldExistIcon, _ := walk.Resources.Icon("/img/warn.ico")

	if err := (Composite{
		AssignTo: &vtec.Composite,
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Composite{
						Layout:        VBox{MarginsZero: true, Margins: Margins{Right: 3}},
						StretchFactor: 2,
						Children: []Widget{
							Composite{
								Layout: HBox{MarginsZero: true},
								Children: []Widget{
									ComboBox{
										AssignTo:      &vtec.urlCb,
										BindingMember: "Value",
										DisplayMember: "Key",
										CurrentIndex:  0,
										//	Model:              utils.KnownConfigUrl(),
										Alignment:          AlignHNearVCenter,
										AlwaysConsumeSpace: true,
										OnCurrentIndexChanged: func() {

										},
									},
									ImageView{
										Image:       "/img/search1.ico",
										ToolTipText: "搜索",
										Alignment:   AlignHFarVCenter,
										OnMouseDown: func(x, y int, button walk.MouseButton) {

											/*	urls := vtec.urlCb.Model().([]*utils.KeyValuePair2)
												ef.ConfigUrl = urls[vtec.urlCb.CurrentIndex()].Value

												if cmd, err := ef.runExportDialog(vtec.mainWin); err != nil {
													mynotify.Error("CAN关键词搜索：" + err.Error())
												} else if cmd == walk.DlgCmdOK {
													vtec.mappingCan(ef.SearchStr, ef.ConfigUrl)
												}*/
										},
									},
								}},
							TableView{
								AssignTo:         &vtec.searchCanTv,
								AlternatingRowBG: true,
								//AlternatingRowBGColor: walk.RGB(239, 239, 239),
								CheckBoxes:       true,
								ColumnsOrderable: true,
								MultiSelection:   false,
								Columns: []TableViewColumn{
									{Title: "#", Width: 50, Frozen: true},
									{Title: "中文名"},
									{Title: "编号", Width: 50},
									{Title: "英文名"},
									{Title: "单位", Width: 50},
								},
								StyleCell: func(style *walk.CellStyle) {
									//	item := vtec.searchCanModel.items[style.Row()]

									/*	if item.Checked {
										if style.Row()%2 == 0 {
											style.BackgroundColor = walk.RGB(159, 215, 255)
										} else {
											style.BackgroundColor = walk.RGB(143, 199, 239)
										}
									}*/

								},
								Model: vtec.searchCanModel,
							},
						},
					},
					Composite{
						Layout:        VBox{MarginsZero: true, Margins: Margins{Left: 3}},
						StretchFactor: 8,
						Children: []Widget{
							Composite{
								Layout: HBox{MarginsZero: true},
								Children: []Widget{
									Label{
										//Text: fmt.Sprintf("%s - %s", vt.OrgName, vt.TypeName),
									},
									HSpacer{},
									ComboBox{
										AssignTo:      &vtec.groupSel,
										BindingMember: "Id",
										DisplayMember: "Name",
										MinSize:       Size{Width: 160},
										OnCurrentIndexChanged: func() {
											/*	vtec.listCan()
												vtec.reloadTargetList()*/
										},
									},
									PushButton{
										Image: "/img/import.ico",
										Text:  "导入",
										OnClicked: func() {

											//判断是否存在分组
											if vtec.groupSel.CurrentIndex() == -1 {
												walk.MsgBox(vtec.mainWin, "", "请选择分组", walk.MsgBoxIconWarning)
												return
											}

											/*	// 分组ID
												ddlCanGroup := vtec.groupSel.Model().([]*dataSource.CanGroupEntity)
												g := ddlCanGroup[vtec.groupSel.CurrentIndex()]

												// 过滤掉已存在的字段
												temp := []*model.CanDetail{}
												for _, v := range vtec.targetCanModel.items {
													if _, isExist := vtec.groupCanMap[v.OutfieldId]; !isExist {
														temp = append(temp, v)
													}
												}

												if len(temp) == 0 {
													walk.MsgBox(vtec.mainWin, "", "无新增字段", walk.MsgBoxIconWarning)
													return
												}

												// 批量插入字段
												err := (&dataSource.VehicleTypeEntity{}).InsertCanFromBigdata(g.Id, temp)
												if err != nil {
													walk.MsgBox(vtec.mainWin, "执行失败", err.Error(), walk.MsgBoxIconError)
													return
												}*/

											walk.MsgBox(vtec.mainWin, "", "执行成功，如需要编辑，请进编辑界面继续操作", walk.MsgBoxIconInformation)
										},
									},
								},
							},
							TableView{
								AssignTo:           &vtec.targetCanTv,
								AlwaysConsumeSpace: true,
								AlternatingRowBG:   true,
								//AlternatingRowBGColor: walk.RGB(239, 239, 239),
								CheckBoxes:       true,
								ColumnsOrderable: true,
								MultiSelection:   true,
								Columns: []TableViewColumn{
									{Title: "#", Width: 50, Frozen: true},
									{Title: "编号", Width: 60, Alignment: AlignFar},
									{Title: "中文名", Width: 120, Alignment: AlignCenter},
									{Title: "别名", Width: 120, Alignment: AlignCenter},
									{Title: "单位", Width: 50, Alignment: AlignCenter},
									{Title: "数据类型", Width: 80, Alignment: AlignCenter},
									{Title: "转换公式", Width: 80, Alignment: AlignCenter},
									{Title: "数值范围", Width: 160, Alignment: AlignCenter},
									{Title: "小数位", Width: 50, Alignment: AlignCenter},
									{Title: "软报警项", Width: 75, Alignment: AlignCenter, FormatFunc: func(value interface{}) string {
										switch value {
										case "0":
											return ""
										case "1":
											return "报警项"
										default:
											return ""
										}
									}},
									{Title: "可分析项", Width: 85, Alignment: AlignCenter, FormatFunc: func(value interface{}) string {
										switch value {
										case 0:
											return ""
										case 1:
											return "可分析"
										default:
											return ""
										}
									}},
									{Title: "排序", Width: 50, Alignment: AlignCenter},
								},
								StyleCell: func(style *walk.CellStyle) {
									//	item := vtec.targetCanModel.items[style.Row()]

									/*if item.Checked {
										if style.Row()%2 == 0 {
											style.BackgroundColor = walk.RGB(159, 215, 255)
										} else {
											style.BackgroundColor = walk.RGB(143, 199, 239)
										}
									}
									switch style.Col() {
									case 1:
										idx := strings.Index(item.Note, "已存在")
										if idx != -1 {
											//style.TextColor = walk.RGB(255, 0, 0)
											style.Image = fieldExistIcon
										}
									}*/
								},
								Model: vtec.targetCanModel,
							},
						},
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true, Margins: Margins{Top: 3}},
				Children: []Widget{
					Label{
						AssignTo: &vtec.searchStatLbl,
						//Text:     fmt.Sprintf("共 %d 项", len(vtec.searchCanModel.items)),
						Font: Font{PointSize: 10},
					},
					HSpacer{},
					ImageView{Image: "/img/warn.ico"},
					Label{Text: "：当前车系下，已存在该字段，将不会导入"},
				},
			},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	/*data, err := (&dataSource.VehicleTypeEntity{TypeId: vt.TypeId}).GetVehicleType()
	if err != nil {
		return vtec, err
	}

	if len(data.CanGroup) > 0 {
		vtec.groupSel.SetModel(data.CanGroup)
		vtec.groupSel.SetCurrentIndex(0)
	}
	*/
	vtec.searchCanModel.parent = vtec
	vtec.targetCanModel.parent = vtec

	return vtec, nil
}

// VehicleTypeAddPage 页面对象
type CanAddPage struct {
	*walk.Composite

	//vehicleType *model.VehicleTypeStats
	mainWin *TabMainWindow

	urlCb *walk.ComboBox
	// 搜索关键词模型
	//	searchModel *model.SearchModel

	// 字段搜索组件
	searchCanTv    *walk.TableView
	searchCanModel *SearchCanTable

	// 目标组件
	groupSel       *walk.ComboBox
	targetCanTv    *walk.TableView
	targetCanModel *TargetCanTable

	// 底部汇总组件
	searchStatLbl *walk.Label

	// 存储当前分组下已存在的Can信息 outfieldId -> *model.CanDetail
	//groupCanMap map[string]*model.CanDetail
}

// 过滤can
type SearchCanTable struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	//items      []*model.CanDetail

	parent *CanAddPage
}

// ResetRows 排序
func (m *SearchCanTable) ResetRows() {

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}

// TargetCanTable 目标can
type TargetCanTable struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	//items      []*model.CanDetail

	parent *CanAddPage
}

// LoadTargetCan 目标Table
func LoadTargetCan() *TargetCanTable {
	m := new(TargetCanTable)
	m.ResetRows()
	return m
}

// NewTargetCanModel
func NewTargetCanModel() *TargetCanTable {
	m := new(TargetCanTable)
	m.ResetRows()
	return m
}

// RowCount Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.

// Value Called by the TableView when it needs the text to display for a given cell.

// SetChecked Called by the TableView when the user toggled the check box of a given row.

// ResetRows
func (m *TargetCanTable) ResetRows() {
	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
