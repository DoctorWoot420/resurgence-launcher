//go:build !minimal
// +build !minimal

package winextras

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "winextras.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtWinExtras_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtWinExtras_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type QWinColorizationChangeEvent struct {
	QWinEvent
}

type QWinColorizationChangeEvent_ITF interface {
	QWinEvent_ITF
	QWinColorizationChangeEvent_PTR() *QWinColorizationChangeEvent
}

func (ptr *QWinColorizationChangeEvent) QWinColorizationChangeEvent_PTR() *QWinColorizationChangeEvent {
	return ptr
}

func (ptr *QWinColorizationChangeEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QWinColorizationChangeEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWinEvent_PTR().SetPointer(p)
	}
}

func PointerFromQWinColorizationChangeEvent(ptr QWinColorizationChangeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinColorizationChangeEvent_PTR().Pointer()
	}
	return nil
}

func NewQWinColorizationChangeEventFromPointer(ptr unsafe.Pointer) (n *QWinColorizationChangeEvent) {
	n = new(QWinColorizationChangeEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QWinColorizationChangeEvent) DestroyQWinColorizationChangeEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWinCompositionChangeEvent struct {
	QWinEvent
}

type QWinCompositionChangeEvent_ITF interface {
	QWinEvent_ITF
	QWinCompositionChangeEvent_PTR() *QWinCompositionChangeEvent
}

func (ptr *QWinCompositionChangeEvent) QWinCompositionChangeEvent_PTR() *QWinCompositionChangeEvent {
	return ptr
}

func (ptr *QWinCompositionChangeEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QWinCompositionChangeEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWinEvent_PTR().SetPointer(p)
	}
}

func PointerFromQWinCompositionChangeEvent(ptr QWinCompositionChangeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinCompositionChangeEvent_PTR().Pointer()
	}
	return nil
}

func NewQWinCompositionChangeEventFromPointer(ptr unsafe.Pointer) (n *QWinCompositionChangeEvent) {
	n = new(QWinCompositionChangeEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QWinCompositionChangeEvent) DestroyQWinCompositionChangeEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWinEvent struct {
	core.QEvent
}

type QWinEvent_ITF interface {
	core.QEvent_ITF
	QWinEvent_PTR() *QWinEvent
}

func (ptr *QWinEvent) QWinEvent_PTR() *QWinEvent {
	return ptr
}

func (ptr *QWinEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QWinEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQWinEvent(ptr QWinEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinEvent_PTR().Pointer()
	}
	return nil
}

func NewQWinEventFromPointer(ptr unsafe.Pointer) (n *QWinEvent) {
	n = new(QWinEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QWinEvent) DestroyQWinEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWinJumpList struct {
	core.QObject
}

type QWinJumpList_ITF interface {
	core.QObject_ITF
	QWinJumpList_PTR() *QWinJumpList
}

func (ptr *QWinJumpList) QWinJumpList_PTR() *QWinJumpList {
	return ptr
}

func (ptr *QWinJumpList) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWinJumpList) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWinJumpList(ptr QWinJumpList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinJumpList_PTR().Pointer()
	}
	return nil
}

func NewQWinJumpListFromPointer(ptr unsafe.Pointer) (n *QWinJumpList) {
	n = new(QWinJumpList)
	n.SetPointer(ptr)
	return
}
func NewQWinJumpList(parent core.QObject_ITF) *QWinJumpList {
	tmpValue := NewQWinJumpListFromPointer(C.QWinJumpList_NewQWinJumpList(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWinJumpList) AddCategory(category QWinJumpListCategory_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList_AddCategory(ptr.Pointer(), PointerFromQWinJumpListCategory(category))
	}
}

func (ptr *QWinJumpList) AddCategory2(title string, items []*QWinJumpListItem) *QWinJumpListCategory {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		return NewQWinJumpListCategoryFromPointer(C.QWinJumpList_AddCategory2(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: titleC, len: C.longlong(len(title))}, func() unsafe.Pointer {
			tmpList := NewQWinJumpListFromPointer(NewQWinJumpListFromPointer(nil).__addCategory_items_newList2())
			for _, v := range items {
				tmpList.__addCategory_items_setList2(v)
			}
			return tmpList.Pointer()
		}()))
	}
	return nil
}

func (ptr *QWinJumpList) Categories() []*QWinJumpListCategory {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWinExtras_PackedList) []*QWinJumpListCategory {
			out := make([]*QWinJumpListCategory, int(l.len))
			tmpList := NewQWinJumpListFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__categories_atList(i)
			}
			return out
		}(C.QWinJumpList_Categories(ptr.Pointer()))
	}
	return make([]*QWinJumpListCategory, 0)
}

//export callbackQWinJumpList_Clear
func callbackQWinJumpList_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinJumpListFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QWinJumpList) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinJumpList) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QWinJumpList) Clear() {
	if ptr.Pointer() != nil {
		C.QWinJumpList_Clear(ptr.Pointer())
	}
}

func (ptr *QWinJumpList) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QWinJumpList_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QWinJumpList) Frequent() *QWinJumpListCategory {
	if ptr.Pointer() != nil {
		return NewQWinJumpListCategoryFromPointer(C.QWinJumpList_Frequent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWinJumpList) Identifier() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWinJumpList_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWinJumpList) Recent() *QWinJumpListCategory {
	if ptr.Pointer() != nil {
		return NewQWinJumpListCategoryFromPointer(C.QWinJumpList_Recent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWinJumpList) SetIdentifier(identifier string) {
	if ptr.Pointer() != nil {
		var identifierC *C.char
		if identifier != "" {
			identifierC = C.CString(identifier)
			defer C.free(unsafe.Pointer(identifierC))
		}
		C.QWinJumpList_SetIdentifier(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: identifierC, len: C.longlong(len(identifier))})
	}
}

func (ptr *QWinJumpList) Tasks() *QWinJumpListCategory {
	if ptr.Pointer() != nil {
		return NewQWinJumpListCategoryFromPointer(C.QWinJumpList_Tasks(ptr.Pointer()))
	}
	return nil
}

//export callbackQWinJumpList_DestroyQWinJumpList
func callbackQWinJumpList_DestroyQWinJumpList(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWinJumpList"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinJumpListFromPointer(ptr).DestroyQWinJumpListDefault()
	}
}

func (ptr *QWinJumpList) ConnectDestroyQWinJumpList(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWinJumpList"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWinJumpList", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWinJumpList", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinJumpList) DisconnectDestroyQWinJumpList() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWinJumpList")
	}
}

func (ptr *QWinJumpList) DestroyQWinJumpList() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinJumpList_DestroyQWinJumpList(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinJumpList) DestroyQWinJumpListDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinJumpList_DestroyQWinJumpListDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinJumpList) __addCategory_items_atList2(i int) *QWinJumpListItem {
	if ptr.Pointer() != nil {
		return NewQWinJumpListItemFromPointer(C.QWinJumpList___addCategory_items_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QWinJumpList) __addCategory_items_setList2(i QWinJumpListItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList___addCategory_items_setList2(ptr.Pointer(), PointerFromQWinJumpListItem(i))
	}
}

func (ptr *QWinJumpList) __addCategory_items_newList2() unsafe.Pointer {
	return C.QWinJumpList___addCategory_items_newList2(ptr.Pointer())
}

func (ptr *QWinJumpList) __categories_atList(i int) *QWinJumpListCategory {
	if ptr.Pointer() != nil {
		return NewQWinJumpListCategoryFromPointer(C.QWinJumpList___categories_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QWinJumpList) __categories_setList(i QWinJumpListCategory_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList___categories_setList(ptr.Pointer(), PointerFromQWinJumpListCategory(i))
	}
}

func (ptr *QWinJumpList) __categories_newList() unsafe.Pointer {
	return C.QWinJumpList___categories_newList(ptr.Pointer())
}

func (ptr *QWinJumpList) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinJumpList___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinJumpList) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinJumpList) __children_newList() unsafe.Pointer {
	return C.QWinJumpList___children_newList(ptr.Pointer())
}

func (ptr *QWinJumpList) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWinJumpList___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWinJumpList) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWinJumpList) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWinJumpList___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWinJumpList) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinJumpList___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinJumpList) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinJumpList) __findChildren_newList() unsafe.Pointer {
	return C.QWinJumpList___findChildren_newList(ptr.Pointer())
}

func (ptr *QWinJumpList) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinJumpList___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinJumpList) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinJumpList) __findChildren_newList3() unsafe.Pointer {
	return C.QWinJumpList___findChildren_newList3(ptr.Pointer())
}

//export callbackQWinJumpList_ChildEvent
func callbackQWinJumpList_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWinJumpListFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWinJumpList) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWinJumpList_ConnectNotify
func callbackQWinJumpList_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinJumpListFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinJumpList) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinJumpList_CustomEvent
func callbackQWinJumpList_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWinJumpListFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWinJumpList) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWinJumpList_DeleteLater
func callbackQWinJumpList_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinJumpListFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWinJumpList) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinJumpList_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWinJumpList_Destroyed
func callbackQWinJumpList_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWinJumpList_DisconnectNotify
func callbackQWinJumpList_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinJumpListFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinJumpList) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinJumpList_Event
func callbackQWinJumpList_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinJumpListFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWinJumpList) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinJumpList_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWinJumpList_EventFilter
func callbackQWinJumpList_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinJumpListFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWinJumpList) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinJumpList_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWinJumpList_MetaObject
func callbackQWinJumpList_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWinJumpListFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWinJumpList) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWinJumpList_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWinJumpList_ObjectNameChanged
func callbackQWinJumpList_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWinExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWinJumpList_TimerEvent
func callbackQWinJumpList_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWinJumpListFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWinJumpList) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpList_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWinJumpListCategory struct {
	ptr unsafe.Pointer
}

type QWinJumpListCategory_ITF interface {
	QWinJumpListCategory_PTR() *QWinJumpListCategory
}

func (ptr *QWinJumpListCategory) QWinJumpListCategory_PTR() *QWinJumpListCategory {
	return ptr
}

func (ptr *QWinJumpListCategory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWinJumpListCategory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWinJumpListCategory(ptr QWinJumpListCategory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinJumpListCategory_PTR().Pointer()
	}
	return nil
}

func NewQWinJumpListCategoryFromPointer(ptr unsafe.Pointer) (n *QWinJumpListCategory) {
	n = new(QWinJumpListCategory)
	n.SetPointer(ptr)
	return
}

// QWinJumpListCategory::Type
//
//go:generate stringer -type=QWinJumpListCategory__Type
type QWinJumpListCategory__Type int64

const (
	QWinJumpListCategory__Custom   QWinJumpListCategory__Type = QWinJumpListCategory__Type(0)
	QWinJumpListCategory__Recent   QWinJumpListCategory__Type = QWinJumpListCategory__Type(1)
	QWinJumpListCategory__Frequent QWinJumpListCategory__Type = QWinJumpListCategory__Type(2)
	QWinJumpListCategory__Tasks    QWinJumpListCategory__Type = QWinJumpListCategory__Type(3)
)

func NewQWinJumpListCategory(title string) *QWinJumpListCategory {
	var titleC *C.char
	if title != "" {
		titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
	}
	tmpValue := NewQWinJumpListCategoryFromPointer(C.QWinJumpListCategory_NewQWinJumpListCategory(C.struct_QtWinExtras_PackedString{data: titleC, len: C.longlong(len(title))}))
	qt.SetFinalizer(tmpValue, (*QWinJumpListCategory).DestroyQWinJumpListCategory)
	return tmpValue
}

func (ptr *QWinJumpListCategory) AddDestination(filePath string) *QWinJumpListItem {
	if ptr.Pointer() != nil {
		var filePathC *C.char
		if filePath != "" {
			filePathC = C.CString(filePath)
			defer C.free(unsafe.Pointer(filePathC))
		}
		return NewQWinJumpListItemFromPointer(C.QWinJumpListCategory_AddDestination(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: filePathC, len: C.longlong(len(filePath))}))
	}
	return nil
}

func (ptr *QWinJumpListCategory) AddItem(item QWinJumpListItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpListCategory_AddItem(ptr.Pointer(), PointerFromQWinJumpListItem(item))
	}
}

