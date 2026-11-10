package main

import (
	"fmt"

	"ecommerce-go/catalogo"
	"ecommerce-go/pedidos"
	"ecommerce-go/reportes"
)

// listar muestra los productos en consola.
func listar(cat []catalogo.Producto) {
	if len(cat) == 0 {
		fmt.Println("No hay productos para mostrar.")
		return
	}
	for _, p := range cat {
		fmt.Printf("%d | %s | $%.2f | stock: %d\n", p.ID, p.Nombre, p.Precio, p.Stock)
	}
}

// agregarProducto recibe un apuntador al catálogo para poder
// reemplazarlo por la nueva versión que devuelve catalogo.Agregar.
func agregarProducto(cat *[]catalogo.Producto) {
	var p catalogo.Producto
	fmt.Print("ID: ")
	fmt.Scan(&p.ID)
	fmt.Print("Nombre (una sola palabra): ")
	fmt.Scan(&p.Nombre)
	fmt.Print("Precio: ")
	fmt.Scan(&p.Precio)
	fmt.Print("Stock: ")
	fmt.Scan(&p.Stock)
	*cat = catalogo.Agregar(*cat, p)
	fmt.Println("Producto agregado.")
}

// crearPedido arma un pedido, calcula el total con reglas y avanza el estado.
func crearPedido(cat *[]catalogo.Producto) {
	items := []pedidos.Item{}
	for {
		var id, cantidad int
		fmt.Print("ID del producto (0 para terminar): ")
		fmt.Scan(&id)
		if id == 0 {
			break
		}
		prod, existe := catalogo.BuscarPorID(*cat, id)
		if !existe {
			fmt.Println("Ese producto no existe.")
			continue
		}
		fmt.Print("Cantidad: ")
		fmt.Scan(&cantidad)
		if cantidad <= 0 || cantidad > prod.Stock {
			fmt.Println("Cantidad no válida o sin stock suficiente.")
			continue
		}
		items = append(items, pedidos.Item{Producto: prod, Cantidad: cantidad})
		*cat = catalogo.DescontarStock(*cat, id, cantidad)
	}

	if len(items) == 0 {
		fmt.Println("El pedido está vacío.")
		return
	}

	subtotal := pedidos.Subtotal(items)
	calcular := pedidos.IVA(15)
	if subtotal > 100 {
		calcular = pedidos.Componer(pedidos.Descuento(10), pedidos.IVA(15))
		fmt.Println("Se aplicó 10% de descuento por compra mayor a $100.")
	}

	estado := pedidos.Pendiente
	fmt.Printf("Subtotal: $%.2f | Total con IVA: $%.2f | Estado: %s\n", subtotal, calcular(subtotal), estado)
	estado = pedidos.Siguiente(estado)
	fmt.Println("Estado actualizado a:", estado)
}

func main() {
	cat := []catalogo.Producto{
		{ID: 1, Nombre: "Laptop", Precio: 650.00, Stock: 5},
		{ID: 2, Nombre: "Mouse", Precio: 15.50, Stock: 20},
		{ID: 3, Nombre: "Teclado", Precio: 35.00, Stock: 2},
	}

	var opcion int
	for {
		fmt.Println("\n===== TIENDA E-COMMERCE =====")
		fmt.Println("1. Agregar producto")
		fmt.Println("2. Listar catálogo")
		fmt.Println("3. Buscar producto")
		fmt.Println("4. Crear pedido")
		fmt.Println("5. Ver productos con stock bajo")
		fmt.Println("6. Ver valor del inventario")
		fmt.Println("0. Salir")
		fmt.Print("Opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			agregarProducto(&cat)
		case 2:
			listar(cat)
		case 3:
			var texto string
			fmt.Print("Buscar: ")
			fmt.Scan(&texto)
			listar(catalogo.Buscar(cat, texto))
		case 4:
			crearPedido(&cat)
		case 5:
			listar(reportes.ProductosBajoStock(cat, 3))
		case 6:
			fmt.Printf("Valor total del inventario: $%.2f\n", reportes.ValorInventario(cat))
		case 0:
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
