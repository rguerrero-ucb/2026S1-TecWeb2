
const btnAccion = document.getElementById("btnAccion");
const contenedor = document.getElementById("contenedor");

// const factorial = (n) => {
//     return (n === 1 || n === 0) ? 1 : n * factorial(n - 1);
// }

const tardar = (ms) => {
    let inicio = new Date().getTime();
    while (new Date().getTime() - inicio < ms) { }
}

const getPersona = async (id) => {
    return new Promise((resolve, reject) => {
        tardar(id);
        let obj = { id: 2000, nombre: "juan" }
        if (obj.id === id) {
            resolve(obj.nombre);
        } else {
            reject("No se encontró la persona");
        }
    });
}

btnAccion.addEventListener("click", async () => {
    console.log("FN1: Inicio de la función");
    getPersona(4000)
        .then((nombre) => {
            console.log("FN1: Fin funcion");
        })
        .catch((error) => {
            console.log("FN1: Error");
        });
    console.log("FN2: Inicio de la función");
    await getPersona(2000)
        .then((nombre) => {
            console.log("FN2: Fin función");
        })
        .catch((error) => {
            console.log("FN2: Error");
        });
});