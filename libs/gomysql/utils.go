package gomysql

import (
	"reflect"
	"strings"
	"github.com/yangjinguang/wechat-server/libs/logger"
	"errors"
	"strconv"
)

type Utils struct {
}

func (u *Utils) FieldParse(d *DB, p reflect.Type) {
	var pFds []*Field
	for i := 0; i < p.NumField(); i++ {
		pField := p.Field(i)
		fName := pField.Tag.Get("mysql")
		if fName == "" {
			fName = pField.Name
		}
		for _, f := range d.Fields {
			if f.Name == fName {
				f.Index = i
			}
		}
		pFds = append(pFds, &Field{Index: i, Name: fName})
	}
	if len(d.Fields) <= 0 {
		d.Fields = pFds
	}
}

func (u *Utils) ValueParse(d *DB, value reflect.Value) {
	elemv := value.Elem()
	var sArr []interface{}
	for i := 0; i < elemv.NumField(); i++ {
		if elemv.Field(i).Kind() == reflect.Struct {
			if elemv.Field(i).Type().Name() == "Time" {
				logger.Debug(elemv.Field)
			}
		}
		sArr = append(sArr, elemv.Field(i).Interface())
	}
	d.Values = sArr
}

func (u *Utils) SqlBuild(d *DB) error {
	var fdsArr []string
	for _, f := range d.Fields {
		fdsArr = append(fdsArr, "`"+f.Name+"`")
	}
	fds := strings.Join(fdsArr, ",")
	switch d.QueryMethod {
	case "select":
		if fds == "" {
			fds = "*"
		}
		d.QuerySql = "select " + fds + " from `" + d.Table + "`"
		if d.QueryWhere != "" {
			d.QuerySql += " where " + d.QueryWhere
		}
		if len(d.QueryLimit) > 0 {
			d.QuerySql += " limit " + strconv.Itoa(d.QueryLimit[0]) + "," + strconv.Itoa(d.QueryLimit[1])
		}
		break
	case "insert":
		var qt []string
		for i := 0; i < len(d.Fields); i++ {
			qt = append(qt, "?")
		}
		logger.Debug("qt", qt)
		d.QuerySql = "insert into `" + d.Table + "` ( " + fds + " ) values ( " + strings.Join(qt, ",") + " )"
		break
	case "replace":
		var qt []string
		for i := 0; i < len(d.Fields); i++ {
			qt = append(qt, "?")
		}
		logger.Debug("qt", qt)
		d.QuerySql = "replace into `" + d.Table + "` ( " + fds + " ) values ( " + strings.Join(qt, ",") + " )"
		break
	case "deleteById":
		d.QuerySql = "delete from `" + d.Table + "` where id = ?"
		break
	case "delete":
		if d.QueryWhere == "" {
			return errors.New("must has where filter")
		}
		d.QuerySql = "delete from `" + d.Table + "` where " + d.QueryWhere
		break
	}
	logger.Debug(d.QuerySql)
	return nil
}
