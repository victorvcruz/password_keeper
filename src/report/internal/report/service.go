package report

type ReportServiceClient interface {
	CreateReport(report ReportRequest) error
	FindReportsByUserId(id string) ([]Report, error)
}

type reportService struct {
	ReportServiceClient
	repository ReportRepositoryClient
}

func NewReportService(_repository ReportRepositoryClient) ReportServiceClient {
	return &reportService{
		repository: _repository,
	}
}

func (r reportService) CreateReport(req ReportRequest) error {

	var report Report
	report.FillFields(req.Action, req.UserId, req.VaultId, req.Description)
	err := r.repository.Create(&report)
	if err != nil {
		return err
	}

	return nil
}

func (r reportService) FindReportsByUserId(id string) ([]Report, error) {

	reports, err := r.repository.FindByUserId(id)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
