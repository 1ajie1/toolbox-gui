<template>
  <div class="terminal-tools-section">
    <div class="section-header">
      <h2>终端工具</h2>
    </div>

    <div class="terminal-tools-grid">
      <div
        v-for="tool in terminalTools"
        :key="tool.id"
        class="tool-card clickable"
        @click="tool.onClick"
      >
        <div class="tool-icon terminal">
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

<script setup>
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faTerminal } from "@fortawesome/free-solid-svg-icons";

// 添加图标到库中
library.add(faTerminal);

// 打开终端函数
const openTerminal = async (type) => {
  try {
    await window.go.main.App.OpenTerminal(type);
  } catch (error) {
    console.error("打开终端失败:", error);
  }
};

// 终端工具数据
const terminalTools = [
  {
    id: "powershell",
    title: "PowerShell",
    icon: "terminal",
    description: "打开 Windows PowerShell 终端，已配置工具环境变量",
    onClick: () => openTerminal("powershell"),
  },
  {
    id: "cmd",
    title: "命令提示符",
    icon: "terminal",
    description: "打开 Windows CMD 终端，已配置工具环境变量",
    onClick: () => openTerminal("cmd"),
  },
];
</script>

<style scoped>
.terminal-tools-section {
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

.terminal-tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  grid-gap: 20px;
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

.terminal {
  background-color: rgba(33, 33, 33, 0.1);
  color: #212121;
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

.tool-card.clickable {
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.tool-card.clickable:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.tool-card.clickable:active {
  transform: translateY(0);
}
</style>
