<script setup lang="ts">
const stats = [
  { label: "Live students", value: "120", note: "Current roster overview" },
  { label: "Pending CAP", value: "18", note: "Awaiting allocation" },
  { label: "Allocated", value: "102", note: "Updated assignments" },
  { label: "Today’s runs", value: "04", note: "Allocation sessions" }
]

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
        <h2>Control student records, CAP allocation, and branch outcomes from one place.</h2>
        <p class="dashboard-hero__text">
          JEEL!ve keeps the student roster, preference tracking, and allocation workflow organized for daily administration.
        </p>

        <div class="home-actions">
          <RouterLink to="/students" class="primary-btn">View Students</RouterLink>
          <RouterLink to="/add" class="back-link">Add Student</RouterLink>
        </div>
      </div>

      <div class="dashboard-hero__panel">
        <div class="dashboard-stats">
          <article v-for="stat in stats" :key="stat.label" class="dashboard-stat-card">
            <span class="student-metric-label">{{ stat.label }}</span>
            <strong class="student-metric-value">{{ stat.value }}</strong>
            <p class="dashboard-stat-note">{{ stat.note }}</p>
          </article>
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
