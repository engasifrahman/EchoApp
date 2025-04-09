package property

type CreatePropertyRequest struct {
    Name        string `json:"name" validate:"required,alpha_spaces_numbers,min=3,max=255"`
    Location    string `json:"location" validate:"required,alpha_spaces_numbers"`
    Description string `json:"description" validate:"omitempty,max=1000"`
}

type UpdatePropertyRequest struct {
    Name        string `json:"name"`
    Location    string `json:"location"`
    Description string `json:"description"`
}
