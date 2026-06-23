// assignment.go — Day 10 Assignment: Notification system using interfaces.
// Notifier interface requires Notify(msg string) error. Four concrete types
// (Email, SMS, Log, Failing) implement it. NotifyAll iterates a []Notifier slice,
// showing how interfaces enable open/closed design — new notifiers can be added
// without modifying existing code.
package main

import (
	"fmt"
)

// Notifier is the abstraction for anything that can deliver a message.
// Any type that implements Notify(msg string) error satisfies this interface
// implicitly, which lets NotifyAll treat all notifiers uniformly.
type Notifier interface {
	Notify(msg string) error
}

// EmailNotifier sends notifications via email (simulated by printing).
type EmailNotifier struct{}

// SMSNotifier sends notifications via SMS (simulated by printing).
type SMSNotifier struct{}

// LogNotifier records notifications to a log (simulated by printing).
type LogNotifier struct{}

// FailingNotifier always fails when notifying; it is used to demonstrate
// error handling in the notification pipeline.
type FailingNotifier struct{}

// Notify implements the Notifier interface for FailingNotifier. It always
// returns an error to simulate a delivery failure. The msg parameter is
// ignored. It also prints the formatted error and returns that same error.
func (e FailingNotifier) Notify(msg string) error {
	NotifyErr := fmt.Errorf("Exception with notification")
	fmt.Printf("%w", NotifyErr)
	return NotifyErr
}

// Notify implements the Notifier interface for EmailNotifier. It prints the
// message prefixed with "EMAIL:" and returns nil to indicate success.
func (e EmailNotifier) Notify(msg string) error {
	fmt.Println("EMAIL:", msg)
	return nil
}

// Notify implements the Notifier interface for SMSNotifier. It prints the
// message prefixed with "SMS:" and returns nil to indicate success.
func (e SMSNotifier) Notify(msg string) error {
	fmt.Println("SMS:", msg)
	return nil
}

// Notify implements the Notifier interface for LogNotifier. It prints the
// message prefixed with "Log:" and returns nil to indicate success.
func (e LogNotifier) Notify(msg string) error {
	fmt.Println("Log:", msg)
	return nil
}

// NotifyAll dispatches msg to every Notifier in the notifiers slice by calling
// each one's Notify method in order. Errors returned by individual notifiers
// are ignored, so a single failing notifier does not stop the rest from being
// invoked.
func NotifyAll(notifiers []Notifier, msg string) {
	for _, i := range notifiers {
		i.Notify(msg)
	}

}

// main builds a slice of different Notifier implementations (including the
// failing one) and broadcasts a single message to all of them via NotifyAll,
// demonstrating interface-based polymorphism.
func main() {
	notifiers := []Notifier{
		FailingNotifier{},
		LogNotifier{},
		SMSNotifier{},
		EmailNotifier{},
	}
	NotifyAll(notifiers, "helloworld")

}
