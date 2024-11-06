<template>
  <div ref="MyChart" :class="props.ChartClass" />
</template>

<script setup>
import { onBeforeUnmount, ref } from 'vue';

const props = defineProps({
  ChartOptions: { Array, default: {} },
  ChartClass: { String, default: 'max-h-full min-h-96 w-full' },
});

const MyChart = ref();

let Chart; // 圖表實體

// 設定視窗縮放重繪圖表
const resizeChart = (options) => () => {
  Chart.setOption(options, true);
  Chart.resize();
};

const setChartOption = () => {
  if (Chart) window.echarts.dispose(Chart);

  Chart = window.echarts.init(MyChart.value); // 初始化圖表
  Chart.setOption(props.ChartOptions);
};
window.emitter.on('SetChartOption', setChartOption);
window.emitter.on('ResizeChart', resizeChart(props.ChartOptions));

onBeforeUnmount(() => {
  window.emitter.off('SetChartOption', setChartOption);
  window.emitter.off('ResizeChart', resizeChart);
});
</script>
