package domain3

type Service interface {
	Domain1Get() error
	Domain1Post() error
	Domain1Patch() error
	Domain1Delete() error
}
