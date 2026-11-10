package catalogo

import "testing"

func TestAgregarNoModificaOriginal(t *testing.T) {
	original := []Producto{{ID: 1, Nombre: "Laptop", Precio: 650, Stock: 5}}
	nuevo := Agregar(original, Producto{ID: 2, Nombre: "Mouse", Precio: 15.5, Stock: 20})

	if len(original) != 1 {
		t.Errorf("el catálogo original cambió: tiene %d productos", len(original))
	}
	if len(nuevo) != 2 {
		t.Errorf("el catálogo nuevo debería tener 2 productos, tiene %d", len(nuevo))
	}
}

func TestDescontarStockNoModificaOriginal(t *testing.T) {
	original := []Producto{{ID: 1, Nombre: "Laptop", Precio: 650, Stock: 5}}
	nuevo := DescontarStock(original, 1, 2)

	if original[0].Stock != 5 {
		t.Errorf("el stock original cambió a %d", original[0].Stock)
	}
	if nuevo[0].Stock != 3 {
		t.Errorf("el stock nuevo debería ser 3, es %d", nuevo[0].Stock)
	}
}