func (ptr *QWinJumpListCategory) AddLink(title string, executablePath string, arguments []string) *QWinJumpListItem {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		var executablePathC *C.char
		if executablePath != "" {
			executablePathC = C.CString(executablePath)
			defer C.free(unsafe.Pointer(executablePathC))
		}
		argumentsC := C.CString(strings.Join(arguments, "¡¦!"))
		defer C.free(unsafe.Pointer(argumentsC))
		return NewQWinJumpListItemFromPointer(C.QWinJumpListCategory_AddLink(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWinExtras_PackedString{data: executablePathC, len: C.longlong(len(executablePath))}, C.struct_QtWinExtras_PackedString{data: argumentsC, len: C.longlong(len(strings.Join(arguments, "¡¦!")))}))
	}
	return nil
}

func (ptr *QWinJumpListCategory) AddLink2(icon gui.QIcon_ITF, title string, executablePath string, arguments []string) *QWinJumpListItem {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		var executablePathC *C.char
		if executablePath != "" {
			executablePathC = C.CString(executablePath)
			defer C.free(unsafe.Pointer(executablePathC))
		}
		argumentsC := C.CString(strings.Join(arguments, "¡¦!"))
		defer C.free(unsafe.Pointer(argumentsC))
		return NewQWinJumpListItemFromPointer(C.QWinJumpListCategory_AddLink2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.struct_QtWinExtras_PackedString{data: titleC, len: C.longlong(len(title))}, C.struct_QtWinExtras_PackedString{data: executablePathC, len: C.longlong(len(executablePath))}, C.struct_QtWinExtras_PackedString{data: argumentsC, len: C.longlong(len(strings.Join(arguments, "¡¦!")))}))
	}
	return nil
}

func (ptr *QWinJumpListCategory) AddSeparator() *QWinJumpListItem {
	if ptr.Pointer() != nil {
		return NewQWinJumpListItemFromPointer(C.QWinJumpListCategory_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWinJumpListCategory) Clear() {
	if ptr.Pointer() != nil {
		C.QWinJumpListCategory_Clear(ptr.Pointer())
	}
}

func (ptr *QWinJumpListCategory) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWinJumpListCategory_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWinJumpListCategory) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinJumpListCategory_IsEmpty(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinJumpListCategory) IsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinJumpListCategory_IsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinJumpListCategory) Items() []*QWinJumpListItem {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWinExtras_PackedList) []*QWinJumpListItem {
			out := make([]*QWinJumpListItem, int(l.len))
			tmpList := NewQWinJumpListCategoryFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__items_atList(i)
			}
			return out
		}(C.QWinJumpListCategory_Items(ptr.Pointer()))
	}
	return make([]*QWinJumpListItem, 0)
}

func (ptr *QWinJumpListCategory) SetTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QWinJumpListCategory_SetTitle(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QWinJumpListCategory) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWinJumpListCategory_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QWinJumpListCategory) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWinJumpListCategory_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWinJumpListCategory) Type() QWinJumpListCategory__Type {
	if ptr.Pointer() != nil {
		return QWinJumpListCategory__Type(C.QWinJumpListCategory_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWinJumpListCategory) DestroyQWinJumpListCategory() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinJumpListCategory_DestroyQWinJumpListCategory(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinJumpListCategory) __items_atList(i int) *QWinJumpListItem {
	if ptr.Pointer() != nil {
		return NewQWinJumpListItemFromPointer(C.QWinJumpListCategory___items_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return nil
}

func (ptr *QWinJumpListCategory) __items_setList(i QWinJumpListItem_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpListCategory___items_setList(ptr.Pointer(), PointerFromQWinJumpListItem(i))
	}
}

func (ptr *QWinJumpListCategory) __items_newList() unsafe.Pointer {
	return C.QWinJumpListCategory___items_newList(ptr.Pointer())
}

type QWinJumpListItem struct {
	ptr unsafe.Pointer
}

type QWinJumpListItem_ITF interface {
	QWinJumpListItem_PTR() *QWinJumpListItem
}

func (ptr *QWinJumpListItem) QWinJumpListItem_PTR() *QWinJumpListItem {
	return ptr
}

func (ptr *QWinJumpListItem) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWinJumpListItem) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWinJumpListItem(ptr QWinJumpListItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinJumpListItem_PTR().Pointer()
	}
	return nil
}

func NewQWinJumpListItemFromPointer(ptr unsafe.Pointer) (n *QWinJumpListItem) {
	n = new(QWinJumpListItem)
	n.SetPointer(ptr)
	return
}

// QWinJumpListItem::Type
//
//go:generate stringer -type=QWinJumpListItem__Type
type QWinJumpListItem__Type int64

const (
	QWinJumpListItem__Destination QWinJumpListItem__Type = QWinJumpListItem__Type(0)
	QWinJumpListItem__Link        QWinJumpListItem__Type = QWinJumpListItem__Type(1)
	QWinJumpListItem__Separator   QWinJumpListItem__Type = QWinJumpListItem__Type(2)
)

func NewQWinJumpListItem(ty QWinJumpListItem__Type) *QWinJumpListItem {
	tmpValue := NewQWinJumpListItemFromPointer(C.QWinJumpListItem_NewQWinJumpListItem(C.longlong(ty)))
	qt.SetFinalizer(tmpValue, (*QWinJumpListItem).DestroyQWinJumpListItem)
	return tmpValue
}

func (ptr *QWinJumpListItem) Arguments() []string {
	if ptr.Pointer() != nil {
		return unpackStringList(cGoUnpackString(C.QWinJumpListItem_Arguments(ptr.Pointer())))
	}
	return make([]string, 0)
}

func (ptr *QWinJumpListItem) Description() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWinJumpListItem_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWinJumpListItem) FilePath() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWinJumpListItem_FilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWinJumpListItem) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QWinJumpListItem_Icon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWinJumpListItem) SetArguments(arguments []string) {
	if ptr.Pointer() != nil {
		argumentsC := C.CString(strings.Join(arguments, "¡¦!"))
		defer C.free(unsafe.Pointer(argumentsC))
		C.QWinJumpListItem_SetArguments(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: argumentsC, len: C.longlong(len(strings.Join(arguments, "¡¦!")))})
	}
}

func (ptr *QWinJumpListItem) SetDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.QWinJumpListItem_SetDescription(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

func (ptr *QWinJumpListItem) SetFilePath(filePath string) {
	if ptr.Pointer() != nil {
		var filePathC *C.char
		if filePath != "" {
			filePathC = C.CString(filePath)
			defer C.free(unsafe.Pointer(filePathC))
		}
		C.QWinJumpListItem_SetFilePath(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: filePathC, len: C.longlong(len(filePath))})
	}
}

func (ptr *QWinJumpListItem) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWinJumpListItem_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QWinJumpListItem) SetTitle(title string) {
	if ptr.Pointer() != nil {
		var titleC *C.char
		if title != "" {
			titleC = C.CString(title)
			defer C.free(unsafe.Pointer(titleC))
		}
		C.QWinJumpListItem_SetTitle(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: titleC, len: C.longlong(len(title))})
	}
}

func (ptr *QWinJumpListItem) SetType(ty QWinJumpListItem__Type) {
	if ptr.Pointer() != nil {
		C.QWinJumpListItem_SetType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QWinJumpListItem) SetWorkingDirectory(workingDirectory string) {
	if ptr.Pointer() != nil {
		var workingDirectoryC *C.char
		if workingDirectory != "" {
			workingDirectoryC = C.CString(workingDirectory)
			defer C.free(unsafe.Pointer(workingDirectoryC))
		}
		C.QWinJumpListItem_SetWorkingDirectory(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: workingDirectoryC, len: C.longlong(len(workingDirectory))})
	}
}

func (ptr *QWinJumpListItem) Title() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWinJumpListItem_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWinJumpListItem) Type() QWinJumpListItem__Type {
	if ptr.Pointer() != nil {
		return QWinJumpListItem__Type(C.QWinJumpListItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWinJumpListItem) WorkingDirectory() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWinJumpListItem_WorkingDirectory(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWinJumpListItem) DestroyQWinJumpListItem() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinJumpListItem_DestroyQWinJumpListItem(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QWinMime struct {
	ptr unsafe.Pointer
}

type QWinMime_ITF interface {
	QWinMime_PTR() *QWinMime
}

func (ptr *QWinMime) QWinMime_PTR() *QWinMime {
	return ptr
}

func (ptr *QWinMime) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QWinMime) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQWinMime(ptr QWinMime_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinMime_PTR().Pointer()
	}
	return nil
}

func NewQWinMimeFromPointer(ptr unsafe.Pointer) (n *QWinMime) {
	n = new(QWinMime)
	n.SetPointer(ptr)
	return
}

type QWinTaskbarButton struct {
	core.QObject
}

type QWinTaskbarButton_ITF interface {
	core.QObject_ITF
	QWinTaskbarButton_PTR() *QWinTaskbarButton
}

func (ptr *QWinTaskbarButton) QWinTaskbarButton_PTR() *QWinTaskbarButton {
	return ptr
}

func (ptr *QWinTaskbarButton) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWinTaskbarButton) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWinTaskbarButton(ptr QWinTaskbarButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinTaskbarButton_PTR().Pointer()
	}
	return nil
}

func NewQWinTaskbarButtonFromPointer(ptr unsafe.Pointer) (n *QWinTaskbarButton) {
	n = new(QWinTaskbarButton)
	n.SetPointer(ptr)
	return
}
func NewQWinTaskbarButton(parent core.QObject_ITF) *QWinTaskbarButton {
	tmpValue := NewQWinTaskbarButtonFromPointer(C.QWinTaskbarButton_NewQWinTaskbarButton(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWinTaskbarButton_ClearOverlayIcon
func callbackQWinTaskbarButton_ClearOverlayIcon(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clearOverlayIcon"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).ClearOverlayIconDefault()
	}
}

func (ptr *QWinTaskbarButton) ConnectClearOverlayIcon(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clearOverlayIcon"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clearOverlayIcon", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clearOverlayIcon", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarButton) DisconnectClearOverlayIcon() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clearOverlayIcon")
	}
}

