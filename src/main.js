import './assets/main.css';
import 'primeicons/primeicons.css';

import { createApp } from 'vue';
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faPersonRunning } from '@fortawesome/free-solid-svg-icons';

import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';

import * as XLSX from 'xlsx/xlsx.mjs';
import moment from 'moment';
import Swal from 'sweetalert2';

import Divider from 'primevue/divider';
// import FileUpload from 'primevue/fileupload';
// import DataTable from 'primevue/datatable';
// import Column from 'primevue/column';
// import Select from 'primevue/select';
// import SelectButton from 'primevue/selectbutton';
// import Dialog from 'primevue/dialog';
// import DatePicker from 'primevue/datepicker';
// import InputNumber from 'primevue/inputnumber';
// import InputText from 'primevue/inputtext';
// import Menubar from 'primevue/menubar';
// import Tag from 'primevue/tag';

import FullCalendar from '@fullcalendar/vue3';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import listPlugin from '@fullcalendar/list';
import interactionPlugin from '@fullcalendar/interaction';

import * as echarts from 'echarts';

import App from './App.vue';
import router from './router';

library.add(faPersonRunning);

const app = createApp(App);

window.XLSX = XLSX;
window.router = router;
window.moment = moment; // 時間格式
window.Swal = Swal;
window.dayGridPlugin = dayGridPlugin;
window.timeGridPlugin = timeGridPlugin;
window.listPlugin = listPlugin;
window.interactionPlugin = interactionPlugin;
window.echarts = echarts;

app.use(router);
app.use(PrimeVue, {
  theme: {
    preset: Aura,
    // options: {
    //   cssLayer: {
    //     name: 'primevue',
    //     order: 'primevue',
    //   },
    // },
  },
});

app.component('primevue-divider', Divider);
// app.component('Button', Button);
// app.component('SelectButton', SelectButton);
// app.component('FileUpload', FileUpload);
// app.component('DataTable', DataTable);
// app.component('Dialog', Dialog);
// app.component('Select', Select);

// app.component('Column', Column);
// app.component('Tag', Tag);
// app.component('DatePicker', DatePicker);
// app.component('InputNumber', InputNumber);
// app.component('InputText', InputText);
// app.component('Menubar', Menubar);
app.component('FullCalendar', FullCalendar);
app.component('font-awesome-icon', FontAwesomeIcon);

app.mount('#app');
