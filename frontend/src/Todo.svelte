<script>
  import { TodoAPI } from "./api/todo";

  export let id;
  export let name;
  export let categoryId;

  let deleted = false;

  function deleteHandler() {
    TodoAPI.deleteTodo(id);
    deleted = true;
  }

  function editHandler() {
    name = prompt("Please enter todo:", name);
    TodoAPI.updateTodo(id, name, categoryId);
  }
</script>

{#if deleted === false}
  <div
    class="h-20 border-2 w-full my-2 flex min-h-min rounded-md hover:bg-gray-100"
    {id}
  >
    <div class="flex flex-row items-center px-3 w-full">
      <div class="justify-self-start grow">
        <p class="text-justify">{name}</p>
      </div>
      <div class="justify-self-end">
        <button class="border-0" on:click={editHandler}>✏️</button>
        <button class="border-0" on:click={deleteHandler}>❌</button>
      </div>
    </div>
  </div>
{/if}
