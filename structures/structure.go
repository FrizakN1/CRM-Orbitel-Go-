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

type User struct {
	ID         int
	Login      string
	Password   string
	Name       string
	Role       int
	Blocked    int
	Department int
}

type Abonent struct {
	ID             int
	Name           string
	Address        string
	Phone          string
	ContractNumber string
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
