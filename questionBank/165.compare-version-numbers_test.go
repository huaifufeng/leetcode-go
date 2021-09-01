package questionBank

import "testing"

func TestCompareVersion(t *testing.T) {
	version1 := "1.01"
	version2 := "1.001"
	res := compareVersion(version1, version2)

	if res != 0 {
		t.Error("比较版本号算法测试1错误")
	}

	version1 = "1.0"
	version2 = "1.0.0"
	res = compareVersion(version1, version2)

	if res != 0 {
		t.Error("比较版本号算法测试2错误")
	}

	version1 = "0.1"
	version2 = "1.1"
	res = compareVersion(version1, version2)

	if res != -1 {
		t.Error("比较版本号算法测试3错误")
	}

	version1 = "1.0.1"
	version2 = "1"
	res = compareVersion(version1, version2)

	if res != 1 {
		t.Error("比较版本号算法测试4错误")
	}

	version1 = "7.5.2.4"
	version2 = "7.5.3"
	res = compareVersion(version1, version2)

	if res != -1 {
		t.Error("比较版本号算法测试5错误")
	}
}
