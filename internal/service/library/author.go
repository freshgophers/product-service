package library

import (
	"context"

	"product-service/internal/domain/author"
)

func (s *Service) ListAuthors(ctx context.Context) (res []author.Response, err error) {
	data, err := s.authorRepository.Select(ctx)
	if err != nil {
		return
	}
	res = author.ParseFromEntities(data)

	return
}

func (s *Service) AddAuthor(ctx context.Context, req author.Request) (res author.Response, err error) {
	data := author.Entity{
		FullName:  &req.FullName,
		Pseudonym: &req.Pseudonym,
		Specialty: &req.Specialty,
	}

	data.ID, err = s.authorRepository.Create(ctx, data)
	if err != nil {
		return
	}
	res = author.ParseFromEntity(data)

	return
}

func (s *Service) GetAuthor(ctx context.Context, id string) (res author.Response, err error) {
	data, err := s.authorRepository.Get(ctx, id)
	if err != nil {
		return
	}
	res = author.ParseFromEntity(data)

	return
}

func (s *Service) UpdateAuthor(ctx context.Context, id string, req author.Request) (err error) {
	data := author.Entity{
		FullName:  &req.FullName,
		Pseudonym: &req.Pseudonym,
		Specialty: &req.Specialty,
	}
	return s.authorRepository.Update(ctx, id, data)
}

func (s *Service) DeleteAuthor(ctx context.Context, id string) (err error) {
	return s.authorRepository.Delete(ctx, id)
}
