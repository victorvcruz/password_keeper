package report

type ServiceClient interface {
	CreateReport(report Request) error
	FindReportsByUserId(id string) ([]Report, error)
}

type service struct {
	ServiceClient
	repository RepositoryClient
}

func NewReportService(_repository RepositoryClient) ServiceClient {
	return &service{
		repository: _repository,
	}
}

func (r *service) CreateReport(req Request) error {

	var report Report
	report.FillFields(req.Action, req.UserId, req.VaultId, req.Description)
	err := r.repository.Create(&report)
	if err != nil {
		return err
	}

	return nil
}

func (r *service) FindReportsByUserId(id string) ([]Report, error) {

	reports, err := r.repository.FindByUserId(id)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
