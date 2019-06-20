package service

import (
	db2 "github.com/crueltycute/tech-db-forum/internal/app/db"
	"github.com/crueltycute/tech-db-forum/internal/models"
	"net/http"
)

func Clear(res http.ResponseWriter, req *http.Request) {
	db := db2.Connection

	db.QueryRow(queryClearDB)
	//return operations.NewClearOK()

	// по логике не подходит, но мне влом делать новый тип респонсов
	models.ErrResponse(res, http.StatusOK, "ok")
	return
}

func Status(res http.ResponseWriter, req *http.Request) {
	db := db2.Connection

	status := &models.Status{}
	err := db.QueryRow(queryGetStatus).Scan(&status.User, &status.Forum, &status.Thread, &status.Post)
	if err != nil {
		panic(err)
	}
	//return operations.NewStatusOK().WithPayload(status)
	models.ResponseObject(res, http.StatusOK, status)
	return
}
