package repository

import (
	"product-service/internal/domain/author"
	"product-service/internal/domain/book"
	"product-service/internal/domain/category"
	"product-service/internal/domain/member"
	"product-service/internal/domain/product"
	"product-service/internal/repository/memory"
	"product-service/internal/repository/postgres"
	"product-service/pkg/store"
)

// Configuration is an alias for a function that will take in a pointer to a Repository and modify it
type Configuration func(r *Repository) error

// Repository is an implementation of the Repository
type Repository struct {
	postgres *store.Database

	Author   author.Repository
	Book     book.Repository
	Member   member.Repository
	Product  product.Repository
	Category category.Repository
}

// New takes a variable amount of Configuration functions and returns a new Repository
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Repository, err error) {
	// Create the repository
	s = &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

// Close closes the repository and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (r *Repository) Close() {
	if r.postgres != nil {
		r.postgres.Client.Close()
	}
}

// WithMemoryStore applies a memory store to the Repository
func WithMemoryStore() Configuration {
	return func(s *Repository) (err error) {
		// Create the memory store, if we needed parameters, such as connection strings they could be inputted here
		s.Author = memory.NewAuthorRepository()
		s.Book = memory.NewBookRepository()
		s.Member = memory.NewMemberRepository()
		s.Category = memory.NewCategoryRepository()
		s.Product = memory.NewProductRepository()

		return
	}
}

// WithPostgresStore applies a postgres store to the Repository
func WithPostgresStore(schema, dataSourceName string) Configuration {
	return func(s *Repository) (err error) {
		// Create the postgres store, if we needed parameters, such as connection strings they could be inputted here
		s.postgres, err = store.NewDatabase(schema, dataSourceName)
		if err != nil {
			return
		}

		err = s.postgres.Migrate()
		if err != nil {
			return
		}

		s.Author = postgres.NewAuthorRepository(s.postgres.Client)
		s.Book = postgres.NewBookRepository(s.postgres.Client)
		s.Member = postgres.NewMemberRepository(s.postgres.Client)
		s.Category = postgres.NewCategoryRepository(s.postgres.Client)
		s.Product = postgres.NewProductRepository(s.postgres.Client)
		return
	}
}
