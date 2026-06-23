// assignment.go — Day 10 Assignment: Notification system using interfaces.
// Notifier interface requires Notify(msg string) error. Four concrete types
// (Email, SMS, Log, Failing) implement it. NotifyAll iterates a []Notifier slice,
// showing how interfaces enable open/closed design — new notifiers can be added
// without modifying existing code.
package main

import (
	"fmt"
)

type Notifier interface {
	Notify(msg string) error
}

type EmailNotifier struct{}
type SMSNotifier struct{}
type LogNotifier struct{}
type FailingNotifier struct{}

func (e FailingNotifier) Notify(msg string) error {
	NotifyErr := fmt.Errorf("Exception with notification")
	fmt.Printf("%w", NotifyErr)
	return NotifyErr
}
func (e EmailNotifier) Notify(msg string) error {
	fmt.Println("EMAIL:", msg)
	return nil
}

func (e SMSNotifier) Notify(msg string) error {
	fmt.Println("SMS:", msg)
	return nil
}

func (e LogNotifier) Notify(msg string) error {
	fmt.Println("Log:", msg)
	return nil
}

func NotifyAll(notifiers []Notifier, msg string) {
	for _, i := range notifiers {
		i.Notify(msg)
	}

}
func main() {
	notifiers := []Notifier{
		FailingNotifier{},
		LogNotifier{},
		SMSNotifier{},
		EmailNotifier{},
	}
	NotifyAll(notifiers, "helloworld")

}
