// 1. Selección de elementos (Traer el HTML al mundo JS)
const objetivo = document.getElementById('objetivo');
const marcador = document.getElementById('marcador');
let puntos = 0;

// 2. Definir la lógica de movimiento
function moverObjetivo() {
    const maxX = 350; // Ancho escenario - ancho objetivo
    const maxY = 350;

    // Generar coordenadas aleatorias
    const x = Math.floor(Math.random() * maxX);
    const y = Math.floor(Math.random() * maxY);

    // Manipulación directa de estilos (DOM)
    objetivo.style.left = x + 'px';
    objetivo.style.top = y + 'px';
}

// 3. El Evento (La escucha de la acción del usuario)
objetivo.addEventListener('click', () => {
    puntos++;
    marcador.innerText = puntos; // Actualizar el texto en el HTML
    moverObjetivo();
});