package databasetemplate

import (
	. "github.com/fanngyuan/mcstorage"
)

type GenericDao interface{
	Storage
	GetAll()(object interface{},err error)
	createTable()
	truncateTable()
}

type GenericDaoImpl struct{
	DatabaseTemplate DatabaseTemplate
}
