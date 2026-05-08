<script setup lang="ts">
import { useRouter } from "vue-router"
import { runCap } from "../services/api"

const router = useRouter()

const metrics = [
  {
    label: "Total Students",
    value: "156",
    note: "+ 12 this week",
    tone: "blue"
  },
  {
    label: "Total Branches",
    value: "24",
    note: "3 departments",
    tone: "green"
  },
  {
    label: "CAP Runs",
    value: "8",
    note: "This month",
    tone: "violet"
  },
  {
    label: "Allocated",
    value: "142",
    note: "91% success rate",
    tone: "amber"
  }
] as const

const students = [
  { name: "Tejas", percentile: "96", branch: "Computer Science", status: "Pending", tone: "pending" },
  {
    name: "Vighnesh Berde",
    percentile: "91",
    branch: "Information Technology",
    status: "Pending",
    tone: "pending"
  },
  { name: "harsh", percentile: "88", branch: "Mechanical Engineering", status: "Pending", tone: "pending" },
  { name: "Priya Sharma", percentile: "95", branch: "Computer Science", status: "Allocated", tone: "success" },
  { name: "Rahul Kumar", percentile: "87", branch: "Electrical Engineering", status: "Allocated", tone: "success" }
] as const

const actions = [
  {
    label: "Add New Student",
    tone: "blue",
    icon: "user-plus",
    action: () => router.push("/add")
  },
  {
    label: "Run CAP",
    tone: "green",
    icon: "play",
    action: async () => {
      await runCap()
    }
  },
  {
    label: "View All Students",
    tone: "violet",
    icon: "users",
    action: () => router.push("/students")
  },
  {
    label: "View CAP History",
    tone: "amber",
    icon: "history",
    action: () => router.push("/students")
  }
] as const
</script>

