function calcular() {
    const operando1 = document.getElementById("operando1").value;
    const operando2 = document.getElementById("operando2").value;

    const operador = document.querySelector("input[name='operador']:checked").value;

    let resultado;

    switch (operador) {
        case "+":
            resultado = parseFloat(operando1) + parseFloat(operando2);
            break;
        case "-":
            resultado = parseFloat(operando1) - parseFloat(operando2);
            break;
        case "*":
            resultado = parseFloat(operando1) * parseFloat(operando2);
            break;
        case "/":
            if (parseFloat(operando2) !== 0) {
                resultado = parseFloat(operando1) / parseFloat(operando2);
            } else {
                resultado = "Error: División por cero";
            }
            break;
        default:
            resultado = "Operador no válido";
    }

    document.getElementById("resultado").innerText = resultado;
}