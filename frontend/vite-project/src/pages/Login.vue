<script setup lang="ts">
import { computed, ref } from "vue"
import { useRoute, useRouter } from "vue-router"
import { login, setToken } from "../services/api"

const route = useRoute()
const router = useRouter()

const email = ref("")
const password = ref("")
const isSubmitting = ref(false)
const errorMessage = ref("")

const redirectTarget = computed(() => {
  const value = route.query.redirect
  return typeof value === "string" && value.length ? value : "/"
})

const handleSubmit = async () => {
  if (!email.value.trim() || !password.value.trim()) {
    errorMessage.value = "Please enter both your email and password."
    return
  }

  isSubmitting.value = true
  errorMessage.value = ""

  try {
    const payload = {
      email: email.value.trim(),
      password: password.value
    }

    const response = await login(payload)
    if (!response?.token) {
      throw new Error("Missing token in login response")
    }

    setToken(response.token)
    await router.replace(redirectTarget.value)
  } catch {
    errorMessage.value = "Login failed. Please check your details and try again."
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <div class="login-shell">
    <section class="login-hero">
      <p class="section-label">Secure access</p>
      <h2>Sign in to JEEL!ve and unlock the admin dashboard.</h2>
      <p class="login-hero__text">
        Use your login credentials to manage students, branch seats, and CAP allocation from one protected workspace.
      </p>

      <div class="login-badges">
        <span class="login-badge">Student records</span>
        <span class="login-badge">CAP allocation</span>
        <span class="login-badge">Branch seats</span>
      </div>
    </section>

    <section class="login-card">
      <div class="login-card__head">
        <p class="login-card__kicker">Admin login</p>
        <h3>Welcome back</h3>
        <p class="login-card__copy">Enter your credentials to continue to the app.</p>
      </div>

      <form class="login-form" @submit.prevent="handleSubmit">
        <label class="field">
          <span>Email</span>
          <input
            v-model="email"
            type="email"
            autocomplete="email"
            placeholder="Enter your email"
          />
        </label>

        <label class="field">
          <span>Password</span>
          <input
            v-model="password"
            type="password"
            autocomplete="current-password"
            placeholder="Enter your password"
          />
        </label>

        <p v-if="errorMessage" class="error login-form__error">{{ errorMessage }}</p>

        <button class="primary-btn login-form__submit" type="submit" :disabled="isSubmitting">
          {{ isSubmitting ? "Signing in..." : "Sign in" }}
        </button>
      </form>
    </section>
  </div>
</template>
