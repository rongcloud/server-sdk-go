package sdk

import (
	"os"
	"testing"
)

func TestRongCloud_SensitiveAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)
	err := rc.SensitiveAdd(
		"7Szq13MKRVortoknTAk7W8",
		"7Szq13MKRVortoknTAk7W8",
		1,
	)
	t.Log(err)
}

func TestRongCloud_SensitiveBatchAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)
	ret, err := rc.SensitiveBatchAdd(
		[]SensitiveWord{
			{
				Type:        "0",
				Word:        "replaceStr1",
				ReplaceWord: "replace1",
			},
			{
				Type:        "1",
				Word:        "blockStr1",
				ReplaceWord: "",
			},
		})
	t.Log(err)
	t.Log(ret)
}

func TestRongCloud_SensitiveGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)
	rep, err := rc.SensitiveGetList()
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_SensitiveRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
	)
	err := rc.SensitiveRemove(
		[]string{"7Szq13MKRVortoknTAk7W8"},
	)
	t.Log(err)
}
