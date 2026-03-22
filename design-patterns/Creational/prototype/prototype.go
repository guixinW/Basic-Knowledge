package prototype

type Cloneable interface {
	Clone() Cloneable
}

type Document struct {
	Title   string
	Content string
	Tags    []string          // 引用类型，必须进行深拷贝
	Meta    map[string]string // 引用类型，必须进行深拷贝
}

func (d *Document) Clone() Cloneable {
	// 初始化新对象，执行值类型的浅拷贝
	clone := &Document{
		Title:   d.Title,
		Content: d.Content,
	}

	// 针对切片执行深拷贝
	if d.Tags != nil {
		clone.Tags = make([]string, len(d.Tags))
		copy(clone.Tags, d.Tags) // 使用内置的 copy 函数
	}

	// 针对映射执行深拷贝
	if d.Meta != nil {
		clone.Meta = make(map[string]string, len(d.Meta))
		for k, v := range d.Meta {
			clone.Meta[k] = v
		}
	}

	return clone
}
