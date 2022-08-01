package mongodb_base

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func GetMongoDB() *ddd_mongodb.MongoDB {
	return restapp.GetMongoDB()
}

type BaseRepository[T ddd.Entity] struct {
	super *ddd_mongodb.Repository[T]
}

func NewBaseRepository[T ddd.Entity](newFunc func() T, collectionName string, opts ...*RepositoryOptions) *BaseRepository[T] {
	options := NewRepositoryOptions()
	options.Merge(opts...)
	coll := options.mongoDB.GetCollection(collectionName)
	return &BaseRepository[T]{
		super: ddd_mongodb.NewRepository[T](newFunc, options.mongoDB, coll),
	}
}

func (u *BaseRepository[T]) CreateMany(ctx context.Context, entity []T, opts ...*ddd_repository.SetOptions) error {
	return u.super.InsertMany(ctx, entity, opts...).GetError()
}

func (u *BaseRepository[T]) UpdateManyById(ctx context.Context, entity []T, opts ...*ddd_repository.SetOptions) error {
	return u.super.UpdateManyById(ctx, entity, opts...).GetError()
}

func (u *BaseRepository[T]) UpdateManyByFilter(ctx context.Context,  tenantId, filter string, data interface{}, opts ...*ddd_repository.SetOptions) error {
	return u.super.UpdateManyByFilter(ctx, tenantId, filter, data, opts...).GetError()
}

func (u *BaseRepository[T]) Create(ctx context.Context, entity T, opts ...*ddd_repository.SetOptions) (T, error) {
	return u.super.Insert(ctx, entity, opts...).Result()
}

func (u *BaseRepository[T]) Update(ctx context.Context, entity T, opts ...*ddd_repository.SetOptions) (T, error) {
	return u.super.Update(ctx, entity, opts...).Result()
}

func (u *BaseRepository[T]) DeleteById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.SetOptions) error {
	return u.super.DeleteById(ctx, tenantId, id, opts...).GetError()
}

func (u *BaseRepository[T]) DeleteAll(ctx context.Context, tenantId string, opts ...*ddd_repository.SetOptions) error {
	return u.super.DeleteAll(ctx, tenantId, opts...).GetError()
}

func (u *BaseRepository[T]) DeleteByMap(ctx context.Context, tenantId string, filterMap map[string]interface{}, opts ...*ddd_repository.SetOptions) error {
	return u.super.DeleteByMap(ctx, tenantId, filterMap, opts...).GetError()
}

func (u *BaseRepository[T]) FindById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.FindOptions) (T, bool, error) {
	return u.super.FindById(ctx, tenantId, id, opts...).Result()
}

func (u *BaseRepository[T]) FindAll(ctx context.Context, tenantId string, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	return u.super.FindAll(ctx, tenantId, opts...)
}

func (u *BaseRepository[T]) FindListByMap(ctx context.Context, tenantId string, filterMap map[string]interface{}, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	return u.super.FindListByMap(ctx, tenantId, filterMap, opts...)
}

func (u *BaseRepository[T]) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...*ddd_repository.FindOptions) *ddd_repository.FindPagingResult[T] {
	return u.super.FindPaging(ctx, query, opts...)
}



type RepositoryOptions struct {
	mongoDB *ddd_mongodb.MongoDB
}

func NewRepositoryOptions() *RepositoryOptions {
	return &RepositoryOptions{}
}

func (o *RepositoryOptions) SetMongoDB(mongoDB *ddd_mongodb.MongoDB) *RepositoryOptions {
	o.mongoDB = mongoDB
	return o
}

func (o *RepositoryOptions) Merge(opts ...*RepositoryOptions) {
	if opts != nil {
		for _, item := range opts {
			if item.mongoDB != nil {
				o.mongoDB = item.mongoDB
			}
		}
	}
	if o.mongoDB == nil {
		o.mongoDB = GetMongoDB()
	}
}