func (ptr *QWinTaskbarButton) ClearOverlayIcon() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_ClearOverlayIcon(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarButton) ClearOverlayIconDefault() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_ClearOverlayIconDefault(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarButton) OverlayAccessibleDescription() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWinTaskbarButton_OverlayAccessibleDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWinTaskbarButton) OverlayIcon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QWinTaskbarButton_OverlayIcon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarButton) Progress() *QWinTaskbarProgress {
	if ptr.Pointer() != nil {
		tmpValue := NewQWinTaskbarProgressFromPointer(C.QWinTaskbarButton_Progress(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWinTaskbarButton_SetOverlayAccessibleDescription
func callbackQWinTaskbarButton_SetOverlayAccessibleDescription(ptr unsafe.Pointer, description C.struct_QtWinExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "setOverlayAccessibleDescription"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(description))
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).SetOverlayAccessibleDescriptionDefault(cGoUnpackString(description))
	}
}

func (ptr *QWinTaskbarButton) ConnectSetOverlayAccessibleDescription(f func(description string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOverlayAccessibleDescription"); signal != nil {
			f := func(description string) {
				(*(*func(string))(signal))(description)
				f(description)
			}
			qt.ConnectSignal(ptr.Pointer(), "setOverlayAccessibleDescription", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOverlayAccessibleDescription", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarButton) DisconnectSetOverlayAccessibleDescription() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOverlayAccessibleDescription")
	}
}

func (ptr *QWinTaskbarButton) SetOverlayAccessibleDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.QWinTaskbarButton_SetOverlayAccessibleDescription(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

func (ptr *QWinTaskbarButton) SetOverlayAccessibleDescriptionDefault(description string) {
	if ptr.Pointer() != nil {
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.QWinTaskbarButton_SetOverlayAccessibleDescriptionDefault(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

//export callbackQWinTaskbarButton_SetOverlayIcon
func callbackQWinTaskbarButton_SetOverlayIcon(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setOverlayIcon"); signal != nil {
		(*(*func(*gui.QIcon))(signal))(gui.NewQIconFromPointer(icon))
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).SetOverlayIconDefault(gui.NewQIconFromPointer(icon))
	}
}

func (ptr *QWinTaskbarButton) ConnectSetOverlayIcon(f func(icon *gui.QIcon)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setOverlayIcon"); signal != nil {
			f := func(icon *gui.QIcon) {
				(*(*func(*gui.QIcon))(signal))(icon)
				f(icon)
			}
			qt.ConnectSignal(ptr.Pointer(), "setOverlayIcon", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setOverlayIcon", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarButton) DisconnectSetOverlayIcon() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setOverlayIcon")
	}
}

func (ptr *QWinTaskbarButton) SetOverlayIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_SetOverlayIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QWinTaskbarButton) SetOverlayIconDefault(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_SetOverlayIconDefault(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QWinTaskbarButton) SetWindow(window gui.QWindow_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_SetWindow(ptr.Pointer(), gui.PointerFromQWindow(window))
	}
}

func (ptr *QWinTaskbarButton) Window() *gui.QWindow {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQWindowFromPointer(C.QWinTaskbarButton_Window(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWinTaskbarButton_DestroyQWinTaskbarButton
func callbackQWinTaskbarButton_DestroyQWinTaskbarButton(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWinTaskbarButton"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).DestroyQWinTaskbarButtonDefault()
	}
}

func (ptr *QWinTaskbarButton) ConnectDestroyQWinTaskbarButton(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWinTaskbarButton"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWinTaskbarButton", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWinTaskbarButton", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarButton) DisconnectDestroyQWinTaskbarButton() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWinTaskbarButton")
	}
}

func (ptr *QWinTaskbarButton) DestroyQWinTaskbarButton() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinTaskbarButton_DestroyQWinTaskbarButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinTaskbarButton) DestroyQWinTaskbarButtonDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinTaskbarButton_DestroyQWinTaskbarButtonDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinTaskbarButton) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinTaskbarButton___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarButton) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinTaskbarButton) __children_newList() unsafe.Pointer {
	return C.QWinTaskbarButton___children_newList(ptr.Pointer())
}

func (ptr *QWinTaskbarButton) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWinTaskbarButton___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarButton) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWinTaskbarButton) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWinTaskbarButton___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWinTaskbarButton) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinTaskbarButton___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarButton) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinTaskbarButton) __findChildren_newList() unsafe.Pointer {
	return C.QWinTaskbarButton___findChildren_newList(ptr.Pointer())
}

func (ptr *QWinTaskbarButton) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinTaskbarButton___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarButton) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinTaskbarButton) __findChildren_newList3() unsafe.Pointer {
	return C.QWinTaskbarButton___findChildren_newList3(ptr.Pointer())
}

//export callbackQWinTaskbarButton_ChildEvent
func callbackQWinTaskbarButton_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWinTaskbarButton) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWinTaskbarButton_ConnectNotify
func callbackQWinTaskbarButton_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinTaskbarButton) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinTaskbarButton_CustomEvent
func callbackQWinTaskbarButton_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWinTaskbarButton) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWinTaskbarButton_DeleteLater
func callbackQWinTaskbarButton_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWinTaskbarButton) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinTaskbarButton_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWinTaskbarButton_Destroyed
func callbackQWinTaskbarButton_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWinTaskbarButton_DisconnectNotify
func callbackQWinTaskbarButton_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinTaskbarButton) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinTaskbarButton_Event
func callbackQWinTaskbarButton_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinTaskbarButtonFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWinTaskbarButton) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinTaskbarButton_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWinTaskbarButton_EventFilter
func callbackQWinTaskbarButton_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinTaskbarButtonFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWinTaskbarButton) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinTaskbarButton_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWinTaskbarButton_MetaObject
func callbackQWinTaskbarButton_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWinTaskbarButtonFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWinTaskbarButton) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWinTaskbarButton_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWinTaskbarButton_ObjectNameChanged
func callbackQWinTaskbarButton_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWinExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWinTaskbarButton_TimerEvent
func callbackQWinTaskbarButton_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWinTaskbarButtonFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWinTaskbarButton) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarButton_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWinTaskbarProgress struct {
	core.QObject
}

type QWinTaskbarProgress_ITF interface {
	core.QObject_ITF
	QWinTaskbarProgress_PTR() *QWinTaskbarProgress
}

func (ptr *QWinTaskbarProgress) QWinTaskbarProgress_PTR() *QWinTaskbarProgress {
	return ptr
}

func (ptr *QWinTaskbarProgress) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWinTaskbarProgress) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWinTaskbarProgress(ptr QWinTaskbarProgress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinTaskbarProgress_PTR().Pointer()
	}
	return nil
}

func NewQWinTaskbarProgressFromPointer(ptr unsafe.Pointer) (n *QWinTaskbarProgress) {
	n = new(QWinTaskbarProgress)
	n.SetPointer(ptr)
	return
}
func NewQWinTaskbarProgress(parent core.QObject_ITF) *QWinTaskbarProgress {
	tmpValue := NewQWinTaskbarProgressFromPointer(C.QWinTaskbarProgress_NewQWinTaskbarProgress(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWinTaskbarProgress_Hide
func callbackQWinTaskbarProgress_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).HideDefault()
	}
}

func (ptr *QWinTaskbarProgress) ConnectHide(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "hide"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "hide", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hide", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectHide() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "hide")
	}
}

func (ptr *QWinTaskbarProgress) Hide() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_Hide(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarProgress) HideDefault() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_HideDefault(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarProgress) IsPaused() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinTaskbarProgress_IsPaused(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinTaskbarProgress) IsStopped() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinTaskbarProgress_IsStopped(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinTaskbarProgress) IsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinTaskbarProgress_IsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinTaskbarProgress) Maximum() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWinTaskbarProgress_Maximum(ptr.Pointer())))
	}
	return 0
}

//export callbackQWinTaskbarProgress_MaximumChanged
func callbackQWinTaskbarProgress_MaximumChanged(ptr unsafe.Pointer, maximum C.int) {
	if signal := qt.GetSignal(ptr, "maximumChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(maximum)))
	}

}

func (ptr *QWinTaskbarProgress) ConnectMaximumChanged(f func(maximum int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "maximumChanged") {
			C.QWinTaskbarProgress_ConnectMaximumChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "maximumChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "maximumChanged"); signal != nil {
			f := func(maximum int) {
				(*(*func(int))(signal))(maximum)
				f(maximum)
			}
			qt.ConnectSignal(ptr.Pointer(), "maximumChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "maximumChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectMaximumChanged() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_DisconnectMaximumChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "maximumChanged")
	}
}

func (ptr *QWinTaskbarProgress) MaximumChanged(maximum int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_MaximumChanged(ptr.Pointer(), C.int(int32(maximum)))
	}
}

func (ptr *QWinTaskbarProgress) Minimum() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWinTaskbarProgress_Minimum(ptr.Pointer())))
	}
	return 0
}

//export callbackQWinTaskbarProgress_MinimumChanged
func callbackQWinTaskbarProgress_MinimumChanged(ptr unsafe.Pointer, minimum C.int) {
	if signal := qt.GetSignal(ptr, "minimumChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(minimum)))
	}

}

func (ptr *QWinTaskbarProgress) ConnectMinimumChanged(f func(minimum int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "minimumChanged") {
			C.QWinTaskbarProgress_ConnectMinimumChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "minimumChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "minimumChanged"); signal != nil {
			f := func(minimum int) {
				(*(*func(int))(signal))(minimum)
				f(minimum)
			}
			qt.ConnectSignal(ptr.Pointer(), "minimumChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "minimumChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectMinimumChanged() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_DisconnectMinimumChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "minimumChanged")
	}
}

func (ptr *QWinTaskbarProgress) MinimumChanged(minimum int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_MinimumChanged(ptr.Pointer(), C.int(int32(minimum)))
	}
}

//export callbackQWinTaskbarProgress_Pause
func callbackQWinTaskbarProgress_Pause(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pause"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).PauseDefault()
	}
}

func (ptr *QWinTaskbarProgress) ConnectPause(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pause"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "pause", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pause", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectPause() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pause")
	}
}

func (ptr *QWinTaskbarProgress) Pause() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_Pause(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarProgress) PauseDefault() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_PauseDefault(ptr.Pointer())
	}
}

//export callbackQWinTaskbarProgress_PausedChanged
func callbackQWinTaskbarProgress_PausedChanged(ptr unsafe.Pointer, paused C.char) {
	if signal := qt.GetSignal(ptr, "pausedChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(paused) != 0)
	}

}

func (ptr *QWinTaskbarProgress) ConnectPausedChanged(f func(paused bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pausedChanged") {
			C.QWinTaskbarProgress_ConnectPausedChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pausedChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pausedChanged"); signal != nil {
			f := func(paused bool) {
				(*(*func(bool))(signal))(paused)
				f(paused)
			}
			qt.ConnectSignal(ptr.Pointer(), "pausedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pausedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectPausedChanged() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_DisconnectPausedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pausedChanged")
	}
}

func (ptr *QWinTaskbarProgress) PausedChanged(paused bool) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_PausedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(paused))))
	}
}

//export callbackQWinTaskbarProgress_Reset
func callbackQWinTaskbarProgress_Reset(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QWinTaskbarProgress) ConnectReset(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "reset"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "reset", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "reset", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectReset() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "reset")
	}
}

func (ptr *QWinTaskbarProgress) Reset() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_Reset(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarProgress) ResetDefault() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_ResetDefault(ptr.Pointer())
	}
}

//export callbackQWinTaskbarProgress_Resume
func callbackQWinTaskbarProgress_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resume"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QWinTaskbarProgress) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resume"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "resume", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resume", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resume")
	}
}

func (ptr *QWinTaskbarProgress) Resume() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_Resume(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarProgress) ResumeDefault() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_ResumeDefault(ptr.Pointer())
	}
}

//export callbackQWinTaskbarProgress_SetMaximum
func callbackQWinTaskbarProgress_SetMaximum(ptr unsafe.Pointer, maximum C.int) {
	if signal := qt.GetSignal(ptr, "setMaximum"); signal != nil {
		(*(*func(int))(signal))(int(int32(maximum)))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).SetMaximumDefault(int(int32(maximum)))
	}
}

func (ptr *QWinTaskbarProgress) ConnectSetMaximum(f func(maximum int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMaximum"); signal != nil {
			f := func(maximum int) {
				(*(*func(int))(signal))(maximum)
				f(maximum)
			}
			qt.ConnectSignal(ptr.Pointer(), "setMaximum", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMaximum", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectSetMaximum() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMaximum")
	}
}

func (ptr *QWinTaskbarProgress) SetMaximum(maximum int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetMaximum(ptr.Pointer(), C.int(int32(maximum)))
	}
}

func (ptr *QWinTaskbarProgress) SetMaximumDefault(maximum int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetMaximumDefault(ptr.Pointer(), C.int(int32(maximum)))
	}
}

//export callbackQWinTaskbarProgress_SetMinimum
func callbackQWinTaskbarProgress_SetMinimum(ptr unsafe.Pointer, minimum C.int) {
	if signal := qt.GetSignal(ptr, "setMinimum"); signal != nil {
		(*(*func(int))(signal))(int(int32(minimum)))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).SetMinimumDefault(int(int32(minimum)))
	}
}

func (ptr *QWinTaskbarProgress) ConnectSetMinimum(f func(minimum int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setMinimum"); signal != nil {
			f := func(minimum int) {
				(*(*func(int))(signal))(minimum)
				f(minimum)
			}
			qt.ConnectSignal(ptr.Pointer(), "setMinimum", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setMinimum", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectSetMinimum() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setMinimum")
	}
}

func (ptr *QWinTaskbarProgress) SetMinimum(minimum int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetMinimum(ptr.Pointer(), C.int(int32(minimum)))
	}
}

func (ptr *QWinTaskbarProgress) SetMinimumDefault(minimum int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetMinimumDefault(ptr.Pointer(), C.int(int32(minimum)))
	}
}

