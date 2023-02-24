package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"project/internal/chat"
	pkg "project/pkg"
)

type chatHandler struct {
	usecase chat.Usecase
}

func writeInLogAndWriter(w http.ResponseWriter, message []byte) {
	log.Printf(string(message))
	w.Write(message)
}

func (u *chatHandler) GetChatHandler(w http.ResponseWriter, r *http.Request) {
	chatID, err := pkg.ParsingIdUrl(r, "chatID")

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	chat, err := u.usecase.GetChatById(chatID)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	jsonChat, err := json.Marshal(chat)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	writeInLogAndWriter(w, jsonChat)
}

func (u *chatHandler) DeleteChatHandler(w http.ResponseWriter, r *http.Request) {
	chatID, err := pkg.ParsingIdUrl(r, "chatID")

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	err = u.usecase.DeleteChatById(chatID)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	jsonError, err := json.Marshal(err)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	writeInLogAndWriter(w, jsonError)
}

func (u *chatHandler) GetAllChatsHandler(w http.ResponseWriter, r *http.Request) {
	allChats, err := u.usecase.GetAllChats()

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	jsonAllChats, err := json.Marshal(allChats)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	writeInLogAndWriter(w, jsonAllChats)
}

func (u *chatHandler) PostChatHandler(w http.ResponseWriter, r *http.Request) {
	newChat, err := u.usecase.CreateChat([]byte(""))

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	jsonNewChat, err := json.Marshal(newChat)

	if err != nil {
		writeInLogAndWriter(w, []byte(err.Error()))
		return
	}

	writeInLogAndWriter(w, jsonNewChat)
}

func NewChatHandler(r *mux.Router, us chat.Usecase) {
	handler := chatHandler{usecase: us}
	chatIdUrl := "/chat/{chatID:[0-9]+}"
	chatUrl := "/chat/"

	r.HandleFunc(chatIdUrl, handler.GetChatHandler).
		Methods("GET")
	r.HandleFunc(chatIdUrl, handler.DeleteChatHandler).
		Methods("DELETE")
	r.HandleFunc(chatUrl, handler.GetAllChatsHandler).
		Methods("GET")
	r.HandleFunc(chatUrl, handler.PostChatHandler).
		Methods("POST")
}
