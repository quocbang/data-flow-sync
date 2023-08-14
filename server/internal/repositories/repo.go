package repositories

type Repositories interface {
	Close() error
	Services
}
