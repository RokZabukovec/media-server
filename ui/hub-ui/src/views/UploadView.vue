<template>
  <div class="upload mt-5">
    <h1>Upload media</h1>
    <form method="post" action="/api/files" enctype="multipart/form-data">
      <input type="file" name="file" class="file-input file-input-bordered file-input-accent w-full max-w-xs" />
      <div class="flex mt-5">
        <button class="btn mr-3" type="reset">Clear</button>
        <button class="btn btn-accent" type="submit">Upload</button>
      </div>
    </form>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
<script setup>
</script>