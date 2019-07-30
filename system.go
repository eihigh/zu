package zu

var (
	systems []System
)

type System interface {
	Update()
}

type systemFunc func()

func (f systemFunc) Update() {
	f()
}

func PushSystemFunc(f func()) {
	systems = append(systems, systemFunc(f))
}

func PushSystems(sys ...System) {
	systems = append(systems, sys...)
}

func PopSystem() {
	systems = systems[:len(systems)-1]
}

func RemoveSystems(sys ...System) {

}
