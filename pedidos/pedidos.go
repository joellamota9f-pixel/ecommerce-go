package pedidos

import "ecommerce-go/catalogo"

// Item representa un producto con la cantidad pedida.
type Item struct {
	Producto catalogo.Producto
	Cantidad int
}

// Estado indica en qué etapa está un pedido.
type Estado string

const (
	Pendiente Estado = "pendiente"
	Pagado    Estado = "pagado"
	Enviado   Estado = "enviado"
)

// Siguiente devuelve el estado que sigue al actual.
func Siguiente(e Estado) Estado {
	switch e {
	case Pendiente:
		return Pagado
	case Pagado:
		return Enviado
	default:
		return e
	}
}

// Subtotal suma el precio por la cantidad de cada item.
func Subtotal(items []Item) float64 {
	total := 0.0
	for _, it := range items {
		total += it.Producto.Precio * float64(it.Cantidad)
	}
	return total
}

// Regla es una función que recibe un total y devuelve otro total.
type Regla func(float64) float64

// Descuento devuelve una regla que resta un porcentaje al total.
func Descuento(porcentaje float64) Regla {
	return func(total float64) float64 {
		return total - total*porcentaje/100
	}
}

// IVA devuelve una regla que suma un porcentaje de impuesto al total.
func IVA(tasa float64) Regla {
	return func(total float64) float64 {
		return total + total*tasa/100
	}
}

// Componer une varias reglas en una sola, aplicándolas en orden.
func Componer(reglas ...Regla) Regla {
	return func(total float64) float64 {
		for _, r := range reglas {
			total = r(total)
		}
		return total
	}
}
