package helpers

import (
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// CountDaysOfLavoura returns quantity of days since Lavoura started
func CountDaysOfLavoura(startTime time.Time) int {
	now := time.Now()
	days := now.Sub(startTime).Hours() / 24
	return int(days)
}

// AddDaysToData returns time.Time added by days
func AddDaysToData(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

// HashAndSaltPassword generates Hash password with salt
func HashAndSaltPassword(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorln(err)
	}
	return string(hash), nil
}

//ComparePasswords compares a bcrypt hashed password with its possible plaintext equivalent.
func ComparePasswords(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		logrus.Errorln(err)
		return false
	}
	return true
}
