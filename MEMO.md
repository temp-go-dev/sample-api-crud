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


#### エラーハンドリング
https://qiita.com/nayuneko/items/3c0b3c0de9e8b27c9548

#### return nilはダメ
https://qiita.com/najeira/items/0bb0acdd7a71fc3f559b