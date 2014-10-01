/*
Copyright 2011-2014 Paul Ruane.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package entities

import (
	"sort"
)

type ValueId uint

type ValueIds []ValueId

func (valueIds ValueIds) Len() int {
	return len(valueIds)
}

func (valueIds ValueIds) Less(i, j int) bool {
	return valueIds[i] < valueIds[j]
}

func (valueIds ValueIds) Swap(i, j int) {
	valueIds[i], valueIds[j] = valueIds[j], valueIds[i]
}

func (valueIds ValueIds) Uniq() ValueIds {
	if len(valueIds) == 0 {
		return valueIds
	}

	sort.Sort(valueIds)
	uniq := ValueIds{valueIds[0]}
	previous := valueIds[0]

	for _, valueId := range valueIds[1:len(valueIds)] {
		if valueId != previous {
			uniq = append(uniq, valueId)
			previous = valueId
		}
	}

	return uniq
}

type Value struct {
	Id   ValueId
	Name string
}

type Values []*Value

func (values Values) Len() int {
	return len(values)
}

func (values Values) Swap(i, j int) {
	values[i], values[j] = values[j], values[i]
}

func (values Values) Less(i, j int) bool {
	return values[i].Name < values[j].Name
}

func (values Values) Contains(searchValue *Value) bool {
	for _, value := range values {
		if value.Id == searchValue.Id {
			return true
		}
	}

	return false
}

func (values Values) ContainsName(name string) bool {
	for _, value := range values {
		if value.Name == name {
			return true
		}
	}

	return false
}

func (values Values) Any(predicate func(*Value) bool) bool {
	for _, value := range values {
		if predicate(value) {
			return true
		}
	}

	return false
}
