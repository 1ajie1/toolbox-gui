<script setup>
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faHome,
  faGlobe,
  faTerminal,
  faTools,
  faQuestionCircle,
  faInfoCircle,
} from "@fortawesome/free-solid-svg-icons";

// 添加图标到库中
library.add(
  faHome,
  faGlobe,
  faTerminal,
  faTools,
  faQuestionCircle,
  faInfoCircle
);

// 首页功能区块
const homeFeatures = [
  {
    id: "network-tools",
    title: "网络工具",
    icon: "globe",
    description:
      "提供各种网络诊断、监控和分析工具，包括Ping、路由追踪、DNS查询等",
  },
  {
    id: "terminal-tools",
    title: "终端工具",
    icon: "terminal",
    description:
      "集成PowerShell和CMD终端，自动配置工具环境变量，让您更高效地使用命令行工具",
  },
  {
    id: "custom-tools",
    title: "自定义工具集成",
    icon: "tools",
    description: "支持将您的自定义脚本和工具整合到工具箱中，统一管理和使用",
  },
];
</script>

<template>
  <div class="home-container">
    <!-- 首页内容 -->
    <div class="home-section">
      <div class="welcome-section">
        <h1>欢迎使用运维工具箱</h1>
        <p class="intro-text">
          这是一个集成了各种网络工具和终端工具的一站式运维解决方案，帮助您高效完成日常运维工作。
        </p>

        <div class="version-info">
          <span class="version">v1.0.0</span>
          <span class="update-date">最后更新: 2025/04/26</span>
        </div>
      </div>

      <div class="features-section">
        <h2>功能特点</h2>
        <div class="features-grid">
          <div
            v-for="feature in homeFeatures"
            :key="feature.id"
            class="feature-card"
          >
            <div class="feature-icon">
              <font-awesome-icon :icon="['fas', feature.icon]" />
            </div>
            <div class="feature-content">
              <h3>{{ feature.title }}</h3>
              <p>{{ feature.description }}</p>
            </div>
          </div>
        </div>
      </div>

      <div class="help-section">
        <h2>使用指南</h2>
        <div class="help-cards">
          <div class="help-card">
            <div class="help-icon">
              <font-awesome-icon :icon="['fas', 'info-circle']" />
            </div>
            <div class="help-content">
              <h3>关于工具箱</h3>
              <p>
                运维工具箱是一个集成多种运维工具的桌面应用，通过简单直观的界面提供网络诊断、终端管理等功能。
              </p>
            </div>
          </div>

          <div class="help-card">
            <div class="help-icon">
              <font-awesome-icon :icon="['fas', 'question-circle']" />
            </div>
            <div class="help-content">
              <h3>如何使用</h3>
              <ul>
                <li>使用左侧菜单切换不同的工具分类</li>
                <li>点击工具卡片启动相应的工具</li>
                <li>终端工具已自动配置环境变量，可直接使用工具目录中的命令</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home-container {
  width: 100%;
}

.home-section {
  background-color: #ffffff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  margin-bottom: 20px;
}

/* 首页样式 */
.welcome-section {
  text-align: center;
  padding: 30px 20px;
  margin-bottom: 30px;
  background: linear-gradient(135deg, #6200ee 0%, #8e44ad 100%);
  color: white;
  border-radius: 8px;
}

.welcome-section h1 {
  font-size: 28px;
  margin-bottom: 15px;
  font-weight: 600;
}

.intro-text {
  font-size: 16px;
  max-width: 700px;
  margin: 0 auto 20px;
  line-height: 1.6;
}

.version-info {
  margin-top: 20px;
  font-size: 13px;
  opacity: 0.8;
}

.version-info .version {
  margin-right: 15px;
  font-weight: 500;
}

.features-section,
.help-section {
  margin-bottom: 30px;
}

.features-section h2,
.help-section h2 {
  font-size: 20px;
  margin-bottom: 20px;
  color: #333;
  position: relative;
  padding-bottom: 10px;
}

.features-section h2:after,
.help-section h2:after {
  content: "";
  position: absolute;
  left: 0;
  bottom: 0;
  width: 40px;
  height: 3px;
  background-color: #6200ee;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  margin-bottom: 30px;
  grid-gap: 20px;
}

.feature-card,
.help-card {
  display: flex;
  background-color: #ffffff;
  border: 1px solid #eee;
  border-radius: 8px;
  padding: 20px;
  cursor: pointer;
  transition: all 0.3s;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}

.feature-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.feature-icon,
.help-icon {
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

.feature-icon {
  background-color: rgba(98, 0, 238, 0.1);
  color: #6200ee;
}

.help-icon {
  background-color: rgba(3, 169, 244, 0.1);
  color: #03a9f4;
}

.feature-content,
.help-content {
  flex: 1;
}

.feature-content h3,
.help-content h3 {
  margin: 0 0 8px 0;
  font-size: 16px;
  color: #333;
  font-weight: 500;
}

.feature-content p,
.help-content p {
  margin: 0;
  font-size: 13px;
  color: #666;
  line-height: 1.5;
}

.help-content ul {
  margin: 0;
  padding-left: 18px;
  font-size: 13px;
  color: #666;
  line-height: 1.5;
}

.help-content li {
  margin-bottom: 5px;
}

.help-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  grid-gap: 20px;
}
</style>
