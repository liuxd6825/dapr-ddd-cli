package service

import (
	"context"
	user_commands2 "github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserDomainService struct {
}

var validate = validator.New()

func (s *UserDomainService) UserCreate(ctx context.Context, cmd *user_commands2.UserCreateCommand) (*model.UserAggregate, error) {
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}

	user := s.NewAggregate()
	err := ddd.CreateAggregate(ctx, user, cmd)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) UserUpdate(ctx context.Context, cmd *user_commands2.UserUpdateCommand) (*model.UserAggregate, error) {
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}
	user := s.NewAggregate()
	if err := ddd.CommandAggregate(ctx, user, cmd); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
	user := s.NewAggregate()
	_, ok, err := ddd.LoadAggregate(ctx, tenantId, id, user)
	return user, ok, err
}

func (s *UserDomainService) AddressCreate(ctx context.Context, cmd *user_commands2.AddressCreateCommand) (*model.UserAggregate, error) {
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}

	user := s.NewAggregate()
	err := ddd.CommandAggregate(ctx, user, cmd)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) AddressUpdate(ctx context.Context, cmd *user_commands2.AddressUpdateCommand) (*model.UserAggregate, error) {
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}
	user := s.NewAggregate()
	if err := ddd.CommandAggregate(ctx, user, cmd); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) DeleteAddress(ctx context.Context, cmd *user_commands2.AddressDeleteCommand) (*model.UserAggregate, error) {
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}
	user := s.NewAggregate()
	if err := ddd.CommandAggregate(ctx, user, cmd); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) NewAggregate() *model.UserAggregate {
	return model.NewUserAggregate()
}
