package property

import (
	"context"
)

type UseCase interface {
	CreateProperty(ctx context.Context, input CreatePropertyRequest) (*Property, error)
	GetProperty(ctx context.Context, id string) (*Property, error)
	ListProperties(ctx context.Context) ([]*Property, error)
	UpdateProperty(ctx context.Context, id string, input UpdatePropertyRequest) error
	DeleteProperty(ctx context.Context, id string) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return &useCase{repo: r}
}

func (u *useCase) CreateProperty(ctx context.Context, input CreatePropertyRequest) (*Property, error) {
	prop := &Property{
		Name:        input.Name,
		Location:    input.Location,
		Description: input.Description,
	}
	return prop, u.repo.Create(ctx, prop)
}

func (u *useCase) GetProperty(ctx context.Context, id string) (*Property, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *useCase) ListProperties(ctx context.Context) ([]*Property, error) {
	return u.repo.List(ctx)
}

func (u *useCase) UpdateProperty(ctx context.Context, id string, input UpdatePropertyRequest) error {
	existing, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	existing.Name = input.Name
	existing.Location = input.Location
	existing.Description = input.Description
	return u.repo.Update(ctx, existing)
}

func (u *useCase) DeleteProperty(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
