<script setup>
import { ref } from "vue";
import Sidebar from "./components/Sidebar.vue";
import Header from "./components/Header.vue";
import Home from "./components/Home.vue";
import NetworkTools from "./components/NetworkTools.vue";
import TerminalTools from "./components/TerminalTools.vue";

// 导航栏折叠状态
const collapsed = ref(false);
const activeMenu = ref("home");

// 切换导航栏折叠状态
const toggleSidebar = () => {
  collapsed.value = !collapsed.value;
};

// 菜单切换
const changeMenu = (menu) => {
  activeMenu.value = menu;
};
</script>

<template>
  <div class="container">
    <!-- 侧边导航 -->
    <Sidebar
      :collapsed="collapsed"
      :activeMenu="activeMenu"
      @toggle-sidebar="toggleSidebar"
      @menu-change="changeMenu"
    />

    <div class="main-content">
      <!-- 顶部标题 -->
      <Header :activeMenu="activeMenu" />

      <div class="content-container">
        <!-- 根据菜单选择渲染不同组件 -->
        <Home v-if="activeMenu === 'home'" />
        <NetworkTools v-if="activeMenu === 'network'" />
        <TerminalTools v-if="activeMenu === 'terminal'" />
      </div>
    </div>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: "Microsoft YaHei", Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.container {
  display: flex;
  height: 100vh;
  background-color: #f5f7f9;
}

.main-content {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}

.content-container {
  background-color: #f5f7f9;
  margin-bottom: 20px;
}
</style>
