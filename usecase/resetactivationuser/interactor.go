package resetactivationuser

import "context"

//go:generate mockery --name Outport -output mocks/

type resetActivationUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ResetActivationUser
func NewUsecase(outputPort Outport) Inport {
	return &resetActivationUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ResetActivationUser
func (r *resetActivationUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	//!

	return res, nil
}
