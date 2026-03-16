const tblBody = document.querySelector("#tbl-body");
const btnActualizar = document.querySelector("#btn-actualizar");
const txtProducto = document.querySelector("#txt-producto");
const txtCantidad = document.querySelector("#txt-cantidad");
const btnAgregar = document.querySelector("#btn-agregar");

const data = [
    { producto: "Manzana", cantidad: 10 },
    { producto: "Banana", cantidad: 5 },
    { producto: "Naranja", cantidad: 8 }
];

const actualizarTabla = () => {
    tblBody.innerHTML = "";
    data.forEach((item) => {
        let row = `<tr>
            <td>${item.producto}</td>
            <td>${item.cantidad}</td>
        </tr>`;
        tblBody.innerHTML += row;
    });
}

btnActualizar.addEventListener("click", actualizarTabla);

btnAgregar.addEventListener("click", () => {
    const producto = txtProducto.value;
    const cantidad = parseInt(txtCantidad.value);
    if (producto && !isNaN(cantidad)) {
        data.push({ producto, cantidad });
        actualizarTabla();
        txtProducto.value = "";
        txtCantidad.value = "";
    } else {
        alert("Por favor, ingrese un producto y una cantidad válida.");
    }
});