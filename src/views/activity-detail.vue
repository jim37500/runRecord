<template>1</template>

<script setup>
import ActivityService from '../services/ActivityService';

const LoadData = () => {
  const routerParams = window.router.currentRoute.value.params;
  ActivityService.GetActivityLaps(108845218, routerParams.type, routerParams.activityid).then((p) => {
    console.log(p)
  });
  console.log(routerParams);
};
LoadData();
</script>
