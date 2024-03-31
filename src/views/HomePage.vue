<template>
  <div class="w-3">
    <Button label="新增跑步紀錄" @click="Add" />
  </div>
  

  <div>
    <DataTable :value="RunRecordData">
      <Column field="Date" header="日期">
        <template #body="slotProps">
          {{ moment(new Date(slotProps.data.Date)).format("YYYY/MM/DD") }}
        </template>
      </Column>
      <Column field="TotalTime" header="時間"></Column>
      <Column field="TotalDistance" header="總里程"></Column>
      <Column field="AvgPace" header="平均配速"></Column>
      <Column field="TrainingType" header="訓練類型"></Column>
      <Column>
        <template #body="slotProps">
          <Button
            @click="ViewRunRecord(slotProps.data)"
            icon="pi pi-eye"
            severity="info"
            class="mr-2"
          />
          <Button
            @click="DeleteRunRecord(slotProps.data)"
            icon="pi pi-trash"
            severity="danger"
          />
        </template>
      </Column>
    </DataTable>
  </div>

  <FileUpload
    accept=".csv"
    mode="basic"
    :chooseLabel="'請上傳Excel檔'"
    :maxFileSize="1000000"
    :customUpload="true"
    :auto="true"
    :showUploadButton="false"
    :showCancelButton="false"
    @uploader="HandleCSVFile"
    ref="MyUpload"
  />

  <Dialog v-model:visible="IsModelVisible" modal header="新增跑步紀錄">
    <div class="flex align-items-center mb-2">
      <div class="text-lg mr-3">日期:</div>
      <Calendar v-model="NowDate" dateFormat="yy-mm-dd" />
      <div class="text-lg mr-3">日期:</div>
      <Calendar v-model="NowDate" dateFormat="yy-mm-dd" />
    </div>
    <div class="flex align-items-center">
      <div class="text-lg mr-3">日期:</div>
      <Calendar v-model="NowDate" dateFormat="yy-mm-dd" />
    </div>
    <div class="flex justify-content-end gap-2 mt-3">
      <Button
        type="button"
        label="Cancel"
        severity="secondary"
        @click="IsModelVisible = false"
      ></Button>
      <Button type="button" label="Save" @click="Save" />
    </div>
  </Dialog>
</template>

<script setup>
import { ref } from "vue";
import RunRecordService from "@/services/RunRecordService";

const { moment } = window; // 時間格式
const MyUpload = ref(null);
const RunRecordData = ref([]);
const NowRunRecordData = ref({});
const IsModelVisible = ref(false);
const NowDate = ref();

const reader = new FileReader();
reader.onload = (o) => {
  const decoder = new TextDecoder();
  const result = decoder.decode(o.target.result);
  const workbook = window.XLSX.read(result, { type: "binary" }); // 讀取CSV
  const res = window.XLSX.utils.sheet_to_json(workbook.Sheets[workbook.SheetNames[0]], {
    raw: true,
  });
  console.log(res);

  NowRunRecordData.value.Detail = [];
  res.forEach((o, index) => {
    if (index < res.length - 1) {
      if (o.Distance > 0.02) {
        NowRunRecordData.value.Detail.push({
          Laps: o.Laps,
          Distance: o.Distance,
          Time: o.Time,
          AvgPace: o["Avg Pace"],
        });
      }
    } else {
      NowRunRecordData.value.TotalDistance = o.Distance;
      NowRunRecordData.value.TotalTime = o.Time;
      NowRunRecordData.value.AvgPace = o["Avg Pace"];
    }
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
  IsModelVisible.value = true;
};

const HandleCSVFile = (event) => {
  reader.readAsArrayBuffer(event.files[0]);
  MyUpload.value.clear();
};

const Save = () => {
  // RunRecordService.AddRunRecord(NowRunRecordData.value.)
  // Date, TotalDistance, TotalTime, AvgPace, Temperature, TrainingType, Detail, Feeling
  console.log("click");
};

const ViewRunRecord = (data) => {
  console.log(data);
};

const DeleteRunRecord = (data) => {
  console.log(data);
};
</script>
