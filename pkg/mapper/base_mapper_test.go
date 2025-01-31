/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mapper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/acmestack/gobatis"
	"github.com/acmestack/gobatis/datasource"
	"github.com/acmestack/gobatis/factory"
	_ "github.com/go-sql-driver/mysql"
)

func connect() factory.Factory {
	return gobatis.NewFactory(
		gobatis.SetMaxConn(100),
		gobatis.SetMaxIdleConn(50),
		gobatis.SetDataSource(&datasource.MysqlDataSource{
			Host:     "localhost",
			Port:     3306,
			DBName:   "test",
			Username: "root",
			Password: "123456",
			Charset:  "utf8",
		}))
}

type TestTable struct {
	TableName gobatis.TableName `test_table`
	Id        int64             `column:"id"`
	Username  string            `column:"username"`
	Password  string            `column:"password"`
}

func TestUserMapperImpl_SelectList(t *testing.T) {
	mgr := gobatis.NewSessionManager(connect())
	userMapper := BaseMapper[TestTable]{SessMgr: mgr}
	queryWrapper := &QueryWrapper[TestTable]{}
	queryWrapper.Eq("username", "user123").Select("username")
	list, err := userMapper.SelectList(queryWrapper)
	if err != nil {
		fmt.Println(err.Error())
	}
	marshal, _ := json.Marshal(list)
	fmt.Println(string(marshal))
}

func TestUserMapperImpl_SelectOne(t *testing.T) {
	mgr := gobatis.NewSessionManager(connect())
	userMapper := BaseMapper[TestTable]{SessMgr: mgr}
	queryWrapper := &QueryWrapper[TestTable]{}
	queryWrapper.Eq("username", "user1").Select("username", "password")
	entity, err := userMapper.SelectOne(queryWrapper)
	if err != nil {
		fmt.Println(err.Error())
	}
	marshal, _ := json.Marshal(entity)
	fmt.Println(string(marshal))
}

func TestUserMapperImpl_SelectCount(t *testing.T) {
	mgr := gobatis.NewSessionManager(connect())
	userMapper := BaseMapper[TestTable]{SessMgr: mgr}
	queryWrapper := &QueryWrapper[TestTable]{}
	queryWrapper.Eq("username", "user123")
	count, err := userMapper.SelectCount(queryWrapper)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}

func TestUserMapperImpl_SelectBatchIds(t *testing.T) {
	mgr := gobatis.NewSessionManager(connect())
	userMapper := BaseMapper[TestTable]{SessMgr: mgr}
	var arr []any
	arr = append(arr, 1)
	arr = append(arr, 103)
	list, err := userMapper.SelectBatchIds(arr)
	if err != nil {
		fmt.Println(err.Error())
	}
	marshal, _ := json.Marshal(list)
	fmt.Println(string(marshal))
}

func TestUserMapperImpl_SelectById(t *testing.T) {
	mgr := gobatis.NewSessionManager(connect())
	userMapper := BaseMapper[TestTable]{SessMgr: mgr}
	entity, err := userMapper.SelectById(103)
	if err != nil {
		fmt.Println(err.Error())
	}
	marshal, _ := json.Marshal(entity)
	fmt.Println(string(marshal))
}
