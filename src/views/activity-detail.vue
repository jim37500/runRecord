<template>
  <div v-if="IsLoading" class="flex justify-center items-center w-full h-full">
    <i class="pi pi-spin pi-spinner" style="font-size: 2rem" />
  </div>
  <div v-else class="w-full">
    <h2 class="text-3xl md:text-4xl font-semibold text-center text-blue-800">
      {{ ActivityData.Name }}
    </h2>
    <div class="w-full grid grid-cols-1 xl:grid-cols-2 gap-8 px-2 py-8 sm:p-8">
      <div class="xl:col-span-2 bg-zinc-50 border-2 border-zinc-100 px-8 py-8 rounded-2xl z-0">
        <div class="font-semibold text-2xl">詳細資料</div>
        <div class="grid sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 mt-5 ml-3">
          <div v-if="ActivityService.SportTypes[ActivityData.SportType]" class="font-semibold text-lg">運動類型： {{ ActivityService.SportTypes[ActivityData.SportType].Name }}</div>
          <div class="font-semibold text-lg">日期： {{ moment(new Date(ActivityData.Date)).format('yyyy-MM-DD') }}</div>
          <div v-if="ActivityData.AverageTemperature" class="font-semibold text-lg">平均溫度： {{ ActivityData.AverageTemperature }} °C</div>
          <div class="font-semibold text-lg">
            距離：
            {{ ActivityService.CalculateDistance(ActivityData.Distance).toFixed(1) }} km
          </div>
          <div class="font-semibold text-lg">
            平均速度：
            {{ ActivityService.CalculateSpeed(ActivityData.AverageSpeed, ActivityData.SportType) }}
            {{ ActivityService.SportTypes[ActivityData.SportType].Unit }}
          </div>
          <div class="font-semibold text-lg">移動時間： {{ ActivityService.CalculateTime(ActivityData.MovingTime) }}</div>
          <div class="font-semibold text-lg">經過時間： {{ ActivityService.CalculateTime(ActivityData.ElapsedTime) }}</div>
          <div class="font-semibold text-lg">爬升： {{ ActivityData.TotalElevationGain }} m</div>
          <div class="font-semibold text-lg">
            最大速度：
            {{ ActivityService.CalculateSpeed(ActivityData.MaxSpeed, ActivityData.SportType) }}
            {{ ActivityService.SportTypes[ActivityData.SportType].Unit }}
          </div>
          <div class="font-semibold text-lg">平均心率： {{ ActivityData.AverageHeartrate }} bpm</div>
          <div class="font-semibold text-lg">最大心率： {{ ActivityData.MaxHeartrate }} bpm</div>
          <div v-if="ActivityData.AverageCadence" class="font-semibold text-lg">平均踏頻： {{ ActivityData.AverageCadence }} spm</div>
          <div v-if="ActivityData.AverageWatts" class="font-semibold text-lg">平均功率： {{ ActivityData.AverageWatts }} W</div>
          <div v-if="ActivityData.MaxWatts" class="font-semibold text-lg">最大功率： {{ ActivityData.MaxWatts }} W</div>
        </div>
      </div>
      <div class="xl:col-span-2 bg-zinc-50 border-2 border-zinc-100 p-3 rounded-2xl z-0">
        <LeafletMap v-if="ActivityData.Polyline" :Polyline="ActivityData.Polyline" :IsLongDistance="ActivityData.SportType === 'Ride'" />
      </div>
      <div class="xl:col-span-2 bg-zinc-50 border-2 border-zinc-100 px-8 py-8 rounded-2xl z-0">
        <primevue-data-table :value="ActivityData.Laps" tableClass="border-zinc-900" :pt="{ tableContainer: { class: '!rounded-lg' } }">
          <template #empty>找不到資料!</template>
          <primevue-column field="Distance" header="圈數">
            <template #body="slotProps">
              {{ slotProps.index + 1 }}
            </template>
          </primevue-column>
          <primevue-column field="Distance" header="距離 (km)">
            <template #body="slotProps">
              {{ (Math.round(slotProps.data.Distance / 100) / 10).toFixed(2) }}
            </template>
          </primevue-column>
          <primevue-column field="AverageSpeed" :header="ActivityData.SportType === 'Ride' ? '平均速度 (km/hr)' : `平均配速 (${ActivityService.SportTypes[ActivityData.SportType].Unit})`">
            <template #body="slotProps">
              {{ ActivityService.CalculateSpeed(slotProps.data.AverageSpeed, ActivityData.SportType) }}
            </template>
          </primevue-column>
          <primevue-column field="MovingTime" header="移動時間 (min)">
            <template #body="slotProps">
              {{ ActivityService.CalculateTime(slotProps.data.MovingTime) }}
            </template>
          </primevue-column>
          <primevue-column field="ElapsedTime" header="經過時間 (min)">
            <template #body="slotProps">
              {{ ActivityService.CalculateTime(slotProps.data.ElapsedTime) }}
            </template>
          </primevue-column>
          <primevue-column field="MaxSpeed" :header="ActivityData.SportType === 'Ride' ? '最大速度 (km/hr)' : `最大配速 (${ActivityService.SportTypes[ActivityData.SportType].Unit})`">
            <template #body="slotProps">
              {{ ActivityService.CalculateSpeed(slotProps.data.MaxSpeed, ActivityData.SportType) }}
            </template>
          </primevue-column>
          <primevue-column field="AverageHeartrate" header="平均心率(bpm)" />
          <primevue-column field="MaxHeartrate" header="最大心率(bpm)" />
          <primevue-column field="AverageCadence" header="平均踏頻(spm)">
            <template #body="slotProps">
              <div v-if="slotProps.data.AverageCadence">
                {{ slotProps.data.AverageCadence }}
              </div>
              <div v-else>-</div>
            </template>
          </primevue-column>
          <primevue-column field="AverageWatts" header="平均功率(W)">
            <template #body="slotProps">
              <div v-if="slotProps.data.AverageWatts">
                {{ slotProps.data.AverageWatts }}
              </div>
              <div v-else>-</div>
            </template>
          </primevue-column>
        </primevue-data-table>
      </div>
      <div class="bg-zinc-50 border-2 border-zinc-100 p-1 sm:p-8 rounded-2xl z-0">
        <ECharts :ChartOptions="ChartSpeedOptions" />
      </div>
      <div class="bg-zinc-50 border-2 border-zinc-100 p-1 sm:p-8 rounded-2xl z-0">
        <ECharts :ChartOptions="ChartHeartrateOptions" />
      </div>
      <div v-if="ActivityData.AverageCadence" class="bg-zinc-50 border-2 border-zinc-100 p-1 sm:p-8 rounded-2xl z-0">
        <ECharts :ChartOptions="ChartCadenceOptions" />
      </div>
      <div v-if="ActivityData.AverageWatts" class="bg-zinc-50 border-2 border-zinc-100 p-1 sm:p-8 rounded-2xl z-0">
        <ECharts :ChartOptions="ChartPowerOptions" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onBeforeUnmount, ref } from 'vue';
