// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/techyang/sqltoy/client"
	"log"
	"math/rand"
	"strings"
	"time"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

type MyMainWindow struct {
	*walk.MainWindow
	model *EnvModel
	lb    *walk.ListBox
	te    *walk.TextEdit
}

func main() {
	client.InitWin()
}
func main1() {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

	//mw := new(MyMainWindow)
	mw := &MyMainWindow{model: NewEnvModel()}
	var openAction, showAboutBoxAction *walk.Action
	var recentMenu *walk.Menu
	//var toggleSpecialModePB *walk.PushButton

	if err, _ := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Walk Actions Example",
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Image:       "../img/open.png",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: mw.openAction_Triggered,
					},
					Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					Separator{},
					Action{
						Text:        "E&xit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&View",
				Items: []MenuItem{
					Action{
						Text:    "Open / Special Enabled",
						Checked: Bind("enabledCB.Visible"),
					},
					Action{
						Text:    "Open Hidden",
						Checked: Bind("openHiddenCB.Visible"),
					},
				},
			},
			Menu{
				Text: "帮助",
				Items: []MenuItem{
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "关于",
						OnTriggered: mw.showAboutBoxAction_Triggered,
						Shortcut:    Shortcut{walk.ModControl, walk.KeyH},
					},
					Separator{},
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "帮助",
						OnTriggered: mw.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				ActionRef{&openAction},
				Menu{
					Text:  "New A",
					Image: "../img/document-new.png",
					Items: []MenuItem{
						Action{
							Text:        "A",
							OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text:        "B",
							OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text:        "C",
							OnTriggered: mw.newAction_Triggered,
						},
					},
					OnTriggered: mw.newAction_Triggered,
				},
				Separator{},
				Menu{
					Text:  "View",
					Image: "../img/document-properties.png",
					Items: []MenuItem{
						Action{
							Text:        "X",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:        "Y",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:        "Z",
							OnTriggered: mw.changeViewAction_Triggered,
						},
					},
				},
				Separator{},
				Action{
					Text:        "Special",
					Image:       "../img/system-shutdown.png",
					Enabled:     Bind("isSpecialMode && enabledCB.Checked"),
					OnTriggered: mw.specialAction_Triggered,
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showAboutBoxAction},
		},
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					ListBox{
						//AssignTo: &mw.lb,
						//	Model: NewEnvModel(),
						//OnCurrentIndexChanged: mw.lb_CurrentIndexChanged,
						//OnItemActivated:       mw.lb_ItemActivated,
					},
					TreeView{
						//AssignTo: &treeView,
						//Model:     createTreeModel(),
					},
					TableView{
						Name:             "tableView", // Name is needed for settings persistence
						AlternatingRowBG: true,
						ColumnsOrderable: true,
						Columns: []TableViewColumn{
							// Name is needed for settings persistence
							{Name: "#", DataMember: "Index"}, // Use DataMember, if names differ
							{Name: "Bar"},
							{Name: "Baz", Format: "%.2f", Alignment: AlignFar},
							{Name: "Quux", Format: "2006-01-02 15:04:05", Width: 150},
						},
						Model: NewFooModel(),
					},
				},
			},
		},
	}.Run()); err != 0 {
		log.Fatal(err)
	}

	/*addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openAction_Triggered)
			recentMenu.Actions().Add(a)
		}
	}*/

	//addRecentFileActions("Foo", "Bar", "Baz")

	mw.Run()
}

func (mw *MyMainWindow) openAction_Triggered() {
	//walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
	dlg := new(walk.FileDialog)
	dlg.Title = "选择文件"
	dlg.Filter = "可执行文件 (*.exe)|*.exe|所有文件 (*.*)|*.*"

	//mw.edit.SetText("") //通过重定向变量设置TextEdit的Text
	if ok, err := dlg.ShowOpen(mw); err != nil {

		//walk.MsgBox(mw, "Open", dlg.FilePaths, walk.MsgBoxIconInformation)
		//mw.edit.AppendText("Error : File Open\r\n")
		return
	} else if !ok {
		//	mw.edit.AppendText("Cancel\r\n")
		return
	}

}

func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialAction_Triggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}
func NewFooModel() *FooModel {
	now := time.Now()

	rand.Seed(now.UnixNano())

	m := &FooModel{items: make([]*Foo, 10)}

	for i := range m.items {
		m.items[i] = &Foo{
			Index: i,
			Bar:   strings.Repeat("*", rand.Intn(5)+1),
			Baz:   rand.Float64() * 1000,
			Quux:  time.Unix(rand.Int63n(now.Unix()), 0),
		}
	}

	return m
}

type FooModel struct {
	walk.SortedReflectTableModelBase
	items []*Foo
}

func (m *FooModel) Items() interface{} {
	return m.items
}

type Foo struct {
	Index int
	Bar   string
	Baz   float64
	Quux  time.Time
}

type EnvItem struct {
	name  string
	value string
}

type EnvModel struct {
	walk.ListModelBase
	items []*EnvItem
}

/*func NewEnvModel() *EnvModel {
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
}*/

func NewEnvModel2() *EnvModel {
	now := time.Now()

	rand.Seed(now.UnixNano())

	m := &EnvModel{items: make([]*EnvItem, 10)}

	for i := range m.items {
		m.items[i] = &EnvItem{
			name:  "abcd",
			value: "abc",
		}
	}

	return m
}

func NewEnvModel() *EnvModel {
	now := time.Now()

	rand.Seed(now.UnixNano())

	m := &EnvModel{items: make([]*EnvItem, 10)}

	for i := range m.items {
		m.items[i] = &EnvItem{

			name:  "abcd",
			value: "abc",
		}
	}

	return m
}

/*func createTreeModel() *walk.TreeViewModel {
	// 创建一个简单的树形视图模型
	root := walk.newfol("Root", "Root node")

	child1 := walk.NewFolderNode("Child1", "First child node")
	root.AddChild(child1)

	child2 := walk.NewFolderNode("Child2", "Second child node")
	root.AddChild(child2)

	// 为第一个子节点添加叶子节点
	leaf1 := walk.NewLeafNode("Leaf1", "Leaf node 1")
	child1.AddChild(leaf1)

	leaf2 := walk.NewLeafNode("Leaf2", "Leaf node 2")
	child1.AddChild(leaf2)

	// 创建树模型并添加根节点
	treeModel := walk.NewTreeViewModel(root)
	return treeModel
}
*/
