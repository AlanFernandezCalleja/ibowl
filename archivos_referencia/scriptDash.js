// Configuración inicial del gráfico
const distanciaCtx = document.getElementById('distanciaChart').getContext('2d');
const distanciaChart = new Chart(distanciaCtx, {
    type: 'line',
    data: {
        labels: [], // Las etiquetas de tiempo se añadirán dinámicamente
        datasets: [{
            label: 'Distancia (cm)',
            data: [], // Los datos se añadirán dinámicamente
            borderColor: '#3498db',
            backgroundColor: 'rgba(52, 152, 219, 0.1)',
            borderWidth: 2,
            fill: true,
            tension: 0.4
        }]
    },
    options: {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
            y: {
                beginAtZero: true,
                title: {
                    display: true,
                    text: 'Centímetros (cm)'
                }
            },
            x: {
                title: {
                    display: true,
                    text: 'Tiempo'
                }
            }
        },
        plugins: {
            legend: {
                position: 'top',
            },
            tooltip: {
                mode: 'index',
                intersect: false,
            }
        },
        interaction: {
            mode: 'nearest',
            axis: 'x',
            intersect: false
        }
    }
});

function alimentar() {
    alert("Comando de alimentación enviado al dispositivo.");
    // Aquí puedes agregar lógica para enviar una solicitud a tu backend PHP
    // fetch("tu-api/alimentar.php").then(...)
}


// Función para obtener datos de la API
function fetchData() {
    fetch('http://localhost/ibowl/archivos_referencia/latest_data.php')
        .then(response => {
            if (!response.ok) {
                throw new Error('Error en la respuesta de la API');
            }
            return response.json();
        })
        .then(data => {
            // Actualizar gráfico de distancia
            const now = new Date().toLocaleTimeString();
            
            // Añadir nuevo dato
            distanciaChart.data.labels.push(now);
            distanciaChart.data.datasets[0].data.push(data.distancia);
            
            // Mantener solo los últimos 20 puntos para mejor visualización
            const maxDataPoints = 20;
            if (distanciaChart.data.labels.length > maxDataPoints) {
                distanciaChart.data.labels.shift();
                distanciaChart.data.datasets[0].data.shift();
            }
            
            // Actualizar gráfico
            distanciaChart.update();
            
            // Actualizar valor actual
            document.getElementById('currentDistance').textContent = `Valor actual: ${data.distancia} cm`;
            document.getElementById('lastUpdate').textContent = `Última actualización: ${now}`;
            
            // Actualizar estado de movimiento
            updateMovementStatus(data.movimiento);
        })
        .catch(error => {
            console.error('Error al obtener datos:', error);
            // Opcional: Mostrar mensaje de error en la interfaz
            document.getElementById('lastUpdate').textContent = 'Error al obtener datos';
            document.getElementById('movimientoText').textContent = 'Error en conexión';
            document.getElementById('movimientoIndicator').classList.add('alert');
            document.getElementById('movimientoIndicator').classList.remove('on', 'off');
        });
}

// Función para actualizar el estado de movimiento
function updateMovementStatus(detected) {
    const indicator = document.getElementById('movimientoIndicator');
    const text = document.getElementById('movimientoText');
    const detectionTime = document.getElementById('lastDetection');
    
    if (detected) {
        indicator.classList.remove('off', 'alert');
        indicator.classList.add('on');
        text.textContent = 'Movimiento detectado!';
        detectionTime.textContent = `Última detección: ${new Date().toLocaleTimeString()}`;
    } else {
        indicator.classList.remove('on', 'alert');
        indicator.classList.add('off');
        text.textContent = 'Inactivo';
    }
}

// Configurar intervalo de actualización
const updateInterval = setInterval(fetchData, 2000); // 2 segundos

// Cargar datos inmediatamente al cargar la página
fetchData();