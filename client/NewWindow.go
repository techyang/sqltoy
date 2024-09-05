package client

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// CustomWindow 结构体定义了新窗口
type CustomWindow struct {
	dialog   *walk.Dialog   // 窗口的实际对象
	lineEdit *walk.LineEdit // 输入框组件
	owner    walk.Form      // 父窗口，用于管理新窗口的拥有关系
}

// NewCustomWindow 创建并返回一个 CustomWindow 的实例
func NewCustomWindow(owner walk.Form) *CustomWindow {
	return &CustomWindow{
		owner: owner,
	}
}

// Run 显示并运行新窗口
func (cw *CustomWindow) Run() {
	// 创建新窗口的对话框
	Dialog{
		AssignTo: &cw.dialog,
		Title:    "新窗口",
		MinSize:  Size{Width: 200, Height: 150},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text: "请输入内容：",
			},
			LineEdit{
				AssignTo: &cw.lineEdit,
			},
			PushButton{
				Text: "确定",
				OnClicked: func() {
					walk.MsgBox(cw.owner, "输入", cw.lineEdit.Text(), walk.MsgBoxIconInformation)
				},
			},
		},
	}.Run(cw.owner)
}
