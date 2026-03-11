function saludar() {
    const titulo = document.getElementById("titulo");
    const valor1 = document.getElementById("valor1");
    const valor2 = document.getElementById("valor2");


    titulo.innerText = "Hola Mundo " + valor1.value;
    titulo.style.color = "red";
    titulo.style.backgroundColor = "yellow";

    valor2.value = valor1.value;

}