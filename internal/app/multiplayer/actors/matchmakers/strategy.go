package matchmakers

import (
	"fmt"

	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer/actors/matchmakers/strategies"
	"github.com/lesta-battleship/matchmaking/pkg/packets"
)

type Strategy interface {
	HandlePacket(senderId string, packet packets.Packet)
	OnExit()

	fmt.Stringer
}

var strategiesMap = map[MatchType]func(*Matchmaker){
	RandomMatch: SetRandom,
	RankedMatch: SetRanked,
	GuildMatch:  SetGuild,
	CustomMatch: SetCustom,
}

func SetStrategy(matchmaker *Matchmaker, matchType MatchType) {
	strategiesMap[matchType](matchmaker)
}

func SetRandom(matchmaker *Matchmaker) {
	matchmaker.ChangeStrategy(&strategies.Random{Matchmaker: matchmaker, Hub: matchmaker.hub, Queue: matchmaker.queue})
}

func SetRanked(matchmaker *Matchmaker) {
	matchmaker.ChangeStrategy(&strategies.Ranked{Matchmaker: matchmaker, Hub: matchmaker.hub, Queue: matchmaker.queue})
}

func SetGuild(matchmaker *Matchmaker) {
	matchmaker.ChangeStrategy(&strategies.Guild{Matchmaker: matchmaker, Hub: matchmaker.hub})
}

func SetCustom(matchmaker *Matchmaker) {
	matchmaker.ChangeStrategy(&strategies.Custom{Matchmaker: matchmaker, Hub: matchmaker.hub})
}
