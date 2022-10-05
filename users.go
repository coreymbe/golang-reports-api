package main

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func (d *database) UserExists(r_name string, r_password string) (*User, error) {
	rows, err := d.db.Query("SELECT * FROM users WHERE username=$1 AND password=$2", r_name, r_password)
	if err != nil {
		logError(err)
		return nil, err
	}

	var u User
	for rows.Next() {
		err = rows.Scan(&u.Name, &u.Password)
		if err != nil {
			return nil, err
		}
	}

	return &u, nil
}
