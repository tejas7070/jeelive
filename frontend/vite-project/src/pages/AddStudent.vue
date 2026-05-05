<script setup lang="ts">
import { addStudent } from "../services/api"
import { useRouter } from "vue-router"
import { reactive } from "vue"
import { BRANCHES } from "../constants/preference"

const router = useRouter()

const formData = reactive({
  name: "",
  percentile: 0,
  preferences: [] as string[]
})

const errors  = reactive({
  name: "",
  percentile: "",
  preferences: ""
})

const resetErrors = () => {
  errors.name = ""
  errors.percentile = ""
  errors.preferences = ""
}
const isSelected = (branch: string) => formData.preferences.includes(branch)

const togglePreference = (branch: string) => {
  if (branch === "ALL") {
    formData.preferences = isSelected("ALL") ? [] : ["ALL"]
    return
  }

  const next = new Set(formData.preferences)

  if (next.has(branch)) {
    next.delete(branch)
  } else {
    next.delete("ALL")
    next.add(branch)
  }

  formData.preferences = Array.from(next)
}

const submit = async () => {
  resetErrors()
 try{
   await addStudent({
     name: formData.name,
     percentile: formData.percentile,
     preferences: formData.preferences
   })

   router.push("/students")
 } catch (e: any) {
   const msg = e.response?.data?.error || "An error occurred"

   const lower = msg.toLowerCase()

   if (lower.includes("name")) {
     errors.name = msg
   } else if (lower.includes("percentile")) {
     errors.percentile = msg
   } else if (lower.includes("preference")) {
     errors.preferences = msg
   } else {
     alert(msg)
   }
 }
}
</script>

<template>
  <div class="page-stack">
    <div class="page-heading">
      <div>
        <p class="section-label">New entry</p>
        <h2>Add Student</h2>
      </div>
    </div>

    <form class="form-card" @submit.prevent="submit">
      <label class="field">
        <span>Name</span>
        <input v-model.trim="formData.name" placeholder="Student name" />
        <small v-if="errors.name" class="error">
          {{ errors.name }}*
        </small>
      </label>

      <label class="field">
        <span>Percentile</span>
        <input
          v-model.number="formData.percentile"
          type="number"
          min="0"
          max="100"
          step="0.01"
          placeholder="98.4"
        />
        <small v-if="errors.percentile" class="error">
          {{ errors.percentile }}*
        </small>
      </label>

      <div class="field preferences-field">
        <span>Preferences</span>
        <p class="preferences-help">
          Pick one or more branches, or use ALL to lock the selection.
        </p>

        <div class="preference-grid" role="group" aria-label="Preferences">
          <label
            v-for="b in [...BRANCHES, 'ALL']"
            :key="b"
            class="preference-option"
            :class="{ 'preference-option--selected': isSelected(b) }"
          >
            <input
              class="preference-option__input"
              type="checkbox"
              :checked="isSelected(b)"
              @change="togglePreference(b)"
            />

            <span class="preference-option__card">
              <span class="preference-option__title">{{ b }}</span>
              <span class="preference-option__note">
                {{ b === "ALL" ? "Use all preferences" : "Toggle branch choice" }}
              </span>
            </span>
          </label>
        </div>
        <small v-if="errors.preferences" class="error">
          {{ errors.preferences }}*
        </small>
      </div>

      <button class="primary-btn" type="submit">Submit</button>
    </form>
  </div>
</template>
