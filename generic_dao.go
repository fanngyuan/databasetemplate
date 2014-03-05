package databasetemplate

import (
	. "github.com/fanngyuan/mcstorage"
	"reflect"
	"strings"
	"strconv"
	"fmt"
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

func (this GenericDaoImpl) GenerateInStatement(values []interface{})(string,error){
	if values==nil || len(values)==0{
		return "()",nil
	}
	kind:=reflect.TypeOf(values[0]).Kind()
	switch kind {
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64,reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64,reflect.Float32,reflect.Float64:
		return this.intStatement(values)
	case reflect.String:
		return this.stringStatement(values)
	default:
		return "()",nil
	}
}

func (this GenericDaoImpl) toStringArray(values []interface{})[]string{
	length:=len(values)
	result:=make([]string, length)
	for index,value:=range(values){
		kind:=reflect.TypeOf(value).Kind()
		switch kind{
		case reflect.String:
			result[index]=value.(string)
			continue
		case reflect.Int:
			result[index]=strconv.Itoa(value.(int))
			continue
		case reflect.Int64:
			result[index]=strconv.Itoa(int(value.(int64)))
			continue
		default:
			continue
		}
	}
	return result
}

func (this GenericDaoImpl) intStatement(values []interface{})(string,error){
	stringValues:=this.toStringArray(values)
	connectString:=strings.Join(stringValues,",")
	result:=fmt.Sprintf("(%s)",connectString)
	return result,nil
}

func (this GenericDaoImpl) stringStatement(values []interface{})(string,error){
	stringValues:=this.toStringArray(values)
	connectString:=strings.Join(stringValues,"','")
	result:=fmt.Sprintf("('%s')",connectString)
	return result,nil
}
