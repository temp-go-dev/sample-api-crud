# MEMO

## 言語メモ

### 外側には、内部キャッシュを見せたくない実装？

type Cache struct {
	*cache
}

type cache struct {
	defaultExpiration time.Duration
	items             map[string]Item
	mu                sync.RWMutex
	onEvicted         func(string, interface{})
	janitor           *janitor
}


### ドットインポート

