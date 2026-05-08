<script setup lang="ts">
import { computed } from "vue"
import { useRoute, useRouter } from "vue-router"
import { clearToken, hasToken, runCap } from "./services/api"

const route = useRoute()
const router = useRouter()

const showShell = computed(() => hasToken() && route.name !== "login")

const handleLogout = () => {
  clearToken()
  router.push({ name: "login" })
}

const handleRunCap = async () => {
  await runCap()
}
</script>

<template>
  <div id="app" :class="['jeel-app', showShell ? 'jeel-app--shell' : 'jeel-app--auth']">
    <template v-if="showShell">
      <div class="jeel-shell">
        <aside class="jeel-sidebar" aria-label="Sidebar navigation">
          <div class="jeel-sidebar__top">
            <div class="jeel-brand">
              <div class="jeel-brand__mark" aria-hidden="true">J!</div>
              <div class="jeel-brand__copy">
                <h1>JEEL!ve</h1>
              </div>
            </div>

            <button class="jeel-icon-button" type="button" aria-label="Toggle navigation">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M4 6h16M4 12h16M4 18h16" />
              </svg>
            </button>
          </div>

          <nav class="jeel-sidebar__nav">
            <RouterLink
              to="/"
              class="jeel-sidebar__link"
              :class="{ 'jeel-sidebar__link--active': route.name === 'home' }"
            >
              <span class="jeel-sidebar__icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <rect x="3" y="3" width="18" height="18" rx="5" />
                  <path d="M7 12h10M12 7v10" />
                </svg>
              </span>
              <span>Dashboard</span>
            </RouterLink>

            <p class="jeel-sidebar__section">Student Management</p>

            <RouterLink to="/students" class="jeel-sidebar__link">
              <span class="jeel-sidebar__icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M16 21v-2a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2" />
                  <circle cx="9.5" cy="7.5" r="3.5" />
                  <path d="M21 21v-2a4 4 0 0 0-3-3.85" />
                  <path d="M16.5 4.7a3.5 3.5 0 0 1 0 6.6" />
                </svg>
              </span>
              <span>Students</span>
            </RouterLink>

            <RouterLink to="/add" class="jeel-sidebar__link">
              <span class="jeel-sidebar__icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M12 5v14M5 12h14" />
                </svg>
              </span>
              <span>Add Student</span>
            </RouterLink>

            <p class="jeel-sidebar__section">CAP Management</p>

            <button class="jeel-sidebar__link jeel-sidebar__button" type="button" @click="handleRunCap">
              <span class="jeel-sidebar__icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M5 12h14" />
                  <path d="m12 5 7 7-7 7" />
                </svg>
              </span>
              <span>Run CAP</span>
            </button>

            <button class="jeel-sidebar__link jeel-sidebar__button" type="button">
              <span class="jeel-sidebar__icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M5 5h14v14H5z" />
                  <path d="M8 9h8M8 13h8" />
                </svg>
              </span>
              <span>CAP History</span>
            </button>

            <p class="jeel-sidebar__section">Configuration</p>

            <button class="jeel-sidebar__link jeel-sidebar__button" type="button">
              <span class="jeel-sidebar__icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M3 19h18" />
                  <path d="M6 16V8" />
                  <path d="M12 16V4" />
                  <path d="M18 16v-6" />
                </svg>
              </span>
              <span>Branches</span>
            </button>

            <button class="jeel-sidebar__link jeel-sidebar__button" type="button">
              <span class="jeel-sidebar__icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path d="M4 19h16" />
                  <path d="M5 15l4-4 4 2 6-8" />
                </svg>
              </span>
              <span>Cutoffs</span>
            </button>

            <button class="jeel-sidebar__link jeel-sidebar__button" type="button">
              <span class="jeel-sidebar__icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <circle cx="12" cy="12" r="3" />
                  <path
                    d="M19.4 15a7.8 7.8 0 0 0 0-6l2-1.2-2-3.4-2.3 1a8.2 8.2 0 0 0-5.1-3l-.3-2.4H10l-.3 2.4a8.2 8.2 0 0 0-5.1 3l-2.3-1-2 3.4L2.2 9a7.8 7.8 0 0 0 0 6l-2 1.2 2 3.4 2.3-1a8.2 8.2 0 0 0 5.1 3l.3 2.4h2.4l.3-2.4a8.2 8.2 0 0 0 5.1-3l2.3 1 2-3.4-2-1.2Z"
                  />
                </svg>
              </span>
              <span>Settings</span>
            </button>
          </nav>

          <section class="jeel-sidebar__status">
            <div class="jeel-status-card">
              <p class="jeel-status-card__title">System Status</p>
              <div class="jeel-status-card__line">
                <span class="jeel-status-card__dot" aria-hidden="true"></span>
                <span class="jeel-status-card__state">Operational</span>
              </div>
              <p class="jeel-status-card__copy">All systems are running smoothly</p>
              <button class="jeel-status-card__link" type="button">View Details</button>
            </div>
          </section>
        </aside>

        <div class="jeel-workspace">
          <header class="jeel-topbar">
            <div class="jeel-topbar__title">JEEL!VE ALLOCATION HUB</div>

            <nav class="jeel-topbar__nav" aria-label="Primary navigation">
              <RouterLink to="/" class="jeel-topbar__link">Home</RouterLink>
              <RouterLink to="/students" class="jeel-topbar__link">Students</RouterLink>
              <RouterLink to="/add" class="jeel-topbar__link">Add Student</RouterLink>
            </nav>

            <div class="jeel-topbar__meta">
              <button class="jeel-notification" type="button" aria-label="Notifications">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
                  <path
                    d="M18 8a6 6 0 0 0-12 0c0 7-3 7-3 9h18c0-2-3-2-3-9"
                  />
                  <path d="M10 19a2 2 0 0 0 4 0" />
                </svg>
                <span class="jeel-notification__dot" aria-hidden="true"></span>
              </button>

              <div class="jeel-user">
                <div class="jeel-user__avatar">A</div>
                <div class="jeel-user__copy">Admin</div>
                <button class="jeel-user__chevron" type="button" aria-label="Account menu">
                  <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="m6 9 6 6 6-6" />
                  </svg>
                </button>
              </div>
            </div>
          </header>

          <main class="jeel-content">
            <RouterView />
          </main>
        </div>
      </div>
    </template>

    <main v-else class="jeel-auth">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.jeel-app {
  min-height: 100vh;
  background: linear-gradient(180deg, #f8fafc 0%, #f4f6f9 100%);
}

.jeel-shell {
  width: 100%;
  min-height: 100vh;
  display: grid;
  grid-template-columns: 290px minmax(0, 1fr);
  background: #f8fafc;
}

.jeel-sidebar {
  position: sticky;
  top: 0;
  height: 100vh;
  display: grid;
  grid-template-rows: auto 1fr auto;
  gap: 16px;
  padding: 14px 14px 16px;
  border-right: 1px solid #e5eaf1;
  background: rgba(255, 255, 255, 0.82);
  backdrop-filter: blur(10px);
  box-shadow: 8px 0 24px rgba(15, 23, 42, 0.03);
}

.jeel-sidebar__top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  min-height: 60px;
  padding: 4px 4px 6px 2px;
}

