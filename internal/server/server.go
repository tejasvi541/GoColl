package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
	"github.com/tejasvi541/GoColl/internal/handlers"
	w "github.com/tejasvi541/GoColl/pkg/webrtc"
)

// Flag values

var (
	// ServerAddr is the address the server listens on.
	address = flag.String("addr" + ":", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()

	if *address == ":" {
		*address = ":8080"
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
		}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket", websocket.New(handlers.StreamWebsocket, 
		websocket.Config{HandshakeTimeout: 10 * time.Second}))
	app.Get("/stream/:ssuid/chat",websocket.New(handlers.StreamChatWebsocket))
	app.Get("/stream/:ssuid/viewer/websocket",websocket.New(handlers.StreamViewerWebsocket))
	app.Static("/", "./assests")

	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)

	go dispatchKeyFrames()
	if *cert != "" {
		return app.ListenTLS(*address, *cert, *key)
	}

	
}

func dispatchKeyFrames(){
		for range time.NeWTicker(3 * time.Second).C {
			for _, room := range w.Rooms {
				room.Peer.DispatchKeyFrame()
		}
	}
}