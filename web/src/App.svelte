<script>
  import { Router, Link, Route } from "svelte-routing";
  import Home from "./Home.svelte";
  import Editor from "./Editor.svelte";
</script>

<main>
  <Router basepath="/app">
    <Route path="/" component={Home} />
    <Route path="edit/:name" component={Editor} />
  </Router>
</main>

<style>
  main {
    overflow: auto;
    width: 70%;
    margin: 0 auto;
  }
</style>
