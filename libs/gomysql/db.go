package gomysql

import (
	"database/sql"
	"reflect"
	"github.com/yangjinguang/wechat-server/libs/logger"
	"strconv"
)

type DB struct {
	*sql.DB
	Table       string
	QueryMethod string
	Fields      []*Field
	Values      []interface{}
	QueryWhere  string
	QueryLimit  []int
	QuerySql    string
}

type Field struct {
	Index int
	Name  string
}

var utils Utils

func (d *DB) T(tableName string) *DB {
	d.Table = tableName
	return d
}

func (d *DB) Select(fields ...string) *DB {
	d.QueryMethod = "select"
	var fds []*Field
	for i, field := range fields {
		fds = append(fds, &Field{Index: i, Name: field})
	}
	d.Fields = fds
	return d
}

func (d *DB) SelectById(id int64) *DB {
	return d.Select().Where("`id` = '" + strconv.FormatInt(id, 10) + "'")
}

func (d *DB) Where(sqlStr string) *DB {
	d.QueryWhere = sqlStr
	return d
}

func (d *DB) Limit(offset int, size int) *DB {
	d.QueryLimit = append(d.QueryLimit, offset, size)
	return d
}

func (d *DB) Insert(result interface{}) (int64, error) {
	resultv := reflect.ValueOf(result)
	d.QueryMethod = "insert"
	utils.FieldParse(d, resultv.Type().Elem())
	utils.ValueParse(d, resultv)
	err := utils.SqlBuild(d)
	if err != nil {
		return 0, err
	}
	logger.Debug(d.QuerySql)
	res, err := d.DB.Exec(d.QuerySql, d.Values...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (d *DB) Replace(result interface{}) error {
	resultv := reflect.ValueOf(result)
	d.QueryMethod = "replace"
	utils.FieldParse(d, resultv.Type().Elem())
	utils.ValueParse(d, resultv)
	err := utils.SqlBuild(d)
	if err != nil {
		return err
	}
	logger.Debug(d.QuerySql)
	_, err = d.DB.Exec(d.QuerySql, d.Values...)
	return err
}

func (d *DB) All(result interface{}) error {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("result argument must be a slice address")
	}

	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem().Elem()
	utils.FieldParse(d, elemt)
	err := utils.SqlBuild(d)
	if err != nil {
		return err
	}
	rows, err := d.DB.Query(d.QuerySql)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		elemp := reflect.New(elemt)
		var vals []interface{}
		for _, f := range d.Fields {
			vals = append(vals, elemp.Elem().Field(f.Index).Addr().Interface())
		}
		err := rows.Scan(vals...)
		if err != nil {
			return err
		}
		slicev = reflect.Append(slicev, elemp.Elem().Addr())
	}
	resultv.Elem().Set(slicev)
	return nil
}

func (d *DB) One(result interface{}) (notFound bool, err error) {
	resultv := reflect.ValueOf(result)
	elemt := resultv.Type().Elem()
	utils.FieldParse(d, elemt)
	err = utils.SqlBuild(d)
	if err != nil {
		return false, err
	}
	row := d.DB.QueryRow(d.QuerySql)
	elemp := reflect.New(elemt)
	var vals []interface{}
	for _, f := range d.Fields {
		vals = append(vals, elemp.Elem().Field(f.Index).Addr().Interface())
	}
	err = row.Scan(vals...)
	if err != nil {
		return true, nil
	}
	resultv.Elem().Set(elemp.Elem())
	return false, nil
}

func (d *DB) DeleteById(id int64) error {
	return d.Where("`id` = '" + strconv.FormatInt(id, 10) + "'").Delete()
}

func (d *DB) Delete() error {
	d.QueryMethod = "delete"
	err := utils.SqlBuild(d)
	if err != nil {
		return err
	}
	_, err = d.DB.Exec(d.QuerySql)
	return err
}
