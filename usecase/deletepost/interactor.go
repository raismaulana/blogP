package deletepost

import (
	"context"

	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/domain/repository"
)

//go:generate mockery --name Outport -output mocks/

type deletePostInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase DeletePost
func NewUsecase(outputPort Outport) Inport {
	return &deletePostInteractor{
		outport: outputPort,
	}
}

// Execute the usecase DeletePost
func (r *deletePostInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		postObj, err := r.outport.FindPostByID(ctx, req.ID)
		if err != nil {
			return apperror.ObjectNotFound.Var(postObj)
		}
		if postObj.UserID != req.UserID && req.Role != "king" {
			return apperror.ProhibitedFromAccessingOtherPeoplesResources
		}
		postObj.DeleteAssociation()
		err = r.outport.DeletePost(ctx, postObj)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
