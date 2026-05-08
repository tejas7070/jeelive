<script setup lang="ts">
import { computed, onMounted, reactive, ref } from "vue"
import { addBranch, deleteBranch, getBranch, getBranches, updateBranch } from "../services/api"
import { getBranchDisplayValue, normalizeBranch, type NormalizedBranch, type RawBranch } from "../utils/branch"

type BranchForm = {
  name: string
  seats: number
  cutoff: number
}

type ModalMode = "add" | "edit" | "view"

const branches = ref<RawBranch[]>([])
const normalizedBranches = computed(() => branches.value.map((branch, index) => normalizeBranch(branch, index)))
const isLoading = ref(true)
const isSaving = ref(false)
const errorMessage = ref("")
const modalOpen = ref(false)
const modalMode = ref<ModalMode>("add")
const selectedBranchId = ref("")
const modalLoading = ref(false)

const form = reactive<BranchForm>({
  name: "",
  seats: 0,
  cutoff: 0
})

const resetForm = () => {
  form.name = ""
  form.seats = 0
  form.cutoff = 0
}

const toBranchArray = (value: unknown) => {
  if (Array.isArray(value)) return value as RawBranch[]
  if (value && typeof value === "object") {
    const record = value as Record<string, any>
    if (Array.isArray(record.data)) return record.data as RawBranch[]
    if (Array.isArray(record.branches)) return record.branches as RawBranch[]
  }
  return []
}

const loadBranches = async () => {
  isLoading.value = true
  errorMessage.value = ""

  try {
    const response = await getBranches()
    branches.value = toBranchArray(response)
  } catch {
    errorMessage.value = "Unable to load branches right now."
  } finally {
    isLoading.value = false
  }
}

const openAddModal = () => {
  modalMode.value = "add"
  selectedBranchId.value = ""
  resetForm()
  modalOpen.value = true
}

const openEditModal = async (branch: NormalizedBranch) => {
  modalMode.value = "edit"
  selectedBranchId.value = branch.id
  modalOpen.value = true
  modalLoading.value = true

  try {
    const response = await getBranch(branch.id)
    const branchData = Array.isArray(response) ? response[0] ?? response : response
    form.name = getBranchDisplayValue(branchData, ["name", "branchName", "title", "branch"], branch.name)
    form.seats = Number(
      getBranchDisplayValue(branchData, ["seats", "totalSeats", "capacity", "strength"], String(branch.seats))
    )
    form.cutoff = Number(
      getBranchDisplayValue(branchData, ["cutoff", "cutoffPercentile", "percentileCutoff"], String(branch.cutoff))
    )
  } finally {
    modalLoading.value = false
  }
}

const openViewModal = async (branch: NormalizedBranch) => {
  modalMode.value = "view"
  selectedBranchId.value = branch.id
  modalOpen.value = true
  modalLoading.value = true

  try {
    const response = await getBranch(branch.id)
    const branchData = Array.isArray(response) ? response[0] ?? response : response
    form.name = getBranchDisplayValue(branchData, ["name", "branchName", "title", "branch"], branch.name)
    form.seats = Number(
      getBranchDisplayValue(branchData, ["seats", "totalSeats", "capacity", "strength"], String(branch.seats))
    )
    form.cutoff = Number(
      getBranchDisplayValue(branchData, ["cutoff", "cutoffPercentile", "percentileCutoff"], String(branch.cutoff))
    )
  } finally {
    modalLoading.value = false
  }
}

const closeModal = () => {
  modalOpen.value = false
  modalLoading.value = false
  selectedBranchId.value = ""
  resetForm()
}

const saveBranch = async () => {
  if (!form.name.trim() || form.seats < 0) {
    return
  }

  isSaving.value = true
  errorMessage.value = ""

  const payload = {
    name: form.name.trim(),
    seats: Number(form.seats),
    cutoff: Number(form.cutoff)
  }

  try {
    if (modalMode.value === "edit" && selectedBranchId.value) {
      await updateBranch(selectedBranchId.value, payload)
    } else {
      await addBranch(payload)
    }

    await loadBranches()
    closeModal()
  } catch {
    errorMessage.value = "Unable to save this branch right now."
  } finally {
    isSaving.value = false
  }
}

