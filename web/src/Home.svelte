<script>
  import { Link } from "svelte-routing";

  function getMarkdown() {
    return fetch("http://localhost:2001/api/list").then((res) => res.json());
  }

  let list = [];
  const listPromise = getMarkdown().then((json) => (list = json));

  getMarkdown();
</script>

<svelte:head>
  <title>Marked</title>
</svelte:head>
<h1>Marked, Markdown Editor 📝</h1>
{#await listPromise}
  <h2>Loading...</h2>
{:then}
  {#if list.length > 0}
    <h2>Start editing 🚀</h2>
    <ul>
      {#each list as markdown}
        <Link to={"edit/" + markdown}><li>{markdown}</li></Link>
      {/each}
    </ul>
  {:else}
    <h2>No markdown found 😢</h2>
  {/if}
{:catch error}
  <h2>Error: {error.message}</h2>
{/await}

<style>
  li {
    color: black !important;
  }
</style>
