package interfaces

type WorldInterface interface {
	AddVillage(name string) error
	DestroyVillage(name string) error
}
