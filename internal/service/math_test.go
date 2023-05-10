package service

import (
	"context"
	"errors"
	"testing"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
	"github.com/kdl-dev/iConText-test-task/internal/repository"
	"github.com/kdl-dev/iConText-test-task/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestIncrement(t *testing.T) {
	var val int64 = 19
	incrementInput := &entity.IncrementDTO{Key: "age", Value: &val}

	tests := []struct {
		name       string
		input      *entity.IncrementDTO
		expections func(context.Context, *mocks.MathOperation)
		err        error
	}{
		{
			name:  "success increment test",
			input: incrementInput,
			expections: func(ctx context.Context, mo *mocks.MathOperation) {
				mo.On("Increment", ctx, incrementInput).Return(&entity.IncrementResult{}, nil)
			},
		},
		{
			name:  "storage error test",
			input: incrementInput,
			expections: func(ctx context.Context, mo *mocks.MathOperation) {
				mo.On("Increment", ctx, incrementInput).Return(nil, errors.New("some error"))
			},
			err: errors.New("service.math.Increment error: some error"),
		},
	}

	for _, test := range tests {
		t.Logf("running: %s", test.name)
		mathOperationRepo := mocks.NewMathOperation(t)
		test.expections(context.Background(), mathOperationRepo)
		svc := NewMathOperationService(&repository.Repository{MathOperation: mathOperationRepo})

		_, err := svc.Increment(context.Background(), test.input)
		if err != nil {
			if test.err != nil {
				assert.EqualError(t, test.err, err.Error())
			} else {
				t.Errorf("expected no error, found: %s", err.Error())
			}
		}

		mathOperationRepo.AssertExpectations(t)
	}
}
