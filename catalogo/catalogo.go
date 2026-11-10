package catalogo

import "strings"

// Producto representa un artículo de la tienda.
type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

// Agregar devuelve un NUEVO catálogo sin modificar el original.
func Agregar(cat []Producto, p Producto) []Producto {
	nuevo := make([]Producto, len(cat), len(cat)+1)
	copy(nuevo, cat)
	return append(nuevo, p)
}

// Filtrar es una función de orden superior: recibe una condición como parámetro.
func Filtrar(cat []Producto, cond func(Producto) bool) []Producto {
	resultado := []Producto{}
	for _, p := range cat {
		if cond(p) {
			resultado = append(resultado, p)
		}
	}
	return resultado
}

// Buscar devuelve los productos cuyo nombre contiene el texto indicado.
func Buscar(cat []Producto, texto string) []Producto {
	return Filtrar(cat, func(p Producto) bool {
		return strings.Contains(strings.ToLower(p.Nombre), strings.ToLower(texto))
	})
}

// BuscarPorID devuelve el producto y un booleano que indica si existe.
func BuscarPorID(cat []Producto, id int) (Producto, bool) {
	for _, p := range cat {
		if p.ID == id {
			return p, true
		}
	}
	return Producto{}, false
}

// DescontarStock devuelve un nuevo catálogo con el stock actualizado.
func DescontarStock(cat []Producto, id, cantidad int) []Producto {
	nuevo := make([]Producto, len(cat))
	for i, p := range cat {
		if p.ID == id {
			p.Stock -= cantidad
		}
		nuevo[i] = p
	}
	return nuevo
}
