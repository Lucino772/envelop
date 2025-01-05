package install

import (
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func parseExports(exports map[string]any, data any) map[string]any {
	exp := make(map[string]any, 0)
	for key, value := range exports {
		formattedKey := cases.Title(language.English, cases.Compact).String(strings.ToLower(key))
		if stringVal, ok := value.(string); ok {
			templ, err := template.New(key).Parse(stringVal)
			if err != nil {
				continue
			}
			var buf strings.Builder
			if err := templ.Execute(&buf, data); err != nil {
				continue
			}
			exp[formattedKey] = buf.String()
		} else {
			exp[formattedKey] = value
		}
	}
	return exp
}
