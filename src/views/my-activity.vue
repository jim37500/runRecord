<template>
  <div class="mx-auto w-10/12">
    <!-- <div class="text-4xl font-semibold my-2 text-blue-800">跑步紀錄</div> -->
    <div class="w-full grid grid-cols-1 xl:grid-cols-2 gap-8">
      <div class="hidden sm:block bg-zinc-50 border-2 border-zinc-100 px-8 py-8 rounded-2xl z-0">
        <FullCalendar ref="MyCalendar" :options="CalendarOptions" />
      </div>
      <div class="flex flex-col bg-zinc-50 border-2 border-zinc-100 rounded-2xl">
        <div class="flex justify-between items-center p-3 pt-8 sm:pt-8 sm:px-8">
          <div class="flex bg-sky-700 rounded py-1 px-0.5">
            <button
              v-for="item in TimeTypes"
              :key="item"
              type="button"
              class="px-3 py-1 mx-0.5 rounded font-semibold sm:text-lg text-white hover:bg-sky-600 transition-colors"
              :class="NowTimeType === item.Value ? 'bg-sky-600' : ''"
              @click="ChangeTimeType(item.Value)"
            >
              {{ item.Name }}
            </button>
          </div>

          <div class="flex bg-sky-700 rounded py-1 px-0.5">
            <div v-for="item in SportTypes" :key="item">
              <button
                v-if="item.Value !== 'All'"
                type="button"
                class="px-1 sm:px-3 py-1 mx-0.5 rounded font-semibold sm:text-lg text-white hover:bg-sky-600 transition-colors"
                :class="ChartSportType.some((type) => type.Value === item.Value) ? 'bg-sky-600' : ''"
                @click="ChangeChartSportType(item)"
              >
                {{ item.Name }}
              </button>
            </div>
          </div>
        </div>
        <div v-if="IsLoading" class="flex justify-center items-center h-full">
          <i class="pi pi-spin pi-spinner" style="font-size: 2rem" />
        </div>
        <div v-else class="flex flex-col flex-1 px-3 pb-3 sm:px-8 justify-center items-center">
          <div class="font-semibold md:text-2xl mb-2 mt-3 sm:mb-0">
            {{ StatisticsOverview }}
          </div>
          <ECharts :ChartOptions="ChartOption" />
          <div class="w-full flex justify-center">
            <div>
              <button type="button" @click="ChangeTime(-1)">
                <i class="pi pi-angle-left" />
              </button>
              <span class="text-lg font-semibold mx-3">{{ NowDisplayTime }}</span>
              <button type="button" @click="ChangeTime(1)">
                <i class="pi pi-angle-right" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-8 text-2xl font-semibold my-2 text-blue-800">近期活動</div>
    <div class="w-full grid lg:grid-cols-2 2xl:grid-cols-3 gap-4 bg-zinc-50 border-2 border-zinc-100 p-3 sm:p-8 rounded-2xl">
      <div class="flex flex-col sm:flex-row col-span-1 lg:col-span-2 2xl:col-span-3 items-center justify-center">
        <div class="flex bg-sky-700 rounded py-1 px-0.5">
          <div v-for="item in SportTypes" :key="item">
            <button
              type="button"
              class="px-1 sm:px-3 py-1 mx-0.5 rounded font-semibold sm:text-lg text-white hover:bg-sky-600 transition-colors"
              :class="RecentSportType.Value === item.Value ? 'bg-sky-600' : ''"
              @click="ChangeRecentSportType(item)"
            >
              {{ item.Name }}
            </button>
          </div>
        </div>
        <primevue-input-group class="sm:ml-6 mt-2 sm:mt-0 flex-1">
          <primevue-input-text v-model="Keyword" placeholder="搜尋活動" />
          <primevue-button icon="pi pi-search" severity="contrast" />
        </primevue-input-group>
      </div>
      <div
        v-for="item in RecentData"
        :key="item"
        class="bg-blue-50 border-2 border-blue-200 cursor-pointer p-2 sm:p-6 rounded-2xl hover:bg-sky-100 hover:border-sky-300 active:bg-cyan-50 active:border-cyan-200 shadow-md"
        @click="GoToRunDetail(item)"
      >
        <div class="flex items-center">
          <div class="hidden sm:flex flex-col justify-center items-center w-2/12 mr-2">
            <font-awesome-icon v-if="item.SportType === 'Run'" icon="fa-solid fa-person-running" class="text-2xl bg-yellow-200 p-4 rounded-full" />
            <font-awesome-icon v-else-if="item.SportType === 'Ride'" icon="fa-solid fa-person-biking" class="text-2xl bg-yellow-200 p-4 rounded-full" />
            <font-awesome-icon v-else icon="fa-solid fa-person-swimming" class="text-2xl bg-yellow-200 p-4 rounded-full" />
            <div class="text-xs font-semibold text-center text-gray-500 mt-2">
              {{ moment(item.Date).format('MMM D, YYYY') }}
            </div>
          </div>
          <div class="w-full sm:w-10/12 flex justify-center">
            <div>
              <div class="text-lg font-semibold text-center">{{ item.Name }}</div>
              <div class="block sm:hidden text-xs text-gray-500 font-semibold text-center">
                {{ moment(item.Date).format('MMM D, YYYY') }}
              </div>
              <div class="flex mt-3">
                <div>
                  <p>距離</p>
                  <p class="font-semibold">{{ ActivityService.CalculateDistance(item.Distance).toFixed(1) }} km</p>
                </div>
                <primevue-divider layout="vertical" />
                <div>
                  <p v-if="item.SportType === 'Ride'">速度</p>
                  <p v-else>配速</p>
                  <p class="font-semibold">
                    {{ ActivityService.CalculateSpeed(item.AverageSpeed, item.SportType) }}
                    {{ ActivityService.SportTypes[item.SportType].Unit }}
                  </p>
                </div>
                <primevue-divider layout="vertical" />
                <div>
                  <p>時間</p>
                  <p class="font-semibold">
                    {{ ActivityService.CalculateTime(item.MovingTime) }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <primevue-paginator
        class="col-span-1 lg:col-span-2 2xl:col-span-3"
        :rows="12"
        :totalRecords="FilteredData.length"
        :pt="{ root: { class: '!bg-zinc-50' } }"
        @page="ChangePaginator"
      ></primevue-paginator>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch, onBeforeUnmount, computed } from 'vue';
