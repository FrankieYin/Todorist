package data

import (
				"github.com/FrankieYin/todo/internal/util"
)

type TodoList struct {
	Data map[int]*TodoItem `json:"data"`
	Order []int `json:"order"`
	Changed bool `json:"changed"`
}

func NewTodoList() *TodoList {
	var data= make(map[int]*TodoItem)
	var order= make([]int, 0)
	return &TodoList{Data: data, Order: order}
}

func (l *TodoList) changed() {
	l.Changed = true
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
				ProjList.DeleteTodo(pTodo)
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
				ProjList.DeleteTodo(pTodo)
			}
		} else {
			return util.InvalidIdError{Id: id}
		}
	}

	l.changed()
	// we have stored the archived todos; now delete them
	return l.DeleteTodo(ids...)
}

func (l *TodoList) AddTodo(pTodo *TodoItem) {
	l.Data[pTodo.Id] = pTodo
	l.Order = append(l.Order, pTodo.Id)
	ProjList.AddTodo(pTodo)
	l.changed()
}

func (l *TodoList) DeleteTodo(ids ...int) error {
	// check all ids are valid
	if id, ok := l.ContainsId(ids...); ok {
		for _, id = range ids {
			pTodo := l.Data[id]
			ProjList.DeleteTodo(pTodo)
			delete(l.Data, id)
			i := l.indexOf(id)
			l.Order = append(l.Order[:i], l.Order[i+1:]...)
		}
	} else {
		return util.InvalidIdError{Id: id}
	}

	l.changed()
	return nil
}

func (l *TodoList) DoTodo(undo bool, ids ...int) error {

	if id, ok := l.ContainsId(ids...); ok {
		for _, id = range ids {
			l.Data[id].Done = !undo
		}
	} else {
		return util.InvalidIdError{Id: id}
	}

	l.changed()
	return nil
}

/**
 WARNING: should be used on archive list ONLY
 */
func (l *TodoList) Merge(another *TodoList) {
	for k, pTodo := range another.Data {
		l.Data[k] = pTodo
	}
	l.changed()
}

func (l *TodoList) ContainsId(ids ...int) (int, bool) {
	for _, id := range ids {
		if _, ok := l.Data[id]; !ok {
			return id, false
		}
	}
	return -1, true
}

func (l *TodoList) indexOf(id int) int {
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

func (l *TodoList) Save() error {
	if !l.Changed {return nil}
	return save(Todos, todoJsonFilename)
}
