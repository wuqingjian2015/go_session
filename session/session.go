package session 

import (
	"time"
	"net/url"
	"net/http"
	"encoding/base64"
	"io"
	"fmt"
	"sync"
	"crypto/rand"
)
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLiftTime int64)
}
type Manager struct {
	cookieName string 
	lock sync.Mutex
	provider Provider 
	maxlifetime int64 
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	 cookie, err := r.Cookie(manager.cookieName)
	 if err != nil || cookie.Value == "" {
		sid := manager.sessionId(); 
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name:manager.cookieName, Value:url.QueryEscape(sid), Path:"/", HttpOnly:true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	 } else {
		 sid, _ := url.QueryUnescape(cookie.Value)
		 session, _ = manager.provider.SessionRead(sid)
	 }
	 return session
}
// func (manager *Manager) SessionRead(w http.ResponseWriter, r *http.Request) (session Session){

// }
func (manager *Manager) GC(){
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func(){manager.GC()})
}
type Session interface {
	Set(key, value interface{}) error 
	Get(key interface{}) interface{}
	Delete(key interface{}) error 
	SessionID() string  
}

var providers = make(map[string]Provider) 

func Register(name string, provider Provider){
	if provider == nil {
		panic("session: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("session: Register called twice for provider " + name)
	}

	fmt.Println("Register:", provider)
	providers[name] = provider 
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error){
	provider, ok := providers[provideName]; 
	if !ok {
		return nil, fmt.Errorf("Session: unknown provide %q (forgotten int64?)", provideName)
	} 
	return &Manager{provider:provider, cookieName: cookieName, maxlifetime:maxlifetime}, nil 
}

func init(){


}