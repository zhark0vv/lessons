package domain

// 2. Value Object - объект-значение.
// Объект-значение не имеет уникального идентификатора и не имеет жизненного цикла

type Address struct {
	Street  string
	City    string
	Country string
}

func NewAddress(street, city, country string) Address {
	return Address{
		Street:  street,
		City:    city,
		Country: country,
	}
}
