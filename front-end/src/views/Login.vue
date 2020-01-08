<template>
    <section class="section section-shaped section-lg my-0">
        <div class="shape shape-style-1 bg-gradient-default">
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <span></span>
        </div>
        <div class="container pt-lg-md">
            <div class="row justify-content-center">
                <div class="col-lg-5">
                    <card type="secondary" shadow
                          header-classes="bg-white pb-5"
                          body-classes="px-lg-5 py-lg-5"
                          class="border-0">
                        <template>
                            <div class="text-muted text-center mb-3">
                                <small>Sign in with</small>
                            </div>
                            <div class="btn-wrapper text-center">
                                <base-button type="neutral">
                                    <img slot="icon" src="img/icons/common/github.svg">
                                    Github
                                </base-button>

                                <base-button type="neutral" v-on:click="googleOAuth">
                                    <img slot="icon" src="img/icons/common/google.svg">
                                    Google
                                </base-button>
                            </div>
                        </template>
                        <template>
                            <div class="text-center text-muted mb-4">
                                <small>Or sign in with credentials</small>
                            </div>
                            <form id="loginForm" role="form">
                                <base-input 
                                    alternative
                                    class="mb-3"
                                    v-bind:class="emailInvalid"
                                    placeholder="Email"
                                    addon-left-icon="ni ni-email-83"
                                    v-model="email"
                                    name="email"
                                    v-bind:value="email"
                                >
                                </base-input>
                                <base-input 
                                    alternative
                                    v-bind:class="passInvalid"
                                    type="password"
                                    placeholder="Password"
                                    addon-left-icon="ni ni-lock-circle-open"
                                    v-model="password"
                                    name="password"
                                    v-bind:value="password"
                                >
                                </base-input>
                                <base-checkbox>
                                    Remember me
                                </base-checkbox>
                                <div class="text-center">
                                    <base-button type="primary" class="my-4" v-on:click="loginHandler">Sign In</base-button>
                                </div>
                            </form>
                        </template>
                    </card>
                    <div class="row mt-3">
                        <div class="col-6">
                            <a href="#" class="text-light">
                                <small>Forgot password?</small>
                            </a>
                        </div>
                        <div class="col-6 text-right">
                            <a href="/#/register" class="text-light">
                                <small>Create new account</small>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
<script>
// TODO create different callbacks based on environment variables
export default {
    data() {
        return {
            email: '',
            password: ''
        }
    },
    computed: {
        emailInvalid: function () {
            return {
                // add email password strength checkig
                "has-danger": this.email.length == 0,
                "has-success": this.email.length > 0
            }
        },
        passInvalid: function () {
            return {
                // add email password strength checkig
                "has-danger": this.password.length == 0,
                "has-success": this.password.length > 0
            }
        },
        credValid: function() {
            return this.email.length > 0 && this.password.length > 0
        }
    },
    methods: {
        googleOAuth() {
            window.location.href = `${process.env.VUE_APP_API_URL}/auth/google/login`
        },
        loginHandler() {
            // throw error if password or username is blank
            fetch(`${process.env.VUE_APP_API_URL}/auth/login`,  {
                method: 'POST', // *GET, POST, PUT, DELETE, etc.
                body: JSON.stringify({email: this.email, password: this.password})
            })
            .then(response => {
                return response.json()
            })
            .then(myJson => {
                console.log(myJson)
            })
            .catch(error => {
                console.log(error)
            })
        }
    }
}
</script>
<style>
</style>
