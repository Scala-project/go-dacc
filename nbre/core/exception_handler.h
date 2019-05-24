// Copyright (C) 2018 go-dacc authors
//
// This file is part of the go-dacc library.
//
// the go-dacc library is free software: you can redistribute it and/or
// modify
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
// along with the go-dacc library.  If not, see
// <http://www.gnu.org/licenses/>.
//
#pragma once
#include "common/common.h"
#include <thread>

namespace neb {
namespace core {

class exception_handler {
public:
  exception_handler() = default;
  ~exception_handler();

  void run();

  void kill();

protected:
  std::unique_ptr<std::thread> m_thread;
};
} // namespace core
} // namespace neb
