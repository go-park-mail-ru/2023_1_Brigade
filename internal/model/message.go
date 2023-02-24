package model

type Message struct {
	Id             uint64 `json:"id"`
	Sender         User   `json:"sender"`
	Receiver       User   `json:"receiver"`
	DateOfDispatch string `json:"date_of_dispatch"`
	Text           string `json:"text"`
}
