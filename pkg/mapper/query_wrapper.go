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
	"github.com/acmestack/gobatis-plus/pkg/constants"
	"github.com/acmestack/gobatis/builder"
)

type QueryWrapper[T any] struct {
	Columns           []string
	SqlBuild          *builder.SQLFragment
	Expression        []any
	LastConditionType string
}

func (queryWrapper *QueryWrapper[T]) Eq(column string, val any) Wrapper[T] {
	queryWrapper.setCondition(column, val, constants.Eq)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) Ne(column string, val any) Wrapper[T] {
	queryWrapper.setCondition(column, val, constants.Ne)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) Gt(column string, val any) Wrapper[T] {
	queryWrapper.setCondition(column, val, constants.Gt)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) Ge(column string, val any) Wrapper[T] {
	queryWrapper.setCondition(column, val, constants.Ge)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) Lt(column string, val any) Wrapper[T] {
	queryWrapper.setCondition(column, val, constants.Lt)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) Le(column string, val any) Wrapper[T] {
	queryWrapper.setCondition(column, val, constants.Le)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) Like(column string, val any) Wrapper[T] {
	s := val.(string)
	queryWrapper.setCondition(column, "%"+s+"%", constants.Like)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) NotLike(column string, val any) Wrapper[T] {
	s := val.(string)
	queryWrapper.setCondition(column, "%"+s+"%", constants.Not+constants.Like)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) LikeLeft(column string, val any) Wrapper[T] {
	s := val.(string)
	queryWrapper.setCondition(column, "%"+s, constants.Like)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) LikeRight(column string, val any) Wrapper[T] {
	s := val.(string)
	queryWrapper.setCondition(column, s+"%", constants.Like)
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) And() Wrapper[T] {
	queryWrapper.Expression = append(queryWrapper.Expression, constants.Eq)
	queryWrapper.LastConditionType = constants.Eq
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) Or() Wrapper[T] {
	queryWrapper.Expression = append(queryWrapper.Expression, constants.Or)
	queryWrapper.LastConditionType = constants.Or
	return queryWrapper
}

func (queryWrapper *QueryWrapper[T]) Select(columns ...string) Wrapper[T] {
	queryWrapper.Columns = append(queryWrapper.Columns, columns...)
	return queryWrapper
}

type ParamValue struct {
	value any
}

func (queryWrapper *QueryWrapper[T]) setCondition(column string, val any, conditionType string) {

	if queryWrapper.LastConditionType != constants.And && queryWrapper.LastConditionType != constants.Or && len(queryWrapper.Expression) > 0 {
		queryWrapper.Expression = append(queryWrapper.Expression, constants.And)
	}

	queryWrapper.Expression = append(queryWrapper.Expression, column)

	queryWrapper.Expression = append(queryWrapper.Expression, conditionType)

	queryWrapper.Expression = append(queryWrapper.Expression, ParamValue{val})
}