<template>
  <div class="jeel-dashboard">
    <section class="jeel-hero">
      <div class="jeel-hero__copy">
        <p class="jeel-eyebrow">Dashboard</p>
        <h2>Welcome back, Admin! 👋</h2>
        <p class="jeel-hero__text">Here&apos;s what&apos;s happening with your allocation system today.</p>
      </div>

      <div class="jeel-date-pill">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
          <rect x="3" y="5" width="18" height="16" rx="4" />
          <path d="M8 3v4M16 3v4M3 10h18" />
        </svg>
        <span>May 15, 2025</span>
      </div>
    </section>

    <section class="jeel-metrics" aria-label="Summary metrics">
      <article
        v-for="metric in metrics"
        :key="metric.label"
        class="jeel-metric"
        :class="`jeel-metric--${metric.tone}`"
      >
        <div class="jeel-metric__icon" aria-hidden="true">
          <svg v-if="metric.tone === 'blue'" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
            <path d="M16 19v-2a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2" />
            <circle cx="9.5" cy="7.5" r="3.5" />
            <path d="M22 19v-2a4 4 0 0 0-3-3.85" />
            <path d="M16.5 4.7a3.5 3.5 0 0 1 0 6.6" />
          </svg>
          <svg v-else-if="metric.tone === 'green'" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
            <path d="M4 20h16" />
            <path d="M6 20V8l6-4 6 4v12" />
            <path d="M9 20v-6h6v6" />
          </svg>
          <svg v-else-if="metric.tone === 'violet'" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
            <circle cx="12" cy="12" r="8" />
            <path d="M12 8v4l3 2" />
            <path d="M19 5l-2 2" />
          </svg>
          <svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
            <path d="M20 12a8 8 0 1 1-5.2-7.5" />
            <path d="M20 4v6h-6" />
          </svg>
        </div>

        <div class="jeel-metric__copy">
          <strong>{{ metric.value }}</strong>
          <span>{{ metric.label }}</span>
          <p>{{ metric.note }}</p>
        </div>
      </article>
    </section>

    <section class="jeel-grid">
      <article class="jeel-card jeel-card--table">
        <div class="jeel-card__head jeel-card__head--split">
          <h3>Recent Students</h3>
          <button class="jeel-card__button" type="button" @click="router.push('/students')">View All</button>
        </div>

        <div class="jeel-table-wrap">
          <table class="jeel-table">
            <thead>
              <tr>
                <th>Name</th>
                <th>Percentile</th>
                <th>Alloted Branch</th>
                <th>Status</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="student in students" :key="student.name">
                <td data-label="Name">
                  <div class="jeel-table__name">{{ student.name }}</div>
                </td>
                <td data-label="Percentile">{{ student.percentile }}</td>
                <td data-label="Alloted Branch">{{ student.branch }}</td>
                <td data-label="Status">
                  <span class="jeel-pill" :class="student.tone === 'success' ? 'jeel-pill--success' : 'jeel-pill--pending'">
                    {{ student.status }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </article>

      <div class="jeel-side">
        <article class="jeel-card jeel-card--actions">
          <h3>Quick Actions</h3>

          <div class="jeel-action-grid">
            <button
              v-for="action in actions"
              :key="action.label"
              class="jeel-quick-action"
              :class="`jeel-quick-action--${action.tone}`"
              type="button"
              @click="action.action()"
            >
              <span class="jeel-quick-action__icon" aria-hidden="true">
                <svg v-if="action.icon === 'user-plus'" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M16 19v-2a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2" />
                  <circle cx="9.5" cy="7.5" r="3.5" />
                  <path d="M19 8v6M16 11h6" />
                </svg>
                <svg v-else-if="action.icon === 'play'" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M8 5v14l11-7-11-7Z" />
                </svg>
                <svg v-else-if="action.icon === 'users'" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M16 19v-2a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2" />
                  <circle cx="9.5" cy="7.5" r="3.5" />
                  <path d="M22 19v-2a4 4 0 0 0-3-3.85" />
                  <path d="M16.5 4.7a3.5 3.5 0 0 1 0 6.6" />
                </svg>
                <svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M4 12a8 8 0 1 0 2.3-5.7" />
                  <path d="M4 4v4h4" />
                  <path d="M12 8v4l3 2" />
                </svg>
              </span>
              <span class="jeel-quick-action__label">{{ action.label }}</span>
            </button>
          </div>
        </article>

        <article class="jeel-card jeel-card--snapshot">
          <div class="jeel-card__head jeel-card__head--split">
            <h3>Latest CAP Run</h3>
            <button class="jeel-text-link" type="button">View All</button>
          </div>

          <div class="jeel-snapshot__top">
            <div>
              <div class="jeel-snapshot__title">CAP Run #8</div>
              <div class="jeel-snapshot__meta">May 15, 2025 at 10:30 AM</div>
            </div>
            <span class="jeel-pill jeel-pill--success">Completed</span>
          </div>

          <div class="jeel-snapshot__summary">
            <div class="jeel-snapshot__item">
              <strong>142</strong>
              <span>Allocated</span>
            </div>
            <div class="jeel-snapshot__item">
              <strong>8</strong>
              <span>Not Allocated</span>
            </div>
            <div class="jeel-snapshot__item">
              <strong>91%</strong>
              <span>Success Rate</span>
            </div>
          </div>

          <button class="jeel-details-btn" type="button">View Detailed Results</button>
        </article>
      </div>
    </section>

    <footer class="jeel-footer">
      <span>© 2025 JEEL!ve Allocation Hub. All rights reserved.</span>
      <span>Version 1.0.0</span>
    </footer>
  </div>
</template>

<style scoped>
.jeel-dashboard {
  display: grid;
  gap: 24px;
  padding: 22px 26px 18px;
}

.jeel-hero {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
  min-height: 112px;
  padding: 20px 24px;
  border: 1px solid #dde5ef;
  border-radius: 20px;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.96) 0%, rgba(250, 252, 255, 0.95) 100%);
  box-shadow: 0 12px 26px rgba(15, 23, 42, 0.04);
}

.jeel-hero__copy {
  display: grid;
  gap: 8px;
}

.jeel-eyebrow {
  margin: 0;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.16em;
  font-size: 0.7rem;
  font-weight: 800;
}

.jeel-hero h2 {
  margin: 0;
  color: #0f172a;
  font-size: clamp(2rem, 2.6vw, 2.6rem);
  line-height: 1.05;
  letter-spacing: -0.055em;
  font-weight: 800;
}

.jeel-hero__text {
  margin: 0;
  color: #667085;
  font-size: 0.98rem;
  line-height: 1.55;
}

