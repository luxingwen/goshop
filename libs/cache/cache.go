package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type Cache struct {
	rw        *sync.RWMutex
	size      int
	entryList *list.List
	items     *sync.Map
}

type entry struct {
	key     interface{}
	value   interface{}
	expired int64
}

func NewCache(size int) *Cache {
	return &Cache{size: size, entryList: list.New(), items: new(sync.Map)}
}

func (lru *Cache) Get(key interface{}) (val interface{}, ok bool) {
	if item0, ok := lru.items.Load(key); ok {
		item := item0.(*list.Element)
		val = item.Value.(*entry).value
		if item.Value.(*entry).expired < time.Now().Unix() {
			lru.removeElement(item)
			return nil, false
		}
		if val != nil {
			lru.entryList.MoveToFront(item)
			return val, true
		}
		lru.removeElement(item)
	}

	mcache := &Mcache{}
	r, err := mcache.GetByKey(key)
	if err != nil {
		return nil, false
	}
	if r.Expired < time.Now().Unix() {
		return nil, false
	}

	return r.Value, true
}

func (lru *Cache) Put(key interface{}, value interface{}, expired int64) {

	mcache := &Mcache{Key: key, Value: value.(string), Expired: expired + time.Now().Unix()}
	err := mcache.Insert()
	if err != nil {
		log.Error("mcache insert err:", err)
	}
	if item0, ok := lru.items.Load(key); ok {
		item := item0.(*list.Element)
		lru.entryList.MoveToFront(item)
		ent := item.Value.(*entry)
		ent.value = value
		ent.expired = time.Now().Unix() + expired
		return
	}
	element := lru.entryList.PushFront(&entry{key: key, value: value, expired: time.Now().Unix() + expired})
	lru.items.Store(key, element)
	if lru.size > 0 && lru.size < lru.entryList.Len() {
		lru.removeOldItem()
	}
}

func (lru *Cache) Remove(key interface{}) {
	if item0, ok := lru.items.Load(key); ok && item0 != nil {
		item := item0.(*list.Element)
		if item != nil {
			lru.removeElement(item)
		}

	}
}

func (lru *Cache) removeOldItem() {
	lastItem := lru.entryList.Back()
	if lastItem != nil {
		lru.removeElement(lastItem)
	}
}

func (lru *Cache) removeElement(item *list.Element) {
	lru.entryList.Remove(item)
	if item.Value == nil {
		return
	}
	ent := item.Value.(*entry)
	lru.items.Delete(ent.key)

}

// Resize changes the cache size.
func (lru *Cache) Resize(size int) (evicted int) {
	diff := lru.entryList.Len() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff && size > 0; i++ {
		lru.removeOldItem()
	}
	lru.size = size
	return diff
}

var defaultCache *Cache = NewCache(1000)

func Get(key interface{}) (r interface{}, flag bool) {
	return defaultCache.Get(key)
}

func Put(key interface{}, value interface{}, expired int64) {
	defaultCache.Put(key, value, expired)
}

var db *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init(username, pwd, host, port, dbname string) *gorm.DB {
	log.Info("init db")
	server := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, pwd, host, port, dbname)

	db0, err := gorm.Open("mysql", server)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db0.DB().SetMaxIdleConns(10)
	db0.LogMode(true)
	db = db0
	db.AutoMigrate(&Mcache{})
	return db
}

//
type Mcache struct {
	gorm.Model
	Key     interface{} `gorm:"column:key;type:varchar(75);unique_index"`
	Value   string      `gorm:"column:value;type:text"`
	Expired int64       `gorm:"column:expired"`
}

func (Mcache) TableName() string {
	return "m_cache"
}

func (m *Mcache) GetByKey(key interface{}) (r *Mcache, err error) {
	r = new(Mcache)
	err = db.Where("`m_cache`.`key` = ?", key).First(&r).Error
	return
}

func (m *Mcache) Insert() (err error) {
	_, err = m.GetByKey(m.Key)
	if err != nil && err.Error() == "record not found" {
		err = db.Create(m).Error
		return
	}
	err = db.Table(m.TableName()).Where("id = ? AND key = ?", m.ID, m.Key).Update(m).Error
	return
}
