package pedidos

import (
	"math"
	"testing"
)

func casiIgual(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

func TestComponer(t *testing.T) {
	calcular := Componer(Descuento(10), IVA(15))
	esperado := 103.5 // 100 - 10% = 90; 90 + 15% = 103.5
	if resultado := calcular(100); !casiIgual(resultado, esperado) {
		t.Errorf("esperado %.2f, obtenido %.2f", esperado, resultado)
	}
}

func TestSiguiente(t *testing.T) {
	if Siguiente(Pendiente) != Pagado {
		t.Error("pendiente debe pasar a pagado")
	}
	if Siguiente(Pagado) != Enviado {
		t.Error("pagado debe pasar a enviado")
	}
	if Siguiente(Enviado) != Enviado {
		t.Error("enviado no debe cambiar")
	}
}
