// Copyright (C) 2017 go-dacc authors
//
// This file is part of the go-dacc library.
//
// the go-dacc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dacc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dacc library.  If not, see <http://www.gnu.org/licenses/>.
//
'use strict';
if (typeof _native_storage_handlers === 'undefined') {
    throw new Error("_native_storage_handlers is undefined.");
}

if (typeof _native_storage_handlers !== 'object') {
    throw new Error("_native_storage_handlers is not an object.");
}

["lcs", "gcs"].forEach(function (val) {
    if (typeof _native_storage_handlers[val] !== 'object') {
        throw new Error("_native_storage_handlers[" + val + "] is not an object.");
    }

    var o = _native_storage_handlers[val];
    if (Object.keys(o).length > 0) {
        throw new Error("_native_storage_handlers[" + val + "] should not have accessible keys.")
    }
});