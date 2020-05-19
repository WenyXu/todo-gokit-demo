package service

import (
	"context"
	db2 "todo/pkg/db"
	"todo/pkg/io"
)

// TodoService describes the service.
type TodoService interface {
	Get(ctx context.Context) (t []io.Todo, err error)
	Add(ctx context.Context, todo io.Todo) (t io.Todo, err error)
	Update(ctx context.Context, todo io.Todo) (t io.Todo, err error)
	Delete(ctx context.Context, id string) (err error)
	GetById(ctx context.Context, id string) (t io.Todo, error error)
}

type basicTodoService struct {
	db db2.Service
}

func (b *basicTodoService) Get(ctx context.Context) (t []io.Todo, err error) {
	t, err = b.db.Get()
	return t, err
}
func (b *basicTodoService) Add(ctx context.Context, todo io.Todo) (t io.Todo, err error) {
	t, err = b.db.Add(todo)
	return t, err
}
func (b *basicTodoService) Update(ctx context.Context, todo io.Todo) (t io.Todo, err error) {
	t, err = b.db.Update(todo)
	return t, err
}
func (b *basicTodoService) Delete(ctx context.Context, id string) (err error) {
	err = b.db.Delete(id)
	return err
}

// NewBasicTodoService returns a naive, stateless implementation of TodoService.
func NewBasicTodoService() TodoService {
	return &basicTodoService{db: db2.NewService()}
}

// New returns a TodoService with all of the expected middleware wired in.
func New(middleware []Middleware) TodoService {
	var svc TodoService = NewBasicTodoService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicTodoService) GetById(ctx context.Context, id string) (t io.Todo, error error) {
	// TODO implement the business logic of GetById
	return t, error
}
