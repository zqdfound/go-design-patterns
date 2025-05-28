package structure

import (
	"fmt"
	"strconv"
)

// Proxy Pattern
// The Proxy Pattern provides a surrogate or placeholder for another object to control access to it.
// It can be used to add an additional layer of control, such as lazy initialization, access control, logging, etc.

// 数据服务接口
type DataService interface {
	GetData(id int) string
}

// 真实数据服务
type DBDataService struct {
}

func (d *DBDataService) GetData(id int) string {
	// 模拟从数据库获取数据
	return "Data from DB for ID: " + strconv.Itoa(id)
}

// 代理数据服务
type ProxyDataService struct {
	realService DataService
	cache       map[int]string
}

func NewProxyDataService() *ProxyDataService {
	return &ProxyDataService{
		realService: &DBDataService{},
		cache:       make(map[int]string),
	}
}

func (p *ProxyDataService) GetData(id int) string {
	// 检查缓存
	if data, found := p.cache[id]; found {
		fmt.Println("Cache hit for ID:", id)
		return data
	}
	// 如果缓存中没有，调用真实服务
	data := p.realService.GetData(id)
	fmt.Println("Cache miss for ID:", id)
	// 将数据存入缓存
	p.cache[id] = data
	return data
}
