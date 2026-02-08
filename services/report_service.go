package services

import (
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) Report() (*repositories.Report, error) {
	return s.repo.ReportSummary()
}
