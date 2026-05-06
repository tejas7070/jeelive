<script setup lang="ts">
import { computed, onMounted, ref } from "vue"
import { BRANCHES } from "../constants/preference"
import { getStats } from "../services/api"

type SeatStat = {
  branch: string
  total: number
  filled: number
  remaining: number
}

const branchStats = ref<SeatStat[]>([])
const totalSeats = ref(0)
const totalFilled = ref(0)
const isLoadingStats = ref(true)
const statsError = ref("")

const loadStats = async () => {
  isLoadingStats.value = true
  statsError.value = ""

  try {
    const { data } = await getStats()
    const stats = Array.isArray(data.stats) ? data.stats : []
    const order = new Map(BRANCHES.map((branch, index) => [branch, index]))

    branchStats.value = [...stats].sort((a, b) => {
      const aIndex = order.has(a.branch) ? order.get(a.branch)! : Number.MAX_SAFE_INTEGER
      const bIndex = order.has(b.branch) ? order.get(b.branch)! : Number.MAX_SAFE_INTEGER

      if (aIndex !== bIndex) return aIndex - bIndex
      return a.branch.localeCompare(b.branch)
    })
    totalSeats.value = Number(data.total_seats) || 0
    totalFilled.value = Number(data.total_filled) || 0
  } catch {
    statsError.value = "Unable to load seat stats right now."
  } finally {
    isLoadingStats.value = false
  }
}

onMounted(loadStats)

const totalAvailable = computed(() => Math.max(totalSeats.value - totalFilled.value, 0))

const stats = computed(() => [
  {
    label: "Total seats",
    value: String(totalSeats.value).padStart(2, "0"),
    note: "Capacity across all branches"
  },
  {
    label: "Total filled",
    value: String(totalFilled.value).padStart(2, "0"),
    note: "Seats already occupied in the college"
  },
  {
    label: "Available seats",
    value: String(totalAvailable.value).padStart(2, "0"),
    note: "Seats still open"
  }
])

const shortcuts = [
  {
    title: "Students",
    copy: "Open the live roster, monitor records, and review allotments."
  },
  {
    title: "Add Student",
    copy: "Create a new entry with percentile and branch preferences."
  },
  {
    title: "CAP Allocation",
    copy: "Run the allocation flow and refresh the latest student outcomes."
  }
]

const activity = [
  "New student records are ready for review.",
  "CAP allocations were refreshed for the latest batch.",
  "Student profiles reflect the latest branch updates."
]
</script>

<template>
  <div class="page-stack home-shell">
    <section class="dashboard-hero">
      <div class="dashboard-hero__copy">
        <p class="section-label">JEEL!ve admin panel</p>
        <h2>Control student records, CAP allocation, and branch seat outcomes from one place.</h2>
        <p class="dashboard-hero__text">
          JEEL!ve keeps the student roster, preference tracking, and branch seat allocation organized for daily administration.
        </p>

        <div class="home-actions">
          <RouterLink to="/students" class="primary-btn">View Students</RouterLink>
          <RouterLink to="/add" class="back-link">Add Student</RouterLink>
        </div>
      </div>

      <div class="dashboard-hero__panel">
        <div class="dashboard-stats">
          <article v-if="isLoadingStats" class="dashboard-stat-card dashboard-stat-card--loading">
            <span class="student-metric-label">Loading stats</span>
            <strong class="student-metric-value">--</strong>
            <p class="dashboard-stat-note">Fetching seat capacity from the API.</p>
          </article>

          <article v-else-if="statsError" class="dashboard-stat-card dashboard-stat-card--loading">
            <span class="student-metric-label">Stats unavailable</span>
            <strong class="student-metric-value">--</strong>
            <p class="dashboard-stat-note">{{ statsError }}</p>
          </article>

          <template v-else>
            <article v-for="stat in stats" :key="stat.label" class="dashboard-stat-card">
              <span class="student-metric-label">{{ stat.label }}</span>
              <strong class="student-metric-value">{{ stat.value }}</strong>
              <p class="dashboard-stat-note">{{ stat.note }}</p>
            </article>
          </template>
        </div>

        <div class="dashboard-status">
          <div class="dashboard-status__head">
            <span class="dashboard-status__dot"></span>
            <span>System status</span>
          </div>
          <p>JEEL!ve is ready for student management and CAP updates.</p>
        </div>
      </div>
    </section>

    <section class="dashboard-card dashboard-card--stats">
      <div class="dashboard-card__head">
        <div>
          <p class="home-feature-card__eyebrow">Seat overview</p>
          <h3>Branch capacity snapshot</h3>
        </div>
      </div>

      <p v-if="statsError" class="dashboard-inline-message dashboard-inline-message--error">
        {{ statsError }}
      </p>

      <div v-else-if="isLoadingStats" class="dashboard-inline-message">
        Loading branch stats...
      </div>

      <div v-else-if="branchStats.length" class="branch-stats-grid">
        <article v-for="stat in branchStats" :key="stat.branch" class="branch-stat-card">
          <div class="branch-stat-card__head">
            <span class="branch-stat-card__title">{{ stat.branch }}</span>
            <span class="branch-stat-card__badge">{{ stat.remaining }} open</span>
          </div>

          <div class="branch-stat-card__metrics">
            <div>
              <span class="student-metric-label">Total</span>
              <strong class="student-metric-value">{{ stat.total }}</strong>
            </div>
            <div>
              <span class="student-metric-label">Total filled</span>
              <strong class="student-metric-value">{{ stat.filled }}</strong>
            </div>
            <div>
              <span class="student-metric-label">Available</span>
              <strong class="student-metric-value">{{ stat.remaining }}</strong>
            </div>
          </div>
        </article>
      </div>

      <p v-else class="dashboard-inline-message">
        No branch stats available yet.
      </p>
    </section>

    <section class="dashboard-grid">
      <article class="dashboard-card">
        <p class="home-feature-card__eyebrow">Quick access</p>
        <h3>Admin shortcuts</h3>
        <div class="dashboard-shortcuts">
          <div v-for="item in shortcuts" :key="item.title" class="dashboard-shortcut">
            <span class="dashboard-shortcut__title">{{ item.title }}</span>
            <p>{{ item.copy }}</p>
          </div>
        </div>
      </article>

      <article class="dashboard-card">
        <p class="home-feature-card__eyebrow">Activity</p>
        <h3>Latest JEEL!ve updates</h3>
        <ul class="dashboard-activity">
          <li v-for="entry in activity" :key="entry">{{ entry }}</li>
        </ul>
      </article>
    </section>
  </div>
</template>
