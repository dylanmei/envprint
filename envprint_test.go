package main

import (
	"os"
	"testing"
)

func Test_print(t *testing.T) {
	cases := []struct {
		label    string
		vars     map[string]string
		template string
		document string
		errexit  bool
		errmsg   string
	}{
		{
			label: "matching vars",
			vars: map[string]string{
				"ENVPRINT_WORLD": "palermo",
			},
			template: "hello ${ENVPRINT_WORLD}",
			document: "hello palermo",
		},
		{
			label:    "no matching vars",
			vars:     map[string]string{},
			template: "hello ${ENVPRINT_WORLD}",
			document: "hello ",
		},
		{
			label:    "default vars",
			vars:     map[string]string{},
			template: "hello ${ENVPRINT_WORLD:-catania}",
			document: "hello catania",
		},
		{
			label:    "missing vars",
			vars:     map[string]string{},
			template: "hello ${ENVPRINT_WORLD:?this is not a thing}",
			document: "",
			errexit:  true,
			errmsg:   "this is not a thing",
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			setEnvVars(c.vars)
			defer unsetEnvVars(c.vars)

			document, err := Print([]byte(c.template))

			if c.errexit {
				if err == nil {
					t.Error("Expected an error but didn't get one")
				}

				if err.Error() != c.errmsg {
					t.Errorf("Expected error message to be '%s' but was '%s'", c.errmsg, err.Error())
				}

				return
			}

			if err != nil {
				t.Errorf("Did not expect and error but got %v", err)
				return
			}

			if string(document) != c.document {
				t.Errorf("Expected document to be '%s' but was '%s'", c.document, document)
			}

		})
	}
}

func setEnvVars(vars map[string]string) {
	for name, value := range vars {
		os.Setenv(name, value)
	}
}

func unsetEnvVars(vars map[string]string) {
	for name, _ := range vars {
		os.Unsetenv(name)
	}
}
