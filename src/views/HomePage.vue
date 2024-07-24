<template>
  <div class="w-3">
    <Button label="新增跑步紀錄" @click="Add" />
  </div>

  <div>
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

  <Dialog v-model:visible="IsModelVisible" modal header="新增跑步紀錄" class="w-8">
    <div class="grid mb-2">
      <div class="col-6 flex align-items-center">
        <div class="text-lg mr-3">日期:</div>
        <Calendar v-model="NowDate" dateFormat="yy-mm-dd" class="w-10" />
      </div>
      <div class="col-6 flex align-items-center">
        <div class="text-lg mr-3">訓練類型:</div>
        <Dropdown v-model="NowTrainingType" :options="TrainingTypeOption" optionLabel="Name" placeholder="請選擇訓練類型" class="w-9" />
      </div>
    </div>
    <div class="flex align-items-center">
      <div class="col-6 flex align-items-center pl-0">
        <div class="text-lg mr-3">溫度:</div>
        <InputNumber v-model="NowTemperature" showButtons class="w-10" />
      </div>
      <div class="col-6 flex align-items-center">
        <div class="text-lg mr-3">訓練體感:</div>
        <Dropdown v-model="NowFeeling" :options="FeelingOption" optionLabel="Name" placeholder="請選擇訓練體感" class="w-9" />
      </div>
    </div>
    <div class="flex justify-content-between gap-2 mt-3">
      <div class="flex align-items-center">
        <FileUpload accept=".csv" mode="basic" name="fakeName" :chooseLabel="'請上傳Excel檔'" :maxFileSize="1000000" :customUpload="true" :auto="true" @uploader="HandleCSVFile" ref="MyUpload" />
        <Tag v-if="NowRunRecordData.AvgPace" icon="pi pi-check" class="ml-2" />
      </div>

      <div>
        <Button type="button" label="Cancel" severity="secondary" @click="IsModelVisible = false" class="mr-2" />
        <Button type="button" label="Save" @click="Save" />
      </div>
    </div>
  </Dialog>
</template>

<script setup>
import { ref } from 'vue';
import RunRecordService from '../services/RunRecordService';
import UtilityService from '../services/UtilityService';

const { moment } = window; // 時間格式
const { Alert, Confirm } = UtilityService; // 時間格式
const MyUpload = ref(null); // 上傳資料
const IsModelVisible = ref(false); // 新增跑步紀錄對話框
const RunRecordData = ref([]); // 全部跑步資料
const NowID = ref(0); // 現在跑步紀錄主鍵
const NowRunRecordData = ref({}); // 現在跑步紀錄資料
const NowDate = ref(); // 現在日期
const NowTemperature = ref(0); // 現在溫度
const NowFeeling = ref(''); // 現在訓練體感
const NowTrainingType = ref(); // 現在訓練類別
const FeelingOption = ref([{ Name: '非常弱' }, { Name: '弱' }, { Name: '普通' }, { Name: '強' }, { Name: '非常強' }]); // 體感選項
const TrainingTypeOption = ref([{ Name: '恢復跑' }, { Name: '一般有氧跑' }, { Name: '中長跑' }, { Name: '馬拉松配速跑' }, { Name: '乳酸閾值跑' }, { Name: '最大攝氧量跑' }]); // 訓練類別選項

const reader = new FileReader();
reader.onload = (o) => {
  console.log(o);
  const decoder = new TextDecoder();
  const result = decoder.decode(o.target.result);
  const workbook = window.XLSX.read(result, { type: 'binary' }); // 讀取CSV
  const res = window.XLSX.utils.sheet_to_json(workbook.Sheets[workbook.SheetNames[0]], {
    raw: true,
  });
  console.log(res);

  NowRunRecordData.value.Detail = [];
  res.forEach((p, index) => {
    if (index < res.length - 1) {
      if (p.Distance < 0.02) return;

      NowRunRecordData.value.Detail.push({
        Laps: p.Laps,
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
  console.log(NowRunRecordData.value);
};

const LoadData = () => {
  RunRecordService.GetRunRecord().then((o) => {
    RunRecordData.value = o;
  });
};
LoadData();

const Add = () => {
  NowID.value = 0;
  NowRunRecordData.value = {}; // 現在跑步紀錄資料
  NowDate.value = new Date(); // 現在日期
  NowTemperature.value = 20; // 現在溫度
  NowFeeling.value = {}; // 現在體感
  NowTrainingType.value = {}; // 現在訓練類別
  IsModelVisible.value = true;
};

const HandleCSVFile = (event) => {
  console.log(event);
  reader.readAsArrayBuffer(event.files[0]);
  MyUpload.value.clear();
};

// 儲存跑步紀錄
const Save = () => {
  if (!NowTrainingType.value.Name) return Alert('請選擇訓練類別', 'error');
  if (!NowFeeling.value.Name) return Alert('請選擇訓練體感', 'error');
  if (!NowRunRecordData.value.AvgPace) return Alert('請上傳跑步紀錄的csv檔', 'error');

  if (!NowID.value) {
    return RunRecordService.AddRunRecord(
      NowDate.value,
      NowRunRecordData.value.TotalDistance,
      NowRunRecordData.value.TotalTime,
      NowRunRecordData.value.AvgPace,
      NowRunRecordData.value.AvgHR,
      NowRunRecordData.value.AvgRunCadence,
      NowRunRecordData.value.AvgGroundContactTime,
      NowRunRecordData.value.AvgStrideLength,
      NowRunRecordData.value.AvgVerticalOscillation,
      NowRunRecordData.value.AvgVerticalRatio,
      NowTemperature.value,
      NowTrainingType.value.Name,
      NowRunRecordData.value.Detail,
      NowFeeling.value.Name,
    )
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
