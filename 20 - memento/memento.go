package memento

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"time"
)

// Memento 备忘录接口
type Memento interface {
	Save(string, []byte)
	Get(string) []byte
}

// Caretaker 备忘录管理
type Caretaker interface {
	Backups(interface{}, ...string) error
	Recovery(interface{}, ...string) error
}

// Originator 备份发起者接口
type Originator interface {
	Update()
}

var _ Memento = (storage)(nil)

// Storage 备份存储
type storage map[string][]byte

// Save .
func (s storage) Save(id string, data []byte) {
	s[id] = data
}

// Get .
func (s storage) Get(id string) []byte {
	return s[id]
}

var _ Caretaker = (*StorageCaretaker)(nil)

// StorageCaretaker .
type StorageCaretaker struct {
	storage Memento
}

// NewStorageCaretaker .
func NewStorageCaretaker() *StorageCaretaker {
	return &StorageCaretaker{
		storage: make(storage, 0),
	}
}

// Backups 备份
func (c *StorageCaretaker) Backups(obj interface{}, fields ...string) (err error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	id := v.FieldByName("ID").String()

	backup := make(map[string]interface{}, len(fields))
	for _, fieldname := range fields {
		fieldValue := v.FieldByName(fieldname)
		backup[fieldname] = fieldValue.Interface()
	}

	data, err := json.Marshal(backup)
	if err != nil {
		return
	}
	c.storage.Save(id, data)
	return
}

// Recovery 恢复
func (c *StorageCaretaker) Recovery(obj interface{}, fields ...string) (err error) {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return errors.New("can only assign values with pointer to struct")
	}

	v = v.Elem()
	id := v.FieldByName("ID").String()

	data := c.storage.Get(id)
	if data == nil {
		return errors.New("Backup does not exist")
	}

	backup := make(map[string]interface{}, 0)
	err = json.Unmarshal(data, &backup)
	if err != nil {
		return err
	}

	for k, val := range backup {
		field := v.FieldByName(k)
		field.Set(reflect.ValueOf(val))
	}
	return
}

// Edit Originator
type Edit struct {
	ID         string
	Content    string
	UpdateTime string

	storageCaretaker Caretaker
}

// NewEdit .
func NewEdit(id, content string, sc Caretaker) *Edit {
	return &Edit{
		ID:               id,
		Content:          content,
		UpdateTime:       time.Now().Format("2006/01/02 15:04:05"),
		storageCaretaker: sc,
	}
}

// Update .
func (e *Edit) Update(content string) {
	e.storageCaretaker.Backups(e, "Content", "UpdateTime")
	defer func() {
		if err := recover(); err != nil {
			log.Println("发生异常回滚")
			err = e.storageCaretaker.Recovery(e, "Content", "UpdateTime")
			if err != nil {
				panic(err)
			}
		}
	}()

	e.Content = content
	e.UpdateTime = time.Now().Format("2006/01/02 15:04:05")
	if len(e.Content) > 10 {
		panic("context is too long")
	}
}
