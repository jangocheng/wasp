package sessions

import (
	"bytes"
	"io"
	"sync"

	"github.com/vx-labs/mqtt-protocol/encoder"
	"github.com/vx-labs/mqtt-protocol/packet"
)

type Session struct {
	ID           string
	ClientID     string
	MountPoint   string
	Lwt          *packet.Publish
	conn         io.WriteCloser
	Encoder      *encoder.Encoder
	Disconnected bool
	mtx          sync.Mutex
	topics       [][]byte
}

func NewSession(c io.WriteCloser, stats encoder.StatRecorder) *Session {
	return &Session{
		conn:    c,
		Encoder: encoder.New(c, encoder.WithStatRecorder(stats)),
	}
}

func (s *Session) Send(publish *packet.Publish) error {
	if len(s.MountPoint) > len(publish.Topic) {
		return nil
	}

	outgoing := &packet.Publish{
		Header:    publish.Header,
		Topic:     TrimMountPoint(s.MountPoint, publish.Topic),
		MessageId: 1,
		Payload:   publish.Payload,
	}
	return s.Encoder.Publish(outgoing)
}

func (s *Session) ProcessConnect(connect *packet.Connect) error {
	s.ClientID = string(connect.ClientId)
	if len(connect.WillTopic) > 0 {
		s.Lwt = &packet.Publish{
			Header:  &packet.Header{Retain: connect.WillRetain, Qos: connect.WillQos},
			Topic:   connect.WillTopic,
			Payload: connect.WillPayload,
		}
	}
	return nil
}
func (s *Session) AddTopic(t []byte) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.topics = append(s.topics, t)
}
func (s *Session) RemoveTopic(t []byte) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for idx := range s.topics {
		if bytes.Equal(s.topics[idx], t) {
			s.topics[idx] = s.topics[len(s.topics)-1]
			s.topics = s.topics[:len(s.topics)-1]
		}
	}
}
func (s *Session) GetTopics() [][]byte {
	return s.topics
}
func (s *Session) Close() error {
	return s.conn.Close()
}
