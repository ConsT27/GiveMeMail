package api

type ApiInterface interface {
	Run() []string
	GetName() string
}
