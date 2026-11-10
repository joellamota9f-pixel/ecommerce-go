# Sistema de Gestion de E-commerce en Go

Aprendizaje Autonomo 1: Seleccion de Sistemas de Gestion Empresarial.

**Autor:** Pedro Joel Lamota Mena.
**Institucion:** Universidad Internacional del Ecuador (UIDE).
**Materia:** Programacion Orientada a Objetos.
**Docente:** Ing. Carlos Andres Arevalo Torres.

## Objetivo

Gestionar el catalogo, los pedidos y los reportes de una tienda en linea mediante un programa de consola desarrollado en Go con programacion funcional.

## Modulos

| Modulo | Paquete | Funcionalidad |
|---|---|---|
| Catalogo | `catalogo` | Alta de productos, busqueda por nombre o ID y control de stock |
| Pedidos | `pedidos` | Subtotal, descuento, IVA y estados del pedido |
| Reportes | `reportes` | Productos con stock bajo y valor del inventario |

## Estructura del proyecto

```
ecommerce-go/
├── go.mod
├── main.go
├── catalogo/
│   ├── catalogo.go
│   └── catalogo_test.go
├── pedidos/
│   ├── pedidos.go
│   └── pedidos_test.go
├── reportes/
│   └── reportes.go
└── docs/
```

## Programacion funcional aplicada

- **Funciones puras:** `Agregar` y `DescontarStock` devuelven un catalogo nuevo sin modificar el original.
- **Funciones como valores:** el tipo `Regla` es una funcion que transforma un total.
- **Funciones de orden superior:** `Filtrar` recibe una condicion como parametro.
- **Composicion:** `Componer` une varias reglas (por ejemplo, descuento mas IVA) en una sola.

## Paquetes utilizados

Solo se usan paquetes de la libreria estandar de Go. No hay paquetes de terceros.

| Paquete | Uso |
|---|---|
| `fmt` | Entrada y salida por consola |
| `strings` | Busqueda de texto sin distinguir mayusculas |
| `math` | Comparacion de decimales en las pruebas |
| `testing` | Pruebas automaticas |

## Como ejecutarlo

Requiere tener Go instalado.

```
git clone https://github.com/joellamota9f-pixel/ecommerce-go.git
cd ecommerce-go
go run .
```

## Pruebas

```
go test ./...
```

## Limitaciones

- Los datos se guardan solo en memoria y se pierden al cerrar el programa.
- No incluye pagos reales ni interfaz web.
- Los nombres de producto deben ser de una sola palabra por el uso de `fmt.Scan`.