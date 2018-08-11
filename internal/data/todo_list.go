package data

import (
	"fmt"

	"github.com/FrankieYin/todo/internal/util"
	"encoding/json"
	"os"
)

type TodoList struct {
	Data map[int]*TodoItem `json:"data"`
	Order []int `json:"order"`
}

func NewTodoList() *TodoList {
	var data= make(map[int]*TodoItem)
	var order= make([]int, 0)
	return &TodoList{Data: data, Order: order}
}

func (l *TodoList) ArchTodo(n int, ids ...int) (*TodoList, error) {

	archived := NewTodoList()

	if len(ids) == 0 { // archive all the tasks done
		for k, pTodo := range l.Data {
			if pTodo.Done {
				ids = append(ids, k)
				n++
				pTodo.ArchId = n
				archived.Data[n] = pTodo
			}
		}
	} else {
		if id, ok := l.ContainsId(ids...); ok {
			for _, id = range ids {
				pTodo := l.Data[id]
				n++
				pTodo.ArchId = n
				archived.Data[n] = pTodo
			}
		} else {
			msg := fmt.Sprintf("Error: found no task with id %d\n", id)
			return nil, util.InvalidIdError{Msg: msg}
		}
	}

	// we have stored the archived todos; now delete them
	err := l.DeleteTodo(ids...)

	return archived, err
}

func (l *TodoList) AddTodo(pTodo *TodoItem) {
	l.Data[pTodo.Id] = pTodo
	l.Order = append(l.Order, pTodo.Id)
}

func (l *TodoList) DeleteTodo(ids ...int) error {
	// check all ids are valid
	if id, ok := l.ContainsId(ids...); ok {
		for _, id = range ids {
			delete(l.Data, id)
			i := l.IndexOf(id)
			l.Order = append(l.Order[:i], l.Order[i+1:]...)
		}
	} else {
		msg := fmt.Sprintf("Error: found no task with id %d\n", id)
		return util.InvalidIdError{Msg: msg}
	}

	return nil
}

func (l *TodoList) DoTodo(ids ...int) error {

	if id, ok := l.ContainsId(ids...); ok {
		for _, id = range ids {
			l.Data[id].Done = true
		}
	} else {
		msg := fmt.Sprintf("Error: found no task with id %d\n", id)
		return util.InvalidIdError{Msg: msg}
	}

	return nil
}

/**
 WARNING: should be used on archive list ONLY
 */
func (l *TodoList) Merge(another *TodoList) {
	for k, pTodo := range another.Data {
		l.Data[k] = pTodo
	}
}

func (l *TodoList) ContainsId(ids ...int) (int, bool) {
	for _, id := range ids {
		if _, ok := l.Data[id]; !ok {
			return id, false
		}
	}
	return -1, true
}

func (l *TodoList) IndexOf(id int) int {
	for i, v := range l.Order {
		if v == id {
			return i
		}
	}
	return -1
}

func (l *TodoList) GetTodoById(id int) *TodoItem {
	if t, ok := l.Data[id]; ok {
		return t
	}
	return nil
}

func (l *TodoList) Save(filename string) error {
	// save todolist
	var b []byte
	var err error
	var f *os.File
	if b, err = json.Marshal(l); err != nil { return err}

	if f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644); err != nil { return err}

	defer f.Close()

	if _, err = f.Write(b); err != nil { return err}

	return nil
}