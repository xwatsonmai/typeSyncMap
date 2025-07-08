# typeSyncMap
使用泛型对sync.Map进行二次封装/Secondary encapsulation of sync.Map using generics

# 特性/Features
- 支持泛型/Supports generics
- 有明确的类型约束，不需要在逻辑代码中进行类型断言/Has clear type constraints, no need for type assertions in logic code
- 与sync.Map的API兼容/Compatible with sync.Map API

# 使用方法
```go
package main

import (
	"github.com/xwatsonmai/typeSyncMap"
	"sync"
	_ "github.com/xwatsonmai/typeSyncMap"
)

func main() {
	data := typeSyncMap.NewTypeSyncMap[string, int]()
}
```
