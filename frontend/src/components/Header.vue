<script setup>
import { defineProps } from "vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faHome, faGlobe, faTerminal } from "@fortawesome/free-solid-svg-icons";

// 添加图标到库中
library.add(faHome, faGlobe, faTerminal);

// 定义组件属性
const props = defineProps({
  activeMenu: String,
});
</script>

<template>
  <div class="header">
    <div class="header-left">
      <font-awesome-icon
        :icon="[
          'fas',
          activeMenu === 'home'
            ? 'home'
            : activeMenu === 'network'
            ? 'globe'
            : 'terminal',
        ]"
        class="header-icon"
      />
      <span class="header-title">
        {{
          activeMenu === "home"
            ? "运维工具箱"
            : activeMenu === "network"
            ? "网络工具"
            : "终端工具"
        }}
      </span>
    </div>
  </div>
</template>

<style scoped>
.header {
  display: flex;
  align-items: center;
  margin-bottom: 30px;
}

.header-left {
  display: flex;
  align-items: center;
}

.header-icon {
  font-size: 24px;
  margin-right: 10px;
  color: #6200ee;
}

.header-title {
  font-size: 24px;
  font-weight: bold;
  color: #333;
}
</style>
