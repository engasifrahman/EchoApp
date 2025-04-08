package property

import "errors"

func ValidateCreateProperty(req CreatePropertyRequest) error {
    if req.Name == "" {
        return errors.New("property name is required")
    }
    if req.Location == "" {
        return errors.New("property location is required")
    }
    return nil
}
