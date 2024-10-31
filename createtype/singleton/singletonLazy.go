package singleton

import "sync"

// SingletonLazy 懒汉式单例
type SingletonLazy struct{}

var (
	lazySingleton *SingletonLazy
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *SingletonLazy {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &SingletonLazy{}
		})
	}
	return lazySingleton
}
