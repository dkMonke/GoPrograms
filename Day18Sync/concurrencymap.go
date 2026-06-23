// concurrencymap.go — Day 18: sync.Map for a concurrent session store.
// Session tracks user login/activity times. sync.Map provides a thread-safe
// map without explicit locking. Functions CreateSession, GetSession,
// UpdateLastSeen, CountActiveSessions, and GetAllSessions demonstrate
// Store, Load, and Range operations. (Note: this file is work-in-progress.)
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Session struct {
		UserID string
		LoginTime time.Time
		LastSeen time.Time
	}

	var sessions sync.Map

	func GetSession(sessionID string) (*Session, bool) {
	val,ok := sessions.Load(sessionId)
	if !ok {
		return nil,false
	}
	return val.(*Session), true
}


func CreateSession(sessionID, userID string) {
	session := &Session{
		userID: userID,
		LoginTime : time.Now(),
		LastSeen : time.Now(),
	}
	sessions.Store(sessionID,session)
	fmt.Printf("Session created: %s -> %s\n",sessionID,userID)

}

func UpdateLastSeen(sessionID string){
	val,ok := sessions.Load(sessionID)
	if !ok {
		return
	}
	session := val.(*Session)
	session.LastSeen = time.Now()
	sessions.Store(sessionID,session)
}

func CountACtiveSessions(withinSeconds int) int {
	count := 0
	now := time.now()

	sessions.Range(func(key,value interface{}) bool {
		session := value.(*Session)
		if now.Sub(session.LastSeen) < time.Duration(withinSeconds)*time.Second {
			count++
		}
		return true
	})
		return count
	}

	func GetAllSessions() map[string]*Session {
		result := make(map[string]*Session)
		sessions.Range(func(key,value interface{}) bool {
			sessionID :=key.(string)
			session := value.(*Session)
			result[sessionID] = session
			return true
		})
		return result
	}

	func main() {
		CreateSession("sess-1","alice")
		CreateSession("sess-2","bob")
		CreateSession("sess-3","vharlie")

		fmt.Println("\n======= Simulating concurrent activity ==\n")

		var wg sync.WaitGroup
		for i:=0;i < 100; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for j :=0
				






	}