.jeel-brand {
  display: flex;
  align-items: center;
  gap: 12px;
}

.jeel-brand__mark {
  display: inline-grid;
  place-items: center;
  width: 38px;
  height: 38px;
  border: 1px solid #dce4ef;
  border-radius: 14px;
  background: linear-gradient(180deg, #fff 0%, #f7f9fc 100%);
  color: #2563eb;
  font-size: 1rem;
  font-weight: 800;
  box-shadow: 0 8px 18px rgba(15, 23, 42, 0.05);
}

.jeel-brand__copy h1 {
  margin: 0;
  color: #0f172a;
  font-size: 1.16rem;
  line-height: 1;
  letter-spacing: -0.04em;
  font-weight: 800;
}

.jeel-icon-button {
  display: inline-grid;
  place-items: center;
  width: 34px;
  height: 34px;
  border: 0;
  border-radius: 12px;
  background: transparent;
  color: #64748b;
}

.jeel-icon-button svg {
  width: 18px;
  height: 18px;
}

.jeel-sidebar__nav {
  display: grid;
  align-content: start;
  gap: 8px;
  overflow: auto;
  padding: 2px 0 8px;
}

.jeel-sidebar__section {
  margin: 18px 8px 6px;
  color: #6b7280;
  text-transform: uppercase;
  letter-spacing: 0.12em;
  font-size: 0.66rem;
  font-weight: 800;
}

.jeel-sidebar__link {
  display: flex;
  align-items: center;
  gap: 12px;
  min-height: 48px;
  padding: 0 14px;
  border: 1px solid transparent;
  border-radius: 14px;
  color: #475569;
  font-size: 0.96rem;
  font-weight: 500;
  background: transparent;
}

.jeel-sidebar__button {
  width: 100%;
  text-align: left;
}

.jeel-sidebar__icon {
  display: inline-grid;
  place-items: center;
  width: 18px;
  color: #64748b;
}

.jeel-sidebar__icon svg {
  width: 18px;
  height: 18px;
}

.jeel-sidebar__link.router-link-active,
.jeel-sidebar__link.router-link-exact-active,
.jeel-sidebar__link--active {
  border-color: #dbe4f4;
  background: linear-gradient(180deg, #f6f9ff 0%, #eef4ff 100%);
  color: #2563eb;
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.65);
}

.jeel-sidebar__link.router-link-active .jeel-sidebar__icon,
.jeel-sidebar__link.router-link-exact-active .jeel-sidebar__icon,
.jeel-sidebar__link--active .jeel-sidebar__icon {
  color: #2563eb;
}

.jeel-sidebar__status {
  padding-top: 8px;
}

.jeel-status-card {
  display: grid;
  gap: 6px;
  padding: 14px;
  border: 1px solid #dbe3ef;
  border-radius: 14px;
  background: #fff;
  box-shadow: 0 10px 20px rgba(15, 23, 42, 0.03);
}

.jeel-status-card__title {
  margin: 0;
  color: #0f172a;
  font-size: 0.95rem;
  font-weight: 800;
}

.jeel-status-card__line {
  display: flex;
  align-items: center;
  gap: 8px;
}

.jeel-status-card__dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #16a34a;
  box-shadow: 0 0 0 5px rgba(34, 197, 94, 0.12);
}

