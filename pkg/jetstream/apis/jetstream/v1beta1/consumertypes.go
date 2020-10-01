package v1beta1

import (
	k8smeta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Consumer is a specification for a Consumer resource
type Consumer struct {
	k8smeta.TypeMeta   `json:",inline"`
	k8smeta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsumerSpec `json:"spec"`
	Status Status       `json:"status"`
}

func (c *Consumer) GetSpec() interface{} {
	return c.Spec
}

// ConsumerSpec is the spec for a Consumer resource
type ConsumerSpec struct {
	StreamName     string `json:"streamName"`
	DeliverPolicy  string `json:"deliverPolicy"`
	OptStartSeq    int    `json:"optStartSeq"`
	OptStartTime   string `json:"optStartTime"`
	DurableName    string `json:"durableName"`
	DeliverSubject string `json:"deliverSubject"`
	AckPolicy      string `json:"ackPolicy"`
	AckWait        string `json:"ackWait"`
	MaxDeliver     int    `json:"maxDeliver"`
	FilterSubject  string `json:"filterSubject"`
	ReplayPolicy   string `json:"replayPolicy"`
	SampleFreq     string `json:"sampleFreq"`
	RateLimitBps   int    `json:"rateLimitBps"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConsumerList is a list of Consumer resources
type ConsumerList struct {
	k8smeta.TypeMeta `json:",inline"`
	k8smeta.ListMeta `json:"metadata"`

	Items []Consumer `json:"items"`
}
