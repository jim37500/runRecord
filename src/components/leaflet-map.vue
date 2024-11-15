<template>
  <div class="w-full h-96">
    <l-map ref="map" v-model:zoom="Zoom" v-model:center="Center" :useGlobalLeaflet="false">
      <l-tile-layer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" layer-type="base" name="Stadia Maps Basemap"></l-tile-layer>
      <l-poly-line :lat-lngs="polylinePath" color="blue" :weight="3" />
    </l-map>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const { polyline, leaflet } = window;
const props = defineProps({
  Polyline: { String, default: '' },
  IsLongDistance: { Boolean, default: false },
});

const map = ref();

const Zoom = ref(props.IsLongDistance ? 10 : 14);
const polylinePath = ref(polyline.decode(props.Polyline));
const Center = ref(leaflet.latLngBounds(polylinePath.value).getCenter());
</script>
