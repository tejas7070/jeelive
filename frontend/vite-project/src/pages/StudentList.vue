<script setup lang="ts">
import { getStudents, runCap, deleteStudent } from "../services/api"
import { computed, onMounted, onUnmounted, ref, watch } from "vue"
import type { Student } from "../types/student"
import { useRouter } from "vue-router"
const isCapRunning = ref(false)
let capTimeout: number | undefined

const students = ref<Student[]>([])

const fetchStudents = async () => {
  students.value = (await getStudents()).data
}

watch(isCapRunning, (running) => {
  document.body.style.overflow = running ? "hidden" : ""
})

onUnmounted(() => {
  if (capTimeout !== undefined) {
    window.clearTimeout(capTimeout)
  }
  document.body.style.overflow = ""
})

const handleRunCap = async () => {
  isCapRunning.value = true

  const start = Date.now()

  try {
    await runCap()
    await fetchStudents()
  } finally {
    // ensure minimum 2 sec loading
    const elapsed = Date.now() - start
    const delay = Math.max(0, 2000 - elapsed)

    capTimeout = window.setTimeout(() => {
      isCapRunning.value = false
    }, delay)
  }
}

const router = useRouter()

const handleView = (id: number) => {
  router.push(`/students/${id}`)
}

const handleEdit = (id: number) => {
  router.push(`/students/${id}/edit`)
}


const handleDelete = async (id: number) => {
  const confirmed = confirm("Are you sure you want to delete this student?")
  if (!confirmed) return
  await deleteStudent(id)
  await fetchStudents()
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
      <button class="primary-btn" @click="handleRunCap" :disabled="isCapRunning" :aria-busy="isCapRunning">
        {{ isCapRunning ? "Running CAP..." : "Run CAP" }}
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
              <th>Actions</th>
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
              <td data-label="Allotted">
                <span class="status-pill status-pill--table">{{ s.allotted || "Pending" }}</span>
              </td>
              <td data-label="Preferences">
                <div class="preference-chips">
                  <span v-for="pref in s.preferences" :key="pref" class="mini-chip">
                    {{ pref }}
                  </span>
                </div>
              </td>
              <td data-label="Actions">
                <div class="action-group">
                  <button class="action-btn action-btn--view" @click="handleView(s.id)">
                    View
                  </button>
                  <button class="action-btn action-btn--edit" @click="handleEdit(s.id)">
                    Edit
                  </button>
                  <button class="action-btn action-btn--delete" @click="handleDelete(s.id)">
                    Delete
                  </button>
                </div>
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
  <div
    v-if="isCapRunning"
    class="modal-overlay"
    role="dialog"
    aria-modal="true"
    aria-labelledby="cap-modal-title"
    aria-describedby="cap-modal-subtitle"
  >
    <div class="modal">
      <div class="modal__spinner" aria-hidden="true"></div>
      <div class="modal__copy">
        <p class="modal-title" id="cap-modal-title">Running CAP Allocation</p>
        <p class="modal-sub" id="cap-modal-subtitle">Processing student preferences and updating the live list.</p>
      </div>
    </div>
  </div>
</template>
