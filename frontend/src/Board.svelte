<script lang="ts">
  import type { Writable } from "svelte/store";
  import { CategoryAPI, ICategory } from "./api/category";
  import { ITodo, TodoAPI } from "./api/todo";
  import Todo from "./Todo.svelte";

  export let category: ICategory;
  export let categoryStore: Writable<ICategory[]>;
  export let todoStore: Writable<ITodo[]>;

  let deleted = false;

  function onDragStart(event: DragEvent, todoId: number) {
    event.dataTransfer.setData("todoId", todoId.toString());
  }

  function onDrop(event: DragEvent, categoryId: number) {
    const todoId = parseInt(event.dataTransfer.getData("todoId"));

    for (let index = 0; index < $todoStore.length; index++) {
      const todo = $todoStore[index];
      if (todoId === todo.id) {
        // update todo store
        todoStore.update((todos) => {
          todos[index].category_id = categoryId;
          return todos;
        });

        // update todo in database
        TodoAPI.updateTodo(todoId, todo.name, categoryId);
        break;
      }
    }
  }

  async function addTodoWrapper(categoryId: number) {
    // prompt user todo name
    let todoName = prompt("Please enter todo:", "");

    // send add todo request
    await TodoAPI.addTodo(todoName, categoryId);

    // update todo store
    todoStore.update((todos) => [
      ...todos,
      {
        id: $todoStore.length + 1,
        name: todoName,
        category_id: categoryId,
      },
    ]);
  }

  async function deleteCategoryWrapper(categoryId: number) {
    CategoryAPI.deleteCategory(categoryId);
    deleted = true;
  }
</script>

{#if deleted === false}
  <div
    class="grid border-2 border-black w-full max-h-[60vh] h-[60vh]"
    on:dragenter={(event) => event.preventDefault()}
    on:dragover={(event) => event.preventDefault()}
    on:drop={(event) => onDrop(event, category.ID)}
  >
    <div class="flex flex-col px-4 items-start max-h-full">
      <div class="flex flex-col w-full items-center mb-3">
        <div class="flex flex-row items-baseline w-full">
          <!-- Header -->
          <div class="flex-grow">
            <h3 class="text-lg font-bold my-3 text-center">
              {category.name || "Unnamed Category"}
            </h3>
          </div>

          <div class="justify-self-end shrink-0">
            <button
              class="border-0"
              on:click={() => deleteCategoryWrapper(category.ID)}>❌</button
            >
          </div>
        </div>

        <!-- Add todo button -->
        <button
          class="flex orange-button w-full"
          on:click={() => addTodoWrapper(category.ID)}>ADD</button
        >
      </div>

      <div
        class="flex flex-col w-full overflow-y-scroll overflow-x-hidden mb-3 h-full max-h-full"
      >
        {#each $todoStore as todo}
          {#if todo.category_id === category.ID}
            <div
              draggable={true}
              on:dragstart={(event) => onDragStart(event, todo.id)}
            >
              <Todo id={todo.id} name={todo.name} categoryId={category.ID} />
            </div>
          {/if}
        {/each}
      </div>
    </div>
  </div>
{/if}
