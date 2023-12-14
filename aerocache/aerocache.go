package aerocache

import (
	"fmt"
	"log"
	"sync"
)

// 负责与外界交互，控制缓存存储和获取的主流程
type Group struct {
	name      string //该片缓存的名称
	getter    Getter
	mainCache cache
}

var (
	RWlock sync.RWMutex              //全局锁
	groups = make(map[string]*Group) //缓存区块
)

// 创建一个group的实例
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("getter is not allowed empty")
	}
	RWlock.Lock()
	defer RWlock.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// 获取
func GetGroup(name string) *Group {
	RWlock.RLock()
	defer RWlock.RLocker()
	if g, ok := groups[name]; ok {
		return g
	}
	panic("no cache in")
	return nil
}

type Getter interface {
	Get(key string) ([]byte, error)
}

// 函数式类型
type GetterFunc func(key string) ([]byte, error)

// 此处为一个回调函数
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// 缓存是否击中?从mainCache中查找?返回:加载缓存
func (g *Group) Get(key string) (ByteView, error) {
	//判断
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	//缓存击中
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[AeroCache] hit")
		return v, nil
	}

	//调用加载
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	log.Println("[AeroCache] Loading")
	return g.getLocally(key)
}

// 获取本地缓存，TODO 分布式缓存
func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err

	}
	value := ByteView{b: cloneBytes(bytes)}
	//loading cache
	g.AddCache(key, value)
	return value, nil
}

func (g *Group) AddCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
