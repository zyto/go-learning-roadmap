package words

import "strings"

func Normalize(word string) string {
	return strings.ToLower(clean(word))
}

func clean(word string) string {
	return strings.TrimSpace(word)
}

//думаю clean не стоит экспортировать потому что эта функция уже используется в более общей функции Normalize и отдельно её использовать за пределами текущего пакета не предполагается
