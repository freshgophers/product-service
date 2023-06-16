package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"

	"product-service/internal/domain/book"
	"product-service/pkg/store"
)

type BookRepository struct {
	db map[string]book.Entity
	sync.RWMutex
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		db: make(map[string]book.Entity),
	}
}

func (r *BookRepository) Select(ctx context.Context) (dest []book.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest = make([]book.Entity, 0, len(r.db))
	for _, data := range r.db {
		dest = append(dest, data)
	}

	return
}

func (r *BookRepository) Create(ctx context.Context, data book.Entity) (dest string, err error) {
	r.Lock()
	defer r.Unlock()

	id := r.generateID()
	data.ID = id
	r.db[id] = data

	return id, nil
}

func (r *BookRepository) Get(ctx context.Context, id string) (dest book.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest, ok := r.db[id]
	if !ok {
		err = store.ErrorNotFound
		return
	}

	return
}

func (r *BookRepository) Update(ctx context.Context, id string, data book.Entity) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		return store.ErrorNotFound
	}
	r.db[id] = data

	return
}

func (r *BookRepository) Delete(ctx context.Context, id string) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		return store.ErrorNotFound
	}
	delete(r.db, id)

	return
}

func (r *BookRepository) generateID() string {
	return uuid.New().String()
}
