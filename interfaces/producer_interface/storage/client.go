package storage

type Client struct {
	Id      int
	Name    string
	Surname string
	Age     int
	Address string
}

type ClientStorage interface {
	GetAllClients() ([]Client, error)
	GetClientsByAge(int) ([]Client, error)
	GetClient(int) (Client, error)
	RemoveClient(int) error
	UpdateClient(Client) error
	CreateClient(Client) error
	// ...
}
