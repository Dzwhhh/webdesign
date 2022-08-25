package framework

import (
	"errors"
	"strings"
)

type Trie struct {
	root *node
}

type node struct {
	segment  string              // 节点对应的uri片段
	handlers []ControllerHandler // 中间件 + 节点绑定的handler
	childs   []*node             // 节点的所有子节点
	isLast   bool                // 是否为根节点
}

func NewTrie() *Trie {
	n := &node{
		segment: "",
		childs:  []*node{},
		isLast:  false,
	}
	return &Trie{
		root: n,
	}
}

// 查找路由函数
func (t *Trie) FindHandler(uri string) []ControllerHandler {
	matchNode := t.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers

}

// 添加路由函数
func (t *Trie) AddRouter(uri string, handlers []ControllerHandler) error {
	n := t.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist:" + uri)
	}
	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点
		childNodes := n.filterChildNodes(segment)
		// 如果有匹配的子节点
		if len(childNodes) > 0 {
			// 选择segment相同的节点
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			// 创建一个当前的node节点
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handlers = handlers
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

		n = objNode
	}
	return nil
}

// 是否为通配符
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func newNode() *node {
	return &node{
		segment: "",
		childs:  []*node{},
		isLast:  false,
	}
}

// 过滤下一层满足segment规则的节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}
	if isWildSegment(segment) {
		return n.childs
	}
	nodes := make([]*node, 0, len(n.childs))
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment { // 这里没看懂
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// 匹配符合的下一层节点
	cnodes := n.filterChildNodes(segment)
	if len(cnodes) == 0 {
		return nil
	}
	// 如果只有一个segment，则是最后一个标记
	if len(segments) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}

	// 如果有多个segment，则递归查找
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}
