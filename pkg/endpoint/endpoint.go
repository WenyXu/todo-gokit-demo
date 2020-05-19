package endpoint

import (
	"context"
	io "todo/pkg/io"
	service "todo/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	T   []io.Todo `json:"t"`
	Err error     `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		t, err := s.Get(ctx)
		return GetResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Todo io.Todo `json:"todo"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	T   io.Todo `json:"t"`
	Err error   `json:"err"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		t, err := s.Add(ctx, req.Todo)
		return AddResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Err
}

// UpdateRequest collects the request parameters for the Update method.
type UpdateRequest struct {
	Todo io.Todo `json:"todo"`
}

// UpdateResponse collects the response parameters for the Update method.
type UpdateResponse struct {
	T   io.Todo `json:"t"`
	Err error   `json:"err"`
}

// MakeUpdateEndpoint returns an endpoint that invokes Update on the service.
func MakeUpdateEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		t, err := s.Update(ctx, req.Todo)
		return UpdateResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateResponse) Failed() error {
	return r.Err
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Err error `json:"err"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.Id)
		return DeleteResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (t []io.Todo, err error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).T, response.(GetResponse).Err
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, todo io.Todo) (t io.Todo, err error) {
	request := AddRequest{Todo: todo}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).T, response.(AddResponse).Err
}

// Update implements Service. Primarily useful in a client.
func (e Endpoints) Update(ctx context.Context, todo io.Todo) (t io.Todo, err error) {
	request := UpdateRequest{Todo: todo}
	response, err := e.UpdateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateResponse).T, response.(UpdateResponse).Err
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (err error) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Err
}

// GetByIdRequest collects the request parameters for the GetById method.
type GetByIdRequest struct {
	Id string `json:"id"`
}

// GetByIdResponse collects the response parameters for the GetById method.
type GetByIdResponse struct {
	T     io.Todo `json:"t"`
	Error error   `json:"error"`
}

// MakeGetByIdEndpoint returns an endpoint that invokes GetById on the service.
func MakeGetByIdEndpoint(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIdRequest)
		t, error := s.GetById(ctx, req.Id)
		return GetByIdResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIdResponse) Failed() error {
	return r.Error
}

// GetById implements Service. Primarily useful in a client.
func (e Endpoints) GetById(ctx context.Context, id string) (t io.Todo, error error) {
	request := GetByIdRequest{Id: id}
	response, err := e.GetByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIdResponse).T, response.(GetByIdResponse).Error
}
