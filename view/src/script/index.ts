import '../style/main.scss'
import "babel-polyfill";
import Vue from "vue";
import App from "./components/app.vue";

new Vue({
    el: "#app",
    render: h => h(App)
});
