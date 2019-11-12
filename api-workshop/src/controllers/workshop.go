package controllers

import (
	"fmt"

	"github.com/igorgabriel/api-workshop/src/helpers"
	"github.com/igorgabriel/api-workshop/src/models"
	"github.com/sirupsen/logrus"
)

// GetWorkshops will return all workshops.
func GetWorkshops() ([]models.Workshop, error) {
	db, err := helpers.DBConn()
	defer db.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	rs, err := db.Query(`SELECT * FROM workshop`)
	defer rs.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	var w models.Workshop
	var ws []models.Workshop
	for rs.Next() {
		err = rs.Scan(&w.ID, &w.Nm, &w.Pl)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}
		ws = append(ws, w)
	}

	return ws, nil

}

// GetWorkshopByID will retrieve workshop ID and returns the workshop.
func GetWorkshopByID(id int) (*models.Workshop, error) {
	db, err := helpers.DBConn()
	defer db.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	rs, err := db.Query("SELECT * FROM workshop WHERE id = ?", id)
	if err != nil {
		logrus.Errorln("erro: ", err)
		return nil, err
	}

	w := models.Workshop{}
	for rs.Next() {
		err = rs.Scan(&w.ID, &w.Nm, &w.Pl)
		if err != nil {
			logrus.Errorln("erro: ", err)
			panic(err.Error())
		}
	}

	return &w, nil
}

// SaveWorkshop will retrieve a workshop and save them
func SaveWorkshop(l models.Workshop) error {
	db, err := helpers.DBConn()
	defer db.Close()
	sql := fmt.Sprintf(`INSERT INTO workshop(nome, palestrante) 
						VALUES(?,?)`)
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	_, err = stmt.Exec(l.Nm, l.Pl)
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	return nil
}

// DeleteWorkshop will retrieve a workshop id and delete them
func DeleteWorkshop(id int) error {
	db, err := helpers.DBConn()
	defer db.Close()

	i := fmt.Sprintf(`DELETE FROM workshop WHERE id=?`)

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

// UpdateWorkshop will retrieve workshop and update them
func UpdateWorkshop(w models.Workshop) error {
	db, err := helpers.DBConn()
	defer db.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	sql := fmt.Sprintf(`UPDATE workshop SET nome = ?, palestrante = ? WHERE id = ?`)

	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	_, err = stmt.Exec(w.Nm, w.Pl, w.ID)
	if err != nil {
		logrus.Errorln("erro: ", err)
		return err
	}

	return nil
}
