package util

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	})
	return log
}
