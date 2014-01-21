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
	QueryList(sql string,mapRow MapRow,params ...interface{})([]interface{},error)
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

func (this *DatabaseTemplateImpl) QueryList(sql string,mapRow MapRow,params ...interface{})([]interface{},error){
	result,err:=this.Conn.Query(sql,params...)
	if err!=nil {
		return nil,err
	}
	var resArray []interface{}
	for result.Next(){
		obj,err:=mapRow(result)
		if err!=nil{
			return nil,err
		}
		resArray=append(resArray,obj)
	}
	return resArray,nil
}

func (this *DatabaseTemplateImpl) QueryObject(sql string,mapRow MapRow,params ...interface{})(object interface{},err error){
	result,error:=this.Conn.Query(sql,params...)
	if error!=nil {
		err=error
		return
	}
	if result.Next(){
		object,err=mapRow(result)
	}
	return
}

func (this *DatabaseTemplateImpl) Exec(sql string,params ...interface{})(err error){
	_,error:=this.Conn.Exec(sql,params...)
	if error!=nil {
		err=error
	}
	return
}
