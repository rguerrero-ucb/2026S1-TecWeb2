// 1. Selección de elementos
const objetivo = document.getElementById('objetivo');
const marcador = document.getElementById('marcador');
localStorage.setItem('puntos', 0)

let persona =  {
    nombre: 'Ronal',
    edad: 25,
    profesion: 'Desarrollador'
}

localStorage.setItem('persona', JSON.stringify(persona))

persona = JSON.parse(localStorage.getItem('persona'))

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
objetivo.addEventListener('mouseover', () => {
    let puntos = parseInt(localStorage.getItem('puntos')) || 0; // Obtener puntos actuales
    puntos++;
    localStorage.setItem('puntos', puntos); // Guardar puntos en localStorage
    marcador.innerText = puntos; // Actualizar el texto en el HTML
    moverObjetivo();
});