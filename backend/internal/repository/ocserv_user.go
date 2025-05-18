package repository

import (
	"context"
	"ocserv/pkg/database"
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
	"slices"
)

type OcservUserRepository struct {
	ocUserRepo  oc.OcservUserServiceInterface
	ocOcctlRepo oc.OcctlServiceInterface
}

type OcservUserRepositoryInterface interface {
	GetUsersWithOnlineAttr(ctx context.Context, pagination *request.Pagination) (*[]oc.OcservUser, int64, error)
}

func NewOcservUserRepository() *OcservUserRepository {
	return &OcservUserRepository{
		ocUserRepo:  oc.NewOcservUserService(database.Get()),
		ocOcctlRepo: oc.NewOcctlService(),
	}
}

func (o *OcservUserRepository) GetUsersWithOnlineAttr(ctx context.Context, pagination *request.Pagination) (*[]oc.OcservUser, int64, error) {
	type result struct {
		online []string
		err    error
	}

	ch := make(chan result, 1)

	go func() {
		online, err := o.ocOcctlRepo.WithContext(ctx).GetOnlineUsers()
		ch <- result{online: online, err: err}
	}()

	users, count, err := o.ocUserRepo.GetUsers(ctx, pagination)
	if err != nil {
		return nil, 0, err
	}

	if count > 0 {
		r := <-ch
		if r.err == nil {
			for i, user := range *users {
				if slices.Contains(r.online, user.Username) {
					(*users)[i].IsOnline = true
				}
			}
		}
	}

	return nil, 0, nil
}
