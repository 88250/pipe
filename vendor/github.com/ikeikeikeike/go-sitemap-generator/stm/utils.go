package stm

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/beevik/etree"
)

// BufferPool is
type BufferPool struct {
	sync.Pool
}

// NewBufferPool is
func NewBufferPool() *BufferPool {
	return &BufferPool{
		Pool: sync.Pool{New: func() interface{} {
			b := bytes.NewBuffer(make([]byte, 256))
			b.Reset()
			return b
		}},
	}
}

// Get is
func (bp *BufferPool) Get() *bytes.Buffer {
	return bp.Pool.Get().(*bytes.Buffer)
}

// Put is
func (bp *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	bp.Pool.Put(b)
}

// SetBuilderElementValue if it will change to struct from map if the future's
// author is feeling a bothersome in this function.
func SetBuilderElementValue(elm *etree.Element, data [][]interface{}, basekey string) (*etree.Element, bool) {
	var child *etree.Element

	key := basekey
	ts, tk := spaceDecompose(elm.Tag)
	_, sk := spaceDecompose(elm.Space)

	if elm.Tag != "" && ts != "" && tk != "" {
		key = fmt.Sprintf("%s:%s", elm.Space, basekey)
	} else if sk != "" {
		key = fmt.Sprintf("%s:%s", sk, basekey)
	}

	var values interface{}
	var found bool
	for _, v := range data {
		if v[0] == basekey {
			values = v[1]
			found = true
			break
		}
	}
	if !found {
		return child, false
	}

	switch value := values.(type) {
	case nil:
	default:
		child = elm.CreateElement(key)
		child.SetText(fmt.Sprint(value))
	case int:
		child = elm.CreateElement(key)
		child.SetText(fmt.Sprint(value))
	case string:
		child = elm.CreateElement(key)
		child.SetText(value)
	case float64, float32:
		child = elm.CreateElement(key)
		child.SetText(fmt.Sprint(value))
	case time.Time:
		child = elm.CreateElement(key)
		child.SetText(value.Format(time.RFC3339))
	case bool:
		_ = elm.CreateElement(fmt.Sprintf("%s:%s", key, key))
	case []int:
		for _, v := range value {
			child = elm.CreateElement(key)
			child.SetText(fmt.Sprint(v))
		}
	case []string:
		for _, v := range value {
			child = elm.CreateElement(key)
			child.SetText(v)
		}
	case []Attr:
		for _, attr := range value {
			child = elm.CreateElement(key)
			for k, v := range attr {
				child.CreateAttr(k, v)
			}
		}
	case Attrs:
		val, attrs := value[0], value[1]

		child, _ = SetBuilderElementValue(elm, URL{[]interface{}{basekey, val}}, basekey)
		switch attr := attrs.(type) {
		case map[string]string:
			for k, v := range attr {
				child.CreateAttr(k, v)
			}
		// TODO: gotta remove below
		case Attr:
			for k, v := range attr {
				child.CreateAttr(k, v)
			}
		}

	case interface{}:
		var childkey string
		if sk == "" {
			childkey = fmt.Sprintf("%s:%s", key, key)
		} else {
			childkey = fmt.Sprint(key)
		}

		switch value := values.(type) {
		case []URL:
			for _, val := range value {
				child := elm.CreateElement(childkey)
				for _, v := range val {
					SetBuilderElementValue(child, val, v[0].(string))
				}
			}
		case URL:
			child := elm.CreateElement(childkey)
			for _, v := range value {
				SetBuilderElementValue(child, value, v[0].(string))
			}
		}
	}
	return child, true
}

// MergeMap TODO: Slow function: It wants to change fast function
func MergeMap(src, dst [][]interface{}) [][]interface{} {
	for _, v := range dst {
		found := false
		for _, vSrc := range src {
			if v[0] == vSrc[0] {
				found = true
				break
			}
		}
		if !found {
			src = append(src, v)
		}
	}
	return src
}

// ToLowerString converts lower strings from including capital or upper strings.
func ToLowerString(befores []string) (afters []string) {
	for _, name := range befores {
		afters = append(afters, strings.ToLower(name))
	}
	return afters
}

// URLJoin TODO: Too slowly
func URLJoin(src string, joins ...string) string {
	var u *url.URL
	lastnum := len(joins)
	base, _ := url.Parse(src)

	for i, j := range joins {
		if !strings.HasSuffix(j, "/") && lastnum > (i+1) {
			j = j + "/"
		}

		u, _ = url.Parse(j)
		base = base.ResolveReference(u)
	}

	return base.String()
}

// spaceDecompose is separating strings for the SetBuilderElementValue
func spaceDecompose(str string) (space, key string) {
	colon := strings.IndexByte(str, ':')
	if colon == -1 {
		return "", str
	}
	return str[:colon], str[colon+1:]
}
