# Copyright 2018- The Pixie Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

homebrew_package 'autoconf'
homebrew_package 'automake'
homebrew_package 'checkstyle'
homebrew_package 'clang-format'
homebrew_package 'libtool'
homebrew_package 'postgresql@14'
homebrew_package 'pyenv'
homebrew_package 'php'

directory node['graalvm-native-image']['path'] do
  owner user
  group root_group
  mode '0755'
  action :create
end
