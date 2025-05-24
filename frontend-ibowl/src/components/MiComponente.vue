<template>
    <div>
        <h2>{{ datosJson.titulo }}</h2>
        <ul>
            <li v-for="item in datosJson.items" :key="item.id">
                {{  item.nombre }} - {{ item.precio }}
            </li>
        </ul>
        <p>Ultima Actualizacion {{ datosJson.metadata.fechaActualizacion }}</p>
    </div>
</template>
<script setup>

    
    defineProps(
        {
            datosJson : {
                type: Object,
                required: true,
            }
        }
    )

</script>