//export callbackQWinTaskbarProgress_SetPaused
func callbackQWinTaskbarProgress_SetPaused(ptr unsafe.Pointer, paused C.char) {
	if signal := qt.GetSignal(ptr, "setPaused"); signal != nil {
		(*(*func(bool))(signal))(int8(paused) != 0)
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).SetPausedDefault(int8(paused) != 0)
	}
}

func (ptr *QWinTaskbarProgress) ConnectSetPaused(f func(paused bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPaused"); signal != nil {
			f := func(paused bool) {
				(*(*func(bool))(signal))(paused)
				f(paused)
			}
			qt.ConnectSignal(ptr.Pointer(), "setPaused", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPaused", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectSetPaused() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPaused")
	}
}

func (ptr *QWinTaskbarProgress) SetPaused(paused bool) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetPaused(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(paused))))
	}
}

func (ptr *QWinTaskbarProgress) SetPausedDefault(paused bool) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetPausedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(paused))))
	}
}

//export callbackQWinTaskbarProgress_SetRange
func callbackQWinTaskbarProgress_SetRange(ptr unsafe.Pointer, minimum C.int, maximum C.int) {
	if signal := qt.GetSignal(ptr, "setRange"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(minimum)), int(int32(maximum)))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).SetRangeDefault(int(int32(minimum)), int(int32(maximum)))
	}
}

func (ptr *QWinTaskbarProgress) ConnectSetRange(f func(minimum int, maximum int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRange"); signal != nil {
			f := func(minimum int, maximum int) {
				(*(*func(int, int))(signal))(minimum, maximum)
				f(minimum, maximum)
			}
			qt.ConnectSignal(ptr.Pointer(), "setRange", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRange", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectSetRange() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRange")
	}
}

func (ptr *QWinTaskbarProgress) SetRange(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetRange(ptr.Pointer(), C.int(int32(minimum)), C.int(int32(maximum)))
	}
}

func (ptr *QWinTaskbarProgress) SetRangeDefault(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetRangeDefault(ptr.Pointer(), C.int(int32(minimum)), C.int(int32(maximum)))
	}
}

//export callbackQWinTaskbarProgress_SetValue
func callbackQWinTaskbarProgress_SetValue(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(ptr, "setValue"); signal != nil {
		(*(*func(int))(signal))(int(int32(value)))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).SetValueDefault(int(int32(value)))
	}
}

func (ptr *QWinTaskbarProgress) ConnectSetValue(f func(value int)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setValue"); signal != nil {
			f := func(value int) {
				(*(*func(int))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "setValue", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setValue", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectSetValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setValue")
	}
}

func (ptr *QWinTaskbarProgress) SetValue(value int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetValue(ptr.Pointer(), C.int(int32(value)))
	}
}

func (ptr *QWinTaskbarProgress) SetValueDefault(value int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetValueDefault(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQWinTaskbarProgress_SetVisible
func callbackQWinTaskbarProgress_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QWinTaskbarProgress) ConnectSetVisible(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVisible"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "setVisible", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVisible", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectSetVisible() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVisible")
	}
}

func (ptr *QWinTaskbarProgress) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QWinTaskbarProgress) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWinTaskbarProgress_Show
func callbackQWinTaskbarProgress_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QWinTaskbarProgress) ConnectShow(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "show"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "show", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "show", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectShow() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "show")
	}
}

func (ptr *QWinTaskbarProgress) Show() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_Show(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarProgress) ShowDefault() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_ShowDefault(ptr.Pointer())
	}
}

//export callbackQWinTaskbarProgress_Stop
func callbackQWinTaskbarProgress_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).StopDefault()
	}
}

func (ptr *QWinTaskbarProgress) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QWinTaskbarProgress) Stop() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_Stop(ptr.Pointer())
	}
}

func (ptr *QWinTaskbarProgress) StopDefault() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_StopDefault(ptr.Pointer())
	}
}

//export callbackQWinTaskbarProgress_StoppedChanged
func callbackQWinTaskbarProgress_StoppedChanged(ptr unsafe.Pointer, stopped C.char) {
	if signal := qt.GetSignal(ptr, "stoppedChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(stopped) != 0)
	}

}

func (ptr *QWinTaskbarProgress) ConnectStoppedChanged(f func(stopped bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stoppedChanged") {
			C.QWinTaskbarProgress_ConnectStoppedChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stoppedChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stoppedChanged"); signal != nil {
			f := func(stopped bool) {
				(*(*func(bool))(signal))(stopped)
				f(stopped)
			}
			qt.ConnectSignal(ptr.Pointer(), "stoppedChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stoppedChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectStoppedChanged() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_DisconnectStoppedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stoppedChanged")
	}
}

func (ptr *QWinTaskbarProgress) StoppedChanged(stopped bool) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_StoppedChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(stopped))))
	}
}

func (ptr *QWinTaskbarProgress) Value() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWinTaskbarProgress_Value(ptr.Pointer())))
	}
	return 0
}

//export callbackQWinTaskbarProgress_ValueChanged
func callbackQWinTaskbarProgress_ValueChanged(ptr unsafe.Pointer, value C.int) {
	if signal := qt.GetSignal(ptr, "valueChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(value)))
	}

}

func (ptr *QWinTaskbarProgress) ConnectValueChanged(f func(value int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "valueChanged") {
			C.QWinTaskbarProgress_ConnectValueChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "valueChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "valueChanged"); signal != nil {
			f := func(value int) {
				(*(*func(int))(signal))(value)
				f(value)
			}
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "valueChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectValueChanged() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "valueChanged")
	}
}

func (ptr *QWinTaskbarProgress) ValueChanged(value int) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_ValueChanged(ptr.Pointer(), C.int(int32(value)))
	}
}

//export callbackQWinTaskbarProgress_VisibilityChanged
func callbackQWinTaskbarProgress_VisibilityChanged(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "visibilityChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	}

}

func (ptr *QWinTaskbarProgress) ConnectVisibilityChanged(f func(visible bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "visibilityChanged") {
			C.QWinTaskbarProgress_ConnectVisibilityChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "visibilityChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "visibilityChanged"); signal != nil {
			f := func(visible bool) {
				(*(*func(bool))(signal))(visible)
				f(visible)
			}
			qt.ConnectSignal(ptr.Pointer(), "visibilityChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "visibilityChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectVisibilityChanged() {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_DisconnectVisibilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "visibilityChanged")
	}
}

func (ptr *QWinTaskbarProgress) VisibilityChanged(visible bool) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_VisibilityChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQWinTaskbarProgress_DestroyQWinTaskbarProgress
func callbackQWinTaskbarProgress_DestroyQWinTaskbarProgress(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWinTaskbarProgress"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).DestroyQWinTaskbarProgressDefault()
	}
}

func (ptr *QWinTaskbarProgress) ConnectDestroyQWinTaskbarProgress(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWinTaskbarProgress"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWinTaskbarProgress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWinTaskbarProgress", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinTaskbarProgress) DisconnectDestroyQWinTaskbarProgress() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWinTaskbarProgress")
	}
}

func (ptr *QWinTaskbarProgress) DestroyQWinTaskbarProgress() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinTaskbarProgress_DestroyQWinTaskbarProgress(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinTaskbarProgress) DestroyQWinTaskbarProgressDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinTaskbarProgress_DestroyQWinTaskbarProgressDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinTaskbarProgress) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinTaskbarProgress___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarProgress) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinTaskbarProgress) __children_newList() unsafe.Pointer {
	return C.QWinTaskbarProgress___children_newList(ptr.Pointer())
}

func (ptr *QWinTaskbarProgress) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWinTaskbarProgress___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarProgress) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWinTaskbarProgress) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWinTaskbarProgress___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWinTaskbarProgress) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinTaskbarProgress___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarProgress) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinTaskbarProgress) __findChildren_newList() unsafe.Pointer {
	return C.QWinTaskbarProgress___findChildren_newList(ptr.Pointer())
}

func (ptr *QWinTaskbarProgress) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinTaskbarProgress___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinTaskbarProgress) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinTaskbarProgress) __findChildren_newList3() unsafe.Pointer {
	return C.QWinTaskbarProgress___findChildren_newList3(ptr.Pointer())
}

//export callbackQWinTaskbarProgress_ChildEvent
func callbackQWinTaskbarProgress_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWinTaskbarProgress) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWinTaskbarProgress_ConnectNotify
func callbackQWinTaskbarProgress_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinTaskbarProgress) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinTaskbarProgress_CustomEvent
func callbackQWinTaskbarProgress_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWinTaskbarProgress) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWinTaskbarProgress_DeleteLater
func callbackQWinTaskbarProgress_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWinTaskbarProgress) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinTaskbarProgress_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWinTaskbarProgress_Destroyed
func callbackQWinTaskbarProgress_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWinTaskbarProgress_DisconnectNotify
func callbackQWinTaskbarProgress_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinTaskbarProgress) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinTaskbarProgress_Event
func callbackQWinTaskbarProgress_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinTaskbarProgressFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWinTaskbarProgress) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinTaskbarProgress_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWinTaskbarProgress_EventFilter
func callbackQWinTaskbarProgress_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinTaskbarProgressFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWinTaskbarProgress) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinTaskbarProgress_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWinTaskbarProgress_MetaObject
func callbackQWinTaskbarProgress_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWinTaskbarProgressFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWinTaskbarProgress) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWinTaskbarProgress_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWinTaskbarProgress_ObjectNameChanged
func callbackQWinTaskbarProgress_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWinExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWinTaskbarProgress_TimerEvent
func callbackQWinTaskbarProgress_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWinTaskbarProgressFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWinTaskbarProgress) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinTaskbarProgress_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWinThumbnailToolBar struct {
	core.QObject
}

type QWinThumbnailToolBar_ITF interface {
	core.QObject_ITF
	QWinThumbnailToolBar_PTR() *QWinThumbnailToolBar
}

func (ptr *QWinThumbnailToolBar) QWinThumbnailToolBar_PTR() *QWinThumbnailToolBar {
	return ptr
}

func (ptr *QWinThumbnailToolBar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWinThumbnailToolBar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWinThumbnailToolBar(ptr QWinThumbnailToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinThumbnailToolBar_PTR().Pointer()
	}
	return nil
}

func NewQWinThumbnailToolBarFromPointer(ptr unsafe.Pointer) (n *QWinThumbnailToolBar) {
	n = new(QWinThumbnailToolBar)
	n.SetPointer(ptr)
	return
}
func NewQWinThumbnailToolBar(parent core.QObject_ITF) *QWinThumbnailToolBar {
	tmpValue := NewQWinThumbnailToolBarFromPointer(C.QWinThumbnailToolBar_NewQWinThumbnailToolBar(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWinThumbnailToolBar) AddButton(button QWinThumbnailToolButton_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_AddButton(ptr.Pointer(), PointerFromQWinThumbnailToolButton(button))
	}
}

func (ptr *QWinThumbnailToolBar) Buttons() []*QWinThumbnailToolButton {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtWinExtras_PackedList) []*QWinThumbnailToolButton {
			out := make([]*QWinThumbnailToolButton, int(l.len))
			tmpList := NewQWinThumbnailToolBarFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__buttons_atList(i)
			}
			return out
		}(C.QWinThumbnailToolBar_Buttons(ptr.Pointer()))
	}
	return make([]*QWinThumbnailToolButton, 0)
}

//export callbackQWinThumbnailToolBar_Clear
func callbackQWinThumbnailToolBar_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QWinThumbnailToolBar) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolBar) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QWinThumbnailToolBar) Clear() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_Clear(ptr.Pointer())
	}
}

