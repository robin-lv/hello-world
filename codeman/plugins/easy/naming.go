package easy

import (
	"fmt"
	"github.com/serenize/snaker"
	"strings"
)

func snakeCase(ss ...string) string {
	return snake(ss)
}
func SnakeCase(ss ...string) string {
	return strings.ToUpper(snake(ss))
}
func camelCase(ss ...string) string {
	return snaker.SnakeToCamelLower(snake(ss))
}
func CamelCase(ss ...string) string {
	return snaker.SnakeToCamel(snake(ss))
}

func Ident(values ...any) string {
	ss := make([]string, len(values))
	for i, v := range values {
		ss[i] = strings.TrimSpace(fmt.Sprint(v))
	}
	return snakeCase(ss...)
}

func snake(ss []string) string {
	j := 0
	for i := range ss {
		ss[i] = strings.TrimSpace(ss[i])
		if ss[i] != "" {
			ss[j] = snaker.CamelToSnake(ss[i])
			j++
		}
	}
	return strings.Join(ss[:j], "_")
}
