package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type LinkTrackerClickService struct {
	client *Client
}

func NewLinkTrackerClickService(c *Client) *LinkTrackerClickService {
	return &LinkTrackerClickService{client: c}
}

func (svc *LinkTrackerClickService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.LinkTrackerClickModel, name)
}

func (svc *LinkTrackerClickService) GetByIds(ids []int) (*types.LinkTrackerClicks, error) {
	l := &types.LinkTrackerClicks{}
	return l, svc.client.getByIds(types.LinkTrackerClickModel, ids, l)
}

func (svc *LinkTrackerClickService) GetByName(name string) (*types.LinkTrackerClicks, error) {
	l := &types.LinkTrackerClicks{}
	return l, svc.client.getByName(types.LinkTrackerClickModel, name, l)
}

func (svc *LinkTrackerClickService) GetByField(field string, value string) (*types.LinkTrackerClicks, error) {
	l := &types.LinkTrackerClicks{}
	return l, svc.client.getByName(types.LinkTrackerClickModel, field, value, l)
}

func (svc *LinkTrackerClickService) GetAll() (*types.LinkTrackerClicks, error) {
	l := &types.LinkTrackerClicks{}
	return l, svc.client.getAll(types.LinkTrackerClickModel, l)
}

func (svc *LinkTrackerClickService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.LinkTrackerClickModel, fields, relations)
}

func (svc *LinkTrackerClickService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.LinkTrackerClickModel, ids, fields, relations)
}

func (svc *LinkTrackerClickService) Delete(ids []int) error {
	return svc.client.delete(types.LinkTrackerClickModel, ids)
}