func (ptr *QWinThumbnailToolBar) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QWinThumbnailToolBar) Count() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWinThumbnailToolBar_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWinThumbnailToolBar) IconicLivePreviewPixmap() *gui.QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPixmapFromPointer(C.QWinThumbnailToolBar_IconicLivePreviewPixmap(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQWinThumbnailToolBar_IconicLivePreviewPixmapRequested
func callbackQWinThumbnailToolBar_IconicLivePreviewPixmapRequested(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconicLivePreviewPixmapRequested"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWinThumbnailToolBar) ConnectIconicLivePreviewPixmapRequested(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconicLivePreviewPixmapRequested") {
			C.QWinThumbnailToolBar_ConnectIconicLivePreviewPixmapRequested(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "iconicLivePreviewPixmapRequested")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconicLivePreviewPixmapRequested"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "iconicLivePreviewPixmapRequested", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconicLivePreviewPixmapRequested", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolBar) DisconnectIconicLivePreviewPixmapRequested() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_DisconnectIconicLivePreviewPixmapRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconicLivePreviewPixmapRequested")
	}
}

func (ptr *QWinThumbnailToolBar) IconicLivePreviewPixmapRequested() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_IconicLivePreviewPixmapRequested(ptr.Pointer())
	}
}

func (ptr *QWinThumbnailToolBar) IconicPixmapNotificationsEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolBar_IconicPixmapNotificationsEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinThumbnailToolBar) IconicThumbnailPixmap() *gui.QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQPixmapFromPointer(C.QWinThumbnailToolBar_IconicThumbnailPixmap(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

//export callbackQWinThumbnailToolBar_IconicThumbnailPixmapRequested
func callbackQWinThumbnailToolBar_IconicThumbnailPixmapRequested(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "iconicThumbnailPixmapRequested"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWinThumbnailToolBar) ConnectIconicThumbnailPixmapRequested(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "iconicThumbnailPixmapRequested") {
			C.QWinThumbnailToolBar_ConnectIconicThumbnailPixmapRequested(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "iconicThumbnailPixmapRequested")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "iconicThumbnailPixmapRequested"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "iconicThumbnailPixmapRequested", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "iconicThumbnailPixmapRequested", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolBar) DisconnectIconicThumbnailPixmapRequested() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_DisconnectIconicThumbnailPixmapRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "iconicThumbnailPixmapRequested")
	}
}

func (ptr *QWinThumbnailToolBar) IconicThumbnailPixmapRequested() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_IconicThumbnailPixmapRequested(ptr.Pointer())
	}
}

func (ptr *QWinThumbnailToolBar) RemoveButton(button QWinThumbnailToolButton_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_RemoveButton(ptr.Pointer(), PointerFromQWinThumbnailToolButton(button))
	}
}

func (ptr *QWinThumbnailToolBar) SetButtons(buttons []*QWinThumbnailToolButton) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_SetButtons(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQWinThumbnailToolBarFromPointer(NewQWinThumbnailToolBarFromPointer(nil).__setButtons_buttons_newList())
			for _, v := range buttons {
				tmpList.__setButtons_buttons_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQWinThumbnailToolBar_SetIconicLivePreviewPixmap
func callbackQWinThumbnailToolBar_SetIconicLivePreviewPixmap(ptr unsafe.Pointer, vqp unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setIconicLivePreviewPixmap"); signal != nil {
		(*(*func(*gui.QPixmap))(signal))(gui.NewQPixmapFromPointer(vqp))
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).SetIconicLivePreviewPixmapDefault(gui.NewQPixmapFromPointer(vqp))
	}
}

func (ptr *QWinThumbnailToolBar) ConnectSetIconicLivePreviewPixmap(f func(vqp *gui.QPixmap)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setIconicLivePreviewPixmap"); signal != nil {
			f := func(vqp *gui.QPixmap) {
				(*(*func(*gui.QPixmap))(signal))(vqp)
				f(vqp)
			}
			qt.ConnectSignal(ptr.Pointer(), "setIconicLivePreviewPixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setIconicLivePreviewPixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolBar) DisconnectSetIconicLivePreviewPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setIconicLivePreviewPixmap")
	}
}

func (ptr *QWinThumbnailToolBar) SetIconicLivePreviewPixmap(vqp gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_SetIconicLivePreviewPixmap(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

func (ptr *QWinThumbnailToolBar) SetIconicLivePreviewPixmapDefault(vqp gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_SetIconicLivePreviewPixmapDefault(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

func (ptr *QWinThumbnailToolBar) SetIconicPixmapNotificationsEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_SetIconicPixmapNotificationsEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

//export callbackQWinThumbnailToolBar_SetIconicThumbnailPixmap
func callbackQWinThumbnailToolBar_SetIconicThumbnailPixmap(ptr unsafe.Pointer, vqp unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setIconicThumbnailPixmap"); signal != nil {
		(*(*func(*gui.QPixmap))(signal))(gui.NewQPixmapFromPointer(vqp))
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).SetIconicThumbnailPixmapDefault(gui.NewQPixmapFromPointer(vqp))
	}
}

func (ptr *QWinThumbnailToolBar) ConnectSetIconicThumbnailPixmap(f func(vqp *gui.QPixmap)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setIconicThumbnailPixmap"); signal != nil {
			f := func(vqp *gui.QPixmap) {
				(*(*func(*gui.QPixmap))(signal))(vqp)
				f(vqp)
			}
			qt.ConnectSignal(ptr.Pointer(), "setIconicThumbnailPixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setIconicThumbnailPixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolBar) DisconnectSetIconicThumbnailPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setIconicThumbnailPixmap")
	}
}

func (ptr *QWinThumbnailToolBar) SetIconicThumbnailPixmap(vqp gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_SetIconicThumbnailPixmap(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

func (ptr *QWinThumbnailToolBar) SetIconicThumbnailPixmapDefault(vqp gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_SetIconicThumbnailPixmapDefault(ptr.Pointer(), gui.PointerFromQPixmap(vqp))
	}
}

func (ptr *QWinThumbnailToolBar) SetWindow(window gui.QWindow_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_SetWindow(ptr.Pointer(), gui.PointerFromQWindow(window))
	}
}

func (ptr *QWinThumbnailToolBar) Window() *gui.QWindow {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQWindowFromPointer(C.QWinThumbnailToolBar_Window(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWinThumbnailToolBar_DestroyQWinThumbnailToolBar
func callbackQWinThumbnailToolBar_DestroyQWinThumbnailToolBar(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWinThumbnailToolBar"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).DestroyQWinThumbnailToolBarDefault()
	}
}

func (ptr *QWinThumbnailToolBar) ConnectDestroyQWinThumbnailToolBar(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWinThumbnailToolBar"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWinThumbnailToolBar", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWinThumbnailToolBar", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolBar) DisconnectDestroyQWinThumbnailToolBar() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWinThumbnailToolBar")
	}
}

func (ptr *QWinThumbnailToolBar) DestroyQWinThumbnailToolBar() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinThumbnailToolBar_DestroyQWinThumbnailToolBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinThumbnailToolBar) DestroyQWinThumbnailToolBarDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinThumbnailToolBar_DestroyQWinThumbnailToolBarDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinThumbnailToolBar) __buttons_atList(i int) *QWinThumbnailToolButton {
	if ptr.Pointer() != nil {
		tmpValue := NewQWinThumbnailToolButtonFromPointer(C.QWinThumbnailToolBar___buttons_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolBar) __buttons_setList(i QWinThumbnailToolButton_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar___buttons_setList(ptr.Pointer(), PointerFromQWinThumbnailToolButton(i))
	}
}

func (ptr *QWinThumbnailToolBar) __buttons_newList() unsafe.Pointer {
	return C.QWinThumbnailToolBar___buttons_newList(ptr.Pointer())
}

func (ptr *QWinThumbnailToolBar) __setButtons_buttons_atList(i int) *QWinThumbnailToolButton {
	if ptr.Pointer() != nil {
		tmpValue := NewQWinThumbnailToolButtonFromPointer(C.QWinThumbnailToolBar___setButtons_buttons_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolBar) __setButtons_buttons_setList(i QWinThumbnailToolButton_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar___setButtons_buttons_setList(ptr.Pointer(), PointerFromQWinThumbnailToolButton(i))
	}
}

func (ptr *QWinThumbnailToolBar) __setButtons_buttons_newList() unsafe.Pointer {
	return C.QWinThumbnailToolBar___setButtons_buttons_newList(ptr.Pointer())
}

func (ptr *QWinThumbnailToolBar) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinThumbnailToolBar___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolBar) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinThumbnailToolBar) __children_newList() unsafe.Pointer {
	return C.QWinThumbnailToolBar___children_newList(ptr.Pointer())
}

func (ptr *QWinThumbnailToolBar) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWinThumbnailToolBar___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolBar) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWinThumbnailToolBar) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWinThumbnailToolBar___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWinThumbnailToolBar) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinThumbnailToolBar___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolBar) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinThumbnailToolBar) __findChildren_newList() unsafe.Pointer {
	return C.QWinThumbnailToolBar___findChildren_newList(ptr.Pointer())
}

func (ptr *QWinThumbnailToolBar) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinThumbnailToolBar___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolBar) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinThumbnailToolBar) __findChildren_newList3() unsafe.Pointer {
	return C.QWinThumbnailToolBar___findChildren_newList3(ptr.Pointer())
}

//export callbackQWinThumbnailToolBar_ChildEvent
func callbackQWinThumbnailToolBar_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWinThumbnailToolBar) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWinThumbnailToolBar_ConnectNotify
func callbackQWinThumbnailToolBar_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinThumbnailToolBar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinThumbnailToolBar_CustomEvent
func callbackQWinThumbnailToolBar_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWinThumbnailToolBar) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWinThumbnailToolBar_DeleteLater
func callbackQWinThumbnailToolBar_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWinThumbnailToolBar) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinThumbnailToolBar_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWinThumbnailToolBar_Destroyed
func callbackQWinThumbnailToolBar_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWinThumbnailToolBar_DisconnectNotify
func callbackQWinThumbnailToolBar_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinThumbnailToolBar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinThumbnailToolBar_Event
func callbackQWinThumbnailToolBar_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinThumbnailToolBarFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWinThumbnailToolBar) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolBar_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWinThumbnailToolBar_EventFilter
func callbackQWinThumbnailToolBar_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinThumbnailToolBarFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWinThumbnailToolBar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolBar_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWinThumbnailToolBar_MetaObject
func callbackQWinThumbnailToolBar_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWinThumbnailToolBarFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWinThumbnailToolBar) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWinThumbnailToolBar_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWinThumbnailToolBar_ObjectNameChanged
func callbackQWinThumbnailToolBar_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWinExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWinThumbnailToolBar_TimerEvent
func callbackQWinThumbnailToolBar_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWinThumbnailToolBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWinThumbnailToolBar) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QWinThumbnailToolButton struct {
	core.QObject
}

type QWinThumbnailToolButton_ITF interface {
	core.QObject_ITF
	QWinThumbnailToolButton_PTR() *QWinThumbnailToolButton
}

func (ptr *QWinThumbnailToolButton) QWinThumbnailToolButton_PTR() *QWinThumbnailToolButton {
	return ptr
}

func (ptr *QWinThumbnailToolButton) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWinThumbnailToolButton) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWinThumbnailToolButton(ptr QWinThumbnailToolButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinThumbnailToolButton_PTR().Pointer()
	}
	return nil
}

func NewQWinThumbnailToolButtonFromPointer(ptr unsafe.Pointer) (n *QWinThumbnailToolButton) {
	n = new(QWinThumbnailToolButton)
	n.SetPointer(ptr)
	return
}
func NewQWinThumbnailToolButton(parent core.QObject_ITF) *QWinThumbnailToolButton {
	tmpValue := NewQWinThumbnailToolButtonFromPointer(C.QWinThumbnailToolButton_NewQWinThumbnailToolButton(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWinThumbnailToolButton_Click
func callbackQWinThumbnailToolButton_Click(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "click"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinThumbnailToolButtonFromPointer(ptr).ClickDefault()
	}
}

func (ptr *QWinThumbnailToolButton) ConnectClick(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "click"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "click", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "click", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolButton) DisconnectClick() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "click")
	}
}

func (ptr *QWinThumbnailToolButton) Click() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_Click(ptr.Pointer())
	}
}

func (ptr *QWinThumbnailToolButton) ClickDefault() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_ClickDefault(ptr.Pointer())
	}
}

