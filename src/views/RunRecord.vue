<template>
  <div class="w-12 md:w-10">
    <div class="text-4xl font-semibold my-2">跑步紀錄</div>

    <div class="">
      <div class="w-6">
        <FullCalendar ref="MyCalendar" :options="CalendarOptions" class="w-12" />
      </div>

      <DataTable :value="RunRecordData">
        <Column field="Date" header="日期">
          <template #body="slotProps">
            {{ moment(new Date(slotProps.data.Date)).format('YYYY/MM/DD') }}
          </template>
        </Column>
        <Column field="TotalTime" header="時間" />
        <Column field="TotalDistance" header="總里程" />
        <Column field="AvgPace" header="平均配速" />
        <Column field="TrainingType" header="訓練類型" />
        <Column>
          <template #body="slotProps">
            <Button @click="ViewRunRecord(slotProps.data)" icon="pi pi-eye" severity="info" class="mr-2" />
            <Button @click="DeleteRunRecord(slotProps.data)" icon="pi pi-trash" severity="danger" />
          </template>
        </Column>
      </DataTable>
    </div>

    <Dialog v-model:visible="IsModelVisible" modal header="新增跑步紀錄" class="w-11">
      <div class="grid mb-2">
        <div class="col-4 flex align-items-center">
          <div class="text-lg mr-3">名稱</div>
          <InputText v-model="NowName" class="w-10" />
        </div>
        <div class="col-4 flex align-items-center">
          <div class="text-lg mr-3">日期:</div>
          <DatePicker v-model="NowDate" dateFormat="yy-mm-dd" class="w-10" />
        </div>
        <div class="col-4 flex align-items-center">
          <div class="text-lg mr-3">訓練類型:</div>
          <Select label="" v-model="NowTrainingType" :options="TrainingTypeOption" optionLabel="Name" placeholder="請選擇訓練類型" class="w-9" />
        </div>
      </div>
      <div class="flex align-items-center">
        <div class="col-6 flex align-items-center pl-0">
          <div class="text-lg mr-3">溫度:</div>
          <InputNumber v-model="NowTemperature" showButtons class="w-10" />
        </div>
        <div class="col-6 flex align-items-center">
          <div class="text-lg mr-3">訓練體感:</div>
          <Select v-model="NowFeeling" :options="FeelingOption" optionLabel="Name" placeholder="請選擇訓練體感" class="w-9" />
        </div>
      </div>

      <DataTable v-if="NowUploadData.length" :value="NowUploadData" showGridlines stripedRows scrollable scrollHeight="400px">
        <Column v-for="o in NowUploadExcelColumn" :key="o" :field="o" :header="o" class="w-full" />
      </DataTable>
      <template #footer>
        <div class="flex justify-content-between w-full gap-2 mt-3">
          <div class="flex align-items-center">
            <FileUpload accept=".csv" mode="basic" name="fakeName" :chooseLabel="'請上傳Excel檔'" :maxFileSize="1000000" :customUpload="true" :auto="true" @uploader="HandleCSVFile" ref="MyUpload" />
            <Tag v-if="NowRunRecordData.AvgPace" icon="pi pi-check" class="ml-2" />
          </div>

          <div>
            <Button type="button" label="Cancel" severity="secondary" @click="IsModelVisible = false" class="mr-2" />
            <Button type="button" label="Save" @click="Save" />
          </div>
        </div>
      </template>
    </Dialog>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import RunRecordService from '../services/RunRecordService';
import UtilityService from '../services/UtilityService';

const { moment } = window; // 時間格式
const { Alert, Confirm } = UtilityService; // 時間格式
const MyUpload = ref(null); // 上傳資料
const MyCalendar = ref(); // 外勤紀錄本體
const IsModelVisible = ref(false); // 新增跑步紀錄對話框
const RunRecordData = ref([]); // 全部跑步資料
const NowID = ref(0); // 現在跑步紀錄主鍵
const NowRunRecordData = ref({}); // 現在跑步紀錄資料
const NowDate = ref(); // 現在日期
const NowName = ref(''); // 現在名稱
const NowTemperature = ref(0); // 現在溫度
const NowFeeling = ref(''); // 現在訓練體感
const NowTrainingType = ref(); // 現在訓練類別
const FeelingOption = ref([{ Name: '非常弱' }, { Name: '弱' }, { Name: '普通' }, { Name: '強' }, { Name: '非常強' }]); // 體感選項
// 訓練類別選項
const TrainingTypeOption = ref([
  { Name: '恢復跑' },
  { Name: '一般有氧跑' },
  { Name: '中長跑' },
  { Name: '馬拉松配速跑' },
  { Name: '乳酸閾值跑' },
  { Name: '最大攝氧量跑' },
  { Name: '400間歇跑' },
  { Name: '600間歇跑' },
  { Name: '800間歇跑' },
]);
const NowUploadData = ref([]); // 上傳資料
const NowUploadExcelColumn = ref([]); // 上傳資料Excel欄位

// const Add = () => {
//   NowID.value = 0;
//   NowRunRecordData.value = {}; // 現在跑步紀錄資料
//   NowDate.value = new Date(); // 現在日期
//   NowTemperature.value = 20; // 現在溫度
//   NowFeeling.value = {}; // 現在體感
//   NowTrainingType.value = {}; // 現在訓練類別
//   NowUploadData.value = [];
//   NowUploadExcelColumn.value = [];

//   IsModelVisible.value = true;
// };

const dateClick = (o) => {
  NowID.value = 0;
  NowName.value = '';
  NowRunRecordData.value = {}; // 現在跑步紀錄資料
  NowDate.value = new Date(o.date); // 現在日期
  NowTemperature.value = 20; // 現在溫度
  NowFeeling.value = {}; // 現在體感
  NowTrainingType.value = {}; // 現在訓練類別
  NowUploadData.value = [];
  NowUploadExcelColumn.value = [];

  IsModelVisible.value = true;
};

