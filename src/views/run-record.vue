<template>
  <div class="mx-auto w-10/12">
    <div class="text-4xl font-semibold my-2 text-blue-800">跑步紀錄</div>

    <div class="w-full grid grid-cols-1 xl:grid-cols-2 gap-8">
      <div class="hidden sm:block bg-zinc-50 border-2 border-zinc-100 px-8 py-8 rounded-2xl z-0">
        <FullCalendar ref="MyCalendar" :options="CalendarOptions" />
      </div>
      <div class="bg-zinc-50 border-2 border-zinc-100 rounded-2xl">
        <div v-if="IsLoading" class="flex justify-center items-center h-full">
          <i class="pi pi-spin pi-spinner" style="font-size: 2rem" />
        </div>
        <div v-else class="flex flex-col p-3 py-8 sm:p-8 items-center">
          <div class="block sm:flex sm:justify-between sm:items-center sm:w-full">
            <div class="font-semibold text-2xl mb-2 sm:mb-0">跑量統計: {{ RunStatistics }} km</div>
            <div class="flex">
              <button
                type="button"
                class="border-r-2 border-r-gray-700 rounded-l px-4 py-1.5 font-semibold text-lg text-white hover:bg-sky-600 transition"
                :class="NowTimeType === 'Week' ? 'bg-sky-600' : 'bg-sky-700'"
                @click="ChangeTimeType('Week')"
              >
                週
              </button>
              <button
                type="button"
                class="px-4 py-1.5 font-semibold text-lg hover:bg-sky-600 text-white transition"
                :class="NowTimeType === 'Month' ? 'bg-sky-600' : 'bg-sky-700'"
                @click="ChangeTimeType('Month')"
              >
                月
              </button>
              <button
                type="button"
                class="border-l-2 border-l-gray-700 rounded-r px-4 py-1.5 font-semibold text-lg hover:bg-sky-600 text-white transition"
                :class="NowTimeType === 'Year' ? 'bg-sky-600' : 'bg-sky-700'"
                @click="ChangeTimeType('Year')"
              >
                年
              </button>
            </div>
          </div>
          <div ref="SnapChart" class="max-h-full min-h-96 w-full" />
          <div class="w-full flex justify-center">
            <div>
              <button type="button" @click="ChangeTime(-1)"><i class="pi pi-angle-left" /></button>
              <span class="text-lg font-semibold mx-3">{{ NowDisplayTime }}</span>
              <button type="button" @click="ChangeTime(1)"><i class="pi pi-angle-right" /></button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-8 text-2xl font-semibold my-2 text-blue-800">近期活動</div>
    <div class="w-full grid lg:grid-cols-2 2xl:grid-cols-3 gap-4 bg-zinc-50 border-2 border-zinc-100 p-3 sm:p-8 rounded-2xl">
      <div
        v-for="item in RunRecordData.slice(0, 7)"
        :key="item"
        class="bg-blue-50 border-2 border-blue-200 cursor-pointer p-2 sm:p-6 rounded-2xl hover:bg-sky-100 hover:border-sky-300 active:bg-cyan-50 active:border-cyan-200 shadow-md"
      >
        <div class="flex items-center">
          <div class="hidden sm:flex flex-col justify-center items-center w-2/12 mr-2">
            <font-awesome-icon icon="fa-solid fa-person-running" class="text-2xl bg-yellow-200 p-4 rounded-full" />
            <div class="text-xs font-semibold text-center text-gray-500 mt-2">{{ moment(item.Date).format('MMM D, YYYY') }}</div>
          </div>
          <div class="w-full sm:w-10/12 flex justify-center">
            <div>
              <div class="text-lg font-semibold text-center">{{ item.Name }}</div>
              <div class="block sm:hidden text-xs text-gray-500 font-semibold text-center">{{ moment(item.Date).format('MMM D, YYYY') }}</div>
              <div class="flex mt-3">
                <div>
                  <p>距離</p>
                  <p class="font-semibold">{{ Math.round(item.Distance / 100) / 10 }} km</p>
                </div>
                <primevue-divider layout="vertical" />
                <div>
                  <p>配速</p>
                  <p class="font-semibold">{{ calculateSpeed(item.AverageSpeed) }} / km</p>
                </div>
                <primevue-divider layout="vertical" />
                <div>
                  <p>時間</p>
                  <p class="font-semibold">{{ calculateTime(item.MovingTime) }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import RunRecordService from '../services/RunRecordService';
// import UtilityService from '../services/UtilityService';

const { moment } = window; // 時間格式
// const { Alert } = UtilityService; // 時間格式
const MyCalendar = ref(); // 外勤紀錄本體
const IsLoading = ref(true); // 是否讀取中
const RunRecordData = ref([]); // 全部跑步資料
const SnapChart = ref(); // 快照圖表
const NowTimeType = ref('Week'); // 現在時間統計類別
const NowDisplayTime = ref(''); // 現在顯示時間
const NowTime = ref(new Date()); // 現在時間
const RunStatistics = ref(0); // 跑量統計

let Chart; // 快照圖表實體

const setTooltipFormatter = (o) => (o[0].value ? `<div class="font-semibold">${o[0].axisValue}</div> 🏃 ${o[0].value}K` : '');

const setLabelFormatter = (o) => (o.data ? `${o.data}` : '');

const chartOption = {
  tooltip: {
    trigger: 'axis',
    formatter: setTooltipFormatter,
    axisPointer: {
      type: 'shadow',
    },
  },
  xAxis: {
    type: 'category',
    data: [],
  },
  yAxis: {
    type: 'value',
  },
  series: [
    {
      data: [],
      type: 'bar',
      label: {
        show: true,
        position: 'inside',
        formatter: setLabelFormatter,
      },
    },
  ],
};

// 取得日期細節
const getDateDetail = (dateTime) => ({
  Year: dateTime.getFullYear(),
  Month: dateTime.getMonth() + 1,
  Date: dateTime.getDate(),
  Weekday: dateTime.getDay() ? dateTime.getDay() : 7,
});

// 設置顯示時間內的資料
const setDisplayTimeData = () => {
  chartOption.yAxis.max = 50;
  if (NowTimeType.value === 'Week') {
    const weekday = NowTime.value.getDay();
    const monday = moment(NowTime.value)
      .subtract(weekday - 1, 'days')
      .format('YYYY/MM/DD');

    const sunday = moment(NowTime.value)
      .add(8 - weekday, 'days')
      .format('YYYY/MM/DD');

    chartOption.xAxis.data = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
    const data = Array(7).fill(0);
    const weeklyData = RunRecordData.value.filter((o) => new Date(o.Date) >= new Date(monday) && new Date(o.Date) <= new Date(sunday));

    RunStatistics.value = 0;
    weeklyData.forEach((o) => {
      const nowDate = getDateDetail(new Date(o.Date));
      const nowDistance = Math.round(o.Distance / 100) / 10;
      data[nowDate.Weekday - 1] += nowDistance;
      RunStatistics.value += nowDistance;
    });

    RunStatistics.value = Math.round(RunStatistics.value * 10) / 10;

    chartOption.series[0].data = data;
    NowDisplayTime.value = `${monday} ~ ${moment(NowTime.value)
      .add(7 - weekday, 'days')
      .format('YYYY/MM/DD')}`;

    return;
  }

  const year = NowTime.value.getFullYear();
  if (NowTimeType.value === 'Month') {
    const month = NowTime.value.getMonth() + 1;
    const monthDays = new Date(year, month, 0).getDate();
    const monthlyData = RunRecordData.value.filter((o) => new Date(o.Date) >= new Date(year, month - 1, 1) && new Date(o.Date) <= new Date(year, month - 1, monthDays));
    const data = Array(monthDays).fill(0);
    chartOption.xAxis.data = [];

    for (let i = 1; i <= monthDays; i += 1) {
      chartOption.xAxis.data.push(`${i}`);
    }

    RunStatistics.value = 0;
    monthlyData.forEach((o) => {
      const nowDate = getDateDetail(new Date(o.Date));
      const nowDistance = Math.round(o.Distance / 100) / 10;
      data[nowDate.Date - 1] = Math.round((data[nowDate.Date - 1] + nowDistance) * 10) / 10;
      RunStatistics.value += nowDistance;
    });

    chartOption.series[0].data = data;
    RunStatistics.value = Math.round(RunStatistics.value * 10) / 10;
    NowDisplayTime.value = `${year}年${month}月`;
  } else {
    chartOption.xAxis.data = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
    chartOption.yAxis.max = 500;
    const yearlyData = RunRecordData.value.filter((o) => new Date(o.Date) >= new Date(year, 0, 1) && new Date(o.Date) <= new Date(year + 1, 0, 1));
    const data = Array(12).fill(0);

    RunStatistics.value = 0;
    yearlyData.forEach((o) => {
      const nowDate = getDateDetail(new Date(o.Date));
      const nowDistance = Math.round(o.Distance / 100) / 10;
      data[nowDate.Month - 1] = Math.round((data[nowDate.Month - 1] + nowDistance) * 10) / 10;
      RunStatistics.value += nowDistance;
    });

    chartOption.series[0].data = data;
    RunStatistics.value = Math.round(RunStatistics.value * 10) / 10;
    NowDisplayTime.value = `${year}年`;
  }
};

// 設定視窗縮放重繪圖表
const setWindowResize = (option) => () => {
  Chart.setOption(option, true);
  Chart.resize();
};

// 設置圖表
const setChartOption = () => {
  if (Chart) window.echarts.dispose(Chart);

  setDisplayTimeData(); // 取得現在顯示時間
  Chart = window.echarts.init(SnapChart.value); // 初始化圖表
  Chart.setOption(chartOption);
  window.onresize = setWindowResize(chartOption); // 設定視窗縮放重繪圖表
};

// 日期點擊事件
// const dateClick = () => {
//   IsModelVisible.value = true;
// };

// 事件點擊事件
const eventClick = (info) => {
  console.log(info.event.extendedProps);
};

const CalendarOptions = reactive({
  plugins: [window.dayGridPlugin, window.timeGridPlugin, window.listPlugin, window.interactionPlugin],
  initialView: 'dayGridMonth', // 預設為月曆
  height: 500,
  firstDay: 1,
  // aspectRatio: 10,
  headerToolbar: {
    start: 'title',
    // center: 'title',
    end: 'today prevYear,prev,next,nextYear',
  },
  dayMaxEventRows: true, // 單一天顯示活動限制
  displayEventTime: false, // 不顯示外勤紀錄時間
  // locale: 'zh-tw', // 語言
  showNonCurrentDates: false, // 只顯示當月日期
  selectable: true, // 活動可選
  events: [],
  // dateClick, // 日期點擊事件
  eventOrder: '-allDay, start', // 是否整天, 起始時間排序
  // moreLinkContent, // 更多外勤紀錄事項文字前置
  // moreLinkDidMount, // 設置更多外勤紀錄事項文字
  eventClick, // 外勤紀錄事項點擊事件
  // dayCellContent, // 每日單元內容
});

// 載入資料
const LoadData = () => {
  IsLoading.value = true;
  // 取得跑步紀錄
  RunRecordService.GetRunRecord(108845218)
    .then((o) => {
      RunRecordData.value = o;
      RunRecordData.value.forEach((p) => {
        const event = p;
        event.start = p.Date;
        event.title = `🏃${Math.round(p.Distance / 100) / 10}K`;
      });

      CalendarOptions.events = RunRecordData.value;
      IsLoading.value = false;
    })
    .then(() => setChartOption());
};
LoadData();

// 改變統計時間類別
const ChangeTimeType = (type) => {
  NowTimeType.value = type;
  setChartOption();
  // setTimeout(() => {
  //   NowTimeType.value = type;
  // }, 100);
};

// 更改時間類別
const ChangeTime = (duration) => {
  const date = NowTime.value.getDate();
  const month = NowTime.value.getMonth();
  if (NowTimeType.value === 'Week') {
    NowTime.value.setDate(NowTime.value.getDate() + duration * 7);
  } else if (NowTimeType.value === 'Month') {
    NowTime.value.setMonth(month + duration);
  } else if (NowTimeType.value === 'Year') {
    NowTime.value.setFullYear(NowTime.value.getFullYear() + duration, month, date);
  }

  setChartOption();
};

// 換算成配速
const calculateSpeed = (speed) => {
  const pace = Math.round((60 / speed) * 100) / 100;
  const decimal = pace - Math.floor(pace);

  return Math.round((Math.floor(pace) + decimal * 0.6) * 100) / 100;
};

// 換算時間(sec -> hr:min:sec)
const calculateTime = (secs) => {
  const seconds = secs % 60;
  let minutes = Math.floor(secs / 60);
  const hours = Math.floor(minutes / 60);

  if (hours > 0) {
    minutes = hours % 60;
    return `${hours}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
  }

  return `${minutes}:${seconds.toString().padStart(2, '0')}`;
};
</script>

<style>
@import '../assets/css/views/runrecord.css';
</style>
