package memory

import (
	"fmt"
	"container/list"
	"session"
	"sync"
	"time"
)


var pder = &Provider{list: list.New()}

type SessionStore struct{
	sid string
	timeAccessed time.Time 
	value map[interface{}] interface{}
}

func (ss *SessionStore) Set(key, value interface{}) error {
	ss.value[key] = value 
	pder.SessionUpdate(ss.sid)
	return nil 
}

func (ss *SessionStore) Get(key interface{}) interface{}{
	pder.SessionUpdate(ss.sid)
	if v, ok := ss.value[key]; ok {
		return v 
	} else {
		return  nil 
	}
}
func (ss *SessionStore) Delete(key interface{}) error {
	delete(ss.value, key)
	pder.SessionUpdate(ss.sid)
	return nil 
}

func (ss *SessionStore) SessionID() string {
	return ss.sid 
}

type Provider struct {
	lock sync.Mutex
	sessions map[string] *list.Element
	list *list.List
}

func (pder *Provider) SessionInit(sid string) (session.Session, error){
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newSess := &SessionStore{sid:sid, timeAccessed: time.Now(), value:v}
	element := pder.list.PushBack(newSess)
	pder.sessions[sid] = element  
	return newSess, nil 
}
func (pder *Provider) SessionRead(sid string) (session.Session, error){
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil 
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err 
	}
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil 
	}
	return nil 
}
func (pder *Provider) SessionGC(maxlifttime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break; 
		}
		if element.Value.(*SessionStore).timeAccessed.Unix() + maxlifttime < time.Now().Unix() {
			pder.list.Remove(element)
			fmt.Println("删除session", element.Value.(*SessionStore).sid)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break; 
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
	}
	return nil 
}
func init(){

	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}