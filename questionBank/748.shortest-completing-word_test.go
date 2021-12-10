package questionBank

import "testing"

func TestShortestCompletingWord(t *testing.T) {
	var licensePlate string
	var words []string
	var res string

	licensePlate = "1s3 PSt"
	words = []string{"step", "steps", "stripe", "stepple"}
	res = shortestCompletingWord(licensePlate, words)
	if res != "steps" {
		t.Error("最短补全词算法测试1错误", res)
	}

	licensePlate = "1s3 456"
	words = []string{"looks", "pest", "stew", "show"}
	res = shortestCompletingWord(licensePlate, words)
	if res != "pest" {
		t.Error("最短补全词算法测试2错误")
	}

	licensePlate = "Ah71752"
	words = []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}
	res = shortestCompletingWord(licensePlate, words)
	if res != "husband" {
		t.Error("最短补全词算法测试3错误")
	}

	licensePlate = "OgEu755"
	words = []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}
	res = shortestCompletingWord(licensePlate, words)
	if res != "enough" {
		t.Error("最短补全词算法测试4错误")
	}

	licensePlate = "iMSlpe4"
	words = []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"}
	res = shortestCompletingWord(licensePlate, words)
	if res != "simple" {
		t.Error("最短补全词算法测试4错误")
	}
}
