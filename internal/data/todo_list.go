package data

import (
	"fmt"
			"github.com/FrankieYin/todo/internal/util"
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

func (l *TodoList) ArchTodo(ids ...int) error {

	n := len(ArchList.Data)

	if len(ids) == 0 { // archive all the tasks done
		for id, pTodo := range l.Data {
			if pTodo.Done {
				ids = append(ids, id)
				n++
				pTodo.ArchId = n
				ArchList.Data[n] = pTodo
				// delete from their corresponding project
				ProjList.GetProject(pTodo.Project).DeleteTodo(id)
			}
		}
	} else {
		if id, ok := l.ContainsId(ids...); ok {
			for _, id = range ids {
				pTodo := l.Data[id]
				n++
				pTodo.ArchId = n
				ArchList.Data[n] = pTodo
				// delete from their corresponding project
				ProjList.GetProject(pTodo.Project).DeleteTodo(id)
			}
		} else {
			msg := fmt.Sprintf("Error: found no task with id %d\n", id)
			return util.InvalidIdError{Msg: msg}
		}
	}

	// we have stored the archived todos; now delete them
	return l.DeleteTodo(ids...)
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

func (l *TodoList) DoTodo(undo bool, ids ...int) error {

	if id, ok := l.ContainsId(ids...); ok {
		for _, id = range ids {
			l.Data[id].Done = !undo
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
