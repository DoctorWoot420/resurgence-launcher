package bhconfig

import (
	"github.com/therecipe/qt/core"
)

// Model Roles.
const (
	ConfigText = int(core.Qt__UserRole) + 1<<iota
)

// Model is the bhconfig model used for bhconfig data.
type Model struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Item                  `property:"items"`

	_ func(*Item) `slot:"addItem"`
	_ func()      `slot:"clear"`
}

func (m *Model) init() {
	m.SetRoles(map[int]*core.QByteArray{
		ConfigText: core.NewQByteArray2("title", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectAddItem(m.addItem)
	m.ConnectClear(m.clear)
}

func (m *Model) rowCount(*core.QModelIndex) int {
	return len(m.Items())
}

func (m *Model) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *Model) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Items()) {
		return core.NewQVariant()
	}

	item := m.Items()[index.Row()]

	switch role {
	case ConfigText:
		return core.NewQVariant1(item.ConfigText)
	default:
		return core.NewQVariant()
	}
}

// addItem adds an item to the model.
func (m *Model) addItem(c *Item) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Items()), len(m.Items()))
	m.SetItems(append(m.Items(), c))
	m.EndInsertRows()
}

func (m *Model) clear() {
	m.BeginResetModel()
	m.SetItems([]*Item{})
	m.EndResetModel()
}

func init() {
	Model_QRegisterMetaType()
	Item_QRegisterMetaType()
}
