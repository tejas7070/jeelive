<script setup lang="ts">
import { computed, onMounted, ref } from "vue"
import { useRoute } from "vue-router"
import { getStudentById } from "../services/api"
import type { Student } from "../types/student"

const route = useRoute()

const student = ref<Student | null>(null)
const loading = ref(false)
const error = ref("")

const studentId = Number(route.params.id)

const initial = computed(() => {
  if (!student.value?.name) return ""
  return student.value.name.charAt(0).toUpperCase()
})

const allottedLabel = computed(() => student.value?.allotted || "Pending")

const fetchStudent = async () => {
  loading.value = true
  error.value = ""

  try {
    student.value = await getStudentById(studentId)
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

      <RouterLink to="/" class="back-link">Back to list</RouterLink>
    </div>

    <div v-if="loading" class="student-detail-state">
      Loading student details...
    </div>

    <div v-else-if="error" class="student-detail-state student-detail-state--error">
      {{ error }}
    </div>

    <div v-else-if="student" class="student-detail-card">
      <div class="student-detail-hero">
        <div class="student-avatar" aria-hidden="true">
          {{ initial }}
        </div>

        <div class="student-identity">
          <p class="section-label">Profile summary</p>
          <h3>{{ student.name }}</h3>
          <p class="student-subtitle">
            Allocation snapshot with percentile, allotted branch, and preferences.
          </p>
        </div>
      </div>

      <div class="student-metric-grid">
        <div class="student-metric-card">
          <span class="student-metric-label">Percentile</span>
          <strong class="student-metric-value">{{ student.percentile }}</strong>
        </div>

        <div class="student-metric-card">
          <span class="student-metric-label">Allotted</span>
          <span class="status-pill">{{ allottedLabel }}</span>
        </div>

        <div class="student-metric-card">
          <span class="student-metric-label">Preferences</span>
          <strong class="student-metric-value">{{ student.preferences.length }}</strong>
        </div>
      </div>

      <div class="student-detail-section">
        <div class="student-detail-section__head">
          <h4>Preferences</h4>
          <p>Selected branches for this student.</p>
        </div>

        <div v-if="student.preferences.length" class="preference-chips">
          <span v-for="p in student.preferences" :key="p" class="mini-chip">
            {{ p }}
          </span>
        </div>
        <p v-else class="student-empty-note">No preferences added yet.</p>
      </div>
    </div>
  </div>
</template>
