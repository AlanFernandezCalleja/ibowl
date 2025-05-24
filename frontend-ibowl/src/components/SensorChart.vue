<template>
    <div class="dashboard-card chart-card">
        <div class="card-header">
            <h2><i class="fas fa-ruler-vertical"></i> Distancia (cm)</h2>
        </div>
        <div class="card-body">
            <canvas ref="chartCanvas" style="width: 100%; height: 300px;"></canvas>
        </div>
        <div class="card-footer">
            <div class="sensor-value">Valor actual: {{ currentValue.toFixed(2) }} cm</div>
        </div>
    </div>
</template>


<script>
import { Chart } from 'chart.js/auto'
import { nextTick } from 'vue'

export default {
    props: {
        chartData: {
            type: Object,
            required: true
        },
        currentValue: {
            type: Number,
            required: true
        }
    },
    data() {
        return {
            chartInstance: null
        }
    },
    watch: {
        chartData: {
            handler(newData) {
                this.safeUpdateChart(newData)
            },
            deep: true
        }
    },
    mounted() {
        this.initChart()
    },
    methods: {
        async initChart() {
            await nextTick() // Esperar a que el DOM esté listo

            try {
                const ctx = this.$refs.chartCanvas.getContext('2d')
                if (!ctx) throw new Error('Canvas context not available')

                this.chartInstance = new Chart(ctx, {
                    type: 'line',
                    data: this.chartData,
                    options: {
                        responsive: true,
                        maintainAspectRatio: false,
                        animation: {
                            duration: 0
                        },
                        scales: {
                            y: {
                                suggestedMin: -1,
                                suggestedMax: 5,
                                title: {
                                    display: true,
                                    text: 'Distancia (cm)'
                                }
                            },
                            x: {
                                title: {
                                    display: true,
                                    text: 'Tiempo'
                                },
                                ticks: {
                                    autoSkip: true,
                                    maxTicksLimit: 5 // Mostrar máximo 5 etiquetas en el eje X
                                }
                            }
                        },
                        plugins: {
                            legend: {
                                display: true,
                                position: 'top'
                            },
                            tooltip: {
                                mode: 'index',
                                intersect: false
                            }
                        },
                        elements: {
                            point: {
                                radius: 3, // Tamaño de los puntos
                                hoverRadius: 5 // Tamaño al hacer hover
                            }
                        }
                    }
                })
            } catch (error) {
                console.error('Error initializing chart:', error)
            }
        },
        safeUpdateChart(newData) {
            if (!this.chartInstance || !newData?.datasets?.length) return

            try {
                // Actualización segura del gráfico
                this.chartInstance.data.labels = [...newData.labels]
                this.chartInstance.data.datasets[0].data = [...newData.datasets[0].data]

                // Actualización optimizada
                requestAnimationFrame(() => {
                    if (this.chartInstance) {
                        this.chartInstance.update('none') // 'none' evita animaciones
                    }
                })
            } catch (error) {
                console.error('Error updating chart:', error)
            }
        }
    },
    beforeUnmount() {
        if (this.chartInstance) {
            this.chartInstance.destroy()
            this.chartInstance = null
        }
    }
}
</script>
