package controllers

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.sandmanbb.com/pivotal/agrows-client/src/helpers"
	"gitlab.sandmanbb.com/pivotal/agrows-client/src/models"
)

// GetLavouras will return all active lavouras.
func GetLavouras(pID int) ([]models.Lavoura, error) {
	db, err := helpers.DBConn()
	defer db.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	rs, err := db.Query(`SELECT l.* FROM lavoura l INNER JOIN area a ON (l.area_id = a.id) WHERE a.produtor_id = ?
						 ORDER BY l.dt_plantio DESC`, pID)
	defer rs.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	var l models.Lavoura
	var ls []models.Lavoura
	for rs.Next() {
		var aID, cID int
		var cAt string
		err = rs.Scan(&l.ID, &aID, &cAt, &l.DPlantio, &l.DColheita, &cID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}

		c, err := GetCultivarByID(cID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}
		l.Cultivar = *c

		a, err := GetAreaByID(aID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}
		l.Area = *a

		ls = append(ls, l)
	}

	return ls, nil

}

// GetActiveLavouras will return all active lavouras.
func GetActiveLavouras(pID int) ([]models.Lavoura, error) {
	db, err := helpers.DBConn()
	defer db.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	rs, err := db.Query(`SELECT l.* FROM lavoura l INNER JOIN area a ON (l.area_id = a.id) WHERE a.produtor_id = ?
						 AND ? BETWEEN dt_plantio AND dt_colheita`, pID, time.Now())
	defer rs.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	var l models.Lavoura
	var ls []models.Lavoura
	for rs.Next() {
		var aID, cID int
		var cAt string
		err = rs.Scan(&l.ID, &aID, &cAt, &l.DPlantio, &l.DColheita, &cID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}

		c, err := GetCultivarByID(cID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}
		l.Cultivar = *c

		a, err := GetAreaByID(aID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}
		l.Area = *a

		ls = append(ls, l)
	}

	return ls, nil

}

// GetLavouraByID will retrieve lavoura ID and returns the lavoura.
func GetLavouraByID(id int) (*models.Lavoura, error) {
	db, err := helpers.DBConn()
	defer db.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	rs, err := db.Query("SELECT * FROM lavoura WHERE id = ?", id)
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	l := models.Lavoura{}
	for rs.Next() {
		var aID, cID int
		var ca string
		err = rs.Scan(&l.ID, &aID, &ca, &l.DPlantio, &l.DColheita, &cID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}

		a, err := GetAreaByID(aID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			return nil, err
		}
		l.Area = *a

		c, err := GetCultivarByID(cID)
		if err != nil {
			logrus.Errorln("erro: ", err)
			return nil, err
		}
		l.Cultivar = *c
	}

	return &l, nil
}

// SaveLavoura will retrieve a lavoura and save them
func SaveLavoura(l models.Lavoura) error {
	db, err := helpers.DBConn()
	defer db.Close()
	sql := fmt.Sprintf(`INSERT INTO lavoura(created_at,area_id,cultivar_id,dt_plantio,dt_colheita) 
						VALUES(?,?,?,?,?)`)
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	_, err = stmt.Exec(time.Now(), l.Area.ID, l.Cultivar.ID, l.DPlantio, l.DColheita)
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	return nil
}

// DeleteLavoura will retrieve a lavoura id and delete them
func DeleteLavoura(id int) error {
	db, err := helpers.DBConn()
	defer db.Close()

	i := fmt.Sprintf(`DELETE FROM lavoura WHERE id=?`)

	stmt, err := db.Prepare(i)
	defer stmt.Close()

	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	return nil
}

// UpdateLavoura will retrieve lavoura and update them
func UpdateLavoura(l models.Lavoura) error {
	db, err := helpers.DBConn()
	defer db.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	sql := fmt.Sprintf(`UPDATE lavoura SET area_id = ?, cultivar_id = ?, dt_plantio = ?,
						dt_colheita = ? WHERE id = ?`)

	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	_, err = stmt.Exec(l.Area.ID, l.Cultivar.ID, l.DPlantio, l.DColheita, l.ID)
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	return nil
}
