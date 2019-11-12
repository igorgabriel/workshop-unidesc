package models

// Area ...
type Area struct {
	ID  int    `json:"id"`
	Nm  string `json:"name"`
	Pid int    `json:"produtorId"`
	Tp  string `json:"type"`
	Ca  string `json:"createdAt"`
	SID string `json:"stationId"`
}

// Estagio ...
type Estagio struct {
	ID   int     `json:"id"`
	Nr   int     `json:"numero"`
	Nm   string  `json:"nome"`
	Dias int     `json:"dias"`
	Mm   float64 `json:"mm"`
}

// Produtor ...
type Produtor struct {
	ID    int    `json:"Id"`
	TId   string `json:"telegramId"`
	Uname string `json:"username"`
	Pwd   string `json:"password"`
	FName string `json:"firstName"`
	LName string `json:"lastName"`
	Farm  string `json:"farm"`
	City  string `json:"city"`
	Est   string `json:"estate"`
}

// Cultivar ...
type Cultivar struct {
	ID         int       `json:"id"`
	PID        int       `json:"produtorId"`
	Nome       string    `json:"nome"`
	Fabricante string    `json:"fabricante"`
	Variacao   string    `json:"variacao"`
	Cultivar   int       `json:"cultivar"`
	Estagios   []Estagio `json:"estagios"`
}

// Lavoura ...
type Lavoura struct {
	ID        int      `json:"id"`
	Area      Area     `json:"area"`
	DPlantio  string   `json:"dtPlantio"`
	DColheita string   `json:"dtColheita"`
	Cultivar  Cultivar `json:"cultivar"`
}

// Irrigacao ...
type Irrigacao struct {
	ID          int     `json:"id"`
	CreatedAt   string  `json:"createdAt"`
	DtIrrigacao string  `json:"dtIrrigacao"`
	IsChuva     bool    `json:"isChuva"`
	LavouraID   int     `json:"lavouraId"`
	VlQtdMm     float64 `json:"vlQtdMm"`
}

// IrrigationAmount ...
type IrrigationAmount struct {
	LavouraID   int     `json:"lavouraId"`
	NmArea      string  `json:"nomeArea"`
	NrEstagio   int     `json:"numeroEstagio"`
	NmEstagio   string  `json:"nomeEstagio"`
	Acumulado   float64 `json:"acumulado"`
	Esperado    float64 `json:"esperado"`
	Previsao    float64 `json:"previsao"`
	DtIrrigacao string  `json:"dtIrrigacao"`
}

// UserMessageRequest ...
type UserMessageRequest struct {
	UserID int `json:"userId"`
}

type Barra struct {
	Estagio string  `json:"estagio"`
	Ideal   float64 `json:"ideal"`
	Atual   float64 `json:"atual"`
	Started bool    `json:"started"`
}

type RelativeValue struct {
	Value    float64 `json:"value"`
	Relative float64 `json:"relative"`
}

type DashboardInfo struct {
	Barras     []Barra       `json:"barras"`
	Irrigacoes []Irrigacao   `json:"irrigacoes"`
	Progress   float64       `json:"progress"`
	Water      RelativeValue `json:"water"`
	Rain       RelativeValue `json:"rain"`
	Forecast   RelativeValue `json:"forecast"`
}

//Credentials ...
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//SignInError ...
type SignInError struct {
	Code  int
	Error error
}

//HeaderAuthorization ...
type HeaderAuthorization struct {
	Authorization string `header:"Authorization"`
}

//User ...
type User struct {
	Name   string `json:"name"`
	Farm   string `json:"farm"`
	City   string `json:"city"`
	Estate string `json:"estate"`
	Token  string `json:"token"`
	Role   string `json:"role"`
}
