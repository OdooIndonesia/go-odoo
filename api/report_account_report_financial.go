package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAccountReportFinancialService struct {
	client *Client
}

func NewReportAccountReportFinancialService(c *Client) *ReportAccountReportFinancialService {
	return &ReportAccountReportFinancialService{client: c}
}

func (svc *ReportAccountReportFinancialService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportAccountReportFinancialModel, name)
}

func (svc *ReportAccountReportFinancialService) GetByIds(ids []int) (*types.ReportAccountReportFinancials, error) {
	r := &types.ReportAccountReportFinancials{}
	return r, svc.client.getByIds(types.ReportAccountReportFinancialModel, ids, r)
}

func (svc *ReportAccountReportFinancialService) GetByName(name string) (*types.ReportAccountReportFinancials, error) {
	r := &types.ReportAccountReportFinancials{}
	return r, svc.client.getByName(types.ReportAccountReportFinancialModel, name, r)
}

func (svc *ReportAccountReportFinancialService) GetByField(field string, value string) (*types.ReportAccountReportFinancials, error) {
	r := &types.ReportAccountReportFinancials{}
	return r, svc.client.getByName(types.ReportAccountReportFinancialModel, field, value, r)
}

func (svc *ReportAccountReportFinancialService) GetAll() (*types.ReportAccountReportFinancials, error) {
	r := &types.ReportAccountReportFinancials{}
	return r, svc.client.getAll(types.ReportAccountReportFinancialModel, r)
}

func (svc *ReportAccountReportFinancialService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportAccountReportFinancialModel, fields, relations)
}

func (svc *ReportAccountReportFinancialService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportFinancialModel, ids, fields, relations)
}

func (svc *ReportAccountReportFinancialService) Delete(ids []int) error {
	return svc.client.delete(types.ReportAccountReportFinancialModel, ids)
}
