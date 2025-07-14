package strategies

import (
	"log"

	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer/actors"
	"github.com/lesta-battleship/matchmaking/pkg/packets"
)

type Guild struct {
	Matchmaker actors.Matchmaker

	Hub actors.Actor
}

func (s *Guild) HandlePacket(senderId string, packet packets.Packet) {
	switch packet := packet.Body.(type) {
	case *packets.GuildWarStarted:
		s.handleCreateRoom(senderId, packet)
	case *packets.JoinSearch:
		s.handleJoinSearch(senderId, packet)
	case *packets.Disconnect:
		s.handleDisconnect(senderId, packet)
	default:
		log.Printf("Matchmaker %q: Got incorrect packet %T from %q", s.Matchmaker.Id(), packet, senderId)
	}
}

func (s *Guild) handleCreateRoom(senderId string, packet *packets.GuildWarStarted) {
	log.Printf("Matchmaker %q: Got packet %T from %q", s.Matchmaker.Id(), packet, senderId)
}

func (s *Guild) handleJoinSearch(senderId string, packet *packets.JoinSearch) {
	s.Matchmaker.AddToQueue(senderId)
}

func (s *Guild) handleDisconnect(senderId string, packet *packets.Disconnect) {
	s.Matchmaker.RemoveFromQueue(senderId)

	s.Hub.GetPacket(senderId, packets.Packet{SenderId: senderId, Body: packet})
}

func (s *Guild) OnExit() {
	s.Hub = nil
	s.Matchmaker = nil
}

func (s *Guild) String() string {
	return "Guild"
}
