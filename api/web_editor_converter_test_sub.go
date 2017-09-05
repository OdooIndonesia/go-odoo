package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WebEditorConverterTestSubService struct {
	client *Client
}

func NewWebEditorConverterTestSubService(c *Client) *WebEditorConverterTestSubService {
	return &WebEditorConverterTestSubService{client: c}
}

func (svc *WebEditorConverterTestSubService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WebEditorConverterTestSubModel, name)
}

func (svc *WebEditorConverterTestSubService) GetByIds(ids []int) (*types.WebEditorConverterTestSubs, error) {
	w := &types.WebEditorConverterTestSubs{}
	return w, svc.client.getByIds(types.WebEditorConverterTestSubModel, ids, w)
}

func (svc *WebEditorConverterTestSubService) GetByName(name string) (*types.WebEditorConverterTestSubs, error) {
	w := &types.WebEditorConverterTestSubs{}
	return w, svc.client.getByName(types.WebEditorConverterTestSubModel, name, w)
}

func (svc *WebEditorConverterTestSubService) GetByField(field string, value string) (*types.WebEditorConverterTestSubs, error) {
	w := &types.WebEditorConverterTestSubs{}
	return w, svc.client.getByName(types.WebEditorConverterTestSubModel, field, value, w)
}

func (svc *WebEditorConverterTestSubService) GetAll() (*types.WebEditorConverterTestSubs, error) {
	w := &types.WebEditorConverterTestSubs{}
	return w, svc.client.getAll(types.WebEditorConverterTestSubModel, w)
}

func (svc *WebEditorConverterTestSubService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WebEditorConverterTestSubModel, fields, relations)
}

func (svc *WebEditorConverterTestSubService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WebEditorConverterTestSubModel, ids, fields, relations)
}

func (svc *WebEditorConverterTestSubService) Delete(ids []int) error {
	return svc.client.delete(types.WebEditorConverterTestSubModel, ids)
}
