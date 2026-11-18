package main

import "fmt"

// Item ist die gemeinsame Schnittstelle für Blätter und Composite-Elemente.
type Item interface {
	Price() float64
	Display()
}

// printable ist eine interne Erweiterung für die rekursive Darstellung.
// Sie bleibt außerhalb der öffentlichen Item-Schnittstelle verborgen.
type printable interface {
	Item
	display(indent int)
}

// printLine kapselt die gemeinsame Ausgabe aller Elemente.
func printLine(indent int, name string, ownPrice, totalPrice float64) {
	fmt.Printf("%s%s: Einzelpreis %.2f €, Gesamtpreis %.2f €\n",
		indentString(indent), name, ownPrice, totalPrice)
}

func indentString(level int) string {
	return fmt.Sprintf("%*s", level*2, "")
}

// Glass ist ein Blatt im Composite Pattern und besitzt keinen weiteren Inhalt.
type Glass struct {
	Name  string
	price float64
}

func (g *Glass) Price() float64 {
	return g.price
}

func (g *Glass) Display() {
	g.display(0)
}

func (g *Glass) display(indent int) {
	printLine(indent, g.Name, g.price, g.Price())
}

// Frame ist ebenfalls ein Blatt mit einem individuellen Preis.
type Frame struct {
	Name  string
	price float64
}

func (f *Frame) Price() float64 {
	return f.price
}

func (f *Frame) Display() {
	f.display(0)
}

func (f *Frame) display(indent int) {
	printLine(indent, f.Name, f.price, f.Price())
}

// Window ist ein Composite und kann beliebig viele Items enthalten.
// Die []Item-Einträge speichern Pointer-Referenzen auf die Originalobjekte.
type Window struct {
	Name  string
	items []Item
}

func (w *Window) Add(item Item) {
	w.items = append(w.items, item)
}

// Price berechnet den Gesamtpreis rekursiv über alle enthaltenen Items.
func (w *Window) Price() float64 {
	total := 0.0
	for _, item := range w.items {
		total += item.Price()
	}
	return total
}

func (w *Window) Display() {
	w.display(0)
}

func (w *Window) display(indent int) {
	// Ein Composite hat keinen eigenen Einzelpreis; sein Gesamtpreis ist die Summe.
	printLine(indent, w.Name, 0, w.Price())
	for _, item := range w.items {
		// Alle konkreten Item-Typen in diesem Prototyp unterstützen die interne Darstellung.
		item.(printable).display(indent + 1)
	}
}

func main() {
	// Demo-Daten: Ein Fenster enthält Rahmen, Glas und ein weiteres Fenster.
	leftWing := &Window{Name: "Linker Fensterflügel"}
	leftWing.Add(&Frame{Name: "Holzrahmen links", price: 145.00})
	leftWing.Add(&Glass{Name: "Dreifachverglasung links", price: 210.50})

	window := &Window{Name: "Panoramafenster"}
	window.Add(&Frame{Name: "Außenrahmen", price: 320.00})
	window.Add(&Glass{Name: "Dreifachverglasung", price: 480.75})
	window.Add(leftWing)

	fmt.Println("Fensterstruktur:")
	window.Display()
	fmt.Printf("\nEndpreis: %.2f €\n", window.Price())
}