.jeel-status-card__state {
  color: #16a34a;
  font-size: 0.9rem;
}

.jeel-status-card__copy {
  margin: 0;
  color: #64748b;
  font-size: 0.86rem;
  line-height: 1.45;
}

.jeel-status-card__link {
  padding: 0;
  border: 0;
  background: transparent;
  color: #2563eb;
  text-align: left;
  font-weight: 600;
}

.jeel-workspace {
  min-width: 0;
  display: grid;
  grid-template-rows: auto 1fr;
}

.jeel-topbar {
  display: grid;
  grid-template-columns: auto 1fr auto;
  align-items: center;
  gap: 28px;
  min-height: 79px;
  padding: 0 26px;
  border-bottom: 1px solid #e5eaf1;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(12px);
}

.jeel-topbar__title {
  color: #64748b;
  font-size: 0.92rem;
  letter-spacing: 0.14em;
  font-weight: 800;
}

.jeel-topbar__nav {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 32px;
  min-width: 0;
}

.jeel-topbar__link {
  color: #64748b;
  font-size: 0.98rem;
  font-weight: 500;
}

.jeel-topbar__link.router-link-active,
.jeel-topbar__link.router-link-exact-active {
  color: #2563eb;
  font-weight: 600;
}

.jeel-topbar__meta {
  display: flex;
  align-items: center;
  gap: 12px;
}

.jeel-notification {
  position: relative;
  display: inline-grid;
  place-items: center;
  width: 34px;
  height: 34px;
  border: 0;
  border-radius: 999px;
  background: transparent;
  color: #64748b;
}

.jeel-notification svg {
  width: 16px;
  height: 16px;
}

.jeel-notification__dot {
  position: absolute;
  top: 10px;
  right: 10px;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #ef4444;
  box-shadow: 0 0 0 2px #fff;
}

.jeel-user {
  display: inline-flex;
  align-items: center;
  gap: 10px;
}

.jeel-user__avatar {
  display: inline-grid;
  place-items: center;
  width: 32px;
  height: 32px;
  border-radius: 999px;
  background: #111827;
  color: #fff;
  font-size: 0.9rem;
  font-weight: 800;
}

.jeel-user__copy {
  color: #334155;
  font-size: 0.96rem;
  font-weight: 500;
}

.jeel-user__chevron {
  display: inline-grid;
  place-items: center;
  width: 24px;
  height: 24px;
  padding: 0;
  border: 0;
  background: transparent;
  color: #64748b;
}

.jeel-user__chevron svg {
  width: 14px;
  height: 14px;
}

.jeel-content {
  min-width: 0;
  padding: 10px;
}

.jeel-auth {
  min-height: 100vh;
}

@media (max-width: 1180px) {
  .jeel-shell {
    grid-template-columns: 1fr;
  }

  .jeel-sidebar {
    position: static;
    height: auto;
  }
}

@media (max-width: 900px) {
  .jeel-topbar {
    grid-template-columns: 1fr;
    justify-items: start;
    gap: 14px;
    padding: 16px 18px;
  }

  .jeel-topbar__nav {
    justify-content: flex-start;
    gap: 18px;
    flex-wrap: wrap;
  }
}

@media (max-width: 640px) {
  .jeel-sidebar {
    padding: 12px;
  }

  .jeel-topbar {
    padding: 14px 12px;
  }

  .jeel-topbar__meta {
    flex-wrap: wrap;
  }
}
</style>
