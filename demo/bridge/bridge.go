package main

import (
	"fmt"
	"io"
	"math/rand"
)

type IDataExporter interface {
	Fetcher(fetcher IDataFetcher)
	Export(sql string, writer io.Writer) error
}

type IDataFetcher interface {
	Fetch(sql string) []interface{}
}

type MysqlDataFetcher struct {
	Config string
}

func (mf *MysqlDataFetcher) Fetch(sql string) []interface{} {
	fmt.Println("Fetch data from mysql source: " + mf.Config)
	rows := make([]interface{}, 0)
	// 插入两个随机数组成的切片，模拟查询要返回的数据集
	rows = append(rows, rand.Perm(10), rand.Perm(10))
	return rows
}

func NewMysqlDataFetcher(configStr string) IDataFetcher {
	return &MysqlDataFetcher{
		Config: configStr,
	}
}

type OracleDataFetcher struct {
	Config string
}

func (of *OracleDataFetcher) Fetch(sql string) []interface{} {
	fmt.Println("Fetch data from oracle source: " + of.Config)
	rows := make([]interface{}, 0)
	// 插入两个随机数组成的切片，模拟查询要返回的数据集
	rows = append(rows, rand.Perm(10), rand.Perm(10))
	return rows
}

func NewOracleDataFetcher(configStr string) IDataFetcher {
	return &OracleDataFetcher{
		configStr,
	}
}

type CsvExporter struct {
	mFetcher IDataFetcher
}

func NewCsvExporter(fetcher IDataFetcher) IDataExporter {
	return &CsvExporter{
		fetcher,
	}
}

func (ce *CsvExporter) Fetcher(fetcher IDataFetcher) {
	ce.mFetcher = fetcher
}

func (ce *CsvExporter) Export(sql string, writer io.Writer) error {
	rows := ce.mFetcher.Fetch(sql)
	fmt.Printf("CsvExporter.Export, got %v rows\n", len(rows))
	for i, v := range rows {
		fmt.Printf("  行号: %d 值: %s\n", i+1, v)
	}
	return nil
}

type JsonExporter struct {
	mFetcher IDataFetcher
}

func NewJsonExporter(fetcher IDataFetcher) IDataExporter {
	return &JsonExporter{
		fetcher,
	}
}

func (je *JsonExporter) Fetcher(fetcher IDataFetcher) {
	je.mFetcher = fetcher
}

func (je *JsonExporter) Export(sql string, writer io.Writer) error {
	rows := je.mFetcher.Fetch(sql)
	fmt.Printf("JsonExporter.Export, got %v rows\n", len(rows))
	for i, v := range rows {
		fmt.Printf("  行号: %d 值: %s\n", i+1, v)
	}
	return nil
}