import ECharts from '../components/e-charts.vue';
import ActivityService from '../services/ActivityService';
// import UtilityService from '../services/UtilityService';

const { moment } = window; // 時間格式
const MyCalendar = ref(); // 外勤紀錄本體
const IsLoading = ref(true); // 是否讀取中
const Activities = ref([]); // 全部活動
const FilteredData = ref([]); // 篩選過的資料
const RecentData = ref([]); // 近期資料
const NowDisplayTime = ref(''); // 現在顯示時間
const NowTime = ref(new Date()); // 現在時間
const StatisticsOverview = ref(''); // 運動統計概覽
const Keyword = ref(''); // 關鍵字

const TimeTypes = ref([
  { Name: '週', Value: 'Week' },
  { Name: '月', Value: 'Month' },
  { Name: '年', Value: 'Year' },
]);
const NowTimeType = ref(TimeTypes.value[0].Value); // 現在時間統計類別
const SportTypes = ref([{ Name: '全部', Value: 'All', Icon: '' }, ...Object.values(ActivityService.SportTypes)]);
const ChartSportType = ref(SportTypes.value.slice(1)); // 圖表運動類別
const RecentSportType = ref(SportTypes.value[0]); // 近期運動類別

const setTooltipFormatter = (o) => (o[0].value ? `<div class="font-semibold">${o[0].axisValue}</div> 🏃 ${o[0].value}K` : '');

const setLabelFormatter = (o) => (o.data ? `${o.data}` : '');

// Add this computed property
const legendStyle = computed(() => {
  const isMobile = window.innerWidth < 640; // sm breakpoint in Tailwind
  return {
    fontSize: isMobile ? 12 : 16,
    fontWeight: 'bold',
  };
});

// Add this computed property for legend layout
const legendConfig = computed(() => {
  const isMobile = window.innerWidth < 640;
  return {
    show: true,
    formatter: (name) => {
      const series = ChartOption.series.find(s => s.name === name);
      if (!series) return name;
      const total = series.data.reduce((sum, val) => sum + (val || 0), 0);
      return `${name}: ${total.toFixed(1)}K`;
    },
    textStyle: legendStyle.value,
    itemWidth: isMobile ? 18 : 25,
    itemHeight: isMobile ? 10 : 14,
    itemGap: isMobile ? 15 : 25,
    padding: isMobile ? [10, 15] : [15, 20],
  };
});

