package schemas

type PushEmail struct {
	To      string
	Subject string
	Body    string
}

type ResponseExample struct {
	Code    string
	Status  string
	Message string
	Data    interface{}
}
