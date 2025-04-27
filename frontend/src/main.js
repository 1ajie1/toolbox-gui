import { createApp } from "vue";
import App from "./App.vue";
import "./style.css";

// 导入 Font Awesome 组件
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

const app = createApp(App);

// 全局注册 Font Awesome 组件
app.component("font-awesome-icon", FontAwesomeIcon);

app.mount("#app");
