package databasetemplate

import (
    "database/sql"
)

type DatabaseTemplateImpl struct{
	Conn *sql.DB
}

type MapRow func(resultSet *sql.Rows)(object interface{},err error)

type DatabaseTemplate interface{
	Query(sql string,mapRow MapRow,params ...interface{})(object interface{},err error)
	Exec(sql string,params ...interface{})(err error)
}

func (this *DatabaseTemplateImpl) Query(sql string,mapRow MapRow,params ...interface{})(object interface{},err error){
	result,error:=this.Conn.Query(sql,params...)
	if error!=nil {
		err=error
		return
	}
	object,err=mapRow(result)
	return
}

func (this *DatabaseTemplateImpl) Exec(sql string,params ...interface{})(err error){
	_,error:=this.Conn.Exec(sql,params...)
	if error!=nil {
		err=error
	}
	return
}
