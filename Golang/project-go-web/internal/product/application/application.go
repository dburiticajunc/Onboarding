package application

type Application interface {
	SetUp() (err error)
	Run() (err error)
}
