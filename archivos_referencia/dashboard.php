<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <title>Dashboard de Sensores</title>
  <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
  <style>
    body {
      font-family: Arial, sans-serif;
      margin: 30px;
      background-color: #f4f4f4;
    }
    h2 {
      color: #333;
    }
    canvas {
      background-color: white;
      border-radius: 10px;
      box-shadow: 0 4px 8px rgba(0,0,0,0.1);
      margin-bottom: 30px;
    }
  </style>
</head>
<body>
  <h2>📊 Nivel de Comida (cm)</h2>
  <canvas id="comidaChart" width="600" height="300"></canvas>

  <h2>🚶 Movimiento Detectado</h2>
  <canvas id="movimientoChart" width="600" height="300"></canvas>

  <script>
    const comidaChartCtx = document.getElementById('comidaChart').getContext('2d');
    const movimientoChartCtx = document.getElementById('movimientoChart').getContext('2d');

    let comidaChart = new Chart(comidaChartCtx, {
      type: 'line',
      data: {
        labels: [],
        datasets: [{
          label: 'Distancia en cm',
          data: [],
          backgroundColor: 'rgba(0, 123, 255, 0.2)',
          borderColor: 'rgba(0, 123, 255, 1)',
          borderWidth: 2,
          tension: 0.4
        }]
      },
      options: { responsive: true, scales: { y: { beginAtZero: true } } }
    });

    let movimientoChart = new Chart(movimientoChartCtx, {
      type: 'bar',
      data: {
        labels: [],
        datasets: [{
          label: 'Movimiento (1=SÍ, 0=NO)',
          data: [],
          backgroundColor: 'rgba(255, 99, 132, 0.6)',
          borderColor: 'rgba(255, 99, 132, 1)',
          borderWidth: 1
        }]
      },
      options: { responsive: true, scales: { y: { beginAtZero: true, max: 1 } } }
    });

    async function actualizarDatos() {
      const res = await fetch('datos.php');
      const data = await res.json();

      comidaChart.data.labels = data.map(row => row.fecha);
      comidaChart.data.datasets[0].data = data.map(row => row.comida_cm);
      comidaChart.update();

      movimientoChart.data.labels = data.map(row => row.fecha);
      movimientoChart.data.datasets[0].data = data.map(row => row.movimiento);
      movimientoChart.update();
    }

    setInterval(actualizarDatos, 5000);
    actualizarDatos();
  </script>
</body>
</html>
