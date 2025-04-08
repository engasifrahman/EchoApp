package property

import (
	"context"
	"gorm.io/gorm"
)

// Repository interface defines expected behavior
type Repository interface {
	Create(ctx context.Context, p *Property) error
	GetByID(ctx context.Context, id string) (*Property, error)
	List(ctx context.Context) ([]*Property, error)
	Update(ctx context.Context, p *Property) error
	Delete(ctx context.Context, id string) error
}

// GormRepository is the GORM-based implementation
type GormRepository struct {
	db *gorm.DB
}

// Constructor
func NewRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

// Create a new property
func (r *GormRepository) Create(ctx context.Context, p *Property) error {
	return r.db.WithContext(ctx).Create(p).Error
}

// GetByID fetches a property by ID
func (r *GormRepository) GetByID(ctx context.Context, id string) (*Property, error) {
	var p Property
	err := r.db.WithContext(ctx).First(&p, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// List fetches all properties
func (r *GormRepository) List(ctx context.Context) ([]*Property, error) {
	var properties []*Property
	err := r.db.WithContext(ctx).Find(&properties).Error
	return properties, err
}

// Update modifies an existing property
func (r *GormRepository) Update(ctx context.Context, p *Property) error {
	return r.db.WithContext(ctx).Save(p).Error
}

// Delete removes a property by ID
func (r *GormRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&Property{}, "id = ?", id).Error
}
