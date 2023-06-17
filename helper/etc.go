package helper

import "strings"

func SpaceRemover(s string) string {
	nilaiWithSpace := strings.ReplaceAll(s, " ", "")
	nilai := strings.ReplaceAll(nilaiWithSpace, "\n", "")
	nil := strings.ReplaceAll(nilai, ".", "")
	return nil
}
