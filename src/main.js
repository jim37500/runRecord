import './assets/main.css';
import 'primeicons/primeicons.css';
import 'leaflet/dist/leaflet.css';
import 'nprogress/nprogress.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faPersonRunning, faPersonBiking, faPersonSwimming, faMagnifyingGlass, faPersonWalking } from '@fortawesome/free-solid-svg-icons';
import { LMap, LGeoJson, LTileLayer, LPolyline } from '@vue-leaflet/vue-leaflet';
import leaflet from 'leaflet';

import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';
import { definePreset } from '@primevue/themes';

import * as XLSX from 'xlsx/xlsx.mjs';
import moment from 'moment';
import Swal from 'sweetalert2';
import mitt from 'mitt';
import polyline from '@mapbox/polyline';
import sha256 from 'sha256';

import Checkbox from 'primevue/checkbox';
import Divider from 'primevue/divider';
import Paginator from 'primevue/paginator';
import InputGroup from 'primevue/inputgroup';
// import FileUpload from 'primevue/fileupload';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
// import Select from 'primevue/select';
import SelectButton from 'primevue/selectbutton';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
// import DatePicker from 'primevue/datepicker';
// import InputNumber from 'primevue/inputnumber';
import InputText from 'primevue/inputtext';
import Tooltip from 'primevue/tooltip';
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

library.add(faPersonRunning, faPersonBiking, faPersonSwimming, faMagnifyingGlass, faPersonWalking);

const pinia = createPinia();
const app = createApp(App);

const MyPreset = definePreset(Aura, {
  semantic: {
    primary: {
      50: '{gray.50}',
      100: '{gray.100}',
      200: '{gray.200}',
      300: '{gray.300}',
      400: '{gray.400}',
      500: '{gray.500}',
      600: '{gray.600}',
      700: '{gray.700}',
      800: '{gray.800}',
      900: '{gray.900}',
      950: '{gray.950}',
    },
  },
});

window.XLSX = XLSX;
window.router = router;
window.moment = moment; // 時間格式
window.Swal = Swal;
window.dayGridPlugin = dayGridPlugin;
window.timeGridPlugin = timeGridPlugin;
window.listPlugin = listPlugin;
window.interactionPlugin = interactionPlugin;
window.echarts = echarts;
window.emitter = mitt();
window.polyline = polyline;
window.leaflet = leaflet;
window.sha256 = sha256;

app.use(pinia);
app.use(router);
app.use(PrimeVue, {
  theme: {
    preset: MyPreset,
    options: {
      //   cssLayer: {
      //     name: 'primevue',
      //     order: 'primevue',
      //   },
    },
  },
});

app.directive('tooltip', Tooltip);

app.component('primevue-checkbox', Checkbox);
app.component('primevue-divider', Divider);
app.component('primevue-paginator', Paginator);
app.component('primevue-button', Button);
app.component('primevue-select-button', SelectButton);
app.component('primevue-input-group', InputGroup);
// app.component('FileUpload', FileUpload);
app.component('primevue-data-table', DataTable);
app.component('primevue-dialog', Dialog);
// app.component('Select', Select);

app.component('primevue-column', Column);
// app.component('Tag', Tag);
// app.component('DatePicker', DatePicker);
// app.component('InputNumber', InputNumber);
app.component('primevue-input-text', InputText);
// app.component('Menubar', Menubar);
app.component('FullCalendar', FullCalendar);
app.component('font-awesome-icon', FontAwesomeIcon);

app.component('l-map', LMap);
app.component('l-tile-layer', LTileLayer);
app.component('l-geo-json', LGeoJson);
app.component('l-poly-line', LPolyline);

app.mount('#app');
