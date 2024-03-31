import './assets/main.css';
import 'primevue/resources/themes/aura-light-green/theme.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import PrimeVue from 'primevue/config';
import * as XLSX from 'xlsx/xlsx.mjs';
import moment from 'moment';

import Button from 'primevue/button';
import FileUpload from 'primevue/fileupload';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Dialog from 'primevue/dialog';
import Calendar from 'primevue/calendar';


import Tag from 'primevue/tag';

const app = createApp(App);

window.XLSX = XLSX;
window.moment = moment; // 時間格式

app.use(router);
app.use(PrimeVue);

app.component('Button', Button);
app.component('FileUpload', FileUpload);
app.component('DataTable', DataTable);
app.component('Column', Column);
app.component('Tag', Tag);
app.component('Dialog', Dialog);
app.component('Calendar', Calendar);

app.mount('#app');
