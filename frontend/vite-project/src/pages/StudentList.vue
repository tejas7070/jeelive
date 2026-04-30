<script setup lang="ts">
import { getStudents, runCap } from "../services/api"
import { computed, onMounted, ref } from "vue"
import type { Student } from "../types/student"

const students = ref<Student[]>([])
const loading = ref(false)

const fetchStudents = async () => {
  students.value = (await getStudents()).data
}

const handleRunCap = async () => {
  loading.value = true
  try {
    await runCap()
    await fetchStudents()
  } finally {
    loading.value = false
  }
}

const studentCount = computed(() => students.value.length)

onMounted(fetchStudents)
</script>

<template>
  <div class="page-stack student-list-shell">
    <div class="page-heading">
      <div>
        <p class="section-label">Live list</p>
        <h2>Students ({{ studentCount }})</h2>
      </div>
      <button class="primary-btn" @click="handleRunCap" :disabled="loading">
        {{ loading ? "Running..." : "Run CAP" }}
      </button>
    </div>

    <div v-if="students.length" class="student-table-shell">

      <div class="student-table-scroll">
        <table class="student-table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Percentile</th>
              <th>Allotted</th>
              <th>Preferences</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="s in students" :key="s.id">
              <td data-label="Name">
                <div class="student-cell__name">{{ s.name }}</div>
              </td>
              <td data-label="Percentile">
                <span class="student-value">{{ s.percentile }}</span>
              </td>
              <td data-label="Preferences">
                <div class="preference-chips">
                  <span v-for="pref in s.preferences" :key="pref" class="mini-chip">
                    {{ pref }}
                  </span>
                </div>
              </td>
              <td data-label="Allotted">
                <span class="status-pill status-pill--table">{{ s.allotted || "Pending" }}</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div v-else class="empty-state">
      <p class="empty-state__title">No students yet</p>
      <p class="empty-state__copy">Add a student to start the allocation flow.</p>
    </div>
  </div>
</template>