//export callbackQWinThumbnailToolButton_Clicked
func callbackQWinThumbnailToolButton_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QWinThumbnailToolButton) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.QWinThumbnailToolButton_ConnectClicked(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "clicked")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolButton) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *QWinThumbnailToolButton) Clicked() {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_Clicked(ptr.Pointer())
	}
}

func (ptr *QWinThumbnailToolButton) DismissOnClick() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolButton_DismissOnClick(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinThumbnailToolButton) Icon() *gui.QIcon {
	if ptr.Pointer() != nil {
		tmpValue := gui.NewQIconFromPointer(C.QWinThumbnailToolButton_Icon(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*gui.QIcon).DestroyQIcon)
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolButton) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolButton_IsEnabled(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinThumbnailToolButton) IsFlat() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolButton_IsFlat(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinThumbnailToolButton) IsInteractive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolButton_IsInteractive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinThumbnailToolButton) IsVisible() bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolButton_IsVisible(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWinThumbnailToolButton) SetDismissOnClick(dismiss bool) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_SetDismissOnClick(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(dismiss))))
	}
}

func (ptr *QWinThumbnailToolButton) SetEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enabled))))
	}
}

func (ptr *QWinThumbnailToolButton) SetFlat(flat bool) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_SetFlat(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(flat))))
	}
}

func (ptr *QWinThumbnailToolButton) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QWinThumbnailToolButton) SetInteractive(interactive bool) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_SetInteractive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(interactive))))
	}
}

func (ptr *QWinThumbnailToolButton) SetToolTip(toolTip string) {
	if ptr.Pointer() != nil {
		var toolTipC *C.char
		if toolTip != "" {
			toolTipC = C.CString(toolTip)
			defer C.free(unsafe.Pointer(toolTipC))
		}
		C.QWinThumbnailToolButton_SetToolTip(ptr.Pointer(), C.struct_QtWinExtras_PackedString{data: toolTipC, len: C.longlong(len(toolTip))})
	}
}

func (ptr *QWinThumbnailToolButton) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QWinThumbnailToolButton) ToolTip() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWinThumbnailToolButton_ToolTip(ptr.Pointer()))
	}
	return ""
}

//export callbackQWinThumbnailToolButton_DestroyQWinThumbnailToolButton
func callbackQWinThumbnailToolButton_DestroyQWinThumbnailToolButton(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QWinThumbnailToolButton"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinThumbnailToolButtonFromPointer(ptr).DestroyQWinThumbnailToolButtonDefault()
	}
}

func (ptr *QWinThumbnailToolButton) ConnectDestroyQWinThumbnailToolButton(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QWinThumbnailToolButton"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QWinThumbnailToolButton", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QWinThumbnailToolButton", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QWinThumbnailToolButton) DisconnectDestroyQWinThumbnailToolButton() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QWinThumbnailToolButton")
	}
}

func (ptr *QWinThumbnailToolButton) DestroyQWinThumbnailToolButton() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinThumbnailToolButton_DestroyQWinThumbnailToolButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinThumbnailToolButton) DestroyQWinThumbnailToolButtonDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinThumbnailToolButton_DestroyQWinThumbnailToolButtonDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWinThumbnailToolButton) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinThumbnailToolButton___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolButton) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinThumbnailToolButton) __children_newList() unsafe.Pointer {
	return C.QWinThumbnailToolButton___children_newList(ptr.Pointer())
}

func (ptr *QWinThumbnailToolButton) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QWinThumbnailToolButton___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolButton) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QWinThumbnailToolButton) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QWinThumbnailToolButton___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QWinThumbnailToolButton) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinThumbnailToolButton___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolButton) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinThumbnailToolButton) __findChildren_newList() unsafe.Pointer {
	return C.QWinThumbnailToolButton___findChildren_newList(ptr.Pointer())
}

func (ptr *QWinThumbnailToolButton) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QWinThumbnailToolButton___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWinThumbnailToolButton) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QWinThumbnailToolButton) __findChildren_newList3() unsafe.Pointer {
	return C.QWinThumbnailToolButton___findChildren_newList3(ptr.Pointer())
}

//export callbackQWinThumbnailToolButton_ChildEvent
func callbackQWinThumbnailToolButton_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWinThumbnailToolButtonFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWinThumbnailToolButton) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWinThumbnailToolButton_ConnectNotify
func callbackQWinThumbnailToolButton_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinThumbnailToolButtonFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinThumbnailToolButton) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinThumbnailToolButton_CustomEvent
func callbackQWinThumbnailToolButton_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
	} else {
		NewQWinThumbnailToolButtonFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWinThumbnailToolButton) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWinThumbnailToolButton_DeleteLater
func callbackQWinThumbnailToolButton_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQWinThumbnailToolButtonFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWinThumbnailToolButton) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QWinThumbnailToolButton_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQWinThumbnailToolButton_Destroyed
func callbackQWinThumbnailToolButton_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQWinThumbnailToolButton_DisconnectNotify
func callbackQWinThumbnailToolButton_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWinThumbnailToolButtonFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWinThumbnailToolButton) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWinThumbnailToolButton_Event
func callbackQWinThumbnailToolButton_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinThumbnailToolButtonFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWinThumbnailToolButton) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolButton_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQWinThumbnailToolButton_EventFilter
func callbackQWinThumbnailToolButton_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWinThumbnailToolButtonFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWinThumbnailToolButton) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QWinThumbnailToolButton_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQWinThumbnailToolButton_MetaObject
func callbackQWinThumbnailToolButton_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQWinThumbnailToolButtonFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWinThumbnailToolButton) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWinThumbnailToolButton_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWinThumbnailToolButton_ObjectNameChanged
func callbackQWinThumbnailToolButton_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtWinExtras_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQWinThumbnailToolButton_TimerEvent
func callbackQWinThumbnailToolButton_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWinThumbnailToolButtonFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWinThumbnailToolButton) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWinThumbnailToolButton_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QtWin struct {
	ptr unsafe.Pointer
}

type QtWin_ITF interface {
	QtWin_PTR() *QtWin
}

func (ptr *QtWin) QtWin_PTR() *QtWin {
	return ptr
}

func (ptr *QtWin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QtWin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQtWin(ptr QtWin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtWin_PTR().Pointer()
	}
	return nil
}

func NewQtWinFromPointer(ptr unsafe.Pointer) (n *QtWin) {
	n = new(QtWin)
	n.SetPointer(ptr)
	return
}
func (ptr *QtWin) DestroyQtWin() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QtWin__HBitmapFormat
//QtWin::HBitmapFormat
type QtWin__HBitmapFormat int64

const (
	QtWin__HBitmapNoAlpha            QtWin__HBitmapFormat = QtWin__HBitmapFormat(0)
	QtWin__HBitmapPremultipliedAlpha QtWin__HBitmapFormat = QtWin__HBitmapFormat(1)
	QtWin__HBitmapAlpha              QtWin__HBitmapFormat = QtWin__HBitmapFormat(2)
)

// QtWin::WindowFlip3DPolicy
//
//go:generate stringer -type=QtWin__WindowFlip3DPolicy
type QtWin__WindowFlip3DPolicy int64

const (
	QtWin__FlipDefault      QtWin__WindowFlip3DPolicy = QtWin__WindowFlip3DPolicy(0)
	QtWin__FlipExcludeBelow QtWin__WindowFlip3DPolicy = QtWin__WindowFlip3DPolicy(1)
	QtWin__FlipExcludeAbove QtWin__WindowFlip3DPolicy = QtWin__WindowFlip3DPolicy(2)
)

func QtWin_ColorizationColor(opaqueBlend *bool) *gui.QColor {
	var opaqueBlendC C.char
	if opaqueBlend != nil {
		opaqueBlendC = C.char(int8(qt.GoBoolToInt(*opaqueBlend)))
		defer func() { *opaqueBlend = int8(opaqueBlendC) != 0 }()
	}
	tmpValue := gui.NewQColorFromPointer(C.QtWin_QtWin_ColorizationColor(&opaqueBlendC))
	qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
	return tmpValue
}

func (ptr *QtWin) ColorizationColor(opaqueBlend *bool) *gui.QColor {
	var opaqueBlendC C.char
	if opaqueBlend != nil {
		opaqueBlendC = C.char(int8(qt.GoBoolToInt(*opaqueBlend)))
		defer func() { *opaqueBlend = int8(opaqueBlendC) != 0 }()
	}
	tmpValue := gui.NewQColorFromPointer(C.QtWin_QtWin_ColorizationColor(&opaqueBlendC))
	qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
	return tmpValue
}

func QtWin_DisableBlurBehindWindow(window gui.QWindow_ITF) {
	C.QtWin_QtWin_DisableBlurBehindWindow(gui.PointerFromQWindow(window))
}

func (ptr *QtWin) DisableBlurBehindWindow(window gui.QWindow_ITF) {
	C.QtWin_QtWin_DisableBlurBehindWindow(gui.PointerFromQWindow(window))
}

func QtWin_DisableBlurBehindWindow2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_DisableBlurBehindWindow2(widgets.PointerFromQWidget(window))
}

func (ptr *QtWin) DisableBlurBehindWindow2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_DisableBlurBehindWindow2(widgets.PointerFromQWidget(window))
}

func QtWin_EnableBlurBehindWindow(window gui.QWindow_ITF, region gui.QRegion_ITF) {
	C.QtWin_QtWin_EnableBlurBehindWindow(gui.PointerFromQWindow(window), gui.PointerFromQRegion(region))
}

func (ptr *QtWin) EnableBlurBehindWindow(window gui.QWindow_ITF, region gui.QRegion_ITF) {
	C.QtWin_QtWin_EnableBlurBehindWindow(gui.PointerFromQWindow(window), gui.PointerFromQRegion(region))
}

func QtWin_EnableBlurBehindWindow2(window gui.QWindow_ITF) {
	C.QtWin_QtWin_EnableBlurBehindWindow2(gui.PointerFromQWindow(window))
}

func (ptr *QtWin) EnableBlurBehindWindow2(window gui.QWindow_ITF) {
	C.QtWin_QtWin_EnableBlurBehindWindow2(gui.PointerFromQWindow(window))
}

func QtWin_EnableBlurBehindWindow3(window widgets.QWidget_ITF, region gui.QRegion_ITF) {
	C.QtWin_QtWin_EnableBlurBehindWindow3(widgets.PointerFromQWidget(window), gui.PointerFromQRegion(region))
}

func (ptr *QtWin) EnableBlurBehindWindow3(window widgets.QWidget_ITF, region gui.QRegion_ITF) {
	C.QtWin_QtWin_EnableBlurBehindWindow3(widgets.PointerFromQWidget(window), gui.PointerFromQRegion(region))
}

func QtWin_EnableBlurBehindWindow4(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_EnableBlurBehindWindow4(widgets.PointerFromQWidget(window))
}

func (ptr *QtWin) EnableBlurBehindWindow4(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_EnableBlurBehindWindow4(widgets.PointerFromQWidget(window))
}

func QtWin_ErrorStringFromHresult(hresult int) string {
	return cGoUnpackString(C.QtWin_QtWin_ErrorStringFromHresult(C.int(int32(hresult))))
}

func (ptr *QtWin) ErrorStringFromHresult(hresult int) string {
	return cGoUnpackString(C.QtWin_QtWin_ErrorStringFromHresult(C.int(int32(hresult))))
}

func QtWin_ExtendFrameIntoClientArea(window gui.QWindow_ITF, left int, top int, right int, bottom int) {
	C.QtWin_QtWin_ExtendFrameIntoClientArea(gui.PointerFromQWindow(window), C.int(int32(left)), C.int(int32(top)), C.int(int32(right)), C.int(int32(bottom)))
}

func (ptr *QtWin) ExtendFrameIntoClientArea(window gui.QWindow_ITF, left int, top int, right int, bottom int) {
	C.QtWin_QtWin_ExtendFrameIntoClientArea(gui.PointerFromQWindow(window), C.int(int32(left)), C.int(int32(top)), C.int(int32(right)), C.int(int32(bottom)))
}

