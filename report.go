package main

type Report struct {
	ID              int    `json:"id"`
	Certname        string `json:"certname"`
	Environment     string `json:"environment"`
	Status          string `json:"status"`
	Time            string `json:"time"`
	TransactionUUID string `json:"transaction_uuid"`
}

func (d *database) GetAllReports() ([]*Report, error) {
	rows, err := d.db.Query("SELECT * FROM reports")
	if err != nil {
		return nil, err
	}

	var reports []*Report
	for rows.Next() {
		var r Report
		err = rows.Scan(&r.ID, &r.Certname, &r.Environment, &r.Status, &r.Time, &r.TransactionUUID)
		if err != nil {
			return nil, err
		}
		reports = append(reports, &r)
	}
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (d *database) AddReport(certname string, environment string, status string, time string, transaction_uuid string) error {
	_, err := d.db.Exec("INSERT INTO reports (certname, environment, status, time, transaction_uuid) VALUES ($1, $2, $3, $4, $5)", certname, environment, status, time, transaction_uuid)
	if err != nil {
		return err
	}

	return nil
}

func (d *database) GetReport(r_ID int) (*Report, error) {
	rows, err := d.db.Query("SELECT * FROM reports WHERE id=$1", r_ID)
	if err != nil {
		return nil, err
	}

	var r Report
	for rows.Next() {
		err = rows.Scan(&r.ID, &r.Certname, &r.Environment, &r.Status, &r.Time, &r.TransactionUUID)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (d *database) RemoveReport(r_ID int) error {
	_, err := d.db.Exec("DELETE FROM reports WHERE id = $1", r_ID)
	if err != nil {
		return err
	}

	return nil
}
