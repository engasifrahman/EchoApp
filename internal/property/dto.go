package property

type CreatePropertyRequest struct {
    Name        string `json:"name"`
    Location    string `json:"location"`
    Description string `json:"description"`
}

type UpdatePropertyRequest struct {
    Name        string `json:"name"`
    Location    string `json:"location"`
    Description string `json:"description"`
}