const CalendarOptions = reactive({
  plugins: [window.dayGridPlugin, window.timeGridPlugin, window.listPlugin, window.interactionPlugin],
  initialView: 'dayGridMonth', // 預設為月曆
  // height: 500, // 設置外勤紀錄高度
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
  dateClick,
  eventOrder: '-allDay, start', // 是否整天, 起始時間排序
  // moreLinkContent, // 更多外勤紀錄事項文字前置
  // moreLinkDidMount, // 設置更多外勤紀錄事項文字
  // eventClick, // 外勤紀錄事項點擊事件
  // dayCellContent, // 每日單元內容
});

const reader = new FileReader();
reader.onload = (o) => {
  const decoder = new TextDecoder();
  const result = decoder.decode(o.target.result);
  const workbook = window.XLSX.read(result, { type: 'binary' }); // 讀取CSV
  NowUploadData.value = window.XLSX.utils.sheet_to_json(workbook.Sheets[workbook.SheetNames[0]], { raw: true });

  const columns = new Set();
  NowUploadData.value.forEach((p) => {
    Object.keys(p).forEach((key) => {
      columns.add(key);
    });
  });

  NowUploadExcelColumn.value = [...columns];

  NowRunRecordData.value.Detail = [];
  NowUploadData.value.forEach((p, index) => {
    if (index < NowUploadData.value.length - 1) {
      if (p.Distance < 0.02) return;

      NowRunRecordData.value.Detail.push({
        Laps: p.Laps ? p.Laps : p.Lap,
        Distance: p.Distance,
        Time: p.Time,
        AvgPace: p['Avg Pace'],
        AvgHR: p['Avg HR'],
        AvgRunCadence: p['Avg Run Cadence'],
        AvgGroundContactTime: p['Avg Ground Contact Time'],
        AvgStrideLength: p['Avg Stride Length'],
        AvgVerticalOscillation: p['Avg Vertical Oscillation'],
        AvgVerticalRatio: p['Avg Vertical Ratio'],
      });
      return;
    }

    NowRunRecordData.value.TotalDistance = p.Distance;
    NowRunRecordData.value.TotalTime = p.Time;
    NowRunRecordData.value.AvgPace = p['Avg Pace'];
    NowRunRecordData.value.AvgHR = p['Avg HR'];
    NowRunRecordData.value.AvgRunCadence = p['Avg Run Cadence'];
    NowRunRecordData.value.AvgGroundContactTime = p['Avg Ground Contact Time'];
    NowRunRecordData.value.AvgStrideLength = p['Avg Stride Length'];
    NowRunRecordData.value.AvgVerticalOscillation = p['Avg Vertical Oscillation'];
    NowRunRecordData.value.AvgVerticalRatio = p['Avg Vertical Ratio'];
  });
};

// 載入資料
const LoadData = () => {
  // 取得跑步紀錄
  RunRecordService.GetRunRecord().then((o) => {
    RunRecordData.value = o;
    RunRecordData.value.forEach((p) => {
      const event = p;
      event.start = p.Date;
      event.title = p.Name ? p.Name : `${p.TotalDistance}K`;
    });
    CalendarOptions.events = RunRecordData.value;
  });
};
LoadData();

// 處理上傳csv檔
const HandleCSVFile = (event) => {
  reader.readAsArrayBuffer(event.files[0]);
  MyUpload.value.clear();
};

// 儲存跑步紀錄
const Save = () => {
  if (!NowTrainingType.value.Name) return Alert('請選擇訓練類別', 'error');
  if (!NowFeeling.value.Name) return Alert('請選擇訓練體感', 'error');
  if (!NowRunRecordData.value.AvgPace) return Alert('請上傳跑步紀錄的csv檔', 'error');

  const record = {
    Name: NowName.value,
    Date: NowDate.value,
    TotalDistance: NowRunRecordData.value.TotalDistance,
    TotalTime: NowRunRecordData.value.TotalTime,
    AvgPace: NowRunRecordData.value.AvgPace,
    AvgHR: NowRunRecordData.value.AvgHR,
    AvgRunCadence: NowRunRecordData.value.AvgRunCadence,
    AvgGroundContactTime: NowRunRecordData.value.AvgGroundContactTime,
    AvgStrideLength: NowRunRecordData.value.AvgStrideLength,
    AvgVerticalOscillation: NowRunRecordData.value.AvgVerticalOscillation,
    AvgVerticalRatio: NowRunRecordData.value.AvgVerticalRatio,
    Temperature: NowTemperature.value,
    TrainingType: NowTrainingType.value.Name,
    Detail: NowRunRecordData.value.Detail,
    Feeling: NowFeeling.value.Name,
  };

  if (!NowID.value) {
    return RunRecordService.AddRunRecord(record)
      .then(() => Alert('新增跑步紀錄成功', 'success'))
      .then(() => {
        LoadData();
        IsModelVisible.value = false;
      })
      .catch(() => Alert('新增跑步紀錄失敗', 'error'));
  }
  return null;
};

const ViewRunRecord = (data) => {
  console.log(data);
};

const DeleteRunRecord = (data) => {
  Confirm(`確定要刪除${moment(new Date(data.Date)).format('YYYY/MM/DD')}-${data.TrainingType}`).then((o) => {
    if (!o.isConfirmed) return;

    RunRecordService.DeleteRunRecord(data.ID)
      .then(() => Alert('刪除成功', 'success'))
      .then(() => LoadData())
      .catch(() => Alert('刪除失敗', 'error'));
  });
};
</script>

<style>
@import '../assets/css/views/runrecord.css';
</style>
