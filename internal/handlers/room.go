package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	gguid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	gguid_id := gguid.New()
	return c.Redirect(fmt.Sprintf("/room/%s", gguid_id.String()))
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	if uuid == "" {
		 c.Status(fiber.StatusBadRequest).SendString("No room UUID provided")
		return nil
	}

	uuid, suuid, _ := createOrGetRoom(uuid)
}

func RoomWebsocket(c *websocket.Conn) {
	// Get the room UUID from the request
	uuid := c.Params("uuid")
	
	if uuid == "" {
		c.Close()
		return
	}

	_, _, room := createOrGetRoom(uuid)

	
}

func createOrGetRoom(uuid string) (string, string, *Room) {
	// Create a new room
	room := NewRoom(uuid)
	// Add the room to the global room list
	rooms[uuid] = room
	// Return the room
	return room.UUID, room.SUUID, room
}