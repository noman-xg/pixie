#include <absl/strings/str_cat.h>
#include <gtest/gtest.h>

#include "src/common/base/file.h"
#include "src/common/testing/testing.h"

namespace pl {

TEST(FileUtils, WriteThenRead) {
  std::string write_val = R"(This is a a file content.
It has two lines.)";

  char tmp_dir_template[] = "/tmp/utils_test_XXXXXX";
  char* tmp_dir = mkdtemp(tmp_dir_template);
  CHECK(tmp_dir != nullptr);
  std::string test_file = absl::StrCat(tmp_dir, "/file");

  EXPECT_OK(WriteFileFromString(test_file, write_val));
  std::string read_val = FileContentsOrDie(test_file);
  EXPECT_EQ(read_val, write_val);
}

}  // namespace pl
