<script setup lang="ts">
import { computed, onMounted, ref } from "vue"
import { getBranches } from "../services/api"
import { normalizeBranch, type RawBranch } from "../utils/branch"

const branches = ref<RawBranch[]>([])
const isLoading = ref(true)
const errorMessage = ref("")

const toBranchArray = (value: unknown) => {
  if (Array.isArray(value)) return value as RawBranch[]
  if (value && typeof value === "object") {
    const record = value as Record<string, any>
    if (Array.isArray(record.data)) return record.data as RawBranch[]
    if (Array.isArray(record.branches)) return record.branches as RawBranch[]
  }
  return []
}

const normalizedBranches = computed(() => branches.value.map((branch, index) => normalizeBranch(branch, index)))

const totalSeats = computed(() => normalizedBranches.value.reduce((sum, branch) => sum + branch.seats, 0))
const allocatedSeats = computed(() => normalizedBranches.value.reduce((sum, branch) => sum + branch.allocatedSeats, 0))
const availableBranches = computed(() => normalizedBranches.value.filter((branch) => branch.status === "Available").length)

const loadBranches = async () => {
  isLoading.value = true
  errorMessage.value = ""

  try {
    const response = await getBranches()
    branches.value = toBranchArray(response)
  } catch {
    errorMessage.value = "Unable to load cutoff overview right now."
  } finally {
    isLoading.value = false
  }
}

onMounted(loadBranches)
</script>

<template>
  <div class="cutoff-page">
    <section class="cutoff-hero">
      <div>
        <p class="cutoff-kicker">Configuration</p>
        <h2>Cutoffs</h2>
        <p class="cutoff-copy">
          Review branch cutoff percentiles, seat availability, and allocation pressure at a glance.
        </p>
      </div>
      <button class="cutoff-ghost" type="button" @click="loadBranches">Refresh</button>
    </section>

    <section class="cutoff-metrics">
      <article class="cutoff-metric">
        <strong>{{ normalizedBranches.length }}</strong>
        <span>Branches</span>
      </article>
      <article class="cutoff-metric">
        <strong>{{ totalSeats }}</strong>
        <span>Total Seats</span>
      </article>
      <article class="cutoff-metric">
        <strong>{{ allocatedSeats }}</strong>
        <span>Allocated Seats</span>
      </article>
      <article class="cutoff-metric">
        <strong>{{ availableBranches }}</strong>
        <span>Available Branches</span>
      </article>
    </section>

    <p v-if="errorMessage" class="cutoff-error">{{ errorMessage }}</p>

    <section class="cutoff-card">
      <div class="cutoff-card__head">
        <div>
          <p class="cutoff-section">Cutoff Overview</p>
          <h3>Branch Cutoffs</h3>
        </div>
      </div>

      <div v-if="isLoading" class="cutoff-empty">Loading cutoff overview...</div>

      <div v-else-if="normalizedBranches.length" class="cutoff-grid">
        <article v-for="branch in normalizedBranches" :key="branch.id" class="cutoff-branch">
          <div class="cutoff-branch__top">
            <div>
              <h4>{{ branch.name }}</h4>
              <p>Cutoff percentile and seat position</p>
            </div>
            <span class="cutoff-pill" :class="branch.status === 'Available' ? 'cutoff-pill--ok' : 'cutoff-pill--full'">
              {{ branch.status }}
            </span>
          </div>

          <div class="cutoff-stats">
            <div>
              <span>Cutoff</span>
              <strong>{{ branch.cutoff }}</strong>
            </div>
            <div>
              <span>Seats</span>
              <strong>{{ branch.seats }}</strong>
            </div>
            <div>
              <span>Allocated</span>
              <strong>{{ branch.allocatedSeats }}</strong>
            </div>
            <div>
              <span>Remaining</span>
              <strong>{{ branch.remainingSeats }}</strong>
            </div>
          </div>
        </article>
      </div>

      <div v-else class="cutoff-empty">
        <p class="cutoff-empty__title">No cutoff data yet</p>
        <p class="cutoff-empty__copy">Branches will appear here once they are created.</p>
      </div>
    </section>
  </div>
