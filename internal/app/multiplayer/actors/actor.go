package actors

import "github.com/lesta-battleship/matchmaking/pkg/packets"

type Actor interface {
	Id() string
	GetPacket(senderId string, packet packets.Packet)
	Start()
	Stop()
}

type Matchmaker interface {
	Actor
	CreateRoom() Actor
	ConnectToRoom(roomId, playerId string)
	AddToQueue(playerId string)
	RemoveFromQueue(playerId string)
}

type EventListener interface {
	Actor
	ConnectTo(matchmaker Matchmaker)
}
