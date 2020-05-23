package gotemplate

import "testing"

const Logo = `

         ___   ___   ____________    ____    ____   
		/  /  /  /  /___   ____/    /   /   /   /
       /  /  /  /      /   /       /   /   /   /
      /  /__/  /      /   /       /   /   /   /_____
     /________/      /___/       /___/   /_________/
			


Author: solate.
Version: {{ .Version }}
`

type Util struct {
	Version string
}

func TestTemplateParse(t *testing.T) {
	str, err := TemplateParse("logo", Logo, &Util{Version: "0.0.1"})
	if err != nil {
		t.Error(err)
	}
	t.Log(str)

	m := map[string]interface{}{
		"Version": "0.0.2",
	}
	str, err = TemplateParse("logo", Logo, m)
	if err != nil {
		t.Error(err)
	}
	t.Log(str)

}
