package repositories

import "database/sql"



type ReportRepository struct{
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository{
	return &ReportRepository{db: db}
}

type ProdukTerlaris struct {
	Nama        string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type Report struct{
	TotalRevenue string `json:"total_revenue"`
	TotalTransaksi string `json:"total_transaksi"`
	ProdukTerlaris ProdukTerlaris `json:"produk_terlaris"`
}

func (repo *ReportRepository) ReportSummary() (*Report, error) {
	var report Report

	err := repo.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0)
		FROM transactions
	`).Scan(&report.TotalRevenue)
	if err != nil {
		return nil, err
	}

	err = repo.db.QueryRow(`
		SELECT COUNT(*)
		FROM transactions
	`).Scan(&report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	err = repo.db.QueryRow(`
		SELECT
			p.name,
			SUM(td.quantity) AS qty_terjual
		FROM transaction_details td
		JOIN product p ON p.id = td.product_id
		GROUP BY p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`).Scan(
		&report.ProdukTerlaris.Nama,
		&report.ProdukTerlaris.QtyTerjual,
	)

	if err == sql.ErrNoRows {
		report.ProdukTerlaris = ProdukTerlaris{}
		return &report, nil
	}
	if err != nil {
		return nil, err
	}

	return &report, nil
}

