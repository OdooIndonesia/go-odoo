package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountBalanceReportService struct {
	client *Client
}

func NewAccountBalanceReportService(c *Client) *AccountBalanceReportService {
	return &AccountBalanceReportService{client: c}
}

func (svc *AccountBalanceReportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountBalanceReportModel, name)
}

func (svc *AccountBalanceReportService) GetByIds(ids []int) (*types.AccountBalanceReports, error) {
	a := &types.AccountBalanceReports{}
	return a, svc.client.getByIds(types.AccountBalanceReportModel, ids, a)
}

func (svc *AccountBalanceReportService) GetByName(name string) (*types.AccountBalanceReports, error) {
	a := &types.AccountBalanceReports{}
	return a, svc.client.getByName(types.AccountBalanceReportModel, name, a)
}

func (svc *AccountBalanceReportService) GetByField(field string, value string) (*types.AccountBalanceReports, error) {
	a := &types.AccountBalanceReports{}
	return a, svc.client.getByName(types.AccountBalanceReportModel, field, value, a)
}

func (svc *AccountBalanceReportService) GetAll() (*types.AccountBalanceReports, error) {
	a := &types.AccountBalanceReports{}
	return a, svc.client.getAll(types.AccountBalanceReportModel, a)
}

func (svc *AccountBalanceReportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountBalanceReportModel, fields, relations)
}

func (svc *AccountBalanceReportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBalanceReportModel, ids, fields, relations)
}

func (svc *AccountBalanceReportService) Delete(ids []int) error {
	return svc.client.delete(types.AccountBalanceReportModel, ids)
}
