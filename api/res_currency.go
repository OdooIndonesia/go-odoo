package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResCurrencyService struct {
	client *Client
}

func NewResCurrencyService(c *Client) *ResCurrencyService {
	return &ResCurrencyService{client: c}
}

func (svc *ResCurrencyService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResCurrencyModel, name)
}

func (svc *ResCurrencyService) GetByIds(ids []int) (*types.ResCurrencys, error) {
	r := &types.ResCurrencys{}
	return r, svc.client.getByIds(types.ResCurrencyModel, ids, r)
}

func (svc *ResCurrencyService) GetByName(name string) (*types.ResCurrencys, error) {
	r := &types.ResCurrencys{}
	return r, svc.client.getByName(types.ResCurrencyModel, name, r)
}

func (svc *ResCurrencyService) GetByField(field string, value string) (*types.ResCurrencys, error) {
	r := &types.ResCurrencys{}
	return r, svc.client.getByName(types.ResCurrencyModel, field, value, r)
}

func (svc *ResCurrencyService) GetAll() (*types.ResCurrencys, error) {
	r := &types.ResCurrencys{}
	return r, svc.client.getAll(types.ResCurrencyModel, r)
}

func (svc *ResCurrencyService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResCurrencyModel, fields, relations)
}

func (svc *ResCurrencyService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResCurrencyModel, ids, fields, relations)
}

func (svc *ResCurrencyService) Delete(ids []int) error {
	return svc.client.delete(types.ResCurrencyModel, ids)
}