func QtWin_ExtendFrameIntoClientArea2(window gui.QWindow_ITF, margins core.QMargins_ITF) {
	C.QtWin_QtWin_ExtendFrameIntoClientArea2(gui.PointerFromQWindow(window), core.PointerFromQMargins(margins))
}

func (ptr *QtWin) ExtendFrameIntoClientArea2(window gui.QWindow_ITF, margins core.QMargins_ITF) {
	C.QtWin_QtWin_ExtendFrameIntoClientArea2(gui.PointerFromQWindow(window), core.PointerFromQMargins(margins))
}

func QtWin_ExtendFrameIntoClientArea3(window widgets.QWidget_ITF, margins core.QMargins_ITF) {
	C.QtWin_QtWin_ExtendFrameIntoClientArea3(widgets.PointerFromQWidget(window), core.PointerFromQMargins(margins))
}

func (ptr *QtWin) ExtendFrameIntoClientArea3(window widgets.QWidget_ITF, margins core.QMargins_ITF) {
	C.QtWin_QtWin_ExtendFrameIntoClientArea3(widgets.PointerFromQWidget(window), core.PointerFromQMargins(margins))
}

func QtWin_ExtendFrameIntoClientArea4(window widgets.QWidget_ITF, left int, top int, right int, bottom int) {
	C.QtWin_QtWin_ExtendFrameIntoClientArea4(widgets.PointerFromQWidget(window), C.int(int32(left)), C.int(int32(top)), C.int(int32(right)), C.int(int32(bottom)))
}

func (ptr *QtWin) ExtendFrameIntoClientArea4(window widgets.QWidget_ITF, left int, top int, right int, bottom int) {
	C.QtWin_QtWin_ExtendFrameIntoClientArea4(widgets.PointerFromQWidget(window), C.int(int32(left)), C.int(int32(top)), C.int(int32(right)), C.int(int32(bottom)))
}

func QtWin_IsCompositionEnabled() bool {
	return int8(C.QtWin_QtWin_IsCompositionEnabled()) != 0
}

func (ptr *QtWin) IsCompositionEnabled() bool {
	return int8(C.QtWin_QtWin_IsCompositionEnabled()) != 0
}

func QtWin_IsCompositionOpaque() bool {
	return int8(C.QtWin_QtWin_IsCompositionOpaque()) != 0
}

func (ptr *QtWin) IsCompositionOpaque() bool {
	return int8(C.QtWin_QtWin_IsCompositionOpaque()) != 0
}

func QtWin_IsWindowExcludedFromPeek(window gui.QWindow_ITF) bool {
	return int8(C.QtWin_QtWin_IsWindowExcludedFromPeek(gui.PointerFromQWindow(window))) != 0
}

func (ptr *QtWin) IsWindowExcludedFromPeek(window gui.QWindow_ITF) bool {
	return int8(C.QtWin_QtWin_IsWindowExcludedFromPeek(gui.PointerFromQWindow(window))) != 0
}

func QtWin_IsWindowExcludedFromPeek2(window widgets.QWidget_ITF) bool {
	return int8(C.QtWin_QtWin_IsWindowExcludedFromPeek2(widgets.PointerFromQWidget(window))) != 0
}

func (ptr *QtWin) IsWindowExcludedFromPeek2(window widgets.QWidget_ITF) bool {
	return int8(C.QtWin_QtWin_IsWindowExcludedFromPeek2(widgets.PointerFromQWidget(window))) != 0
}

func QtWin_IsWindowPeekDisallowed(window gui.QWindow_ITF) bool {
	return int8(C.QtWin_QtWin_IsWindowPeekDisallowed(gui.PointerFromQWindow(window))) != 0
}

func (ptr *QtWin) IsWindowPeekDisallowed(window gui.QWindow_ITF) bool {
	return int8(C.QtWin_QtWin_IsWindowPeekDisallowed(gui.PointerFromQWindow(window))) != 0
}

func QtWin_IsWindowPeekDisallowed2(window widgets.QWidget_ITF) bool {
	return int8(C.QtWin_QtWin_IsWindowPeekDisallowed2(widgets.PointerFromQWidget(window))) != 0
}

func (ptr *QtWin) IsWindowPeekDisallowed2(window widgets.QWidget_ITF) bool {
	return int8(C.QtWin_QtWin_IsWindowPeekDisallowed2(widgets.PointerFromQWidget(window))) != 0
}

func QtWin_MarkFullscreenWindow(window gui.QWindow_ITF, fullscreen bool) {
	C.QtWin_QtWin_MarkFullscreenWindow(gui.PointerFromQWindow(window), C.char(int8(qt.GoBoolToInt(fullscreen))))
}

func (ptr *QtWin) MarkFullscreenWindow(window gui.QWindow_ITF, fullscreen bool) {
	C.QtWin_QtWin_MarkFullscreenWindow(gui.PointerFromQWindow(window), C.char(int8(qt.GoBoolToInt(fullscreen))))
}

func QtWin_MarkFullscreenWindow2(window widgets.QWidget_ITF, fullscreen bool) {
	C.QtWin_QtWin_MarkFullscreenWindow2(widgets.PointerFromQWidget(window), C.char(int8(qt.GoBoolToInt(fullscreen))))
}

func (ptr *QtWin) MarkFullscreenWindow2(window widgets.QWidget_ITF, fullscreen bool) {
	C.QtWin_QtWin_MarkFullscreenWindow2(widgets.PointerFromQWidget(window), C.char(int8(qt.GoBoolToInt(fullscreen))))
}

func QtWin_RealColorizationColor() *gui.QColor {
	tmpValue := gui.NewQColorFromPointer(C.QtWin_QtWin_RealColorizationColor())
	qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
	return tmpValue
}

func (ptr *QtWin) RealColorizationColor() *gui.QColor {
	tmpValue := gui.NewQColorFromPointer(C.QtWin_QtWin_RealColorizationColor())
	qt.SetFinalizer(tmpValue, (*gui.QColor).DestroyQColor)
	return tmpValue
}

func QtWin_ResetExtendedFrame(window gui.QWindow_ITF) {
	C.QtWin_QtWin_ResetExtendedFrame(gui.PointerFromQWindow(window))
}

func (ptr *QtWin) ResetExtendedFrame(window gui.QWindow_ITF) {
	C.QtWin_QtWin_ResetExtendedFrame(gui.PointerFromQWindow(window))
}

func QtWin_ResetExtendedFrame2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_ResetExtendedFrame2(widgets.PointerFromQWidget(window))
}

func (ptr *QtWin) ResetExtendedFrame2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_ResetExtendedFrame2(widgets.PointerFromQWidget(window))
}

func QtWin_SetCompositionEnabled(enabled bool) {
	C.QtWin_QtWin_SetCompositionEnabled(C.char(int8(qt.GoBoolToInt(enabled))))
}

func (ptr *QtWin) SetCompositionEnabled(enabled bool) {
	C.QtWin_QtWin_SetCompositionEnabled(C.char(int8(qt.GoBoolToInt(enabled))))
}

func QtWin_SetCurrentProcessExplicitAppUserModelID(id string) {
	var idC *C.char
	if id != "" {
		idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
	}
	C.QtWin_QtWin_SetCurrentProcessExplicitAppUserModelID(C.struct_QtWinExtras_PackedString{data: idC, len: C.longlong(len(id))})
}

func (ptr *QtWin) SetCurrentProcessExplicitAppUserModelID(id string) {
	var idC *C.char
	if id != "" {
		idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
	}
	C.QtWin_QtWin_SetCurrentProcessExplicitAppUserModelID(C.struct_QtWinExtras_PackedString{data: idC, len: C.longlong(len(id))})
}

func QtWin_SetWindowDisallowPeek(window gui.QWindow_ITF, disallow bool) {
	C.QtWin_QtWin_SetWindowDisallowPeek(gui.PointerFromQWindow(window), C.char(int8(qt.GoBoolToInt(disallow))))
}

func (ptr *QtWin) SetWindowDisallowPeek(window gui.QWindow_ITF, disallow bool) {
	C.QtWin_QtWin_SetWindowDisallowPeek(gui.PointerFromQWindow(window), C.char(int8(qt.GoBoolToInt(disallow))))
}

func QtWin_SetWindowDisallowPeek2(window widgets.QWidget_ITF, disallow bool) {
	C.QtWin_QtWin_SetWindowDisallowPeek2(widgets.PointerFromQWidget(window), C.char(int8(qt.GoBoolToInt(disallow))))
}

func (ptr *QtWin) SetWindowDisallowPeek2(window widgets.QWidget_ITF, disallow bool) {
	C.QtWin_QtWin_SetWindowDisallowPeek2(widgets.PointerFromQWidget(window), C.char(int8(qt.GoBoolToInt(disallow))))
}

func QtWin_SetWindowExcludedFromPeek(window gui.QWindow_ITF, exclude bool) {
	C.QtWin_QtWin_SetWindowExcludedFromPeek(gui.PointerFromQWindow(window), C.char(int8(qt.GoBoolToInt(exclude))))
}

func (ptr *QtWin) SetWindowExcludedFromPeek(window gui.QWindow_ITF, exclude bool) {
	C.QtWin_QtWin_SetWindowExcludedFromPeek(gui.PointerFromQWindow(window), C.char(int8(qt.GoBoolToInt(exclude))))
}

func QtWin_SetWindowExcludedFromPeek2(window widgets.QWidget_ITF, exclude bool) {
	C.QtWin_QtWin_SetWindowExcludedFromPeek2(widgets.PointerFromQWidget(window), C.char(int8(qt.GoBoolToInt(exclude))))
}

func (ptr *QtWin) SetWindowExcludedFromPeek2(window widgets.QWidget_ITF, exclude bool) {
	C.QtWin_QtWin_SetWindowExcludedFromPeek2(widgets.PointerFromQWidget(window), C.char(int8(qt.GoBoolToInt(exclude))))
}

func QtWin_SetWindowFlip3DPolicy(window gui.QWindow_ITF, policy QtWin__WindowFlip3DPolicy) {
	C.QtWin_QtWin_SetWindowFlip3DPolicy(gui.PointerFromQWindow(window), C.longlong(policy))
}

func (ptr *QtWin) SetWindowFlip3DPolicy(window gui.QWindow_ITF, policy QtWin__WindowFlip3DPolicy) {
	C.QtWin_QtWin_SetWindowFlip3DPolicy(gui.PointerFromQWindow(window), C.longlong(policy))
}

func QtWin_SetWindowFlip3DPolicy2(window widgets.QWidget_ITF, policy QtWin__WindowFlip3DPolicy) {
	C.QtWin_QtWin_SetWindowFlip3DPolicy2(widgets.PointerFromQWidget(window), C.longlong(policy))
}

func (ptr *QtWin) SetWindowFlip3DPolicy2(window widgets.QWidget_ITF, policy QtWin__WindowFlip3DPolicy) {
	C.QtWin_QtWin_SetWindowFlip3DPolicy2(widgets.PointerFromQWidget(window), C.longlong(policy))
}

func QtWin_StringFromHresult(hresult int) string {
	return cGoUnpackString(C.QtWin_QtWin_StringFromHresult(C.int(int32(hresult))))
}

func (ptr *QtWin) StringFromHresult(hresult int) string {
	return cGoUnpackString(C.QtWin_QtWin_StringFromHresult(C.int(int32(hresult))))
}

func QtWin_TaskbarActivateTab(window gui.QWindow_ITF) {
	C.QtWin_QtWin_TaskbarActivateTab(gui.PointerFromQWindow(window))
}

func (ptr *QtWin) TaskbarActivateTab(window gui.QWindow_ITF) {
	C.QtWin_QtWin_TaskbarActivateTab(gui.PointerFromQWindow(window))
}

func QtWin_TaskbarActivateTab2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_TaskbarActivateTab2(widgets.PointerFromQWidget(window))
}

func (ptr *QtWin) TaskbarActivateTab2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_TaskbarActivateTab2(widgets.PointerFromQWidget(window))
}

func QtWin_TaskbarActivateTabAlt(window gui.QWindow_ITF) {
	C.QtWin_QtWin_TaskbarActivateTabAlt(gui.PointerFromQWindow(window))
}

