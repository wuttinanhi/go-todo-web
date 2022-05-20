<script lang="ts">
  import { TodoAPI, ITodo } from "./api/todo";

  export let todo: ITodo;

  let deleted = false;

  function deleteHandler() {
    TodoAPI.deleteTodo(todo.id);
    deleted = true;
  }

  function editHandler() {
    todo.name = prompt("Please enter todo:", todo.name);
    TodoAPI.updateTodo(todo.id, todo.name, todo.category_id);
  }

  function onDragStart(event: DragEvent) {
    event.dataTransfer.setData("todoId", todo.id.toString());
  }
</script>

{#if deleted === false}
  <div
    class="flex h-20 border-2 w-full my-2 min-h-min rounded-md hover:bg-gray-100 bg-white shrink-0"
    draggable={true}
    on:dragstart={(event) => onDragStart(event)}
  >
    <div class="flex flex-row items-center px-3 w-full">
      <div class="justify-self-start grow">
        <p class="text-justify">{todo.name}</p>
      </div>
      <div class="justify-self-end">
        <button class="border-0" on:click={editHandler}>✏️</button>
        <button class="border-0" on:click={deleteHandler}>❌</button>
      </div>
    </div>
  </div>
{/if}
