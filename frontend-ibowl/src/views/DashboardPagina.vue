<template>
  <div class="dashboard-container">
    <header class="dashboard-header">
      <h1>Panel de Control</h1>
      <div class="header-actions">
        <button class="btn-refresh">
          <i class="fas fa-sync-alt"></i> Actualizar
        </button>
        <div class="user-menu">
          <i class="fas fa-user-circle"></i>
        </div>
      </div>
    </header>

    <div class="dashboard-grid">
      <!-- Sección de tarjetas resumen -->
      <section class="cards-section">
        <DashboardCards />
      </section>

      <!-- Gráfico principal -->
      <section class="main-chart">
        <div class="chart-header">
          <h2>Rendimiento Mensual</h2>
          <select class="chart-filter">
            <option>Últimos 7 días</option>
            <option>Este mes</option>
            <option>Este año</option>
          </select>
        </div>
        <BarChart />
      </section>

      <!-- Gráfico secundario -->
      <section class="secondary-chart">
        <div class="chart-header">
          <h2>Movimientos Recientes</h2>
        </div>
        <MovimientoChart />
      </section>

      <!-- Espacio para componentes adicionales -->
      <section class="additional-section">
        <h2>Actividad Reciente</h2>
        <!-- Aquí puedes agregar más componentes -->
      </section>
    </div>
  </div>
</template>

<script>
import DashboardCards from '@/components/DashboardCards.vue';
import BarChart from '@/components/BarChart.vue';
import MovimientoChart from '@/components/MovimientoChart.vue';

export default {
  name: 'DashboardPage',
  components: {
    DashboardCards,
    BarChart,
    MovimientoChart
  }
}
</script>

<style scoped>
.dashboard-container {
  min-height: 100vh;
  background-color: #f5f7fa;
  padding: 20px;
}

.dashboard-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  padding-bottom: 15px;
  border-bottom: 1px solid #e0e6ed;
}

.dashboard-header h1 {
  font-size: 1.8rem;
  color: #2c3e50;
  margin: 0;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 15px;
}

.btn-refresh {
  background-color: #ffffff;
  border: 1px solid #e0e6ed;
  border-radius: 6px;
  padding: 8px 15px;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  gap: 8px;
}

.btn-refresh:hover {
  background-color: #f8f9fa;
  border-color: #d3d9df;
}

.user-menu {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #ffffff;
  border-radius: 50%;
  border: 1px solid #e0e6ed;
  cursor: pointer;
  font-size: 1.2rem;
  color: #6e8efb;
}

.dashboard-grid {
  display: grid;
  grid-template-columns: repeat(12, 1fr);
  gap: 20px;
}

.cards-section {
  grid-column: span 12;
}

.main-chart {
  grid-column: span 8;
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  padding: 20px;
}

.secondary-chart {
  grid-column: span 4;
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  padding: 20px;
}

.additional-section {
  grid-column: span 12;
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  padding: 20px;
  min-height: 200px;
}

.chart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.chart-header h2 {
  font-size: 1.2rem;
  color: #2c3e50;
  margin: 0;
}

.chart-filter {
  padding: 8px 12px;
  border-radius: 6px;
  border: 1px solid #e0e6ed;
  background-color: #ffffff;
  font-size: 0.9rem;
}

/* Responsive design */
@media (max-width: 1200px) {
  .main-chart {
    grid-column: span 12;
  }
  
  .secondary-chart {
    grid-column: span 12;
  }
}

@media (max-width: 768px) {
  .dashboard-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 15px;
  }
  
  .header-actions {
    width: 100%;
    justify-content: space-between;
  }
}
</style>