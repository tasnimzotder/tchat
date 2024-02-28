package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/tasnimzotder/tchat/server/models"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (s *ServerAPI) messageHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Failed to close connection: %v", err)
			return
		}

		// remove connection
		delete(s.Connections, conn.RemoteAddr().String())
	}(conn)

	// read message
	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			break
		}

		var msg models.Message
		err = json.Unmarshal(msgBytes, &msg)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			break
		}

		// store connection
		s.Connections[conn.RemoteAddr().String()] = ConnStruct{
			Conn:   conn,
			UserID: msg.SenderID,
		}

		// ping message
		if msg.MessageType == "init" {
			log.Printf("Init message received from %s", msg.SenderID)
			continue
		}

		//	send message to recipient if online
		recipientConn, err := s.getRecipientConn(msg.RecipientID)
		if err != nil {
			log.Printf("Failed to get recipient connection: %v", err)
			continue
		}

		err = sendWSMessage(recipientConn, msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			continue
		}
	}
}

func sendWSMessage(recipientConn ConnStruct, msg models.Message) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = recipientConn.Conn.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	return nil
}

func (s *ServerAPI) getRecipientConn(recipientID string) (ConnStruct, error) {
	//	return where UserID is recipientID
	for _, v := range s.Connections {
		if v.UserID == recipientID {
			return v, nil
		}
	}

	//	return empty ConnStruct and error
	return ConnStruct{}, errors.New("recipient is not online")
}