func (ptr *QtWin) TaskbarActivateTabAlt(window gui.QWindow_ITF) {
	C.QtWin_QtWin_TaskbarActivateTabAlt(gui.PointerFromQWindow(window))
}

func QtWin_TaskbarActivateTabAlt2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_TaskbarActivateTabAlt2(widgets.PointerFromQWidget(window))
}

func (ptr *QtWin) TaskbarActivateTabAlt2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_TaskbarActivateTabAlt2(widgets.PointerFromQWidget(window))
}

func QtWin_TaskbarAddTab(window gui.QWindow_ITF) {
	C.QtWin_QtWin_TaskbarAddTab(gui.PointerFromQWindow(window))
}

func (ptr *QtWin) TaskbarAddTab(window gui.QWindow_ITF) {
	C.QtWin_QtWin_TaskbarAddTab(gui.PointerFromQWindow(window))
}

func QtWin_TaskbarAddTab2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_TaskbarAddTab2(widgets.PointerFromQWidget(window))
}

func (ptr *QtWin) TaskbarAddTab2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_TaskbarAddTab2(widgets.PointerFromQWidget(window))
}

func QtWin_TaskbarDeleteTab(window gui.QWindow_ITF) {
	C.QtWin_QtWin_TaskbarDeleteTab(gui.PointerFromQWindow(window))
}

func (ptr *QtWin) TaskbarDeleteTab(window gui.QWindow_ITF) {
	C.QtWin_QtWin_TaskbarDeleteTab(gui.PointerFromQWindow(window))
}

func QtWin_TaskbarDeleteTab2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_TaskbarDeleteTab2(widgets.PointerFromQWidget(window))
}

func (ptr *QtWin) TaskbarDeleteTab2(window widgets.QWidget_ITF) {
	C.QtWin_QtWin_TaskbarDeleteTab2(widgets.PointerFromQWidget(window))
}

func QtWin_WindowFlip3DPolicy(window gui.QWindow_ITF) QtWin__WindowFlip3DPolicy {
	return QtWin__WindowFlip3DPolicy(C.QtWin_QtWin_WindowFlip3DPolicy(gui.PointerFromQWindow(window)))
}

func (ptr *QtWin) WindowFlip3DPolicy(window gui.QWindow_ITF) QtWin__WindowFlip3DPolicy {
	return QtWin__WindowFlip3DPolicy(C.QtWin_QtWin_WindowFlip3DPolicy(gui.PointerFromQWindow(window)))
}

func QtWin_WindowFlip3DPolicy2(window widgets.QWidget_ITF) QtWin__WindowFlip3DPolicy {
	return QtWin__WindowFlip3DPolicy(C.QtWin_QtWin_WindowFlip3DPolicy2(widgets.PointerFromQWidget(window)))
}

func (ptr *QtWin) WindowFlip3DPolicy2(window widgets.QWidget_ITF) QtWin__WindowFlip3DPolicy {
	return QtWin__WindowFlip3DPolicy(C.QtWin_QtWin_WindowFlip3DPolicy2(widgets.PointerFromQWidget(window)))
}

func init() {
	qt.ItfMap["winextras.QWinColorizationChangeEvent_ITF"] = QWinColorizationChangeEvent{}
	qt.ItfMap["winextras.QWinCompositionChangeEvent_ITF"] = QWinCompositionChangeEvent{}
	qt.ItfMap["winextras.QWinEvent_ITF"] = QWinEvent{}
	qt.ItfMap["winextras.QWinJumpList_ITF"] = QWinJumpList{}
	qt.FuncMap["winextras.NewQWinJumpList"] = NewQWinJumpList
	qt.ItfMap["winextras.QWinJumpListCategory_ITF"] = QWinJumpListCategory{}
	qt.FuncMap["winextras.NewQWinJumpListCategory"] = NewQWinJumpListCategory
	qt.EnumMap["winextras.QWinJumpListCategory__Custom"] = int64(QWinJumpListCategory__Custom)
	qt.EnumMap["winextras.QWinJumpListCategory__Recent"] = int64(QWinJumpListCategory__Recent)
	qt.EnumMap["winextras.QWinJumpListCategory__Frequent"] = int64(QWinJumpListCategory__Frequent)
	qt.EnumMap["winextras.QWinJumpListCategory__Tasks"] = int64(QWinJumpListCategory__Tasks)
	qt.ItfMap["winextras.QWinJumpListItem_ITF"] = QWinJumpListItem{}
	qt.FuncMap["winextras.NewQWinJumpListItem"] = NewQWinJumpListItem
	qt.EnumMap["winextras.QWinJumpListItem__Destination"] = int64(QWinJumpListItem__Destination)
	qt.EnumMap["winextras.QWinJumpListItem__Link"] = int64(QWinJumpListItem__Link)
	qt.EnumMap["winextras.QWinJumpListItem__Separator"] = int64(QWinJumpListItem__Separator)
	qt.ItfMap["winextras.QWinMime_ITF"] = QWinMime{}
	qt.ItfMap["winextras.QWinTaskbarButton_ITF"] = QWinTaskbarButton{}
	qt.FuncMap["winextras.NewQWinTaskbarButton"] = NewQWinTaskbarButton
	qt.ItfMap["winextras.QWinTaskbarProgress_ITF"] = QWinTaskbarProgress{}
	qt.FuncMap["winextras.NewQWinTaskbarProgress"] = NewQWinTaskbarProgress
	qt.ItfMap["winextras.QWinThumbnailToolBar_ITF"] = QWinThumbnailToolBar{}
	qt.FuncMap["winextras.NewQWinThumbnailToolBar"] = NewQWinThumbnailToolBar
	qt.ItfMap["winextras.QWinThumbnailToolButton_ITF"] = QWinThumbnailToolButton{}
	qt.FuncMap["winextras.NewQWinThumbnailToolButton"] = NewQWinThumbnailToolButton
	qt.ItfMap["winextras.QtWin_ITF"] = QtWin{}
	qt.FuncMap["winextras.QtWin_ColorizationColor"] = QtWin_ColorizationColor
	qt.FuncMap["winextras.QtWin_DisableBlurBehindWindow"] = QtWin_DisableBlurBehindWindow
	qt.FuncMap["winextras.QtWin_DisableBlurBehindWindow2"] = QtWin_DisableBlurBehindWindow2
	qt.FuncMap["winextras.QtWin_EnableBlurBehindWindow"] = QtWin_EnableBlurBehindWindow
	qt.FuncMap["winextras.QtWin_EnableBlurBehindWindow2"] = QtWin_EnableBlurBehindWindow2
	qt.FuncMap["winextras.QtWin_EnableBlurBehindWindow3"] = QtWin_EnableBlurBehindWindow3
	qt.FuncMap["winextras.QtWin_EnableBlurBehindWindow4"] = QtWin_EnableBlurBehindWindow4
	qt.FuncMap["winextras.QtWin_ErrorStringFromHresult"] = QtWin_ErrorStringFromHresult
	qt.FuncMap["winextras.QtWin_ExtendFrameIntoClientArea"] = QtWin_ExtendFrameIntoClientArea
	qt.FuncMap["winextras.QtWin_ExtendFrameIntoClientArea2"] = QtWin_ExtendFrameIntoClientArea2
	qt.FuncMap["winextras.QtWin_ExtendFrameIntoClientArea3"] = QtWin_ExtendFrameIntoClientArea3
	qt.FuncMap["winextras.QtWin_ExtendFrameIntoClientArea4"] = QtWin_ExtendFrameIntoClientArea4
	qt.FuncMap["winextras.QtWin_IsCompositionEnabled"] = QtWin_IsCompositionEnabled
	qt.FuncMap["winextras.QtWin_IsCompositionOpaque"] = QtWin_IsCompositionOpaque
	qt.FuncMap["winextras.QtWin_IsWindowExcludedFromPeek"] = QtWin_IsWindowExcludedFromPeek
	qt.FuncMap["winextras.QtWin_IsWindowExcludedFromPeek2"] = QtWin_IsWindowExcludedFromPeek2
	qt.FuncMap["winextras.QtWin_IsWindowPeekDisallowed"] = QtWin_IsWindowPeekDisallowed
	qt.FuncMap["winextras.QtWin_IsWindowPeekDisallowed2"] = QtWin_IsWindowPeekDisallowed2
	qt.FuncMap["winextras.QtWin_MarkFullscreenWindow"] = QtWin_MarkFullscreenWindow
	qt.FuncMap["winextras.QtWin_MarkFullscreenWindow2"] = QtWin_MarkFullscreenWindow2
	qt.FuncMap["winextras.QtWin_RealColorizationColor"] = QtWin_RealColorizationColor
	qt.FuncMap["winextras.QtWin_ResetExtendedFrame"] = QtWin_ResetExtendedFrame
	qt.FuncMap["winextras.QtWin_ResetExtendedFrame2"] = QtWin_ResetExtendedFrame2
	qt.FuncMap["winextras.QtWin_SetCompositionEnabled"] = QtWin_SetCompositionEnabled
	qt.FuncMap["winextras.QtWin_SetCurrentProcessExplicitAppUserModelID"] = QtWin_SetCurrentProcessExplicitAppUserModelID
	qt.FuncMap["winextras.QtWin_SetWindowDisallowPeek"] = QtWin_SetWindowDisallowPeek
	qt.FuncMap["winextras.QtWin_SetWindowDisallowPeek2"] = QtWin_SetWindowDisallowPeek2
	qt.FuncMap["winextras.QtWin_SetWindowExcludedFromPeek"] = QtWin_SetWindowExcludedFromPeek
	qt.FuncMap["winextras.QtWin_SetWindowExcludedFromPeek2"] = QtWin_SetWindowExcludedFromPeek2
	qt.FuncMap["winextras.QtWin_SetWindowFlip3DPolicy"] = QtWin_SetWindowFlip3DPolicy
	qt.FuncMap["winextras.QtWin_SetWindowFlip3DPolicy2"] = QtWin_SetWindowFlip3DPolicy2
	qt.FuncMap["winextras.QtWin_StringFromHresult"] = QtWin_StringFromHresult
	qt.FuncMap["winextras.QtWin_TaskbarActivateTab"] = QtWin_TaskbarActivateTab
	qt.FuncMap["winextras.QtWin_TaskbarActivateTab2"] = QtWin_TaskbarActivateTab2
	qt.FuncMap["winextras.QtWin_TaskbarActivateTabAlt"] = QtWin_TaskbarActivateTabAlt
	qt.FuncMap["winextras.QtWin_TaskbarActivateTabAlt2"] = QtWin_TaskbarActivateTabAlt2
	qt.FuncMap["winextras.QtWin_TaskbarAddTab"] = QtWin_TaskbarAddTab
	qt.FuncMap["winextras.QtWin_TaskbarAddTab2"] = QtWin_TaskbarAddTab2
	qt.FuncMap["winextras.QtWin_TaskbarDeleteTab"] = QtWin_TaskbarDeleteTab
	qt.FuncMap["winextras.QtWin_TaskbarDeleteTab2"] = QtWin_TaskbarDeleteTab2
	qt.FuncMap["winextras.QtWin_WindowFlip3DPolicy"] = QtWin_WindowFlip3DPolicy
	qt.FuncMap["winextras.QtWin_WindowFlip3DPolicy2"] = QtWin_WindowFlip3DPolicy2
	qt.EnumMap["winextras.QtWin__HBitmapNoAlpha"] = int64(QtWin__HBitmapNoAlpha)
	qt.EnumMap["winextras.QtWin__HBitmapPremultipliedAlpha"] = int64(QtWin__HBitmapPremultipliedAlpha)
	qt.EnumMap["winextras.QtWin__HBitmapAlpha"] = int64(QtWin__HBitmapAlpha)
	qt.EnumMap["winextras.QtWin__FlipDefault"] = int64(QtWin__FlipDefault)
	qt.EnumMap["winextras.QtWin__FlipExcludeBelow"] = int64(QtWin__FlipExcludeBelow)
	qt.EnumMap["winextras.QtWin__FlipExcludeAbove"] = int64(QtWin__FlipExcludeAbove)
}