const handleDelete = async (branch: NormalizedBranch) => {
  const confirmed = window.confirm(`Delete ${branch.name}?`)
  if (!confirmed) return

  isSaving.value = true
  errorMessage.value = ""

  try {
    await deleteBranch(branch.id)
    await loadBranches()
  } catch {
    errorMessage.value = "Unable to delete this branch right now."
  } finally {
    isSaving.value = false
  }
}

const totalBranches = computed(() => normalizedBranches.value.length)
const totalSeats = computed(() =>
  normalizedBranches.value.reduce((sum, branch) => sum + branch.seats, 0)
)
const allocatedSeats = computed(() =>
  normalizedBranches.value.reduce((sum, branch) => sum + branch.allocatedSeats, 0)
)
const remainingSeats = computed(() =>
  normalizedBranches.value.reduce((sum, branch) => sum + branch.remainingSeats, 0)
)

onMounted(loadBranches)
</script>

<template>
  <div class="branch-page">
    <section class="branch-hero">
      <div>
        <p class="branch-kicker">Configuration</p>
        <h2>Branches</h2>
        <p class="branch-copy">
          Manage branches, seat capacity, and cutoff thresholds from one clean workspace.
        </p>
      </div>

      <button class="branch-primary" type="button" @click="openAddModal">Add Branch</button>
    </section>

    <section class="branch-metrics" aria-label="Branch summary">
      <article class="branch-metric">
        <strong>{{ totalBranches }}</strong>
        <span>Total Branches</span>
      </article>
      <article class="branch-metric">
        <strong>{{ totalSeats }}</strong>
        <span>Total Seats</span>
      </article>
      <article class="branch-metric">
        <strong>{{ allocatedSeats }}</strong>
        <span>Allocated Seats</span>
      </article>
      <article class="branch-metric">
        <strong>{{ remainingSeats }}</strong>
        <span>Remaining Seats</span>
      </article>
    </section>

    <p v-if="errorMessage" class="branch-error">{{ errorMessage }}</p>

    <section class="branch-card">
      <div class="branch-card__head">
        <div>
          <p class="branch-section">Branch Overview</p>
          <h3>Branches Table</h3>
        </div>
        <button class="branch-ghost" type="button" @click="loadBranches">Refresh</button>
      </div>

      <div class="branch-table-toolbar">
        <p class="branch-table-toolbar__copy">
          Showing <strong>{{ normalizedBranches.length }}</strong> branches with live seat allocation data.
        </p>
        <div class="branch-table-toolbar__chips">
          <span class="branch-toolbar-chip">{{ totalSeats }} seats</span>
          <span class="branch-toolbar-chip branch-toolbar-chip--soft">{{ allocatedSeats }} allocated</span>
          <span class="branch-toolbar-chip branch-toolbar-chip--soft">{{ remainingSeats }} remaining</span>
        </div>
      </div>

      <div v-if="isLoading" class="branch-empty">Loading branches...</div>

      <div v-else-if="normalizedBranches.length" class="branch-table-wrap">
        <table class="branch-table">
          <thead>
            <tr>
              <th>Branch Name</th>
              <th>Total Seats</th>
              <th>Allocated Seats</th>
              <th>Remaining Seats</th>
              <th>Cutoff</th>
              <th>Status</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="branch in normalizedBranches" :key="branch.id">
              <td data-label="Branch Name">
                <div class="branch-name-block">
                  <div class="branch-name">{{ branch.name }}</div>
                  <div class="branch-name-block__meta">Branch ID: {{ branch.id }}</div>
                </div>
              </td>
              <td data-label="Total Seats">
                <span class="branch-figure">{{ branch.seats }}</span>
              </td>
              <td data-label="Allocated Seats">
                <span class="branch-figure branch-figure--allocated">{{ branch.allocatedSeats }}</span>
              </td>
              <td data-label="Remaining Seats">
                <span class="branch-figure branch-figure--remaining">{{ branch.remainingSeats }}</span>
              </td>
              <td data-label="Cutoff">
                <span class="branch-cutoff">{{ branch.cutoff }}</span>
              </td>
              <td data-label="Status">
                <span class="branch-pill" :class="branch.status === 'Available' ? 'branch-pill--ok' : 'branch-pill--full'">
                  {{ branch.status }}
                </span>
              </td>
              <td data-label="Actions">
                <div class="branch-actions">
                  <button class="branch-action-btn branch-action-btn--view" type="button" @click="openViewModal(branch)">
                    <svg aria-hidden="true" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                      <path d="M2.5 12s3.5-6.5 9.5-6.5 9.5 6.5 9.5 6.5-3.5 6.5-9.5 6.5S2.5 12 2.5 12Z" />
                      <circle cx="12" cy="12" r="2.5" />
                    </svg>
                  </button>
                  <button class="branch-action-btn branch-action-btn--edit" type="button" @click="openEditModal(branch)">
                    <svg aria-hidden="true" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                      <path d="M4 20h4l10-10-4-4L4 16v4Z" />
                      <path d="M14 6l4 4" />
                    </svg>
                  </button>
                  <button class="branch-action-btn branch-action-btn--danger" type="button" @click="handleDelete(branch)">
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

      <div v-else class="branch-empty">
        <p class="branch-empty__title">No branches yet</p>
        <p class="branch-empty__copy">Add your first branch to start managing seats and cutoffs.</p>
      </div>
    </section>

    <div v-if="modalOpen" class="branch-modal" role="dialog" aria-modal="true">
      <div class="branch-modal__backdrop" @click="closeModal"></div>
      <div class="branch-modal__panel">
        <div class="branch-modal__head">
          <div>
            <p class="branch-section">
              {{ modalMode === 'add' ? 'Add Branch' : modalMode === 'edit' ? 'Edit Branch' : 'Branch Details' }}
            </p>
            <h3>{{ modalMode === 'add' ? 'Create a branch' : form.name || 'Branch' }}</h3>
          </div>
          <button class="branch-ghost" type="button" @click="closeModal">Close</button>
        </div>

        <div v-if="modalLoading" class="branch-empty">Loading branch details...</div>

        <form v-else class="branch-form" @submit.prevent="saveBranch">
          <label class="branch-field">
            <span>Name</span>
            <input v-model.trim="form.name" type="text" placeholder="Computer Science" :disabled="modalMode === 'view'" />
          </label>

          <label class="branch-field">
            <span>Seats</span>
            <input v-model.number="form.seats" type="number" min="0" step="1" :disabled="modalMode === 'view'" />
          </label>

          <label class="branch-field">
            <span>Cutoff</span>
            <input v-model.number="form.cutoff" type="number" min="0" max="100" step="0.01" :disabled="modalMode === 'view'" />
          </label>

          <div class="branch-modal__stats">
            <div>
              <span>Remaining Seats</span>
              <strong>{{ Math.max(form.seats - (normalizedBranches.find((item) => item.id === selectedBranchId)?.allocatedSeats ?? 0), 0) }}</strong>
            </div>
            <div>
              <span>Status</span>
              <strong>
                {{ Math.max(form.seats - (normalizedBranches.find((item) => item.id === selectedBranchId)?.allocatedSeats ?? 0), 0) > 0 ? 'Available' : 'Full' }}
              </strong>
            </div>
          </div>

          <div class="branch-modal__actions">
            <button class="branch-ghost" type="button" @click="closeModal">Cancel</button>
            <button v-if="modalMode !== 'view'" class="branch-primary" type="submit" :disabled="isSaving">
              {{ isSaving ? 'Saving...' : modalMode === 'add' ? 'Create Branch' : 'Save Changes' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.branch-page {
  display: grid;
  gap: 18px;
  padding: 22px 26px 18px;
}

.branch-hero,
.branch-card {
  border: 1px solid #dde5ef;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.96);
  box-shadow: 0 12px 26px rgba(15, 23, 42, 0.04);
}

.branch-hero {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 18px;
  padding: 22px 24px;
}

.branch-kicker,
.branch-section {
  margin: 0 0 8px;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.16em;
  font-size: 0.7rem;
  font-weight: 800;
}

.branch-hero h2,
.branch-card__head h3,
.branch-modal__head h3 {
  margin: 0;
  color: #0f172a;
  letter-spacing: -0.04em;
  font-weight: 800;
}

.branch-hero h2 {
  font-size: clamp(2rem, 3vw, 2.7rem);
}

.branch-copy {
  margin: 8px 0 0;
  color: #667085;
  line-height: 1.65;
}

.branch-primary,
.branch-ghost,
.branch-icon-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-height: 40px;
  padding: 0 16px;
  border-radius: 999px;
  font-weight: 700;
}

.branch-primary {
  border: 1px solid #0f172a;
  background: #0f172a;
  color: #fff;
}

.branch-ghost,
.branch-icon-btn {
  border: 1px solid #dbe4ef;
  background: #fff;
  color: #334155;
}

.branch-metrics {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 16px;
}

.branch-metric {
  display: grid;
  gap: 6px;
  min-height: 110px;
  padding: 18px;
  border: 1px solid #dde5ef;
  border-radius: 18px;
  background: linear-gradient(180deg, #fff 0%, #f9fbff 100%);
}

.branch-metric strong {
  color: #0f172a;
  font-size: 2rem;
  line-height: 1;
  letter-spacing: -0.05em;
  font-weight: 800;
}

.branch-metric span {
  color: #334155;
  font-size: 0.95rem;
  font-weight: 600;
}

.branch-error {
  margin: 0;
  color: #b91c1c;
  font-weight: 600;
}

.branch-card {
  display: grid;
  gap: 16px;
  padding: 20px;
}

.branch-card__head,
.branch-modal__head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 14px;
}

.branch-table-wrap {
  overflow: hidden;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  background: #fff;
}

.branch-table {
  width: 100%;
  border-collapse: collapse;
}

.branch-table thead th {
  padding: 16px;
  text-align: left;
  color: #64748b;
  background: linear-gradient(180deg, #fff 0%, #fafcff 100%);
  border-bottom: 1px solid #e8edf4;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  font-size: 0.66rem;
  font-weight: 800;
}

.branch-table tbody td {
  padding: 10px;
  border-bottom: 1px solid #edf1f7;
  color: #334155;
  vertical-align: middle;
}

.branch-table tbody tr:last-child td {
  border-bottom: 0;
}

.branch-table tbody tr {
  transition:
    background 0.18s ease,
    transform 0.18s ease,
    box-shadow 0.18s ease;
}

.branch-table tbody tr:hover {
  background: #fcfdff;
}

.branch-table tbody tr:hover td {
  background: #fcfdff;
}

.branch-name {
  color: #0f172a;
  font-weight: 700;
}

.branch-name-block {
  display: grid;
  gap: 4px;
}

.branch-name-block__meta {
  color: #94a3b8;
  font-size: 0.82rem;
}

.branch-figure {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 52px;
  min-height: 30px;
  padding: 0 10px;
  border-radius: 999px;
  background: #eff6ff;
  border: 1px solid #dbeafe;
  color: #2563eb;
  font-weight: 700;
}

.branch-figure--allocated {
  background: #ecfdf3;
  border-color: #ccefd6;
  color: #15803d;
}

.branch-figure--remaining {
  background: #f8fafc;
  border-color: #e2e8f0;
  color: #334155;
}

.branch-cutoff {
  display: inline-flex;
  align-items: center;
  min-height: 30px;
  padding: 0 10px;
  border-radius: 999px;
  background: #f5f3ff;
  border: 1px solid #e9ddff;
  color: #7c3aed;
  font-weight: 700;
}

.branch-pill {
  display: inline-flex;
  align-items: center;
  min-height: 28px;
  padding: 5px 10px;
  border-radius: 999px;
  font-size: 0.8rem;
  font-weight: 700;
}

.branch-pill--ok {
  border: 1px solid #ccefd6;
  background: #ebf9ef;
  color: #15803d;
}

.branch-pill--full {
  border: 1px solid #dbeafe;
  background: #eff6ff;
  color: #2563eb;
}

.branch-actions {
  display: inline-grid;
  grid-auto-flow: column;
  gap: 8px;
  justify-content: end;
  align-items: center;
}

.branch-action-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  min-height: 34px;
  padding: 0 12px;
  border: 1px solid #dbe4ef;
  border-radius: 999px;
  background: #fff;
  color: #334155;
  font-size: 0.84rem;
  font-weight: 700;
  white-space: nowrap;
  box-shadow: 0 6px 14px rgba(15, 23, 42, 0.03);
  transition:
    transform 0.18s ease,
    background 0.18s ease,
    border-color 0.18s ease,
    color 0.18s ease,
    box-shadow 0.18s ease;
}

.branch-action-btn svg {
  width: 14px;
  height: 14px;
}

.branch-action-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 10px 18px rgba(15, 23, 42, 0.06);
}

