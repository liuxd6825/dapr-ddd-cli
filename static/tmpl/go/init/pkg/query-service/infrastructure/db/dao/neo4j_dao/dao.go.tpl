package neo4j_dao

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_neo4j"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Dao[T ddd.Entity] struct {
	dao *ddd_neo4j.Neo4jDao[T]
}

func NewDao[T ddd.Entity](labels []string, opts ...*RepositoryOptions) *Dao[T] {
	options := NewRepositoryOptions()
	options.Merge(opts...)
	return &Dao[T]{
		dao: ddd_neo4j.Neo4jDao[T](options.driver, ddd_neo4j.NewReflectBuilder(labels...)),
	}
}

func (u *Dao[T]) Insert(ctx context.Context, entity T, opts ...*ddd_repository.SetOptions) error {
	return u.dao.Insert(ctx, entity, opts...).GetError()
}

func (u *Dao[T]) InsertMany(ctx context.Context, entity []T, opts ...*ddd_repository.SetOptions) error {
	return u.dao.InsertMany(ctx, entity, opts...).GetError()
}

func (u *Dao[T]) Update(ctx context.Context, entity T, opts ...*ddd_repository.SetOptions) error {
	return u.dao.Update(ctx, entity, opts...).GetError()
}

func (u *Dao[T]) UpdateMany(ctx context.Context, entity []T, opts ...*ddd_repository.SetOptions) error {
	return u.dao.UpdateMany(ctx, entity, opts...).GetError()
}

func (u *Dao[T]) DeleteById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.SetOptions) error {
	return u.dao.DeleteById(ctx, tenantId, id, opts...)
}

func (u *Dao[T]) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...*ddd_repository.SetOptions) error {
	return u.dao.DeleteByIds(ctx, tenantId, ids)
}

func (u *Dao[T]) DeleteByFilter(ctx context.Context, tenantId string, filter string, opts ...*ddd_repository.SetOptions) error {
	return u.dao.DeleteByFilter(ctx, tenantId, filter, opts...)
}

func (u *Dao[T]) DeleteAll(ctx context.Context, tenantId string, opts ...*ddd_repository.SetOptions) error {
	return u.dao.DeleteAll(ctx, tenantId, opts...)
}

func (u *Dao[T]) FindById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.FindOptions) (T, bool, error) {
	return u.dao.FindById(ctx, tenantId, id, opts...)
}

func (u *Dao[T]) FindByIds(ctx context.Context, tenantId string, ids []string, opts ...*ddd_repository.FindOptions) ([]T, bool, error) {
	return u.dao.FindByIds(ctx, tenantId, ids, opts...)
}

func (u *Dao[T]) FindAll(ctx context.Context, tenantId string, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	return u.dao.FindAll(ctx, tenantId, opts...)
}

func (u *Dao[T]) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...*ddd_repository.FindOptions) *ddd_repository.FindPagingResult[T] {
	return u.dao.FindPaging(ctx, query, opts...)
}

type RepositoryOptions struct {
	driver neo4j.Driver
}

func NewRepositoryOptions() *RepositoryOptions {
	return &RepositoryOptions{}
}

func (o *RepositoryOptions) SetDriver(driver neo4j.Driver) *RepositoryOptions {
	o.driver = driver
	return o
}

func (o *RepositoryOptions) Merge(opts ...*RepositoryOptions) *RepositoryOptions {
	if opts != nil {
		for _, item := range opts {
			if item.driver != nil {
				o.driver = item.driver
			}
		}
	}
	if o.driver == nil {
		o.driver = restapp.GetNeo4j()
	}
	return o
}

func NewSession(isWrite bool) ddd_repository.Session {
	return ddd_neo4j.NewSession(isWrite, restapp.GetNeo4j())
}
