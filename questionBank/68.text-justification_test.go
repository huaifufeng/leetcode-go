package questionBank

import "testing"

func TestFullJustify(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	res := fullJustify(words, maxWidth)
	if len(res) != 3 {
		t.Error("文本左右对齐算法测试1错误", res)
	}

	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth = 16
	res = fullJustify(words, maxWidth)
	if len(res) != 3 {
		t.Error("文本左右对齐算法测试2错误", res)
	}

	words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth = 20
	res = fullJustify(words, maxWidth)
	if len(res) != 6 {
		t.Error("文本左右对齐算法测试3错误", res)
	}

	words = []string{"Imagination", "is", "more", "important", "than", "knowledge."}
	maxWidth = 14
	res = fullJustify(words, maxWidth)
	if len(res) != 4 {
		t.Error("文本左右对齐算法测试4错误", res)
	}

}
