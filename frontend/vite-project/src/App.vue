<template>
  <div id="app" class="app-shell">
    <header class="app-header">
      <div class="brand">
        <span class="brand-mark">J</span>
        <div>
          <p class="eyebrow">Student management</p>
          <h1>JEEL!ve</h1>
        </div>
      </div>

      <nav class="app-nav" aria-label="Primary">
        <router-link to="/" class="nav-link">Students</router-link>
        <router-link to="/add" class="nav-link">Add Student</router-link>
      </nav>
    </header>

    <main class="app-main">
      <section class="content-panel">
        <RouterView />
      </section>
    </main>
  </div>
</template>
