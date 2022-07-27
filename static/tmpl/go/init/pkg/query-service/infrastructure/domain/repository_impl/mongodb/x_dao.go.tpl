package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)


type Dao[T ddd.Entity] struct {
	dao *ddd_mongodb.Repository[T]
}

func newDao[T ddd.Entity](newFunc func() T, collectionName string, opts ...*DaoOptions) *Dao[T] {
	options := NewDaoOptions()
	options.Merge(opts...)
	coll := options.mongoDB.GetCollection(collectionName)
	return &Dao[T]{
		dao: ddd_mongodb.NewRepository[T](newFunc, options.mongoDB, coll),
	}
}

func (d *Dao[T]) Insert(ctx context.Context, entity T, opts ...*ddd_repository.SetOptions) (T, error) {
	return d.dao.Insert(ctx, entity, opts...).Result()
}

func (d *Dao[T]) InsertMany(ctx context.Context, entity *[]T, opts ...*ddd_repository.SetOptions) error {
	return d.dao.InsertMany(ctx, entity, opts...).GetError()
}

func (d *Dao[T]) Update(ctx context.Context, entity T, opts ...*ddd_repository.SetOptions) (T, error) {
	return d.dao.Update(ctx, entity, opts...).Result()
}

func (d *Dao[T]) UpdateManyById(ctx context.Context, entity []T, opts ...*ddd_repository.SetOptions) error {
	return d.dao.UpdateManyById(ctx, entity, opts...).GetError()
}

func (d *Dao[T]) UpdateManyByFilter(ctx context.Context,  tenantId, filter string, data interface{}, opts ...*ddd_repository.SetOptions) error {
	return d.dao.UpdateManyByFilter(ctx, tenantId, filter, data, opts...).GetError()
}

func (d *Dao[T]) DeleteById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.SetOptions) error {
	return d.dao.DeleteById(ctx, tenantId, id, opts...).GetError()
}

func (d *Dao[T]) DeleteAll(ctx context.Context, tenantId string, opts ...*ddd_repository.SetOptions) error {
	return d.dao.DeleteAll(ctx, tenantId, opts...).GetError()
}

func (d *Dao[T]) DeleteByMap(ctx context.Context, tenantId string, filterMap map[string]interface{}, opts ...*ddd_repository.SetOptions) error {
	return d.dao.DeleteByMap(ctx, tenantId, filterMap, opts...).GetError()
}

func (d *Dao[T]) FindById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.FindOptions) (T, bool, error) {
	return d.dao.FindById(ctx, tenantId, id, opts...).Result()
}

func (d *Dao[T]) FindAll(ctx context.Context, tenantId string, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	return d.dao.FindAll(ctx, tenantId, opts...)
}

func (d *Dao[T]) FindListByMap(ctx context.Context, tenantId string, filterMap map[string]interface{}, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	return d.dao.FindListByMap(ctx, tenantId, filterMap, opts...)
}

func (d *Dao[T]) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...*ddd_repository.FindOptions) *ddd_repository.FindPagingResult[T] {
	return d.dao.FindPaging(ctx, query, opts...)
}

type DaoOptions struct {
	mongoDB *ddd_mongodb.MongoDB
}

func NewDaoOptions() *DaoOptions {
	return &DaoOptions{}
}

func (o *DaoOptions) SetMongoDB(mongoDB *ddd_mongodb.MongoDB) *DaoOptions {
	o.mongoDB = mongoDB
	return o
}

func (o *DaoOptions) Merge(opts ...*DaoOptions) {
	if opts != nil {
		for _, item := range opts {
			if item.mongoDB != nil {
				o.mongoDB = item.mongoDB
			}
		}
	}
	if o.mongoDB == nil {
		o.mongoDB = restapp.GetMongoDB()
	}
}