.branch-action-btn--view {
  color: #2563eb;
  border-color: rgba(37, 99, 235, 0.16);
  background: #f8fbff;
}

.branch-action-btn--edit {
  color: #334155;
}

.branch-action-btn--danger {
  color: #dc2626;
  border-color: rgba(239, 68, 68, 0.18);
  background: #fff7f7;
}

.branch-table-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 14px;
  padding: 0 4px;
}

.branch-table-toolbar__copy {
  margin: 0;
  color: #667085;
  font-size: 0.93rem;
  line-height: 1.5;
}

.branch-table-toolbar__copy strong {
  color: #0f172a;
  font-weight: 800;
}

.branch-table-toolbar__chips {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.branch-toolbar-chip {
  display: inline-flex;
  align-items: center;
  min-height: 32px;
  padding: 0 12px;
  border: 1px solid #dbe4ef;
  border-radius: 999px;
  background: #fff;
  color: #334155;
  font-size: 0.86rem;
  font-weight: 700;
}

.branch-toolbar-chip--soft {
  background: #f8fafc;
}

.branch-empty {
  padding: 28px 8px;
  color: #667085;
}

.branch-empty__title {
  margin: 0 0 6px;
  color: #0f172a;
  font-size: 1.05rem;
  font-weight: 700;
}

.branch-empty__copy {
  margin: 0;
}

.branch-modal {
  position: fixed;
  inset: 0;
  z-index: 40;
  display: grid;
  place-items: center;
  padding: 20px;
}

.branch-modal__backdrop {
  position: absolute;
  inset: 0;
  background: rgba(15, 23, 42, 0.28);
  backdrop-filter: blur(4px);
}

.branch-modal__panel {
  position: relative;
  z-index: 1;
  width: min(680px, 100%);
  display: grid;
  gap: 18px;
  padding: 22px;
  border: 1px solid #dde5ef;
  border-radius: 22px;
  background: #fff;
  box-shadow: 0 24px 60px rgba(15, 23, 42, 0.18);
}

.branch-form {
  display: grid;
  gap: 14px;
}

.branch-field {
  display: grid;
  gap: 8px;
}

.branch-field span {
  color: #475569;
  font-size: 0.9rem;
  font-weight: 700;
}

.branch-field input {
  width: 100%;
  min-height: 46px;
  padding: 0 14px;
  border: 1px solid #dbe4ef;
  border-radius: 14px;
  background: #fff;
  color: #0f172a;
}

.branch-field input:disabled {
  background: #f8fafc;
  color: #475569;
}

.branch-modal__stats {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 12px;
  padding: 14px;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  background: #f8fafc;
}

.branch-modal__stats span {
  display: block;
  color: #64748b;
  font-size: 0.82rem;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  font-weight: 800;
}

.branch-modal__stats strong {
  color: #0f172a;
  font-size: 1rem;
}

.branch-modal__actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

@media (max-width: 1180px) {
  .branch-metrics {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 760px) {
  .branch-page {
    padding: 16px;
  }

  .branch-hero,
  .branch-card__head,
  .branch-table-toolbar,
  .branch-modal__head {
    flex-direction: column;
    align-items: flex-start;
  }

  .branch-metrics,
  .branch-modal__stats {
    grid-template-columns: 1fr;
  }

  .branch-table thead {
    display: none;
  }

  .branch-table,
  .branch-table tbody,
  .branch-table tr,
  .branch-table td {
    display: block;
    width: 100%;
  }

  .branch-table tbody tr {
    border-top: 1px solid #edf1f7;
  }

  .branch-table tbody td {
    display: flex;
    justify-content: space-between;
    gap: 16px;
    padding: 12px 16px;
  }

  .branch-table-toolbar__chips,
  .branch-actions {
    width: 100%;
    justify-content: flex-start;
    grid-auto-flow: row;
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }

  .branch-action-btn {
    justify-content: center;
    width: 100%;
  }

  .branch-table tbody td::before {
    content: attr(data-label);
    color: #64748b;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    font-size: 0.66rem;
    font-weight: 800;
  }
}
</style>
