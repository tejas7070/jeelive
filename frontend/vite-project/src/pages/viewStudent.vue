<script setup lang="ts">
import { computed, onMounted, ref } from "vue"
import { useRoute } from "vue-router"
import { getStudentById } from "../services/api"

type StudentPayload = {
  _id?: string
  createdOn?: number
  id?: string
  modifiedOn?: number
  name: string
  percentile: number
  preferences: string[]
}

const route = useRoute()

const student = ref<StudentPayload | null>(null)
const loading = ref(false)
const error = ref("")

const studentId = route.params.id as string

const normalizedStudent = computed(() => {
  if (!student.value) return null
  return student.value
})

const initial = computed(() => {
  const name = normalizedStudent.value?.name
  if (!name) return ""
  return name.charAt(0).toUpperCase()
})

const formattedCreatedOn = computed(() => {
  const value = normalizedStudent.value?.createdOn
  if (!value) return "-"
  const millis = value > 1_000_000_000_000 ? value : value * 1000
  return new Date(millis).toLocaleString()
})

const formattedModifiedOn = computed(() => {
  const value = normalizedStudent.value?.modifiedOn
  if (!value) return "-"
  const millis = value > 1_000_000_000_000 ? value : value * 1000
  return new Date(millis).toLocaleString()
})

const fetchStudent = async () => {
  loading.value = true
  error.value = ""

  try {
    const response = await getStudentById(studentId)
    student.value = Array.isArray(response) ? response[0] ?? null : response
  } catch (e: any) {
    error.value = e.response?.data?.error || "Failed to load student"
  } finally {
    loading.value = false
  }
}

onMounted(fetchStudent)
</script>

<template>
  <div class="page-stack student-detail-shell">
    <div class="page-heading student-detail-heading">
      <div>
        <p class="section-label">Student profile</p>
        <h2>Student Details</h2>
      </div>

      <RouterLink to="/students" class="back-link">Back to list</RouterLink>
    </div>

    <div v-if="loading" class="student-detail-state">
      Loading student details...
    </div>

    <div v-else-if="error" class="student-detail-state student-detail-state--error">
      {{ error }}
    </div>

    <div v-else-if="normalizedStudent" class="student-detail-card">
      <div class="student-detail-hero">
        <div class="student-avatar" aria-hidden="true">
          {{ initial }}
        </div>

        <div class="student-identity">
          <p class="section-label">Profile summary</p>
          <h3>{{ normalizedStudent.name }}</h3>
          <p class="student-subtitle">
            Allocation snapshot with percentile, timestamps, and preferences.
          </p>
        </div>
      </div>

      <div class="student-metric-grid">
        <div class="student-metric-card">
          <span class="student-metric-label">Student ID</span>
          <span class="status-pill status-pill--table">
            {{ normalizedStudent.id || "-" }}
          </span>
        </div>
        <div class="student-metric-card">
          <span class="student-metric-label">Record ID</span>
          <span class="status-pill status-pill--table">
            {{ normalizedStudent._id || "-" }}
          </span>
        </div>
        <div class="student-metric-card">
          <span class="student-metric-label">Percentile</span>
          <strong class="student-metric-value">{{ normalizedStudent.percentile }}</strong>
        </div>

        <div class="student-metric-card">
          <span class="student-metric-label">Preferences</span>
          <strong class="student-metric-value">{{ normalizedStudent.preferences.length }}</strong>
        </div>

        <div class="student-metric-card">
          <span class="student-metric-label">Created on</span>
          <strong class="student-metric-value">{{ formattedCreatedOn }}</strong>
        </div>
        <div class="student-metric-card">
          <span class="student-metric-label">Modified on</span>
          <strong class="student-metric-value">{{ formattedModifiedOn }}</strong>
        </div>
      </div>

      <div class="student-detail-section">
        <div class="student-detail-section__head">
          <h4>Preferences</h4>
          <p>Selected branches for this student.</p>
        </div>

        <div v-if="normalizedStudent.preferences.length" class="preference-chips">
          <span v-for="pref in normalizedStudent.preferences" :key="pref" class="mini-chip">
            {{ pref }}
          </span>
        </div>
        <p v-else class="student-empty-note">No preferences added yet.</p>
      </div>
    </div>
  </div>
</template>
