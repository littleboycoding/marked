<script>
  import Markdown from "svelte-markdown";
  import { onMount, onDestroy, afterUpdate } from "svelte";
  import { navigate } from "svelte-routing";

  export let name;

  function togglePreview(event) {
    if (event.key === "Escape") {
      if (preview) {
        preview = false;
      } else {
        if (content === originalContent) {
          return navigate("/app");
        } else if (confirm("Are you sure want to leave without saving ?")) {
          return navigate("/app");
        }
      }
    }

    if (event.ctrlKey && event.which === 83) {
      originalContent = content;
      saving = true;
      saveContent(name, content).then(() => (saving = false));
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
  let preview = false;
  let editor;
  let saving = false;

  afterUpdate(() => {
    if (!preview && editor) editor.focus();
  });
</script>

<svelte:head>
  <title>{name}</title>
</svelte:head>
<div class="container">
  {#if saving}
    <div class="dim" />
  {/if}
  {#await contentPromise}
    <h1>Loading...</h1>
  {:then}
    <div class="markdown" class:show={preview}>
      <Markdown source={content} />
    </div>
    <textarea
      placeholder="Insert markdown here"
      class:show={!preview}
      bind:this={editor}
      bind:value={content}
    />
  {:catch error}
    <h1>Error: {error.message}</h1>
  {/await}
</div>

<style>
  .container > :not(.dim) {
    display: none;
  }
  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
  .show {
    display: block !important;
    animation: fadeIn 0.4s;
  }
  .dim {
    background-color: rgba(0, 0, 0, 0.1);
    width: 100vw;
    height: 100vh;
    position: absolute;
    left: 0;
    top: 0;
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
