<template>
  <div id="app" class="app-shell">
    <header class="app-header">
      <div class="brand">
        <span class="brand-mark">J</span>
        <div>
          <p class="eyebrow">JEEL!ve allocation hub</p>
          <h1>JEEL!ve</h1>
        </div>
      </div>

      <nav class="app-nav" aria-label="Primary">
        <router-link to="/" class="nav-link nav-link--home">Home</router-link>
        <router-link to="/students" class="nav-link nav-link--students">Students</router-link>
        <router-link to="/add" class="nav-link nav-link--add">Add Student</router-link>
      </nav>
    </header>

    <main class="app-main">
      <section class="content-panel">
        <RouterView />
      </section>
    </main>
  </div>
</template>
