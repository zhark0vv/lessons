package message

import (
	kafkago "github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type Message struct {
	kafkago.Message
}

func NewFromRaw(m kafkago.Message) Message {
	return Message{m}
}

func FromProto(p proto.Message) (*Message, error) {
	data, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	return &Message{kafkago.Message{Value: data}}, nil
}

func (m *Message) AsProto(p proto.Message) error {
	return proto.Unmarshal(m.Value, p)
}

func (m *Message) Raw() kafkago.Message {
	return m.Message
}
