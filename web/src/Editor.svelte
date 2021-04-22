<script>
  import Markdown from "svelte-markdown";
  import { onMount, onDestroy } from "svelte";

  export let name;

  function togglePreview(event) {
    if (event.ctrlKey && event.which === 83) {
      originalContent = content;
      saveContent(name, content)
        .then((res) => res.text())
        .then((text) => console.log(text));
      event.preventDefault();
      return false;
    }

    if (event.key === "Enter" && event.ctrlKey) {
      preview = !preview;
    }
  }

  onMount(() => {
    document.body.addEventListener("keydown", togglePreview);
  });

  onDestroy(() => {
    document.body.removeEventListener("keydown", togglePreview);
  });

  function saveContent(name, content) {
    const data = new FormData();
    data.append("name", name);
    data.append("content", content);
    return fetch("http://localhost:2001/api/write", {
      method: "POST",
      body: data,
    });
  }

  function getContent(name) {
    return fetch("http://localhost:2001/api/read/" + name)
      .then((res) => res.json())
      .then((json) => json.Content);
  }

  const contentPromise = getContent(name).then((ctn) => {
    originalContent = ctn;
    content = ctn;
  });

  let originalContent = "";
  let content = "";
  let preview = true;
</script>

<svelte:head>
  <title>{name + (content !== originalContent ? "*" : "")}</title>
</svelte:head>
<div class="container">
  {#await contentPromise}
    <h1>Loading...</h1>
  {:then}
    <div class:show={preview}><Markdown source={content} /></div>
    <textarea
      placeholder="Insert markdown here"
      class:show={!preview}
      autofocus
      bind:value={content}
    />
  {:catch error}
    <h1>Error: {error.message}</h1>
  {/await}
</div>

<style>
  .container > * {
    display: none;
  }
  .show {
    display: block;
  }
  textarea {
    overflow: auto;
    font-size: 18px;
    width: 100%;
    height: 100vh;
    outline: 0;
    padding: 20px 0;
    resize: none;
    border: none;
    margin: 0;
  }
</style>