</template>

<style scoped>
.cutoff-page {
  display: grid;
  gap: 18px;
  padding: 22px 26px 18px;
}

.cutoff-hero,
.cutoff-card {
  border: 1px solid #dde5ef;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.96);
  box-shadow: 0 12px 26px rgba(15, 23, 42, 0.04);
}

.cutoff-hero {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 18px;
  padding: 22px 24px;
}

.cutoff-kicker,
.cutoff-section {
  margin: 0 0 8px;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.16em;
  font-size: 0.7rem;
  font-weight: 800;
}

.cutoff-hero h2,
.cutoff-card__head h3,
.cutoff-branch h4 {
  margin: 0;
  color: #0f172a;
  letter-spacing: -0.04em;
  font-weight: 800;
}

.cutoff-hero h2 {
  font-size: clamp(2rem, 3vw, 2.7rem);
}

.cutoff-copy {
  margin: 8px 0 0;
  color: #667085;
  line-height: 1.65;
}

.cutoff-ghost {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-height: 40px;
  padding: 0 16px;
  border: 1px solid #dbe4ef;
  border-radius: 999px;
  background: #fff;
  color: #334155;
  font-weight: 700;
}

.cutoff-metrics {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 16px;
}

.cutoff-metric {
  display: grid;
  gap: 6px;
  min-height: 110px;
  padding: 18px;
  border: 1px solid #dde5ef;
  border-radius: 18px;
  background: linear-gradient(180deg, #fff 0%, #f9fbff 100%);
}

.cutoff-metric strong {
  color: #0f172a;
  font-size: 2rem;
  line-height: 1;
  letter-spacing: -0.05em;
  font-weight: 800;
}

.cutoff-metric span {
  color: #334155;
  font-size: 0.95rem;
  font-weight: 600;
}

.cutoff-error {
  margin: 0;
  color: #b91c1c;
  font-weight: 600;
}

.cutoff-card {
  display: grid;
  gap: 16px;
  padding: 20px;
}

.cutoff-card__head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 14px;
}

.cutoff-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 14px;
}

.cutoff-branch {
  display: grid;
  gap: 14px;
  padding: 18px;
  border: 1px solid #e2e8f0;
  border-radius: 18px;
  background: #fff;
}

.cutoff-branch__top {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
}

.cutoff-branch__top p {
  margin: 6px 0 0;
  color: #667085;
  font-size: 0.9rem;
}

.cutoff-pill {
  display: inline-flex;
  align-items: center;
  min-height: 28px;
  padding: 5px 10px;
  border-radius: 999px;
  font-size: 0.8rem;
  font-weight: 700;
}

.cutoff-pill--ok {
  border: 1px solid #ccefd6;
  background: #ebf9ef;
  color: #15803d;
}

.cutoff-pill--full {
  border: 1px solid #dbeafe;
  background: #eff6ff;
  color: #2563eb;
}

.cutoff-stats {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 10px;
}

.cutoff-stats div {
  display: grid;
  gap: 4px;
  padding: 12px;
  border: 1px solid #edf1f7;
  border-radius: 14px;
  background: #f8fafc;
}

.cutoff-stats span {
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.08em;
  font-size: 0.66rem;
  font-weight: 800;
}

.cutoff-stats strong {
  color: #0f172a;
  font-size: 1.05rem;
}

.cutoff-empty {
  padding: 28px 8px;
  color: #667085;
}

.cutoff-empty__title {
  margin: 0 0 6px;
  color: #0f172a;
  font-size: 1.05rem;
  font-weight: 700;
}

.cutoff-empty__copy {
  margin: 0;
}

@media (max-width: 1180px) {
  .cutoff-metrics,
  .cutoff-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 760px) {
  .cutoff-page {
    padding: 16px;
  }

  .cutoff-hero,
  .cutoff-card__head,
  .cutoff-branch__top {
    flex-direction: column;
    align-items: flex-start;
  }

  .cutoff-metrics,
  .cutoff-grid,
  .cutoff-stats {
    grid-template-columns: 1fr;
  }
}
</style>
