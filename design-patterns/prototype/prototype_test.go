package prototype

import (
	"reflect"
	"testing"
)

func TestDocument_Clone(t *testing.T) {
	// 1. 初始化一个成本较高的基础原型
	prototypeDocument := &Document{
		Title:   "System Architecture",
		Content: "Initial draft...",
		Tags:    []string{"architecture", "design"},
		Meta:    map[string]string{"author": "admin", "status": "draft"},
	}

	// 2. 通过克隆创建新对象
	cloneInterface := prototypeDocument.Clone()
	doc1, ok := cloneInterface.(*Document)

	if !ok {
		t.Fatalf("Clone() 未返回 *Document 类型")
	}

	// 3. 验证对象级别的内存隔离（指针地址不同）
	if prototypeDocument == doc1 {
		t.Fatalf("克隆失败：新对象与原型对象指向相同的内存地址")
	}

	// 4. 对克隆对象进行微调（修改值类型与引用类型）
	doc1.Title = "System Architecture - v1.0"
	doc1.Tags = append(doc1.Tags, "v1.0")
	doc1.Meta["status"] = "published"

	// 5. 断言验证：值类型字段是否隔离
	if prototypeDocument.Title == doc1.Title {
		t.Errorf("值类型隔离失败：修改克隆对象影响了原型对象. 原型 Title: %s", prototypeDocument.Title)
	}

	// 6. 断言验证：引用类型（切片）是否隔离
	if reflect.DeepEqual(prototypeDocument.Tags, doc1.Tags) {
		t.Errorf("切片隔离失败：原型 Tags %v 与克隆 Tags %v 不应相同", prototypeDocument.Tags, doc1.Tags)
	}

	if len(prototypeDocument.Tags) != 2 {
		t.Errorf("切片隔离失败：原型 Tags 长度被意外修改，期望值 2，实际值 %d", len(prototypeDocument.Tags))
	}

	// 7. 断言验证：引用类型（映射）是否隔离
	if prototypeDocument.Meta["status"] == doc1.Meta["status"] {
		t.Errorf("映射隔离失败：原型 Meta status 被意外修改为 %s", prototypeDocument.Meta["status"])
	}
}
