<script setup>
import { ref, computed, onMounted, onUnmounted } from "vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faNetworkWired,
  faSearch,
  faBolt,
  faMapMarkerAlt,
  faFileAlt,
  faChartBar,
} from "@fortawesome/free-solid-svg-icons";

// 添加图标到库中
library.add(
  faNetworkWired,
  faSearch,
  faBolt,
  faMapMarkerAlt,
  faFileAlt,
  faChartBar
);

// 窗口宽度响应式
const windowWidth = ref(window.innerWidth);

// 计算每行显示的工具数量
const toolsPerRow = computed(() => {
  if (windowWidth.value < 768) {
    return 1;
  } else if (windowWidth.value < 1024) {
    return 2;
  } else if (windowWidth.value < 1440) {
    return 3;
  } else {
    return 4;
  }
});

// 监听窗口大小变化
const handleResize = () => {
  windowWidth.value = window.innerWidth;
};

// 组件挂载时添加事件监听
onMounted(() => {
  window.addEventListener("resize", handleResize);
});

// 组件卸载时移除事件监听
onUnmounted(() => {
  window.removeEventListener("resize", handleResize);
});

// 网络工具数据
const networkTools = [
  {
    id: "analyse",
    title: "抓包分析",
    icon: "chart-bar",
    description: "网络数据包抓取和分析工具",
  },
];

// 工具点击处理
const handleToolClick = (tool) => {
  console.log(`点击了工具: ${tool.title}`);
  // 这里可以添加打开具体工具的逻辑
};
</script>

<template>
  <div class="network-tools-section">
    <div class="section-header">
      <h2>网络工具</h2>
      <a href="#" class="view-all">所有网络工具 >></a>
    </div>

    <div
      class="tools-grid"
      :style="{ gridTemplateColumns: `repeat(${toolsPerRow}, 1fr)` }"
    >
      <div
        v-for="tool in networkTools"
        :key="tool.id"
        class="tool-card"
        @click="handleToolClick(tool)"
      >
        <div class="tool-icon" :class="tool.id">
          <font-awesome-icon :icon="['fas', tool.icon]" />
        </div>
        <div class="tool-content">
          <h3>{{ tool.title }}</h3>
          <p>{{ tool.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.network-tools-section {
  background-color: #ffffff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  margin-bottom: 20px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.section-header h2 {
  font-size: 18px;
  color: #333;
  margin: 0;
}

.view-all {
  color: #6200ee;
  text-decoration: none;
  font-size: 14px;
}

.tools-grid {
  display: grid;
  grid-gap: 20px;
  transition: grid-template-columns 0.3s ease;
}

.tool-card {
  display: flex;
  background-color: #ffffff;
  border: 1px solid #eee;
  border-radius: 8px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.tool-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.tool-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  margin-right: 16px;
  border-radius: 8px;
  font-size: 24px;
  flex-shrink: 0;
}

.tool-content {
  flex: 1;
}

.tool-content h3 {
  margin: 0 0 8px 0;
  font-size: 16px;
  color: #333;
  font-weight: 500;
}

.tool-content p {
  margin: 0;
  font-size: 13px;
  color: #666;
  line-height: 1.5;
}

/* 针对特定工具的颜色定制 */
.ping {
  background-color: rgba(152, 144, 227, 0.1);
  color: #6563ff;
}

.trace {
  background-color: rgba(3, 169, 244, 0.1);
  color: #03a9f4;
}

.speed {
  background-color: rgba(255, 152, 0, 0.1);
  color: #ff9800;
}

.ip {
  background-color: rgba(76, 175, 80, 0.1);
  color: #4caf50;
}

.dns {
  background-color: rgba(244, 67, 54, 0.1);
  color: #f44336;
}

.analyse {
  background-color: rgba(156, 39, 176, 0.1);
  color: #9c27b0;
}
</style>
