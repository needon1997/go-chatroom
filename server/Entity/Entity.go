package Entity

const (
	LOGIN = iota
	REGISTER
	MSG
	NOTICE
	OFFLINE
	ONLINE
	ONLINE_MAP
)
const (
	OK        = 20001
	ERROR     = 20002
	NO_ACCESS = 20003
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}
type Message struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}
type LoginResponse struct {
	UserInfo  User              `json:"userinfo"`
	OnlineMap map[string]string `json:"onlineMap"`
}
type ResultType struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
	Data   string `json:"Data"`
}
