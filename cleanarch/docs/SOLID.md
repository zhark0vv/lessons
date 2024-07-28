
# Принципы SOLID

## Таблица принципов

| Инициал | Представляет                | Название                                        | Понятие                                                                                                                                                               |
|---------|-----------------------------|-------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| S       | SRP                         | Принцип единственной ответственности (single responsibility principle) | У класса должна быть только одна причина для изменения. Как подметил Р. Мартин: «Модуль должен отвечать за одного и только за одного актора.»                                                               |
| O       | OCP                         | Принцип открытости/закрытости (open-closed principle) | Программные сущности должны быть открыты для расширения, но закрыты для модификации.                                                                                  |
| L       | LSP                         | Принцип подстановки Лисков (Liskov substitution principle) | Функции, которые используют базовый тип, должны иметь возможность использовать подтипы базового типа, не зная об этом.                                                |
| I       | ISP                         | Принцип разделения интерфейса (interface segregation principle) | Много интерфейсов, специально предназначенных для клиентов, лучше, чем один интерфейс общего назначения.                                                              |
| D       | DIP                         | Принцип инверсии зависимостей (dependency inversion principle) | Зависимость на Абстракциях. Нет зависимости на что-то конкретное.                                                                                                     |

## Реализация принципов SOLID в Go

### Принцип единственной ответственности (SRP)
Принцип единственной ответственности реализуется путем разделения кода на небольшие, высокоспециализированные пакеты и структуры. Каждый пакет или структура должна отвечать за одну конкретную задачу или функциональность.

**Пример:**
```go
package user

// domain.go
type User struct {
    ID    int
    Name  string
    Email string
}

// service.go

type UserService struct{}

func (u *UserService) Save(u *User) error {
    // Code to save user
    return nil
}
```

### Принцип открытости/закрытости (OCP)
Для реализации этого принципа в Go, можно использовать интерфейсы и структуры.
Интерфейсы позволяют добавлять новые функциональности без изменения существующего кода.

**Пример:**
```go
type PaymentProcessor interface {
    ProcessPayment(amount float64) error
}

type CreditCardProcessor struct{}

func (c *CreditCardProcessor) ProcessPayment(amount float64) error {
    // Process credit card payment
    return nil
}

type PayPalProcessor struct{}

func (p *PayPalProcessor) ProcessPayment(amount float64) error {
    // Process PayPal payment
    return nil
}
```

### Принцип подстановки Лисков (LSP)
Принцип подстановки Лисков в Go реализуется с помощью интерфейсов.
Подтипы должны соответствовать интерфейсу базового типа, чтобы их можно было использовать взаимозаменяемо.

**Пример:**
```go
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func CalculateArea(s Shape) float64 {
    return s.Area()
}
```

### Принцип разделения интерфейса (ISP)
Принцип разделения интерфейса реализуется путем создания небольших интерфейсов, которые предназначены для конкретных клиентов, вместо одного общего интерфейса.

**Пример:**
```go
type Printer interface {
    Print() error
}

type Scanner interface {
    Scan() error
}

type MultiFunctionDevice interface {
    Printer
    Scanner
}

type OfficePrinter struct{}

func (o OfficePrinter) Print() error {
    // Print document
    return nil
}

func (o OfficePrinter) Scan() error {
    // Scan document
    return nil
}
```

### Принцип инверсии зависимостей (DIP)
Принцип инверсии зависимостей реализуется путем зависимостей на абстракциях (интерфейсах), а не на конкретных реализациях.

**Пример:**
```go
type Database interface {
    Connect() error
}

type MySQLDatabase struct{}

func (m MySQLDatabase) Connect() error {
    // Connect to MySQL database
    return nil
}

type App struct {
    DB Database
}

func (a *App) Initialize(db Database) {
    a.DB = db
}

func main() {
    var db Database = MySQLDatabase{}
    app := App{}
    app.Initialize(db)
    app.DB.Connect()
}
```

Следуя этим принципам, можно создавать гибкие, поддерживаемые и масштабируемые программы на Go.
