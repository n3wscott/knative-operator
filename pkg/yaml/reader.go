package yaml

import (
	"bufio"
	"fmt"
	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
	"strings"
)

type ConfigFile struct {
	Path      string
	Resources map[string]unstructured.Unstructured
}

func (c *ConfigFile) Read() error {
	if c.Path == "" {
		return fmt.Errorf("no config file provided")
	}
	f, err := os.Open(c.Path)
	if err != nil {
		fmt.Print(err)
	}
	lines := 0
	var ob []byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		l = strings.TrimSpace(l)
		if l == "" || strings.HasPrefix(l, "#") {
			// skip over blanks and comments.
			continue
		}
		if l == "---" {
			if lines > 0 {
				if err := c.Parse(ob); err != nil {
					return err
				}
				ob = []byte{}
				lines = 0
			}
		} else {
			ob = append(ob, scanner.Bytes()...)
			ob = append(ob, []byte("\n")...)
			lines++
		}
	}
	return nil
}

func (c *ConfigFile) Parse(b []byte) error {
	u := unstructured.Unstructured{}
	if err := yaml.Unmarshal(b, &u); err != nil {
		return err
	}
	if err := c.Store(&u); err != nil {
		return err
	}
	return nil
}

func (c *ConfigFile) Store(u *unstructured.Unstructured) error {
	if c.Resources == nil {
		c.Resources = make(map[string]unstructured.Unstructured)
	}
	key := Key(u)
	_, found := c.Resources[key]
	if found {
		return fmt.Errorf("object redefined for key %q", key)
	}
	c.Resources[key] = *u
	return nil
}

func Key(u *unstructured.Unstructured) string {
	kvn := fmt.Sprintf("%s%s/%s/%s", u.GetKind(), u.GetAPIVersion(), u.GetName(), u.GetNamespace())
	return kvn
}
