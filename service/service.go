package service

import "example/data-access/repo"

func GetAccount(accountId string) *repo.Account {
	return repo.GetAccount(accountId)
}
