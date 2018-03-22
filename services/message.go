package services

type Message struct {
	SenderId 		string					`json:"mac"`
	UserHash 		string					`json:"mac"`
	Type 			string					`json:"type"`
	Payload 		interface{}				`json:"payload"`
}

func NewMessage(si string, t string, p interface{}) *Message {
	m := new(Message)
	m.SenderId = si
	m.Type = t
	m.Payload = p

	return m
}