import ECharts from '../components/e-charts.vue';
import LeafletMap from '../components/leaflet-map.vue';
import ActivityService from '../services/ActivityService';

const { moment } = window; // 時間格式
const IsLoading = ref(false); // 是否在載入中
const ActivityData = ref({}); // 活動資料
// const Zoom = ref(2);

// 座標時間格式
const timeFormatter = (o) => ActivityService.CalculateTime(o);
// 座標速度格式
const speedFormatter = (o) => ActivityService.CalculateSpeed(o, ActivityData.value.SportType);
// 鼠標速度格式
const tooltipSpeedFormatter = (o) => `(${ActivityService.CalculateTime(o.value[0])}, ${ActivityService.CalculateSpeed(o.value[1], ActivityData.value.SportType)})`;
// 其他鼠標格式
const tooltipOthersFormatter = (o) => `(${ActivityService.CalculateTime(o.value[0])}, ${o.value[1]})`;
// 決定速度圖表Y軸名稱
const determineSpeedYAxisName = (type) => `${type === 'Ride' ? '平均速度' : '平均配速'} (${ActivityService.SportTypes[type].Unit})`;
// 圖表速度選項
const ChartSpeedOptions = {
  grid: {
    bottom: 90, // 增加以防會跟X軸名重疊
  },
  tooltip: {
    formatter: tooltipSpeedFormatter,
  },
  xAxis: {
    type: 'value',
    name: '移動時間',
    nameLocation: 'center',
    nameGap: 30,
    axisLabel: { formatter: timeFormatter },
  },
  yAxis: {
    type: 'value',
    axisLabel: { formatter: speedFormatter },
  },
  dataZoom: [
    {
      type: 'slider',
      show: true,
      start: 0,
      end: 100,
      bottom: 10,
      labelFormatter: timeFormatter,
    },
  ],
  series: [
    {
      data: [],
      type: 'line',
      areaStyle: {},
    },
  ],
};

