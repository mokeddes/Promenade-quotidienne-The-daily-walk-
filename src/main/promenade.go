package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var s = true
var z = true
var act = true
var act2 = true
var wg sync.WaitGroup
var wg2 sync.WaitGroup

func fenetre(b string) {
	if s {
		s = false
		color(b)
		println("toutes les fenetre sont fermé par ", b)
		rand.Seed(time.Now().UTC().UnixNano())
		time.Sleep(time.Duration(rand.Intn(20)+3) * time.Second)
	}
}

func ventil(b string) {
	if z {
		z = false
		color(b)
		println("toutes les ventil sont fermé par ", b)
		rand.Seed(time.Now().UTC().UnixNano())
		time.Sleep(time.Duration(rand.Intn(20)+3) * time.Second)
	}
}

func preparation(b string) {
	color(b)
	fmt.Println(b, " a commencé ")
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Duration(rand.Intn(7)+3) * time.Second)
	color(b)
	fmt.Println(b, " met ses lunettes")
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Duration(rand.Intn(7)+3) * time.Second)
	color(b)
	fmt.Println(b, "met sa ceinture")
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Duration(rand.Intn(7)+3) * time.Second)
	// fermer la fenetre
	fenetre(b)
	ventil(b)
	color(b)
	fmt.Println(b, " met son tel et ses clés dans la poche")
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
}

func makeshoes(b string) {
	fmt.Println(b, " met ses chaussures")
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Duration(rand.Intn(7)+3) * time.Second)
}

func alice(wg *sync.WaitGroup, wg2 *sync.WaitGroup) {
	start := time.Now()
	preparation("alice")
	end := time.Now()
	z := end.Sub(start)
	color("Alice")
	fmt.Println("Alice a passée", z)
	wg.Done()
	alarme()
	color("Alice")
	start = time.Now()
	makeshoes("Alice")
	end = time.Now()
	z = end.Sub(start)
	color("Alice")
	fmt.Println("Alice a passée", z, " à mettre ses chaussures")
	wg2.Done()
	fermport()

}

func bob(wg *sync.WaitGroup, wg2 *sync.WaitGroup) {
	start := time.Now()
	preparation("bob")
	end := time.Now()
	z := end.Sub(start)
	color("bob")
	fmt.Println("bob a passé", z)
	wg.Done()
	alarme()
	color("bob")
	start = time.Now()
	makeshoes("bob")
	end = time.Now()
	z = end.Sub(start)
	color("bob")
	fmt.Println("Bob a passée", z, " à mettre ses chaussures")
	wg2.Done()
	fermport()

}

func alarme() {
	wg.Wait()
	if act == true {
		act = false
		fmt.Print("\x1B[37m")
		fmt.Println("Alarme en marche ! #CHIKOUR")
		fmt.Println("Verrouillage des portes dans 60 seconde")
		go compterebour()
	}
}
func fermport() {
	wg2.Wait()
	if act2 {
		act2 = false
		fmt.Print("\x1B[37m")
		fmt.Println("Fermeture des portes #On est Dehors !")

	}
}
func compterebour() {
	for i := 1; i < 7; i++ {
		time.Sleep(10 * time.Second)
		fmt.Print("\x1B[37m")
		println("il vous reste ", (6-i)*10, " secondes")
	}
	fmt.Print("\x1B[37m")
	println("Porte Verrouillée")
}

func color(moi string) {
	if moi == "bob" {
		fmt.Print("\x1B[34m")
	} else {
		fmt.Print("\x1B[32m")
	}
}

func main() {
	wg.Add(2)
	wg2.Add(2)
	fmt.Println("Allons nous premener")
	time.Sleep(time.Duration(rand.Intn(5)+3) * time.Second)
	go alice(&wg, &wg2)
	go bob(&wg, &wg2)
	time.Sleep(120 * time.Second)
}