const ChartOption = {
  tooltip: {
    trigger: 'axis',
    formatter: setTooltipFormatter,
    axisPointer: {
      type: 'shadow',
    },
  },
  legend: legendConfig.value,
  xAxis: {
    type: 'category',
    data: [],
  },
  yAxis: {
    type: 'value',
  },
  series: [
    {
      name: '跑步',
      data: [],
      type: 'bar',
      stack: 'total',
      label: {
        show: true,
        position: 'inside',
        formatter: setLabelFormatter,
      },
    },
    {
      name: '騎車',
      data: [],
      type: 'bar',
      stack: 'total',
      label: {
        show: true,
        position: 'inside',
        formatter: setLabelFormatter,
      },
    },
    {
      name: '游泳',
      data: [],
      type: 'bar',
      stack: 'total',
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

const calculateWeeklyData = () => {
  const weekday = NowTime.value.getDay() ? NowTime.value.getDay() : 7;
  const monday = moment(NowTime.value)
    .subtract(weekday - 1, 'days')
    .format('YYYY/MM/DD');

  // 星期日 24:00
  const sunday = moment(NowTime.value)
    .add(8 - weekday, 'days')
    .format('YYYY/MM/DD');

  ChartOption.xAxis.data = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
  StatisticsOverview.value = '';
  const sportTypeStatistics = [];
  ChartSportType.value.forEach((sportType) => {
    const data = Array(7).fill(0);
    const weeklyData = Activities.value.filter((o) => o.SportType === sportType.Value && new Date(o.Date) >= new Date(monday) && new Date(o.Date) <= new Date(sunday));
    // 計算圖表資料
    weeklyData.forEach((o) => {
      const nowDate = getDateDetail(new Date(o.Date)); // 取得日期細節
      const nowDistance = ActivityService.CalculateDistance(o.Distance);
      data[nowDate.Weekday - 1] = Math.round((nowDistance + data[nowDate.Weekday - 1]) * 10) / 10;
    });
    ChartOption.series.find((o) => o.name === sportType.Name).data = data;
  });
  StatisticsOverview.value = sportTypeStatistics.join(' | ');

  NowDisplayTime.value = `${monday} ~ ${moment(NowTime.value)
    .add(7 - weekday, 'days')
    .format('YYYY/MM/DD')}`;
};

const calculateMonthlyData = () => {
  const year = NowTime.value.getFullYear();
  const month = NowTime.value.getMonth() + 1;
  const monthDays = new Date(year, month, 0).getDate();

  ChartOption.xAxis.data = [];
  for (let i = 1; i <= monthDays; i += 1) {
    ChartOption.xAxis.data.push(`${i}`);
  }

  ChartSportType.value.forEach((sportType) => {
    const data = Array(monthDays).fill(0);
    const monthlyData = Activities.value.filter((o) => o.SportType === sportType.Value && new Date(o.Date) >= new Date(year, month - 1, 1) && new Date(o.Date) <= new Date(year, month - 1, monthDays));
    monthlyData.forEach((o) => {
      const nowDate = getDateDetail(new Date(o.Date));
      const nowDistance = ActivityService.CalculateDistance(o.Distance);
      data[nowDate.Date - 1] = Math.round((data[nowDate.Date - 1] + nowDistance) * 10) / 10;
    });
    ChartOption.series.find((o) => o.name === sportType.Name).data = data;
  });

  NowDisplayTime.value = `${year}年${month}月`;
};

const calculateYearlyData = () => {
  const year = NowTime.value.getFullYear();
  ChartOption.xAxis.data = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

  ChartSportType.value.forEach((sportType) => {
    const yearlyData = Activities.value.filter((o) => o.SportType === sportType.Value && new Date(o.Date) >= new Date(year, 0, 1) && new Date(o.Date) <= new Date(year + 1, 0, 1));
    const data = Array(12).fill(0);

    yearlyData.forEach((o) => {
      const nowDate = getDateDetail(new Date(o.Date));
      const nowDistance = ActivityService.CalculateDistance(o.Distance);
      data[nowDate.Month - 1] = Math.round((data[nowDate.Month - 1] + nowDistance) * 10) / 10;
    });
    ChartOption.series.find((o) => o.name === sportType.Name).data = data;
  });

  NowDisplayTime.value = `${year}年`;
};

// 設置顯示時間內的資料
const setDisplayTimeData = () => {
  // 清空圖表資料
  ChartOption.series.forEach((o) => {
    o.data = [];
  });
  // 設置圖表Y軸最大值
  // ChartOption.yAxis.max = ChartYMaxValue[ChartSportType.value.Value][NowTimeType.value];
  // 計算圖表資料
  if (NowTimeType.value === 'Week') {
    calculateWeeklyData();
  } else if (NowTimeType.value === 'Month') {
    calculateMonthlyData();
  } else {
    calculateYearlyData();
  }
};

// 重新渲染圖表
const RerenderECharts = () => {
  setDisplayTimeData(); // 取得現在顯示時間
  window.emitter.emit('SetChartOption');
};

// 變換圖表運動類型
const ChangeChartSportType = (sportType) => {
  const index = ChartSportType.value.findIndex((type) => type.Value === sportType.Value);
  if (index === -1) {
    ChartSportType.value.push(sportType);
  } else {
    ChartSportType.value.splice(index, 1);
  }
  // Prevent deselecting all sports
  if (ChartSportType.value.length === 0) {
    ChartSportType.value.push(sportType);
  }
  RerenderECharts();
};

// 變換近期運動類型
const ChangeRecentSportType = (sportType) => {
  RecentSportType.value = sportType;

  if (sportType.Value === 'All') {
    FilteredData.value = Activities.value;
  } else {
    FilteredData.value = Activities.value.filter((o) => o.SportType === sportType.Value);
  }

  RecentData.value = FilteredData.value.slice(0, 12);
};

// 關鍵字搜尋活動
const Search = () => {
  RecentSportType.value = SportTypes.value[0];

  if (Keyword.value) {
    FilteredData.value = Activities.value.filter((o) => o.Name.includes(Keyword.value));
  } else {
    FilteredData.value = Activities.value;
  }

  RecentData.value = FilteredData.value.slice(0, 12);
};

watch(() => Keyword.value, Search);

// 導向跑步活動細節
const GoToRunDetail = (activity) => {
  window.router.push(`activity/${activity.ID}`);
};

// 事件點擊事件
const eventClick = (info) => {
  window.router.push(`activity/${info.event.extendedProps.ID}`);
};

// 改變近期活動分頁
const ChangePaginator = (o) => {
  RecentData.value = FilteredData.value.slice(o.first, o.first + 12);
};

// 日曆選項
const CalendarOptions = reactive({
  plugins: [window.dayGridPlugin, window.timeGridPlugin, window.listPlugin, window.interactionPlugin],
  initialView: 'dayGridMonth', // 預設為月曆
  height: 580,
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
  eventOrder: '-allDay, start', // 是否整天, 起始時間排序
  eventClick, // 活動點擊事件
});

// 載入資料
const LoadData = () => {
  IsLoading.value = true;
  // 取得跑步紀錄
  ActivityService.GetActivities(108845218)
    .then((o) => {
      o.forEach((p) => {
        const event = p;
        event.start = p.Date;
        event.title = `${ActivityService.SportTypes[event.SportType].Icon}${ActivityService.CalculateDistance(p.Distance).toFixed(1)}K`;
      });

      Activities.value = o;
      Activities.value.sort((a, b) => new Date(b.Date) - new Date(a.Date));
      CalendarOptions.events = Activities.value;
      FilteredData.value = Activities.value;
      RecentData.value = Activities.value.slice(0, 12);
      IsLoading.value = false;
    })
    .then(() => RerenderECharts());
};
LoadData();

// 改變統計時間類別
const ChangeTimeType = (type) => {
  NowTimeType.value = type;
  RerenderECharts();
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

  RerenderECharts();
};

// Add resize handler to update legend when window size changes
const ResizeChart = () => {
  ChartOption.legend = legendConfig.value;
  window.emitter.emit('ResizeChart');
};

window.addEventListener('resize', ResizeChart);

onBeforeUnmount(() => {
  window.removeEventListener('resize', ResizeChart);
});
</script>

<style>
@import '../assets/css/views/my-activity.css';
</style>
