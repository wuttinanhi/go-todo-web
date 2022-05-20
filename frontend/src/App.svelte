<script lang="ts">
  import { CategoryAPI, ICategory } from "./api/category";
  import { ITodo, TodoAPI } from "./api/todo";
  import { writable } from "svelte/store";
  import Board from "./Board.svelte";

  const categoryStore = writable<ICategory[]>([]);
  const todoStore = writable<ITodo[]>([]);

  async function addCategoryWrapper() {
    // prompt user category name
    let categoryName = prompt("Please enter category name:", "");

    // send add category request
    await CategoryAPI.addCategory(categoryName);

    // update category store
    categoryStore.update((categories) => [
      ...categories,
      {
        ID: categories.length + 1,
        name: categoryName,
      },
    ]);
  }

  async function fetcher() {
    const categories = await CategoryAPI.getCategory();
    const todos = await TodoAPI.getTodo();

    categoryStore.set(categories);
    todoStore.set(todos);
  }

  fetcher();
</script>

<main>
  <h1 class="text-3xl font-bold my-5 text-center text-orange-500">Todo App</h1>

  <div class="flex justify-center my-5">
    <button class="flex orange-button" on:click={addCategoryWrapper}>
      Add Category
    </button>
  </div>

  <div
    class="w-full h-screen grid grid-cols-4 gap-3 content-between justify-around align-middle "
  >
    {#each $categoryStore as category}
      <Board {category} {categoryStore} {todoStore} />
    {/each}
  </div>
</main>

<style lang="postcss" global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;

  .orange-button {
    @apply bg-white text-orange-600 border-2 rounded-md font-bold text-center justify-center border-orange-600 py-2 mx-5 px-5;
  }

  .orange-button:hover {
    @apply bg-orange-600 text-white;
  }
</style>
