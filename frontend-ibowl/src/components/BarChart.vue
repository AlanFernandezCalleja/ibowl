<template>
    <div class="chart-container">
      <canvas ref="chartCanvas"></canvas>
      <button @click="toggleAutoUpdate" class="update-button">
        {{ autoUpdate ? 'Detener' : 'Iniciar' }} actualización
      </button>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue';
  import { 
    Chart, 
    BarController, 
    CategoryScale, 
    LinearScale, 
    BarElement,
    Legend,
    Title,
    Tooltip
  } from 'chart.js';
  
  // Registra componentes de Chart.js
  Chart.register(
    BarController, 
    CategoryScale, 
    LinearScale, 
    BarElement,
    Legend,
    Title,
    Tooltip
  );
  
  const chartCanvas = ref(null);
  let myChart = null;
  const autoUpdate = ref(false);
  let updateInterval = null;
  
  // Datos iniciales
  const initialData = {
    labels: ['Enero', 'Febrero', 'Marzo', 'Abril', 'Mayo'],
    datasets: [{
      label: 'Ventas 2023',
      backgroundColor: [
        'rgba(255, 99, 132, 0.7)',
        'rgba(54, 162, 235, 0.7)',
        'rgba(255, 206, 86, 0.7)',
        'rgba(75, 192, 192, 0.7)',
        'rgba(153, 102, 255, 0.7)'
      ],
      borderColor: [
        'rgba(255, 99, 132, 1)',
        'rgba(54, 162, 235, 1)',
        'rgba(255, 206, 86, 1)',
        'rgba(75, 192, 192, 1)',
        'rgba(153, 102, 255, 1)'
      ],
      borderWidth: 1
    }]
  };
  
  // Opciones del gráfico
  const chartOptions = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top',
      },
      title: {
        display: true,
        text: 'Ventas Mensuales (Actualización cada 3s)'
      },
      tooltip: {
        callbacks: {
          title: (context) => `Mes: ${context[0].label}`,
          label: (context) => `Valor: ${context.raw}`
        }
      }
    },
    scales: {
      y: {
        beginAtZero: true,
        max: 100 // Límite máximo del eje Y
      }
    }
  };
  
  // Función para generar datos aleatorios
  function generateRandomData() {
    return initialData.labels.map(() => 
      Math.floor(Math.random() * 90) + 10 // Valores entre 10 y 100
    );
  }
  
  // Actualiza los datos del gráfico
  function updateChartData() {
    if (myChart) {
      myChart.data.datasets[0].data = generateRandomData();
      myChart.update(); 
    }
  }
  
  // Alternar actualización automática
  function toggleAutoUpdate() {
    autoUpdate.value = !autoUpdate.value;
    
    if (autoUpdate.value) {
      updateInterval = setInterval(updateChartData, 1000); 
      updateChartData(); 
    } else {
      clearInterval(updateInterval);
    }
  }
  
  onMounted(() => {
    // Inicializa el gráfico con datos
    initialData.datasets[0].data = generateRandomData();
    
    myChart = new Chart(
      chartCanvas.value.getContext('2d'),
      {
        type: 'bar',
        data: initialData,
        options: chartOptions
      }
    );
  });
  
  onBeforeUnmount(() => {
    // Limpieza al desmontar el componente
    if (myChart) {
      myChart.destroy();
    }
    if (updateInterval) {
      clearInterval(updateInterval);
    }
  });
  </script>
  
  <style scoped>
/* Estilos copiados y adaptados de DashboardCards.vue */
.card {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  padding: 20px;
  color: #333;
  transition: transform 0.2s ease-in-out;
  display: flex; /* Para organizar el título, gráfico y botón verticalmente */
  flex-direction: column; /* Dirección vertical */
  align-items: center; /* Centrar contenido horizontalmente */
  max-width: 600px; /* Ancho máximo para la tarjeta del gráfico */
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
  color: inherit; /* Hereda el color de .card o .card-chart-style */
  width: 100%; /* Ocupa todo el ancho para alineación */
  justify-content: flex-start; /* Alinea el título a la izquierda */
}

.card h3 i {
  margin-right: 10px;
  font-size: 1.3em;
}

/* Estilo específico para la tarjeta del gráfico, puedes elegir uno de los card-colorX */
.card-chart-style {
  /* Usaremos un gradiente similar al card-color3 como ejemplo */
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  color: #2c3e50; /* Color de texto para este fondo */
}
.card-chart-style h3 i { 
  color: #3498db; /* Un color de icono que combine */
}


.chart-wrapper {
  position: relative;
  width: 100%; /* El wrapper ocupa todo el ancho de la tarjeta */
  height: 300px; /* Define una altura para el gráfico o ajústala según necesites */
  margin-bottom: 20px; /* Espacio antes del botón */
}

/* El canvas debe llenar el chart-wrapper */
canvas {
  width: 100% !important;
  height: 100% !important;
}

.update-button {
  margin-top: 0; /* Ajustado ya que el chart-wrapper tiene margin-bottom */
  padding: 10px 20px; /* Un poco más grande */
  background-color: #3498db; /* Color que combine con card-chart-style */
  color: white;
  border: none;
  border-radius: 5px; /* Bordes más redondeados */
  cursor: pointer;
  transition: background-color 0.3s, box-shadow 0.3s;
  font-weight: bold;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.update-button:hover {
  background-color: #2980b9;
  box-shadow: 0 4px 8px rgba(0,0,0,0.15);
}

/* Asegúrate de tener Font Awesome si usas las clases fas */
/* Puedes añadir el CDN en tu index.html:
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css">
*/
</style>