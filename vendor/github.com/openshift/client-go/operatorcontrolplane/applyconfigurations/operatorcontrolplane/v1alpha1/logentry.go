// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LogEntryApplyConfiguration represents an declarative configuration of the LogEntry type for use
// with apply.
type LogEntryApplyConfiguration struct {
	Start   *v1.Time     `json:"time,omitempty"`
	Success *bool        `json:"success,omitempty"`
	Reason  *string      `json:"reason,omitempty"`
	Message *string      `json:"message,omitempty"`
	Latency *v1.Duration `json:"latency,omitempty"`
}

// LogEntryApplyConfiguration constructs an declarative configuration of the LogEntry type for use with
// apply.
func LogEntry() *LogEntryApplyConfiguration {
	return &LogEntryApplyConfiguration{}
}

// WithStart sets the Start field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Start field is set to the value of the last call.
func (b *LogEntryApplyConfiguration) WithStart(value v1.Time) *LogEntryApplyConfiguration {
	b.Start = &value
	return b
}

// WithSuccess sets the Success field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Success field is set to the value of the last call.
func (b *LogEntryApplyConfiguration) WithSuccess(value bool) *LogEntryApplyConfiguration {
	b.Success = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *LogEntryApplyConfiguration) WithReason(value string) *LogEntryApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *LogEntryApplyConfiguration) WithMessage(value string) *LogEntryApplyConfiguration {
	b.Message = &value
	return b
}

// WithLatency sets the Latency field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Latency field is set to the value of the last call.
func (b *LogEntryApplyConfiguration) WithLatency(value v1.Duration) *LogEntryApplyConfiguration {
	b.Latency = &value
	return b
}
