package main

import (
	"fmt"
	"math"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Order struct {
	UserEmail string
	Products  []Product
}

// single Responsibility Principle (SRP):A class or module should have only one reason to change

func (o *Order) CalculateTotalPrice() float64 {
	var total float64

	for _, product := range o.Products {

		total += product.Price * float64(product.Quantity)
	}
	return total
}

func (o *Order) NotifyCustomer() {
	fmt.Println("Your Order received successfully")

}

// Open/Closed Principle (OCP):open for extension but closed for modification

type Shape interface {
	Area() float64
}

type Rect struct {
	Width  float64
	Height float64
}

func (r Rect) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radious float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radious * c.Radious
}

// open for exteention

func CalculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0

	for _, shape := range shapes {
		totalArea += shape.Area()

	}
	return totalArea
}

// Liskov Substitution Principle (LSP):derived classes maintain the behavior expected from the base class.

func PrintArea(s Shape) {
	fmt.Printf("LSP: Total Area:%.2f\n", s.Area())

}

//Interface Segregation Principle (ISP):Clients should not be forced to depend on interfaces they do not use.

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}
type Multifunction interface {
	Printer
	Scanner
}

type SimplePrinter struct{}
type SimpleScanner struct{}

func (s *SimplePrinter) Print() {

	fmt.Println("ISP: Printing...............")
}

func (s *SimpleScanner) Scan() {
	fmt.Println("ISP: Scanning...............")
}

// Dependency Inversion Principle (DIP):High-level modules should not depend on low-level modules.

// MessageSender interface serves as an abstraction that defines the contract for sending messages
type MessageSender interface {
	SendMessage(message string)
}

type EmailSender struct{}

func (e *EmailSender) SendMessage(message string) {
	fmt.Println("Sending email:", message)
}

type SmsSender struct{}

func (s *SmsSender) SendMessage(message string) {
	fmt.Println("Send Message:", message)
}

//Instead of depending directly on concrete implementations of message senders (such as EmailSender or SMSsender), the NotificationService depends on the abstract MessageSender interface.
//High-Level Module (NotificationService)

type NotificationService struct {
	sender MessageSender
}

func (n *NotificationService) SendNotification(message string) {
	n.sender.SendMessage(message)
}

func main() {

	Product1 := Product{Name: "Rahul", Quantity: 2, Price: 1050}

	Order := Order{
		UserEmail: "rahulreza920",
		Products:  []Product{Product1},
	}
	//SRP
	fmt.Println(Order.CalculateTotalPrice())
	Order.NotifyCustomer()

	//OCP
	Shape := []Shape{
		Rect{Width: 2, Height: 4},
		Circle{Radious: 3},
	}

	totalArea := CalculateTotalArea(Shape)
	fmt.Printf("Total Area:%.2f\n", totalArea)

	//LSP:objects of different subclasses (Rectangle and Circle) to be used interchangeably through their shared superclass interface (Shape).

	Rect := &Rect{Width: 3, Height: 4}
	circle := &Circle{Radious: 5}
	PrintArea(Rect)
	PrintArea(circle)

	//ISP:creating smaller, specific interfaces rather than a single large one.

	printer := SimplePrinter{}
	scanner := SimpleScanner{}

	printer.Print()
	scanner.Scan()

	//DIP

	email := &EmailSender{}

	sms := &SmsSender{}

	emailNotification := &NotificationService{sender: email}
	smsNotification := &NotificationService{sender: sms}

	emailNotification.SendNotification("DIP: Hello, this is an email notification!")
	smsNotification.SendNotification("DIP: Hello, this is an email notification!")
}
