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

// Session represents a single user session and tracks when the user logged in
// and when they were last observed to be active. Instances are stored as
// pointers in the sessions map so that updates (such as refreshing LastSeen)
// are visible to all callers holding the same pointer.
type Session struct {
	UserID    string    // identifier of the user who owns this session
	LoginTime time.Time // time the session was first created
	LastSeen  time.Time // time of the user's most recent observed activity
}

// sessions is the process-wide, concurrency-safe store of active sessions,
// keyed by session ID (string) with *Session values. It uses sync.Map so that
// many goroutines can read and write sessions without explicit locking.
var sessions sync.Map

// GetSession looks up the session associated with sessionID.
// It returns the stored *Session and true when a session exists, or nil and
// false when no session is found. The returned pointer refers to the live
// session in the store, so callers should not mutate it without coordination.
func GetSession(sessionID string) (*Session, bool) {
	val, ok := sessions.Load(sessionID)
	if !ok {
		return nil, false
	}
	return val.(*Session), true
}

// CreateSession creates a new Session for userID and stores it under sessionID,
// overwriting any existing session with the same ID. Both LoginTime and
// LastSeen are initialized to the current time. It prints a confirmation line
// to standard output as a side effect.
func CreateSession(sessionID, userID string) {
	session := &Session{
		UserID:    userID,
		LoginTime: time.Now(),
		LastSeen:  time.Now(),
	}
	sessions.Store(sessionID, session)
	fmt.Printf("Session created: %s -> %s\n", sessionID, userID)

}

// UpdateLastSeen refreshes the LastSeen timestamp of the session identified by
// sessionID to the current time. If no session exists for sessionID, the call
// is a no-op. The session pointer is re-stored after mutation; because Session
// values are stored by pointer, the update is visible to other holders of the
// same pointer.
func UpdateLastSeen(sessionID string) {
	val, ok := sessions.Load(sessionID)
	if !ok {
		return
	}
	session := val.(*Session)
	session.LastSeen = time.Now()
	sessions.Store(sessionID, session)
}

// CountActiveSessions returns the number of sessions whose LastSeen falls
// within the last withinSeconds seconds relative to the current time. It scans
// every entry in the store via sync.Map.Range; the count reflects the sessions
// observed during the scan and may not be a precise instant-in-time snapshot
// under concurrent updates.
func CountActiveSessions(withinSeconds int) int {
	count := 0
	now := time.Now()

	sessions.Range(func(key, value interface{}) bool {
		session := value.(*Session)
		if now.Sub(session.LastSeen) < time.Duration(withinSeconds)*time.Second {
			count++
		}
		return true
	})
	return count
}

// GetAllSessions returns a plain map containing every session currently in the
// store, keyed by session ID. The returned map is a freshly allocated snapshot,
// but its *Session values are the live pointers from the store, so mutating a
// session through the returned map also mutates the stored session.
func GetAllSessions() map[string]*Session {
	result := make(map[string]*Session)
	sessions.Range(func(key, value interface{}) bool {
		sessionID := key.(string)
		session := value.(*Session)
		result[sessionID] = session
		return true
	})
	return result
}

// main demonstrates concurrent use of the sync.Map-backed session store. It
// seeds three sessions, then launches 100 reader goroutines and 5 updater
// goroutines that concurrently look up and refresh randomly chosen sessions.
// After waiting for all goroutines to finish, it prints the final state of
// every session and the number of sessions active in the last 5 seconds.
func main() {
	CreateSession("sess-1", "alice")
	CreateSession("sess-2", "bob")
	CreateSession("sess-3", "vharlie")

	fmt.Println("\n======= Simulating concurrent activity ==\n")

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				sessionID := fmt.Sprintf("sess-%d", rand.Intn(3)+1)
				session, ok := GetSession(sessionID)
				if ok {
					//simulating user reading data
					_ = session.UserID
				}
				time.Sleep(time.Millisecond)
			}
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				sessionID := fmt.Sprintf("sess-%d", rand.Intn(3)+1)
				UpdateLastSeen(sessionID)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("\n=== Final state ==\n")
	sessions.Range(func(key, value interface{}) bool {
		sessionID := key.(string)
		session := value.(*Session)
		fmt.Printf("%s : user=%s, last_seen=%v ago\n", sessionID, session.UserID, time.Since(session.LastSeen).Round(time.Millisecond))
		return true
	})
	fmt.Printf("\nActive sessions (last 5s): %d\n", CountActiveSessions(5))
}
