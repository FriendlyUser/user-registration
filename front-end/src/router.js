import Vue from "vue";
import Router from "vue-router";
import AppHeader from "./layout/AppHeader";
import AppFooter from "./layout/AppFooter";
import Components from "./views/Components.vue";
import Landing from "./views/Landing.vue";
import RefLogin from "./views/RefLogin.vue";
import RefRegister from "./views/RefRegister.vue";
import Profile from "./views/Profile.vue";

// customized components
import Login from "./views/Login.vue";
import Register from "./views/Register.vue";
import CustomHeader from "./layout/CustomAppHeader.vue";
import CustomFooter from "./layout/CustomAppFooter.vue";
Vue.use(Router);

export default new Router({
  linkExactActiveClass: "active",
  routes: [
    {
      path: "/",
      name: "components",
      components: {
        header: AppHeader,
        default: Components,
        footer: AppFooter
      }
    },
    {
      path: "/landing",
      name: "landing",
      components: {
        header: AppHeader,
        default: Landing,
        footer: AppFooter
      }
    },
    {
      path: "/ref/login",
      name: "reflogin",
      components: {
        header: AppHeader,
        default: RefLogin,
        footer: AppFooter
      }
    },
    {
      path: "/ref/register",
      name: "refregister",
      components: {
        header: AppHeader,
        default: RefRegister,
        footer: AppFooter
      }
    },
    {
      path: "/login",
      name: "login",
      components: {
        header: CustomHeader,
        default: Login,
        footer: CustomFooter
      }
    },
    {
      path: "/register",
      name: "register",
      components: {
        header: CustomHeader,
        default: Register,
        footer: CustomFooter
      }
    },
    {
      path: "/profile",
      name: "profile",
      components: {
        header: AppHeader,
        default: Profile,
        footer: AppFooter
      }
    }
  ],
  scrollBehavior: to => {
    if (to.hash) {
      return { selector: to.hash };
    } else {
      return { x: 0, y: 0 };
    }
  }
});
