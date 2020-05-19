package db

import "todo/pkg/io"

// Service interface
type Service interface {
	Get() (t []io.Todo, err error)
	Add(todo io.Todo) (t io.Todo, err error)
	Update(todo io.Todo) (t io.Todo, err error)
	Delete(id string)(err error)
}

func mapToSlice(m map[string]io.Todo) []io.Todo {
	s := make([]io.Todo, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// service struct  
type svc struct {
	todoList map[string]io.Todo
}

func (s *svc) Get()(t []io.Todo, err error){
	return mapToSlice(s.todoList),nil
}

func (s *svc) Add(todo io.Todo) (t io.Todo, err error){
	s.todoList[todo.ID]=todo
	return todo,nil
}

func (s *svc) Update(todo io.Todo) (t io.Todo, err error){
	s.todoList[todo.ID]=todo
	return todo,nil
}

func (s *svc) Delete(id string) (err error){
	delete(s.todoList,id)
	return nil
}

func NewService() Service{
	return &svc{todoList: map[string]io.Todo{}}
}
