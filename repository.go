package main

import (
	"context"
	"db/ent"
	"db/ent/tourproduct"
	"db/ent/user"
)

type UserRepository struct {
	Client *ent.UserClient
}

type TourProductRepository struct {
	Client *ent.TourProductClient
}

func (repo *UserRepository) FindAll() []*ent.User {
	result := repo.Client.Query().WithProducts().AllX(context.TODO())
	return result
}

func (repo *TourProductRepository) FindAll() []*ent.TourProduct {
	result := repo.Client.Query().WithManager().AllX(context.TODO())
	return result
}

func (repo *TourProductRepository) FindAllByManager(managerID string) []*ent.TourProduct {
	result := repo.Client.Query().Where(tourproduct.HasManagerWith(user.ID(managerID))).WithManager().AllX(context.TODO())
	return result
}
