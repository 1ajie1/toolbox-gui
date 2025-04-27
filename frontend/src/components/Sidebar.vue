<script setup>
import { defineProps, defineEmits } from "vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faTools,
  faGlobe,
  faTerminal,
  faHome,
} from "@fortawesome/free-solid-svg-icons";

// 添加图标到库中
library.add(faTools, faGlobe, faTerminal, faHome);

// 定义组件属性
const props = defineProps({
  collapsed: Boolean,
  activeMenu: String,
});

// 定义事件
const emit = defineEmits(["toggle-sidebar", "menu-change"]);

// 切换导航栏折叠状态
const toggleSidebar = () => {
  emit("toggle-sidebar");
};

// 菜单切换
const changeMenu = (menu) => {
  emit("menu-change", menu);
};
</script>

<template>
  <div class="sidebar" :class="{ collapsed: collapsed }">
    <div class="logo">
      <font-awesome-icon
        :icon="['fas', 'tools']"
        class="logo-icon"
        v-show="!collapsed"
      />
      <span class="logo-text" v-show="!collapsed">运维工具箱</span>
      <div class="toggle-btn" @click="toggleSidebar">
        {{ collapsed ? ">" : "<" }}
      </div>
    </div>
    <div
      class="menu-item"
      :class="{ active: activeMenu === 'home' }"
      @click="changeMenu('home')"
    >
      <font-awesome-icon :icon="['fas', 'home']" class="menu-icon" />
      <span class="menu-text" v-show="!collapsed">首页</span>
    </div>
    <div
      class="menu-item"
      :class="{ active: activeMenu === 'network' }"
      @click="changeMenu('network')"
    >
      <font-awesome-icon :icon="['fas', 'globe']" class="menu-icon" />
      <span class="menu-text" v-show="!collapsed">网络工具</span>
    </div>
    <div
      class="menu-item"
      :class="{ active: activeMenu === 'terminal' }"
      @click="changeMenu('terminal')"
    >
      <font-awesome-icon :icon="['fas', 'terminal']" class="menu-icon" />
      <span class="menu-text" v-show="!collapsed">终端工具</span>
    </div>
  </div>
</template>

<style scoped>
.sidebar {
  width: 220px;
  background-color: #ffffff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
  padding: 20px 0;
  display: flex;
  flex-direction: column;
  position: relative;
  transition: width 0.3s ease;
}

.sidebar.collapsed {
  width: 70px;
}

.logo {
  display: flex;
  align-items: center;
  padding: 0 20px;
  margin-bottom: 30px;
}

.logo-icon {
  font-size: 24px;
}

.logo-text {
  font-size: 18px;
  font-weight: bold;
  margin-left: 10px;
  color: #6200ee;
  white-space: nowrap;
}

.menu-item {
  display: flex;
  align-items: center;
  padding: 12px 20px;
  cursor: pointer;
  transition: all 0.3s;
}

.menu-item:hover {
  background-color: #f0f0f7;
}

.menu-item.active {
  background-color: #f0f0f7;
  border-left: 3px solid #6200ee;
}

.menu-icon {
  font-size: 20px;
  margin-right: 10px;
}

.menu-text {
  white-space: nowrap;
}

.toggle-btn {
  margin-left: 12px;
  width: 32px;
  height: 32px;
  color: #6200ee;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  font-weight: bold;
  font-size: 18px;
  border: none;
  outline: none;
  background: #fff;
  transition: background 0.2s, box-shadow 0.2s, color 0.2s, margin 0.2s;
}

.toggle-btn:hover {
  color: #fff;
  background: #6200ee;
  box-shadow: 0 4px 16px rgba(98, 0, 238, 0.18);
}

.sidebar.collapsed .logo {
  justify-content: center;
}

.sidebar.collapsed .toggle-btn {
  margin-left: 0;
}
</style>
