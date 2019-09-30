package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/Masterminds/sprig"
)

func main() {
	const filePath = "options_generated.go"

	data, err := ioutil.ReadFile("options.json")
	if err != nil {
		log.Fatal(err)
	}
	types := make(map[string]string)
	err = json.Unmarshal(data, &types)
	if err != nil {
		log.Fatal(err)
	}

	// Format input tasks.json
	data, err = json.MarshalIndent(types, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("options.json", data, 0664)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, struct {
		Data map[string]string
	}{
		types,
	})
	if err != nil {
		log.Fatal(err)
	}
}

var tmpl = template.Must(template.New("").Funcs(sprig.HermeticTxtFuncMap()).Parse(`// Code generated by go generate; DO NOT EDIT.
package define

// OptionTypeMap is the type map for options.
var OptionTypeMap = map[string]string{
{{- range $k, $v := .Data }}
	"{{ $k }}": "{{ $v }}",
{{- end }}
}

{{- range $k, $v := .Data }}
// With{{ $k | camelcase }} will apply {{ $k }} value to Options
func With{{ $k | camelcase }}(v {{ $v }}) *Option {
	return &Option{
		Key: "{{ $k }}",
		Value: v,
	}
}
{{- end }}
`))
