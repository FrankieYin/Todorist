package data

import "fmt"

type TodoList struct {
	Data map[int]*TodoItem `json:"data"`
	Order []int `json:"order"`
}

type InvalidIdError struct {
	msg string
}

func (e InvalidIdError) Error() string {
	return e.msg
}

func NewTodoList() *TodoList {
	var data= make(map[int]*TodoItem)
	var order= make([]int, 0)
	return &TodoList{Data: data, Order: order}
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
			for i, v := range l.Order {
				if v == id {
					l.Order = append(l.Order[:i], l.Order[i+1:]...)
				}
			}
		}
	} else {
		msg := fmt.Sprintf("todo del error: found no task with id %d\n", id)
		return InvalidIdError{msg: msg}
	}

	return nil
}

func (l *TodoList) DoneTodo(ids ...int) error {

	if id, ok := l.ContainsId(ids...); ok {
		for _, id = range ids {
			l.Data[id].Done = true
		}
	} else {
		msg := fmt.Sprintf("todo done error: found no task with id %d\n", id)
		return InvalidIdError{msg: msg}
	}

	return nil
}

func (l *TodoList) ContainsId(ids ...int) (int, bool) {
	for _, id := range ids {
		if _, ok := l.Data[id]; !ok {
			return id, false
		}
	}
	return -1, true
}