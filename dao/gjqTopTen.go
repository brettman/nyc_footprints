package dao

import "fmt"

// TopTenNames - get the top ten names as a test
func (r *GeoJSONRepo) TopTenNames() {
	db := r.dbctx.Connect()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	rows, err := db.Query("select bin, name, base_bbl, lstmoddate from geojson where name is not null limit $1;", 10)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var bin string
		var name string
		var baseBbl string
		var lstmoddate string
		err = rows.Scan(&bin, &name, &baseBbl, &lstmoddate)
		if err != nil {
			// panic(err)
			fmt.Println(err.Error())
		}
		fmt.Println(bin, name, baseBbl, lstmoddate)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
