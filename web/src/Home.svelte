<script>
  import { Link, navigate } from "svelte-routing";

  function getMarkdown() {
    return fetch("http://localhost:2001/api/list").then((res) => res.json());
  }

  let list = [];
  const listPromise = getMarkdown().then((json) => (list = json));

  getMarkdown();

  function createFile() {
    let filename = prompt("New file name");

    if (filename !== null) {
      navigate("/app/edit/" + filename + ".md");
    }
  }
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
  {:else}
    <h2>No markdown found 😢</h2>
  {/if}
  <ul>
    <a on:click={createFile} href="#"><li>Create new</li></a>
    {#each list as markdown}
      <Link to={"edit/" + markdown}><li>{markdown}</li></Link>
    {/each}
  </ul>
{:catch error}
  <h2>Error: {error.message}</h2>
{/await}
<footer>
  <p>
    Create with ❤ by <a href="https://github.com/littleboycoding"
      >@littleboycoding</a
    >
  </p>
  <p>
    Source code is on <a href="https://github.com/littleboycoding/marked"
      >GitHub</a
    >
  </p>
</footer>

<style>
  li {
    color: black !important;
  }
  footer {
    background-color: white;
    position: fixed;
    bottom: 0;
    left: 0;
    text-align: center;
    width: 100%;
  }
</style>
