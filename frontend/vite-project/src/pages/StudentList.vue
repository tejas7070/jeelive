<script setup lang="ts">
import { getStudents, runCap, deleteStudent } from "../services/api"
import { computed, onMounted, onUnmounted, ref, watch } from "vue"
import type { Student } from "../types/student"
import { useRouter } from "vue-router"
const isCapRunning = ref(false)
let capTimeout: number | undefined

const students = ref<Student[]>([])
const isLoading = ref(false)
const errorMessage = ref("")

const fetchStudents = async () => {
  isLoading.value = true
  errorMessage.value = ""

  try {
    students.value = (await getStudents()).data
  } catch {
    errorMessage.value = "Unable to load students right now."
  } finally {
    isLoading.value = false
  }
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

const handleView = (id: string) => {
  router.push(`/students/${id}`)
}

const handleEdit = (id: string) => {
  router.push(`/students/${id}/edit`)
}


const handleDelete = async (id: string) => {
  const confirmed = confirm("Are you sure you want to delete this student?")
  if (!confirmed) return
  await deleteStudent(id)
  await fetchStudents()
}

const studentCount = computed(() => students.value.length)
const allocatedCount = computed(
  () => students.value.filter((student) => Boolean(student.allotedBranch)).length
)
const pendingCount = computed(
  () => students.value.filter((student) => !student.allotedBranch).length
)

const formatStatus = (status?: string) => {
  if (!status) return "Pending"
  return status.charAt(0).toUpperCase() + status.slice(1).toLowerCase()
}

const getStatusClass = (status?: string) => {
  const normalized = (status || "Pending").trim().toLowerCase()

  if (normalized === "allocated" || normalized === "allotted") {
    return "student-status-pill--allocated"
  }

  if (normalized === "pending") {
    return "student-status-pill--pending"
  }

  return "student-status-pill--neutral"
}

onMounted(fetchStudents)
</script>

<template>
  <div class="page-stack student-list-shell">
    <div class="page-heading">
      <div>
        <p class="section-label">Live list</p>
        <h2>Students ({{ studentCount }})</h2>
        <p class="student-list-subtitle">
          Review student preferences, branch allocation, and CAP status from one place.
        </p>
      </div>
      <button class="primary-btn" @click="handleRunCap" :disabled="isCapRunning" :aria-busy="isCapRunning">
        {{ isCapRunning ? "Running CAP..." : "Run CAP" }}
      </button>
    </div>

    <div class="student-table-shell">
      <div class="student-table-bar">
        <p class="student-table-bar__copy">
          Showing <strong>{{ studentCount }}</strong> students
        </p>
        <div class="student-table-bar__hint">
          <span class="student-table-bar__chip">{{ allocatedCount }} allocated</span>
          <span class="student-table-bar__chip student-table-bar__chip--soft">{{ pendingCount }} pending</span>
        </div>
      </div>

      <p v-if="errorMessage" class="student-list-error">{{ errorMessage }}</p>

      <div v-if="isLoading" class="empty-state">
        <p class="empty-state__title">Loading students...</p>
        <p class="empty-state__copy">Fetching the latest allocation data.</p>
      </div>

      <div v-else-if="students.length" class="student-table-scroll">
        <table class="student-table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Percentile</th>
              <th>Allotted Branch</th>
              <th>Preferences</th>
              <th>Status</th>
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
              <td data-label="Allotted Branch">
                <span class="student-branch-pill">
                  {{ s.allottedBranch || "Pending" }}
                </span>
              </td>
              <td data-label="Preferences">
                <div class="preference-chips">
                  <span v-for="pref in s.preferences" :key="pref" class="mini-chip">
                    {{ pref }}
                  </span>
                </div>
              </td>
              <td data-label="Status">
                <span class="status-pill status-pill--table" :class="getStatusClass(s.status)">
                  {{ formatStatus(s.status) }}
                </span>
              </td>
              <td data-label="Actions">
                <div class="action-group">
                  <button
                    class="action-btn action-btn--view"
                    type="button"
                    aria-label="View student"
                    title="View"
                    @click="handleView(s.id)"
                  >
                    <svg aria-hidden="true" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                      <path d="M2.5 12s3.5-6.5 9.5-6.5 9.5 6.5 9.5 6.5-3.5 6.5-9.5 6.5S2.5 12 2.5 12Z" />
                      <circle cx="12" cy="12" r="2.5" />
                    </svg>
                  </button>
                  <button
                    class="action-btn action-btn--edit"
                    type="button"
                    aria-label="Edit student"
                    title="Edit"
                    @click="handleEdit(s.id)"
                  >
                    <svg aria-hidden="true" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                      <path d="M4 20h4l10-10-4-4L4 16v4Z" />
                      <path d="M14 6l4 4" />
                    </svg>
                  </button>
                  <button
                    class="action-btn action-btn--delete"
                    type="button"
                    aria-label="Delete student"
                    title="Delete"
                    @click="handleDelete(s.id)"
                  >
                    <svg aria-hidden="true" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                      <path d="M4 7h16" />
                      <path d="M6 7l1 13h10l1-13" />
                      <path d="M9 7V4h6v3" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-else class="empty-state">
        <p class="empty-state__title">No students yet</p>
        <p class="empty-state__copy">Add a student to start the allocation flow.</p>
      </div>
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

<style scoped>
.student-list-subtitle {
  margin: 8px 0 0;
  color: var(--text-soft);
  line-height: 1.6;
}

.student-table-bar__copy {
  margin: 0;
  color: var(--text-soft);
  font-size: 0.92rem;
}

.student-table-bar__copy strong {
  color: var(--text);
  font-weight: 800;
}

.student-table-bar__hint {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.student-table-bar__chip {
  display: inline-flex;
  align-items: center;
  min-height: 32px;
  padding: 0 12px;
  border: 1px solid var(--line);
  border-radius: 999px;
  background: #fff;
  color: var(--text);
  font-size: 0.85rem;
  font-weight: 700;
}

.student-table-bar__chip--soft {
  background: #f8fafc;
  color: var(--text-soft);
}

.student-table-scroll {
  overflow: auto;
}

.student-table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0;
  padding: 10px;
}

.student-table thead th {
  position: sticky;
  top: 0;
  z-index: 1;
  text-align: center;
  vertical-align: middle;
  padding: 14px 12px;
  font-size: 0.72rem;
  letter-spacing: 0.12em;
  text-transform: uppercase;
  color: var(--text-soft);
  background: #fff;
  border-bottom: 1px solid var(--line);
}

.student-table tbody td {
  padding: 14px 12px;
  vertical-align: middle;
  border-bottom: 1px solid #edf1f7;
  background: #fff;
  text-align: center;
}

.student-table tbody tr:hover td {
  background: #fafbff;
}

.student-table tbody tr:last-child td {
  border-bottom: 0;
}

.student-cell__name {
  color: var(--text);
  font-weight: 700;
}

.student-value {
  color: var(--text);
  font-weight: 600;
}

.preference-chips {
  justify-content: center;
}

.student-list-error {
  margin: 0;
  color: #b91c1c;
  font-weight: 600;
}

.student-branch-pill,
.student-status-pill--pending,
.student-status-pill--allocated,
.student-status-pill--neutral {
  display: inline-flex;
  align-items: center;
  min-height: 30px;
  padding: 6px 10px;
  border-radius: 999px;
  font-size: 0.82rem;
  font-weight: 700;
}

.student-branch-pill {
  border: 1px solid #dbeafe;
  background: #eff6ff;
  color: #1d4ed8;
}

.student-status-pill--pending {
  border-color: #fde68a;
  background: #fffbeb;
  color: #b45309;
}

.student-status-pill--allocated {
  border-color: #ccefd6;
  background: #ecfdf3;
  color: #15803d;
}

.student-status-pill--neutral {
  border-color: #dbeafe;
  background: #eff6ff;
  color: #2563eb;
}

.action-group {
  display: flex;
  justify-content: center;
  gap: 8px;
}

.action-btn {
  display: inline-grid;
  place-items: center;
  width: 38px;
  height: 38px;
  padding: 0;
  border-radius: 999px;
  border: 1px solid #dbe4ef;
  background: #fff;
  color: #334155;
  box-shadow: 0 6px 14px rgba(15, 23, 42, 0.04);
  transition:
    transform 0.18s ease,
    border-color 0.18s ease,
    background 0.18s ease,
    color 0.18s ease,
    box-shadow 0.18s ease;
}

.action-btn svg {
  width: 16px;
  height: 16px;
}

.action-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 10px 18px rgba(15, 23, 42, 0.08);
}

.action-btn--view {
  color: #2563eb;
  border-color: rgba(37, 99, 235, 0.16);
  background: #f8fbff;
}

.action-btn--edit {
  color: #334155;
}

.action-btn--delete {
  color: #dc2626;
  border-color: rgba(239, 68, 68, 0.18);
  background: #fff7f7;
}

.student-branch-pill,
.student-status-pill--pending,
.student-status-pill--allocated,
.student-status-pill--neutral {
  justify-content: center;
  text-align: center;
}

@media (max-width: 640px) {
  .student-table thead {
    display: none;
  }

  .student-table,
  .student-table tbody,
  .student-table tr,
  .student-table td {
    display: block;
    width: 100%;
  }

  .student-table tbody tr {
    border-bottom: 1px solid #edf1f7;
  }

  .student-table tbody td {
    display: flex;
    justify-content: space-between;
    gap: 14px;
    text-align: right;
  }

  .student-table tbody td::before {
    content: attr(data-label);
    color: var(--text-faint);
    text-transform: uppercase;
    letter-spacing: 0.12em;
    font-size: 0.68rem;
    font-weight: 700;
  }

  .student-table tbody td > * {
    margin-left: auto;
  }

  .action-group {
    margin-left: auto;
  }
}
</style>