// 圖表心率選項
const ChartHeartrateOptions = {
  grid: {
    bottom: 90, // 增加以防會跟X軸名重疊
  },
  tooltip: {
    formatter: tooltipOthersFormatter,
  },
  xAxis: {
    type: 'value',
    name: '移動時間',
    nameLocation: 'center',
    nameGap: 30,
    axisLabel: { formatter: timeFormatter },
  },
  yAxis: {
    type: 'value',
    name: '平均心率 (bpm)',
  },
  dataZoom: [
    {
      type: 'slider',
      show: true,
      start: 0,
      end: 100,
      bottom: 10,
      labelFormatter: timeFormatter,
    },
  ],
  series: [
    {
      data: [],
      type: 'line',
      areaStyle: {},
      lineStyle: { color: '#FF2D2D' },
      itemStyle: { color: '#FF2D2D' },
    },
  ],
};

// 圖表踏頻選項
const ChartCadenceOptions = {
  grid: {
    bottom: 90, // 增加以防會跟X軸名重疊
  },
  tooltip: {
    formatter: tooltipOthersFormatter,
  },
  xAxis: {
    type: 'value',
    name: '移動時間',
    nameLocation: 'center',
    nameGap: 30,
    axisLabel: { formatter: timeFormatter },
  },
  yAxis: {
    type: 'value',
    name: '平均踏頻 (spm)',
  },
  dataZoom: [
    {
      type: 'slider',
      show: true,
      start: 0,
      end: 100,
      bottom: 10,
      labelFormatter: timeFormatter,
    },
  ],
  series: [
    {
      data: [],
      type: 'line',
      areaStyle: {},
      lineStyle: { color: '#01B468' },
      itemStyle: { color: '#01B468' },
    },
  ],
};

// 圖表功率選項
const ChartPowerOptions = {
  grid: {
    bottom: 90, // 增加以防會跟X軸名重疊
  },
  tooltip: {
    formatter: tooltipOthersFormatter,
  },
  xAxis: {
    type: 'value',
    name: '移動時間',
    nameLocation: 'center',
    nameGap: 30,
    axisLabel: { formatter: timeFormatter },
  },
  yAxis: {
    type: 'value',
    name: '平均功率 (W)',
  },
  dataZoom: [
    {
      type: 'slider',
      show: true,
      start: 0,
      end: 100,
      bottom: 10,
      labelFormatter: timeFormatter,
    },
  ],
  series: [
    {
      data: [],
      type: 'line',
      areaStyle: {},
      lineStyle: { color: '#949449' },
      itemStyle: { color: '#949449' },
    },
  ],
};

