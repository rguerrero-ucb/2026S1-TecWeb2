const suma = (a, b) => a + b;
const resta = (a, b) => a - b;
const multiplicacion = (a, b) => a * b;

// const duplicar = a => a * 2;

// console.log(suma(2,3));
// console.log(duplicar(5));

// const sumarYDuplicar = (a, b) => duplicar(suma(a, b));

// function sumarYDuplicar2(a, b) {
//     const suma = suma(a, b);
//     return duplicar(suma);
// }

function operar(a, b, operacion) {
    return operacion(a, b);
}

console.log(operar(2, 3, suma));
console.log(operar(2, 3, resta));
console.log(operar(2, 3, multiplicacion));