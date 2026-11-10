package reportes

import "ecommerce-go/catalogo"

// ProductosBajoStock devuelve los productos con stock igual o menor al límite.
func ProductosBajoStock(cat []catalogo.Producto, limite int) []catalogo.Producto {
	return catalogo.Filtrar(cat, func(p catalogo.Producto) bool {
		return p.Stock <= limite
	})
}

// ValorInventario calcula el valor total del inventario (precio x stock).
func ValorInventario(cat []catalogo.Producto) float64 {
	total := 0.0
	for _, p := range cat {
		total += p.Precio * float64(p.Stock)
	}
	return total
}
