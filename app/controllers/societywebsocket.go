package controllers

import (
	"golang.org/x/net/websocket"
	"github.com/revel/revel"
  "chaokaset-go/app/chatroom"
	"chaokaset-go/app/chatroomearth"
	"chaokaset-go/app/chatroomwater"
	"chaokaset-go/app/chatroompest"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) Room(user string) revel.Result {
	return c.Render(user)
}

func (c WebSocket) RoomSocket(user string, ws *websocket.Conn) revel.Result {
	// Join the room.
	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	chatroom.Join(user) //เมื่อ user เข้าฟังก์ชัน
	defer chatroom.Leave(user) //เมื่อ user จบฟังก์ชัน

	// Send down the archive. ข้อความเก่า
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroom.
	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			chatroom.Say(user, msg)
		}
	}
	return nil
}

func (c WebSocket) RoomEarth(user string, ws *websocket.Conn) revel.Result {
	// Join the room.
	subscription := chatroomearth.Subscribe()
	defer subscription.Cancel()

	chatroomearth.Join(user) //เมื่อ user เข้าฟังก์ชัน
	defer chatroomearth.Leave(user) //เมื่อ user จบฟังก์ชัน

	// Send down the archive. ข้อความเก่า
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroomearth.
	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			chatroomearth.Say(user, msg)
		}
	}
	return nil
}

func (c WebSocket) RoomWater(user string, ws *websocket.Conn) revel.Result {
	// Join the room.
	subscription := chatroomwater.Subscribe()
	defer subscription.Cancel()

	chatroomwater.Join(user) //เมื่อ user เข้าฟังก์ชัน
	defer chatroomwater.Leave(user) //เมื่อ user จบฟังก์ชัน

	// Send down the archive. ข้อความเก่า
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroomwater.
	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			chatroomwater.Say(user, msg)
		}
	}
	return nil
}

func (c WebSocket) RoomPest(user string, ws *websocket.Conn) revel.Result {
	// Join the room.
	subscription := chatroompest.Subscribe()
	defer subscription.Cancel()

	chatroompest.Join(user) //เมื่อ user เข้าฟังก์ชัน
	defer chatroompest.Leave(user) //เมื่อ user จบฟังก์ชัน

	// Send down the archive. ข้อความเก่า
	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	// Now listen for new events from either the websocket or the chatroompest.
	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// They disconnected.
				return nil
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return nil
			}

			// Otherwise, say something.
			chatroompest.Say(user, msg)
		}
	}
	return nil
}
