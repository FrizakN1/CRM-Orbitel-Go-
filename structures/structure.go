package structures

type Session struct {
	Hash   string
	User   User
	Date   string
	Exists bool
}

type Setting struct {
	Address string
	Port    string
	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	DbName  string
}

type Role struct {
	ID   int
	Name string
}

type User struct {
	ID         int
	Login      string
	Password   string
	Name       string
	Role       Role
	Blocked    int
	Department Department
}

type Abonent struct {
	ID                int
	Name              string
	RegisteredAddress string
	ActualAddress     string
	IPAddress         string
	Phone             string
	ContractNumber    string
	PassportSeries    string
	PassportNumber    string
}

type Application struct {
	ID          int
	Abonent     Abonent
	Description string
	Notes       string
	Executor    User
	Status      Status
	Date        string
	Department  Department
	Priority    Priority
	Creator     User
}

type Status struct {
	ID   int
	Name string
}

type Priority struct {
	ID   int
	Name string
}

type Department struct {
	ID   int
	Name string
}

type Event struct {
	ID          int
	Name        string
	User        User
	Date        string
	Application Application
	Comment     string
}

type House struct {
	ID              int
	Name            string
	Internet        int
	TV              int
	Telephony       int
	NameMC          string
	AddressMC       string
	ChairmanName    string
	ChairmanContact string
	Agreement       int
	Power           float64
}
