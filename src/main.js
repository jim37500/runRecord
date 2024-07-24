import './assets/main.css';
import 'primevue/resources/themes/aura-light-green/theme.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';

import { createApp } from 'vue';

import PrimeVue from 'primevue/config';
import * as XLSX from 'xlsx/xlsx.mjs';
import moment from 'moment';
import Swal from 'sweetalert2';

import Button from 'primevue/button';
import FileUpload from 'primevue/fileupload';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Dropdown from 'primevue/dropdown';
import Dialog from 'primevue/dialog';
import Calendar from 'primevue/calendar';
import InputNumber from 'primevue/inputnumber';
import InputText from 'primevue/inputtext';

import Tag from 'primevue/tag';

import App from './App.vue';
import router from './router';

const app = createApp(App);

window.XLSX = XLSX;
window.moment = moment; // 時間格式
window.Swal = Swal;

app.use(router);
app.use(PrimeVue);

app.component('Button', Button);
app.component('FileUpload', FileUpload);
app.component('DataTable', DataTable);
app.component('Dialog', Dialog);
app.component('Dropdown', Dropdown);

app.component('Column', Column);
app.component('Tag', Tag);
app.component('Calendar', Calendar);
app.component('InputNumber', InputNumber);
app.component('InputText', InputText);

app.mount('#app');
