package mime

import (
	"fmt"
	"regexp"

	stringd "github.com/kimbeejay/dry/string"
)

const regex = `^(?i)([a-z]{3,})\/([0-9a-z\-\+\.]{3,})(?:;(\w+)\=([a-z0-9\_\-]+))?$`

// Look more https://www.iana.org/assignments/media-types/media-types.xhtml
type Type struct {
	registry   Registry
	name       string
	parameters map[string]string
}

func Parse(s string) (*Type, error) {
	if stringd.IsEmpty(s) {
		return nil, fmt.Errorf("could not Parse MIME-Type: `%s` is empty", s)
	}

	re := regexp.MustCompile(regex).FindAllStringSubmatch(s, -2)
	if re == nil || len(re) != 1 || len(re[0]) != 5 {
		return nil, fmt.Errorf("illegal MIME-Type format: %s", s)
	}

	t := new(Type)
	if r, er := RegistryOf(re[0][1]); er != nil {
		return nil, er
	} else {
		t.registry = *r
		t.name = re[0][2]
	}

	if !stringd.IsEmpty(re[0][3]) &&
		!stringd.IsEmpty(re[0][4]) {
		t.parameters = map[string]string{
			re[0][3]: re[0][4],
		}
	}

	return t, nil
}

func (t *Type) Registry() Registry {
	return t.registry
}

func (t *Type) Name() string {
	return t.name
}

func (t *Type) Parameters() map[string]string {
	return t.parameters
}