// 計算圖表資料
const calculateChartData = () => {
  const seriesData = { Speed: [], Heartrate: [], Cadence: [], Power: [] };
  let totalMovingTime = 0;

  // Y軸最小值
  const yAxisMin = {
    Speed: ActivityData.value.Laps[0].AverageSpeed,
    Heartrate: ActivityData.value.Laps[0].AverageHeartrate,
    Cadence: ActivityData.value.Laps[0].AverageCadence,
    Power: ActivityData.value.Laps[0].AverageWatts,
  };

  ActivityData.value.Laps.forEach((o) => {
    // 速度
    yAxisMin.Speed = Math.min(yAxisMin.Speed, o.AverageSpeed);
    seriesData.Speed.push([totalMovingTime, o.AverageSpeed]);
    seriesData.Speed.push([totalMovingTime + o.MovingTime, o.AverageSpeed]);
    // 心率
    yAxisMin.Heartrate = Math.min(yAxisMin.Heartrate, o.AverageHeartrate);
    seriesData.Heartrate.push([totalMovingTime, o.AverageHeartrate]);
    seriesData.Heartrate.push([totalMovingTime + o.MovingTime, o.AverageHeartrate]);

    // 踏頻
    if (o.AverageCadence) {
      yAxisMin.Cadence = Math.min(yAxisMin.Cadence, o.AverageCadence);
      seriesData.Cadence.push([totalMovingTime, o.AverageCadence]);
      seriesData.Cadence.push([totalMovingTime + o.MovingTime, o.AverageCadence]);
    }

    // 功率
    if (o.AverageWatts) {
      yAxisMin.Power = Math.min(yAxisMin.Power, o.AverageCadence);
      seriesData.Power.push([totalMovingTime, o.AverageWatts]);
      seriesData.Power.push([totalMovingTime + o.MovingTime, o.AverageWatts]);
    }

    totalMovingTime += o.MovingTime;
  });

  // 速度
  ChartSpeedOptions.xAxis.max = totalMovingTime;
  ChartSpeedOptions.yAxis.name = determineSpeedYAxisName(ActivityData.value.SportType);
  ChartSpeedOptions.yAxis.min = yAxisMin.Speed - 1;
  ChartSpeedOptions.series[0].data = seriesData.Speed;

  // 心率
  ChartHeartrateOptions.xAxis.max = totalMovingTime;
  ChartHeartrateOptions.yAxis.min = Math.round((yAxisMin.Heartrate - 30) / 10) * 10;
  ChartHeartrateOptions.series[0].data = seriesData.Heartrate;

  // 踏頻
  ChartCadenceOptions.xAxis.max = totalMovingTime;
  ChartCadenceOptions.yAxis.min = Math.round((yAxisMin.Cadence - 10) / 10) * 10;
  ChartCadenceOptions.series[0].data = seriesData.Cadence;

  // 功率
  ChartPowerOptions.xAxis.max = totalMovingTime;
  ChartPowerOptions.yAxis.min = Math.round((yAxisMin.Power - 30) / 10) * 10;
  ChartPowerOptions.series[0].data = seriesData.Power;
};

const LoadData = () => {
  IsLoading.value = true;
  const routerParams = window.router.currentRoute.value.params;
  ActivityService.GetActivityLaps(108845218, routerParams.activityid)
    .then((p) => {
      ActivityData.value = p;
      ActivityData.value.Laps = ActivityData.value.Laps.filter((o) => o.Distance);
      IsLoading.value = false;
    })
    .then(() => {
      calculateChartData(); // 計算圖表速度資料
      window.emitter.emit('SetChartOption');
    });
};
LoadData();

const ResizeChart = () => {
  window.emitter.emit('ResizeChart');
};

window.addEventListener('resize', ResizeChart);

onBeforeUnmount(() => window.removeEventListener('resize', ResizeChart));
</script>
