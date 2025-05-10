package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Car struct {
	Number string
	Slot   int
}

type ParkingLot struct {
	capacity int
	slots    map[int]*Car
	carMap   map[string]int
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		slots:    make(map[int]*Car),
		carMap:   make(map[string]int),
	}
}

func (p *ParkingLot) Park(carNumber string) {
	for i := 1; i <= p.capacity; i++ {
		if _, occupied := p.slots[i]; !occupied {
			p.slots[i] = &Car{Number: carNumber, Slot: i}
			p.carMap[carNumber] = i
			fmt.Printf("Allocated slot number: %d\n", i)
			return
		}
	}
	fmt.Println("Sorry, parking lot is full")
}

func (p *ParkingLot) Leave(carNumber string, hours int) {
	if slot, found := p.carMap[carNumber]; found {
		delete(p.slots, slot)
		delete(p.carMap, carNumber)
		charge := 10
		if hours > 2 {
			charge += (hours - 2) * 10
		}
		fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", carNumber, slot, charge)
	} else {
		fmt.Printf("Registration number %s not found\n", carNumber)
	}
}

func (p *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")
	var slots []int
	for slot := range p.slots {
		slots = append(slots, slot)
	}
	sort.Ints(slots)
	for _, slot := range slots {
		fmt.Printf("%d %s\n", slot, p.slots[slot].Number)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var parkingLot *ParkingLot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		command := strings.ToLower(parts[0])

		switch command {
		case "create_parking_lot":
			n, _ := strconv.Atoi(parts[1])
			parkingLot = NewParkingLot(n)
		case "park":
			if parkingLot != nil {
				parkingLot.Park(parts[1])
			}
		case "leave":
			if parkingLot != nil {
				hours, _ := strconv.Atoi(parts[2])
				parkingLot.Leave(parts[1], hours)
			}
		case "status":
			if parkingLot != nil {
				parkingLot.Status()
			}
		}
	}
}
