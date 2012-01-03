/*
Copyright 2011 Paul Ruane.

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

package main

import (
	"os"
	"path/filepath"
)

func databasePath() string {
	path, err := os.Getenverror("TMSU_DB")
	if err == nil { return path }

	//TODO Windows support
	homePath, err := os.Getenverror("HOME")
	if err != nil { panic("Could not determine home directory: environment variable 'HOME' does not exist.") }

	return filepath.Join(homePath, defaultDatabaseName)
}

const defaultDatabaseName = ".tmsu/db"