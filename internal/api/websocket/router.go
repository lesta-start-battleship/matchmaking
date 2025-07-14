package websocket

import (
	"github.com/lesta-battleship/matchmaking/internal/api/websocket/middlewares"
	"github.com/lesta-battleship/matchmaking/internal/app/multiplayer"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router gin.IRouter,
	ApiGatewayrUrl string,
	websocketServer *WebsocketServer,
	engine *multiplayer.Engine,
) {
	addCors := middlewares.AddCorsMiddleware(ApiGatewayrUrl)
	handleError := middlewares.HandleErrorMiddleware()
	requireToken := middlewares.CheckTokenMiddleware()

	randomHandler := HandleGetJoinRandom(websocketServer, engine)
	rankedHandler := HandleGetJoinRanked(websocketServer, engine)
	customHandler := HandleGetJoinCustom(websocketServer, engine)

	router.Use(handleError, addCors, requireToken)
	router.GET("/matchmaking/random", randomHandler)
	router.GET("/matchmaking/ranked", rankedHandler)
	router.GET("/matchmaking/custom", customHandler)
}
