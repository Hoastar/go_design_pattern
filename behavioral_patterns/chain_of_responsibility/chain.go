package chain

// 敏感词过滤场景

// SensitiveWordFilter 敏感词过滤器，判定是否是敏感词
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// SensitiveWordFilterChain 职责链
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

// AddFilter 添加一个过滤器
func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

// Filter 过滤方法
func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		if filter.Filter(content) {
			// 如果发现敏感直接返回
			return true
		}
	}
	return false
}

// AdSensitiveWordFilter 广告过滤
type AdSensitiveWordFilter struct{}

// Filter 实现过滤算法
func (f *AdSensitiveWordFilter) Filter(content string) bool {
	// TODO: 实现算法
	return true
}

// PoliticalWordFilter 政治敏感过滤
type PoliticalWordFilter struct{}

// Filter 实现过滤算法
func (f *PoliticalWordFilter) Filter(content string) bool {
	// TODO: 实现算法
	return true
}
