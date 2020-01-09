<template>
    <div>

        <div class="position-relative">
            <!-- shape Hero -->
            <section class="section-shaped my-0">
                <div class="shape shape-style-1 shape-default shape-skew">
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                </div>
                <div class="container shape-container d-flex">
                    <div class="col px-0">
                        <div class="row">
                            <div class="col-lg-6">
                                <h1 class="display-3 text-white">User registration system
                                    <span>with login using google</span>
                                </h1>
                                <p class="lead  text-white">This is an example of how to use vue and golang to create a simple user registration system.</p>
                                <div class="btn-wrapper">
                                    <base-button tag="a"
                                                 href="https://github.com/FriendlyUser/user-registration"
                                                 class="mb-3 mb-sm-0"
                                                 type="info"
                                                 icon="fa fa-code">
                                        Code
                                    </base-button>
                                    <base-button tag="a"
                                                 href="https://github.com/FriendlyUser/user-registration/releases"
                                                 class="mb-3 mb-sm-0"
                                                 type="white"
                                                 icon="ni ni-cloud-download-95">
                                        Releases
                                    </base-button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
            <!-- 1st Hero Variation -->
        </div>
        <section class="section section-lg pt-lg-0 mt--200">
            <div class="container">
                <div class="row justify-content-center">
                    <div class="col-lg-12">
                        <div class="row row-grid">
                            <div class="col-lg-4">
                                <card class="border-0" hover shadow body-classes="py-5">
                                    <icon name="ni ni-check-bold" type="primary" rounded class="mb-4">
                                    </icon>
                                    <h6 class="text-primary text-uppercase">Built with Argon</h6>
                                    <p class="description mt-3">Argon is a great free UI package based on Bootstrap 4
                                        that includes the most important components and features.</p>
                                    <div>
                                        <badge type="primary" rounded>design</badge>
                                        <badge type="primary" rounded>system</badge>
                                        <badge type="primary" rounded>creative</badge>
                                    </div>
                                    <base-button tag="a" href="#" type="primary" class="mt-4">
                                        Learn more
                                    </base-button>
                                </card>
                            </div>
                            <div class="col-lg-4">
                                <card class="border-0" hover shadow body-classes="py-5">
                                    <icon name="ni ni-istanbul" type="success" rounded class="mb-4">
                                    </icon>
                                    <h6 class="text-success text-uppercase">Extend From</h6>
                                    <p class="description mt-3">Simple User Registration Example built on postgres, vue and golang.</p>
                                    <div>
                                        <badge type="success" rounded>business</badge>
                                        <badge type="success" rounded>vision</badge>
                                        <badge type="success" rounded>success</badge>
                                    </div>
                                    <base-button tag="a" href="#" type="success" class="mt-4">
                                        Learn more
                                    </base-button>
                                </card>
                            </div>
                            <div class="col-lg-4">
                                <card class="border-0" hover shadow body-classes="py-5">
                                    <icon name="ni ni-planet" type="warning" rounded class="mb-4">
                                    </icon>
                                    <h6 class="text-warning text-uppercase">Prepare Launch</h6>
                                    <p class="description mt-3">Built on open source technolgies with components and features.</p>
                                    <div>
                                        <badge type="warning" rounded>marketing</badge>
                                        <badge type="warning" rounded>product</badge>
                                        <badge type="warning" rounded>launch</badge>
                                    </div>
                                    <base-button tag="a" href="#" type="warning" class="mt-4">
                                        Learn more
                                    </base-button>
                                </card>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
        <section class="section section section-shaped my-0 overflow-hidden">
            <div class="container">
                <h3>List of Users</h3>
                <!-- https://bootstrap-vue.js.org/docs/components/table Update at some point if I feel like it-->
                <div class="row justify-content-center">
                    <b-table hover :items="this.users"></b-table>
                </div>
            </div>
        </section>
    </div>
</template>
<script>
import { BTable } from "bootstrap-vue/esm/components/table";
export default {
    name: "home",
    components: {
        BTable
    },
    data() {
        return {
            email: '',
            password: '',
            users: []
        }
    },
    created() {
        console.log(process.env)
        fetch(`${process.env.VUE_APP_API_URL}/users`, {
            method: 'GET'
        })
        .then(response => {
            return response.json()
        })
        .then(myJson => {
            this.users = myJson
        })
    },
    methods: {
    }
};
</script>
<style>
</style>