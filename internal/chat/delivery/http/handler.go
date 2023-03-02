package http

//type chatHandler struct {
//	usecase chat.Usecase
//}
//
//func (u *chatHandler) GetChatHandler(w http.ResponseWriter, r *http.Request) {
//
//	//chatID := http_utils.ParsingIdUrl(r, "chatID")
//	//response := u.usecase.GetChatById(context.Background(), chatID)
//	//http_utils.SendJsonResponse(w, response)
//
//	//if err != nil {
//	//	w.WriteHeader(http.StatusInternalServerError)
//	//	return
//	//}
//	//
//	//var response []byte
//	//err = json.Unmarshal(response, user)
//	//
//	//if err != nil {
//	//	w.WriteHeader(http.StatusInternalServerError)
//	//	return
//	//}
//
//	//w.WriteHeader(http.StatusOK)
//	//w.Write(response)
//
//	//chatID, err := http_utils.ParsingIdUrl(r, "chatID")
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//
//	//chat, err := u.usecase.GetChatById(r.Context(), chatID)
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//
//	//jsonChat, err := json.Marshal(chat)
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//
//	//w.Write(jsonChat)
//}
//
//func (u *chatHandler) DeleteChatHandler(w http.ResponseWriter, r *http.Request) {
//
//	//chatID := http_utils.ParsingIdUrl(r, "chatID")
//	//response := u.usecase.DeleteChatById(context.Background(), chatID)
//	//http_utils.SendJsonResponse(w, response)
//
//	//chatID, err := http_utils.ParsingIdUrl(r, "chatID")
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//
//	//err = u.usecase.DeleteChatById(r.Context(), chatID)
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//
//	//jsonError, err := json.Marshal(err)
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//
//	//w.Write(jsonError)
//
//}
//
//func (u *chatHandler) GetAllChatsHandler(w http.ResponseWriter, r *http.Request) {
//	//
//	//response := u.usecase.GetAllChats(context.Background())
//	//http_utils.SendJsonResponse(w, response)
//
//	//allChats, err := u.usecase.GetAllChats(r.Context())
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//
//	//jsonAllChats, err := json.Marshal(allChats)
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//	//
//	//w.Write(jsonAllChats)
//}
//
//func (u *chatHandler) PostChatHandler(w http.ResponseWriter, r *http.Request) {
//
//	//response := u.usecase.CreateChat(context.Background(), r.Body)
//	//http_utils.SendJsonResponse(w, response)
//
//	//fmt.Println("POST HANDLER")
//	//chatik := model.Chat{1, "vanya", "2 nov", nil, nil}
//	//
//	//ch, _ := json.Marshal(chatik)
//	//
//	//err := u.usecase.CreateChat(r.Context(), ch)
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//
//	//jsonNewChat, err := json.Marshal(newChat)
//	//
//	//if err != nil {
//	//	w.Write([]byte(err.Error()))
//	//	return
//	//}
//
//	//w.Write([]byte(""))
//}
//
//func NewChatHandler(r *mux.Router, us chat.Usecase) {
//	handler := chatHandler{usecase: us}
//	chatIdUrl := "/chat/{chatID:[0-9]+}"
//	chatUrl := "/chat/"
//
//	r.HandleFunc(chatIdUrl, handler.GetChatHandler).
//		Methods("GET")
//	r.HandleFunc(chatIdUrl, handler.DeleteChatHandler).
//		Methods("DELETE")
//	r.HandleFunc(chatUrl, handler.GetAllChatsHandler).
//		Methods("GET")
//	r.HandleFunc(chatUrl, handler.PostChatHandler).
//		Methods("POST")
//}
