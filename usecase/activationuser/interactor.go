package activationuser

import "context"

//go:generate mockery --name Outport -output mocks/

type activationUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase ActivationUser
func NewUsecase(outputPort Outport) Inport {
	return &activationUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase ActivationUser
func (r *activationUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	//!

	return res, nil
}
