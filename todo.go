package main

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"sync"
	"time"
)

type Todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TodoManager struct {
	todos []Todo
	m     sync.Mutex
}

func NewTodoManager() TodoManager {
	return TodoManager{
		todos: make([]Todo, 0),
		m:     sync.Mutex{},
	}
}

func (tm *TodoManager) GetAll() []Todo {
	return tm.todos
}

type CreateTodoRequest struct {
	Name string `json:"name"`
}

type EditTodoRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (tm *TodoManager) Create(createTodoRequest CreateTodoRequest) Todo {
	tm.m.Lock()
	defer tm.m.Unlock()

	newTodo := Todo{
		ID:   strconv.FormatInt(time.Now().UnixMilli(), 10),
		Name: createTodoRequest.Name,
	}

	tm.todos = append(tm.todos, newTodo)

	return newTodo
}

func (tm *TodoManager) Remove(ID string) error {
	tm.m.Lock()
	defer tm.m.Unlock()

	index := -1

	for i, t := range tm.todos {
		if t.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return echo.ErrNotFound
	}

	tm.todos = append(tm.todos[:index], tm.todos[index+1:]...)

	return nil
}

func (tm *TodoManager) Edit(ID string, Name string) error {
	tm.m.Lock()
	defer tm.m.Unlock()

	index := -1

	for i, t := range tm.todos {
		if t.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return echo.ErrNotFound
	}

	tm.todos[index].Name = Name

	return nil
}
