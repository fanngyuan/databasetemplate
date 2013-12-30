package databasetemplate

import (
	. "github.com/fanngyuan/mcstorage"
)

type GenericDao interface{
	Storage
	GetAll()(object interface{},err error)
	CreateTable()
	TruncateTable()
}

type GenericDaoImpl struct{
	DatabaseTemplate DatabaseTemplate
}
