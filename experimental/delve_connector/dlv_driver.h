/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#pragma once

#include <string>

#include <jsonrpcpp.hpp>

#include "src/common/base/base.h"

class DelveDriver {
 public:
  DelveDriver() = default;
  ~DelveDriver();

  ::px::Status Connect(std::string host, int port);
  void Init();
  void Close();

  void CreateBreakpoint(std::string_view symbol);
  void Continue();
  Json Eval(std::string_view expr);

  // TODO(oazizi): Make private.
  jsonrpcpp::response_ptr SendRequest(std::string_view method, std::string_view params);

 private:
  int sockfd_ = -1;
};
