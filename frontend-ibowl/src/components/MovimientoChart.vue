<template>
  <div class="card card-movimiento-style">
    <h3><i class="fas fa-ruler-combined"></i> Detección de Movimiento (Simulado)</h3>
    <div class="chart-wrapper">
      <canvas ref="movimientoChartCanvas"></canvas>
    </div>
    <p class="status-text">Actualizando cada 3 segundos...</p>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';
import {
  Chart,
  LineController,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Legend,
  Title,
  Tooltip,Filler 
} from 'chart.js';

// Registrar componentes de Chart.js necesarios para el gráfico de líneas
Chart.register(
  LineController,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Legend,
  Title,
  Tooltip,Filler 
);

const movimientoChartCanvas = ref(null);
let myChart = null;
let updateInterval = null;

const MAX_DATA_POINTS = 20; // Número máximo de puntos de datos a mostrar

// Datos iniciales y configuración del gráfico
const chartData = {
  labels: [], // Las etiquetas se generarán dinámicamente (tiempo)
  datasets: [{
    label: 'Distancia (cm)',
    data: [], // Los datos se generarán dinámicamente
    borderColor: 'rgb(75, 192, 192)',
    backgroundColor: 'rgba(75, 192, 192, 0.2)',
    tension: 0.3, // Para suavizar la línea
    fill: true, // Rellenar el área bajo la línea
    pointRadius: 4,
    pointBackgroundColor: 'rgb(75, 192, 192)',
  }]
};

const chartOptions = {
  filler: {
      propagate: true // Opciones del plugin
    },
  responsive: true,
  maintainAspectRatio: false,
  scales: {
    y: {
      beginAtZero: true,
      min: 0, // Mínimo del eje Y
      max: 11, // Máximo del eje Y (un poco más de 10 para margen)
      title: {
        display: true,
        text: 'Distancia (cm)',
        color: '#333',
        font: {
          size: 14
        }
      },
      grid: {
        color: 'rgba(200, 200, 200, 0.3)'
      },
      ticks: {
        color: '#555',
        stepSize: 1 // Incremento de 1 en 1
      }
    },
    x: {
      max:6,
      title: {
        display: true,
        text: 'Tiempo',
        color: '#333',
        font: {
          size: 14
        }
      },
      grid: {
        display: false
      },
      ticks: {
        color: '#555',
        maxRotation: 0, // Evitar rotación de etiquetas si son muchas
        autoSkip: true, // Saltar algunas etiquetas si hay demasiadas
        maxTicksLimit: 10 // Mostrar un máximo de 10 etiquetas en el eje X
      }
    }
  },
  plugins: {
    legend: {
      display: true,
      position: 'top',
      labels: {
        color: '#333',
        boxWidth: 6,
        padding: 15
      }
    },
    tooltip: {
      mode: 'index',
      intersect: false,
      backgroundColor: 'rgba(0,0,0,0.7)',
      titleColor: '#fff',
      bodyColor: '#fff',
    }
  }
};

// Función para generar un valor de distancia aleatorio entre 1 y 10
function getRandomDistance() {
  return Math.floor(Math.random() * 10) + 1;
}

function getLastDistance(){
  
}

// Función para actualizar los datos del gráfico
function updateChartData() {
  if (!myChart) return;

  const now = new Date();
  // Formato simple para la etiqueta de tiempo
  const newLabel = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}:${now.getSeconds().toString().padStart(2, '0')}`;
  const newDataPoint = getRandomDistance();

  // Añadir nuevo dato y etiqueta
  myChart.data.labels.push(newLabel);
  myChart.data.datasets[0].data.push(newDataPoint);

  // Mantener solo MAX_DATA_POINTS en el gráfico
  if (myChart.data.labels.length > MAX_DATA_POINTS) {
    myChart.data.labels.shift(); // Elimina la etiqueta más antigua
    myChart.data.datasets[0].data.shift(); // Elimina el dato más antiguo
  }

  myChart.update(); // 'quiet' o 'none' para evitar animaciones si son muchas actualizaciones
}

onMounted(() => {
  // Llenar con algunos datos iniciales para que no empiece vacío
  for (let i = 0; i < MAX_DATA_POINTS / 2; i++) {
    const pastTime = new Date(Date.now() - (MAX_DATA_POINTS / 2 - i) * 3000);
    chartData.labels.push(`${pastTime.getHours().toString().padStart(2, '0')}:${pastTime.getMinutes().toString().padStart(2, '0')}:${pastTime.getSeconds().toString().padStart(2, '0')}`);
    chartData.datasets[0].data.push(getRandomDistance());
  }

  if (movimientoChartCanvas.value) {
    myChart = new Chart(
      movimientoChartCanvas.value.getContext('2d'),
      {
        type: 'line',
        data: chartData,
        options: chartOptions
      }
    );
  }

  // Iniciar la actualización automática
  updateInterval = setInterval(updateChartData, 1500);
});

onBeforeUnmount(() => {
  if (updateInterval) {
    clearInterval(updateInterval);
  }
  if (myChart) {
    myChart.destroy();
  }
});
</script>

<style scoped>
/* Estilos base de la tarjeta (puedes copiarlos de BarChart.vue o DashboardCards.vue si quieres centralizarlos) */
.card {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  padding: 20px;
  color: #333;
  transition: transform 0.2s ease-in-out;
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 700px; /* Ancho máximo para la tarjeta del gráfico */
  margin: 20px auto; /* Centrar la tarjeta del gráfico */
}

.card:hover {
  transform: translateY(-5px);
}

.card h3 {
  margin-top: 0;
  margin-bottom: 15px;
  font-size: 1.2em;
  display: flex;
  align-items: center;
  color: inherit;
  width: 100%;
  justify-content: flex-start;
}

.card h3 i {
  margin-right: 10px;
  font-size: 1.3em;
}

/* Estilo específico para la tarjeta del gráfico de movimiento */
.card-movimiento-style {
  /* Ejemplo de gradiente - puedes elegir otro o un color sólido */
  background: linear-gradient(135deg, #89f7fe 0%, #66a6ff 100%);
  color: #2c3e50;
}
.card-movimiento-style h3 i {
  color: #1e90ff; /* Azul dodger para el icono */
}

.chart-wrapper {
  position: relative;
  width: 100%;
  height: 350px; /* Altura para el gráfico de líneas */
  margin-bottom: 10px; /* Menos espacio antes del texto de estado */
}

canvas {
  width: 100% !important;
  height: 100% !important;
}

.status-text {
  font-size: 0.9em;
  color: #555;
  margin-top: 10px;
}

/* Asegúrate de tener Font Awesome si usas las clases fas */
/* Ejemplo: <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css"> en index.html */
</style>