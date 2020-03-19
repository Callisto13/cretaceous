package dinosaur

type Theropod interface {
	Roar() error
	Run() error
	Eat(prey Sauropod) (bool, error)
}

type Sauropod interface {
	Escape() error
	Eat() (bool, error)
}
