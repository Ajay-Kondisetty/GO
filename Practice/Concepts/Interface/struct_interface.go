package main

import "fmt"

type BaseTest struct {
}

type test interface {
	getID() int
	getName() string
	getAge() int
	getAddress() Address
}

type Address struct {
	Street  string
	Zipcode int
	City    string
	State   string
	Country string
}
type Worker struct {
	ID      int
	Name    string
	Age     int
	Address Address
}

type Manager struct {
	ID      int
	Name    string
	Age     int
	Address Address
	Workers map[int]Worker
}

func (w *Worker) getID() int {
	return w.ID
}

func (w *Worker) getName() string {
	return w.Name
}

func (w *Worker) getAge() int {
	return w.Age
}

func (w *Worker) getAddress() Address {
	return w.Address
}

func (m *Manager) getID() int {
	return m.ID
}

func (m *Manager) getName() string {
	return m.Name
}

func (m *Manager) getAge() int {
	return m.Age
}

func (m *Manager) getAddress() Address {
	return m.Address
}

func (m *Manager) getWorkers() map[int]Worker {
	return m.Workers
}

func (m *Manager) getWorker(workerID int) Worker {
	if worker, ok := m.Workers[workerID]; ok {
		return worker
	}

	return Worker{}
}

func (m *Manager) setWorker(worker Worker) map[int]Worker {
	if len(m.Workers) == 0 {
		m.Workers = make(map[int]Worker)
	}
	m.Workers[worker.ID] = worker

	return m.Workers
}
func main() {
	newWorker := new(Worker)
	newWorker.ID = 1
	newWorker.Age = 27
	newWorker.Name = "Name1"
	newWorker.Address = Address{
		Street:  "test street",
		City:    "test city",
		Zipcode: 123456,
		State:   "test state",
		Country: "test country",
	}
	fmt.Println(newWorker.ID, newWorker.Name, newWorker.Age, newWorker.Address)

	newManger := new(Manager)
	newManger.ID = 1
	newManger.Age = 37
	newManger.Name = "Man1"
	newManger.Address = Address{
		Street:  "man test street",
		City:    "man test city",
		Zipcode: 63452,
		State:   "man test state",
		Country: "man test country",
	}
	_ = newManger.setWorker(*newWorker)

	fmt.Println(newManger.ID, newManger.Name, newManger.Age, newManger.Address)
	fmt.Println(newManger.getWorker(newWorker.ID))
	fmt.Println(newManger.getWorkers())

	w := Worker{
		Name: "AAAAAA",
		ID:   2,
		Age:  22,
		Address: Address{
			Street:  "street",
			Zipcode: 99099,
			Country: "USA",
			State:   "New york",
			City:    "London",
		},
	}

	m := Manager{
		Name: "BBBBB",
		ID:   3,
		Age:  33,
		Address: Address{
			Street:  "street1",
			Zipcode: 99889,
			Country: "Canada",
			State:   "New jersey",
			City:    "BZA",
		},
		Workers: map[int]Worker{
			1: *newWorker,
			2: w,
		},
	}

	t := new(BaseTest)
	t.testing(&w)
	t.testing(&m)
	fmt.Println(m.getWorker(2))
	fmt.Println(m.getWorkers())
}

func (bs *BaseTest) testing(ti test) {
	fmt.Println(ti.getID(), ti.getName(), ti.getAge(), ti.getAddress())
}
