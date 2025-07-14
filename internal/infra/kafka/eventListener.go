package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer/actors"
	"github.com/lesta-battleship/matchmaking/pkg/packets"
	"github.com/segmentio/kafka-go"
)

const GuildWarTopic = "guild_war_confirm"

type KafkaEventListener struct {
	id string

	reader     *kafka.Reader
	matchmaker actors.Matchmaker
}

func NewKafkaEventListener(id, addr string) *KafkaEventListener {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{addr},
		Topic:     GuildWarTopic,
		Partition: 0,
		MaxBytes:  10e6,
	})

	return &KafkaEventListener{
		id: id,

		reader:     reader,
		matchmaker: nil,
	}
}

func (l *KafkaEventListener) Id() string {
	return l.id
}

func (l *KafkaEventListener) ConnectTo(matchmaker actors.Matchmaker) {
	l.matchmaker = matchmaker
}

func (l *KafkaEventListener) GetPacket(senderId string, packet packets.Packet) {
	l.matchmaker.GetPacket(senderId, packet)

	log.Printf("KafkaEventListener %q: Received packet %T from %q", l.id, packet.Body, packet.SenderId)
}

func (l *KafkaEventListener) Start() {
	defer func() {
		l.Stop()
	}()

	log.Printf("KafkaEventListener %q: Started", l.id)

	for {
		message, err := l.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("KafkaEventListener %q: %v", l.id, err)

			return
		}

		guildWar := &packets.GuildWarStarted{}
		json.Unmarshal(message.Value, guildWar)
		packet := packets.Packet{SenderId: l.id, Type: guildWar.Type(), Body: guildWar}

		l.GetPacket(l.id, packet)
	}
}

func (l *KafkaEventListener) Stop() {
	if l.reader != nil {
		l.reader.Close()
	}
	l.reader = nil

	l.matchmaker = nil

	log.Printf("KafkaEventListener %q: Stopped", l.id)
}
