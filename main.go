// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/techyang/sqltoy/client"
	"math/rand"
	"strings"
	"time"
)

import (
	"github.com/lxn/walk"
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
