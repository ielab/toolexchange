package toolexchange

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/patrickmn/go-cache"
	"time"
)

type Item struct {
	Data       map[string]string `json:"data" form:"data" binding:"required"`
	Referrer   string            `json:"referrer" form:"referrer" binding:"required"`
	Expiration time.Time         `json:"expiration"`
}

type Exchanger struct {
	cache              *cache.Cache
	expirationInterval time.Duration
	cleanupInterval    time.Duration
}

func NewExchanger() *Exchanger {
	e := 5 * time.Minute
	c := 10 * time.Minute
	return &Exchanger{
		cache:              cache.New(e, c),
		expirationInterval: e,
		cleanupInterval:    c,
	}
}

func (e *Exchanger) PutItem(item Item) string {
	token := generateToken()
	item.Expiration = time.Now().Add(e.expirationInterval)
	e.cache.Set(token, item, cache.DefaultExpiration)
	return token
}

func (e *Exchanger) GetItem(token string) (Item, bool) {
	if item, ok := e.cache.Get(token); ok {
		return item.(Item), true
	}
	return Item{}, false
}

func generateToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return hex.EncodeToString(sha256.New().Sum(b))[:64]
}
