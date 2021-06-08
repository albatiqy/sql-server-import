package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/albatiqy/gopoh"
	"github.com/albatiqy/gopoh/provider/dblib"
	mssqlTool "github.com/albatiqy/gopoh/tool/generator/store/mssql"
	mysqlTool "github.com/albatiqy/gopoh/tool/generator/store/mysql"

	//"github.com/albatiqy/gopoh/provider/dblib"

	"context"
	"sql-server-import/mssql"

	_ "github.com/go-sql-driver/mysql" // harus diimport kalo ga, ga bisa konek
)

func main() {
	//generateMsSql()
	//createMySqlTable()
	importTable()
}

func generateMySql() {
	mysqlTool.Generate("ptk")
}

func generateMsSql() {
	mssqlTool.Generate("dbo", "ptk")
}

func createMySqlTable() {
	fields, err := mssqlTool.GetFields("dbo", "ptk")
	if err != nil {
		log.Fatalf("db error: %s", err)
	}
	mysqlTool.CreateTableFromMssql("ptk", fields, nil)
}

func importTable() {
	srcDbQuery := mssql.GetDbQuery()

	records, err := srcDbQuery.RawSelect(context.Background(), "SELECT TOP 100 * FROM dbo.ptk", gopoh.Data{})
	if err != nil {
		log.Fatalf("db error: %s", err)
	}

	mysqlFieldMap, err := mysqlTool.GetFields("ptk")
	if err != nil {
		log.Fatalf("db error: %s", err)
	}

	arrayLen := len(mysqlFieldMap) // escape string??====================
	fns := make([]func(interface{}, mysqlTool.Field)string, arrayLen)

	if len(records) > 0 {
		record := gopoh.Record{}
		for key, value := range records[0] {
		  record[key] = value
		}
		if data, err := srcDbQuery.Indirect(record); err != nil {
			log.Fatalf("db error: %s", err)
		} else {
			for attr, val := range data {
				if mysqlMap, ok := mysqlFieldMap[attr]; !ok {
					log.Fatal("tujuan kolom mysql: " + attr + " tidak ditemukan")
				} else {
					nullable := mysqlMap.Nullable()
					switch mysqlMap.DataType {
					case "varchar", "char":
						switch val.(type) {
						case dblib.MsSqlUniqueIdentifier:
							fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
								castV := v.(dblib.MsSqlUniqueIdentifier)
								return "'" + castV.String() + "'" // escape ??
							}
						case dblib.MsSqlNullUniqueIdentifier:
							if nullable {
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(dblib.MsSqlNullUniqueIdentifier)
									if !castV.Valid {
										return "NULL"
									} else {
										return "'" + strings.ReplaceAll(castV.UniqueIdentifier.String(), "'", "\\'") + "'"
									}
								}
							} else { // kalo tidka nul errorkan pas scanning
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
							}
						case dblib.NullString:
							if nullable {
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(dblib.NullString)
									if !castV.Valid {
										return "NULL"
									} else {
										return "'" + strings.ReplaceAll(castV.String, "'", "\\'") + "'"
									}
								}
							} else {
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
							}
						case string:
							fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
								castV := v.(string)
								return "'" + strings.ReplaceAll(castV, "'", "\\'") + "'" // escape ??
							}
						default:
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
						}
					case "datetime", "timestamp", "date":
						switch mysqlMap.DataType {
						case "datetime", "timestamp":
							switch val.(type) {
							case dblib.NullTime:
								if nullable {
									fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
										castV := v.(dblib.NullTime)
										if !castV.Valid {
											return "NULL"
										} else {
											return "'" + castV.Time.Format("2006-01-02 15:04:05") + "'"
										}
									}
								} else {
									log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
								}
							case time.Time:
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(time.Time)
									return "'" + castV.Format("2006-01-02 15:04:05") + "'" // escape ??
								}
							default:
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
							}
						case "date":
							switch val.(type) {
							case dblib.NullTime:
								if nullable {
									fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
										castV := v.(dblib.NullTime)
										if !castV.Valid {
											return "NULL"
										} else {
											return "'" + castV.Time.Format("2006-01-02") + "'"
										}
									}
								} else {
									log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
								}
							case time.Time:
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(time.Time)
									return "'" + castV.Format("2006-01-02") + "'" // escape ??
								}
							default:
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
							}
						}
					case "int": // unsigned????
						if mysqlMap.Unsigned {
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tipe unsigned tidak didefinisikan")
						} else {
							switch val.(type) {
							case dblib.NullInt32:
								if nullable {
									fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
										castV := v.(dblib.NullInt32)
										if !castV.Valid {
											return "NULL"
										} else {
											return strconv.FormatInt(int64(castV.Int32), 10)
										}
									}
								} else {
									log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
								}
							case int: // int->dependent machine??
							fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
								castV := v.(int)
								return strconv.FormatInt(int64(castV), 10) // escape ??
							}
							case int32:
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(int32)
									return strconv.FormatInt(int64(castV), 10) // escape ??
								}
							default:
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
							}
						}
					case "tinyint":
						if mysqlMap.Unsigned {
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tipe unsigned tidak didefinisikan")
						} else {
							switch val.(type) {
							case dblib.NullInt32: // hati2 overflow
								if nullable {
									fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
										castV := v.(dblib.NullInt32)
										if !castV.Valid {
											return "NULL"
										} else {
											return strconv.FormatInt(int64(castV.Int32), 10)
										}
									}
								} else {
									log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
								}
							case byte:
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(byte)
									return strconv.FormatInt(int64(castV), 10) // escape ??
								}
							default:
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
							}
						}
					case "smallint":
						if mysqlMap.Unsigned {
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tipe unsigned tidak didefinisikan")
						} else {
							switch val.(type) {
							case dblib.NullInt32:
								if nullable {
									fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
										castV := v.(dblib.NullInt32)
										if !castV.Valid {
											return "NULL"
										} else {
											return strconv.FormatInt(int64(castV.Int32), 10)
										}
									}
								} else {
									log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
								}
							case int16:
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(int16)
									return strconv.FormatInt(int64(castV), 10) // escape ??
								}
							default:
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
							}
						}
					case "bigint":
						if mysqlMap.Unsigned {
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tipe unsigned tidak didefinisikan")
						} else {
							switch castVal := val.(type) {
							case dblib.NullInt64:
								if nullable {
									fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
										castV := v.(dblib.NullInt64)
										if !castV.Valid {
											return "NULL"
										} else {
											return strconv.FormatInt(castVal.Int64, 10)
										}
									}
								} else {
									log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
								}
							case int64:
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(int64)
									return strconv.FormatInt(castV, 10) // escape ??
								}
							default:
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
							}
						}
					case "decimal", "float":
						/*
						prec := strings.Split(mysqlMap.SizeString,",")
						fp, err := strconv.ParseInt(prec[1], 10, 32)
						if err != nil {
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", float/decimal error")
						}
						*/
						if mysqlMap.Unsigned {
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tipe unsigned tidak didefinisikan")
						} else {
							switch val.(type) {
							case dblib.NullFloat64:
								if nullable {
									fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
										castV := v.(dblib.NullFloat64)
										if !castV.Valid {
											return "NULL"
										} else {
											return strconv.FormatFloat(castV.Float64, 'g', 32, 32)
										}
									}
								} else {
									log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
								}
							case float32:
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(float32)
									return strconv.FormatFloat(float64(castV), 'g', 32, 32) // escape ??
								}
							default:
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
							}
						}
					case "double":
						/*
						prec := strings.Split(mysqlMap.SizeString,",")
						fp, err := strconv.ParseInt(prec[1], 10, 32)
						if err != nil {
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", float/decimal error")
						}
						*/
						if mysqlMap.Unsigned {
							log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tipe unsigned tidak didefinisikan")
						} else {
							switch val.(type) {
							case dblib.NullFloat64:
								if nullable {
									fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
										castV := v.(dblib.NullFloat64)
										if !castV.Valid {
											return "NULL"
										} else {
											return strconv.FormatFloat(castV.Float64, 'g', 64, 64)
										}
									}
								} else {
									log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak nullable")
								}
							case float64:
								fns[mysqlMap.Ordinal] = func(v interface{}, f mysqlTool.Field) string {
									castV := v.(float64)
									return strconv.FormatFloat(float64(castV), 'g', 64, 64) // escape ??
								}
							default:
								log.Fatal("kolom mysql: " + attr + ", tipe: " + mysqlMap.DataType + ", tidak cocok dengan: " + reflect.TypeOf(val).String())
							}
						}
					default:
						log.Fatal("kolom mysql tipe: " + mysqlMap.DataType + ", belum didefinisikan")
					}
				}
			}
		}
	}

	var strRows []string
	for _, record := range records {
		if data, err := srcDbQuery.Indirect(record); err != nil {
			log.Fatalf("db error: %s", err)
		} else {
			strCols := make([]string, arrayLen)
			for attr, val := range data {
				mysqlMap := mysqlFieldMap[attr]
				strCols[mysqlMap.Ordinal] = fns[mysqlMap.Ordinal](val, mysqlMap)
			}
			strRows = append(strRows, "("+strings.Join(strCols, ",")+")")
		}
	}

	strQueryPrefix := "INSERT INTO ptk VALUES "

	mysql := dblib.GetMySql()

	for _, row := range strRows {
		log.Print(strQueryPrefix + row)
		if _, err := mysql.Exec(context.Background(), strQueryPrefix + row, nil); err != nil {
			log.Fatalf("db error: %s", err)
		}
	}

}