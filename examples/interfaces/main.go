package interfaces

import "fmt"

type Storage interface {
	Save(data string) error
}

type FileStorage struct{}

func (fs FileStorage) Save(data string) error {
	fmt.Println("Saving data to file:", data)
	return nil
}

type Service struct {
	storage FileStorage
}

type ProcessResponse struct {
	Success bool
	Message string
}

func (s *Service) Process(data string) ProcessResponse {
	err := s.storage.Save(data)

	if err != nil {
		return ProcessResponse{Success: false, Message: err.Error()}
	}

	return ProcessResponse{Success: true, Message: "Data saved successfully"}

}