.jeel-date-pill {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  min-width: 166px;
  min-height: 44px;
  padding: 0 16px;
  border: 1px solid #dbe4ef;
  border-radius: 14px;
  background: #fff;
  color: #334155;
  font-size: 0.95rem;
  font-weight: 600;
  box-shadow: 0 8px 18px rgba(15, 23, 42, 0.03);
}

.jeel-date-pill svg {
  width: 16px;
  height: 16px;
  color: #475569;
}

.jeel-metrics {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 16px;
}

.jeel-metric {
  display: grid;
  grid-template-columns: 72px minmax(0, 1fr);
  align-items: center;
  gap: 16px;
  min-height: 120px;
  padding: 16px 18px;
  border: 1px solid #dde5ef;
  border-radius: 18px;
  background: #fff;
}

.jeel-metric--blue {
  background: linear-gradient(180deg, #fafcff 0%, #f4f8ff 100%);
}

.jeel-metric--green {
  background: linear-gradient(180deg, #fbfffc 0%, #f3fbf5 100%);
}

.jeel-metric--violet {
  background: linear-gradient(180deg, #fcfaff 0%, #f6f1ff 100%);
}

.jeel-metric--amber {
  background: linear-gradient(180deg, #fffdf8 0%, #fff7ea 100%);
}

.jeel-metric__icon {
  display: inline-grid;
  place-items: center;
  width: 54px;
  height: 54px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.8);
}

.jeel-metric--blue .jeel-metric__icon {
  color: #2563eb;
  background: rgba(191, 219, 254, 0.45);
}

.jeel-metric--green .jeel-metric__icon {
  color: #16a34a;
  background: rgba(187, 247, 208, 0.48);
}

.jeel-metric--violet .jeel-metric__icon {
  color: #7c3aed;
  background: rgba(221, 214, 254, 0.52);
}

.jeel-metric--amber .jeel-metric__icon {
  color: #f59e0b;
  background: rgba(254, 240, 138, 0.42);
}

.jeel-metric__icon svg {
  width: 18px;
  height: 18px;
}

.jeel-metric__copy {
  display: grid;
  gap: 4px;
}

.jeel-metric__copy strong {
  color: #0f172a;
  font-size: 2.1rem;
  line-height: 1;
  letter-spacing: -0.05em;
  font-weight: 800;
}

.jeel-metric__copy span {
  color: #334155;
  font-size: 0.96rem;
  font-weight: 600;
}

.jeel-metric__copy p {
  margin: 0;
  color: #16a34a;
  font-size: 0.86rem;
  font-weight: 600;
}

.jeel-grid {
  display: grid;
  grid-template-columns: minmax(0, 1.53fr) minmax(340px, 0.95fr);
  gap: 16px;
}

.jeel-side {
  display: grid;
  gap: 16px;
}

.jeel-card {
  padding: 18px 18px 16px;
  border: 1px solid #dde5ef;
  border-radius: 18px;
  background: #fff;
  box-shadow: 0 10px 24px rgba(15, 23, 42, 0.03);
}

.jeel-card--table {
  display: grid;
  gap: 12px;
}

.jeel-card__head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.jeel-card__head h3 {
  margin: 0;
  color: #0f172a;
  font-size: 1.12rem;
  font-weight: 800;
}

.jeel-card__button {
  min-height: 38px;
  padding: 0 14px;
  border: 1px solid #dbe4ef;
  border-radius: 999px;
  background: #fff;
  color: #0f172a;
  font-weight: 700;
}

.jeel-table-wrap {
  overflow: hidden;
  border: 1px solid #e2e8f0;
  border-radius: 14px;
}

.jeel-table {
  width: 100%;
  border-collapse: collapse;
}

.jeel-table thead th {
  padding: 14px 16px;
  text-align: left;
  color: #64748b;
  background: #fff;
  border-bottom: 1px solid #e8edf4;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  font-size: 0.68rem;
  font-weight: 800;
}

.jeel-table tbody td {
  padding: 17px 16px;
  border-bottom: 1px solid #edf1f7;
  color: #334155;
  font-size: 0.95rem;
}

.jeel-table tbody tr:last-child td {
  border-bottom: 0;
}

.jeel-table__name {
  color: #0f172a;
  font-weight: 700;
}

.jeel-pill {
  display: inline-flex;
  align-items: center;
  min-height: 28px;
  padding: 5px 10px;
  border-radius: 999px;
  font-size: 0.8rem;
  font-weight: 700;
}

.jeel-pill--pending {
  border: 1px solid #dbeafe;
  background: #eff6ff;
  color: #2563eb;
}

.jeel-pill--success {
  border: 1px solid #ccefd6;
  background: #ebf9ef;
  color: #15803d;
}

.jeel-actions {
  display: flex;
  gap: 8px;
}

.jeel-action-btn {
  display: inline-grid;
  place-items: center;
  width: 34px;
  height: 34px;
  padding: 0;
  border: 1px solid #dbe4ef;
  border-radius: 10px;
  background: #fff;
  color: #475569;
}

.jeel-action-btn svg {
  width: 13px;
  height: 13px;
}

.jeel-card--actions,
.jeel-card--snapshot {
  display: grid;
  gap: 16px;
}

.jeel-action-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 14px;
}

.jeel-quick-action {
  display: inline-flex;
  align-items: center;
  justify-content: flex-start;
  gap: 12px;
  min-height: 48px;
  padding: 0 16px;
  border: 1px solid #dbe4ef;
  border-radius: 14px;
  background: #fff;
  font-weight: 700;
}

.jeel-quick-action--blue {
  color: #2563eb;
}

.jeel-quick-action--green {
  color: #16a34a;
}

.jeel-quick-action--violet {
  color: #7c3aed;
}

.jeel-quick-action--amber {
  color: #ea580c;
}

.jeel-quick-action__icon {
  display: inline-grid;
  place-items: center;
  width: 14px;
  height: 14px;
}

.jeel-quick-action__icon svg {
  width: 14px;
  height: 14px;
}

.jeel-quick-action__label {
  font-size: 0.93rem;
  font-weight: 600;
}

.jeel-text-link {
  padding: 0;
  border: 0;
  background: transparent;
  color: #2563eb;
  font-weight: 700;
}

.jeel-snapshot__top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.jeel-snapshot__title {
  color: #0f172a;
  font-size: 0.98rem;
  font-weight: 800;
}

.jeel-snapshot__meta {
  margin-top: 4px;
  color: #94a3b8;
  font-size: 0.87rem;
}

.jeel-snapshot__summary {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 10px;
  padding: 12px 10px;
  border: 1px solid #dfece4;
  border-radius: 12px;
  background: linear-gradient(180deg, #f0fbf3 0%, #edf9ef 100%);
}

.jeel-snapshot__item {
  display: grid;
  gap: 3px;
  justify-items: center;
  text-align: center;
}

.jeel-snapshot__item strong {
  color: #0f172a;
  font-size: 1.1rem;
  font-weight: 800;
}

.jeel-snapshot__item span {
  color: #475569;
  font-size: 0.84rem;
}

.jeel-details-btn {
  min-height: 40px;
  border: 1px solid #dbe4ef;
  border-radius: 14px;
  background: #fff;
  color: #334155;
  font-weight: 700;
}

.jeel-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 0 4px;
  color: #94a3b8;
  font-size: 0.88rem;
}

@media (max-width: 1180px) {
  .jeel-metrics,
  .jeel-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 760px) {
  .jeel-dashboard {
    padding: 16px;
  }

  .jeel-hero {
    align-items: flex-start;
    flex-direction: column;
  }

  .jeel-metrics,
  .jeel-action-grid,
  .jeel-snapshot__summary {
    grid-template-columns: 1fr;
  }

  .jeel-table thead {
    display: none;
  }

  .jeel-table,
  .jeel-table tbody,
  .jeel-table tr,
  .jeel-table td {
    display: block;
    width: 100%;
  }

  .jeel-table tbody tr {
    border-top: 1px solid #edf1f7;
  }

  .jeel-table tbody td {
    display: flex;
    justify-content: space-between;
    gap: 16px;
    padding: 12px 16px;
  }

  .jeel-table tbody td::before {
    content: attr(data-label);
    color: #64748b;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    font-size: 0.66rem;
    font-weight: 800;
  }

  .jeel-footer {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
