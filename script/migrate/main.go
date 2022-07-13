package main

import (
	"github.com/sirupsen/logrus"

	"github.com/jiazhen-lin/indexer/application/common/db"
	"github.com/jiazhen-lin/indexer/application/common/model"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	// init db
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	// migrate up table
	db.AutoMigrate(&model.Block{})
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.TransactionLog{})

	logrus.Info("migrate up!!")
}
