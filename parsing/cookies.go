package parsing

import (
	"github.com/xanderazuaje/xocket/colors"
	"net/http"
	"strconv"
	"time"
)

type ExpectedCookie struct {
	Name        string
	Value       string
	Quoted      bool
	Path        string
	Domain      string
	Expires     time.Time
	RawExpires  string `yaml:"raw-expires"`
	MaxAge      *int   `yaml:"max-age"`
	Secure      *bool
	HttpOnly    *bool          `yaml:"http-only"`
	SameSite    *http.SameSite `yaml:"same-site"`
	Partitioned *bool
	Raw         string
	Unparsed    []string
}

func printPropertyDiff(count int, k, v1, v2 string) {
	colors.Log(
		"\t@r*(DIFF) at cookie[%d]\n\t\tkey @b*(%s)\n\t\texpected -> %s\n\t\tgot -> %s",
		count,
		k,
		v1,
		v2,
	)
}

func (c *ExpectedCookie) PrintDifference(count int, rc *http.Cookie) bool {
	ok := true
	if rc == nil {
		colors.Log("\t@r*(DIFF) cookie named '%s' doesn't exist in response\n", c.Name)
		return false
	}
	if c.Name != rc.Name {
		printPropertyDiff(count, "name", c.Name, rc.Name)
		ok = false
	}
	if c.Value != "" && c.Value != rc.Value {
		printPropertyDiff(count, "value", c.Value, rc.Value)
		ok = false
	}
	if c.Path != "" && c.Path != rc.Path {
		printPropertyDiff(count, "path", c.Path, rc.Path)
		ok = false
	}
	if c.Domain != "" && c.Domain != rc.Domain {
		printPropertyDiff(count, "domain", c.Domain, rc.Domain)
		ok = false
	}
	if !c.Expires.IsZero() && c.Expires.Compare(rc.Expires) != 0 {
		printPropertyDiff(count, "expires", c.Expires.String(), rc.Expires.String())
		ok = false
	}
	if c.MaxAge != nil && *c.MaxAge != rc.MaxAge {
		printPropertyDiff(count, "domain", c.Domain, rc.Domain)
		ok = false
	}
	if c.Secure != nil && *c.Secure != rc.Secure {
		printPropertyDiff(
			count,
			"secure",
			strconv.FormatBool(*c.Secure),
			strconv.FormatBool(rc.Secure),
		)
		ok = false
	}
	if c.HttpOnly != nil && *c.HttpOnly != rc.HttpOnly {
		printPropertyDiff(
			count,
			"httponly",
			strconv.FormatBool(*c.HttpOnly),
			strconv.FormatBool(rc.HttpOnly),
		)
		ok = false
	}
	if c.SameSite != nil && *c.SameSite != rc.SameSite {
		printPropertyDiff(
			count,
			"samesite",
			strconv.FormatInt(int64(*c.SameSite), 10),
			strconv.FormatInt(int64(rc.SameSite), 10),
		)
		ok = false
	}
	if c.Partitioned != nil && *c.Partitioned != rc.Partitioned {
		printPropertyDiff(
			count,
			"partitioned",
			strconv.FormatBool(*c.Partitioned),
			strconv.FormatBool(rc.Partitioned),
		)
		ok = false
	}
	if c.Raw != "" && c.Raw != rc.Raw {
		printPropertyDiff(count, "raw", c.Raw, rc.Raw)
		ok = false
	}
	return ok
}
