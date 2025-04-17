package ai

import (
	"context"
	"errors"
)

type Suggestion struct {
	ID        string `json:"id"`
	ResumeID  string `json:"resume_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type Service struct {
	// Add AI model dependencies here
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateSuggestions(ctx context.Context, resumeContent string) ([]Suggestion, error) {
	// Implement AI suggestion generation logic
	return nil, errors.New("not implemented")
}

func (s *Service) GenerateContent(ctx context.Context, prompt string) (string, error) {
	// Implement content generation logic
	return "", errors.New("not implemented")
}
