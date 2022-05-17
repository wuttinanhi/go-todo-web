<script lang="ts">
  // import counter store
  import { counter, increase, reset } from "./counter";

  export let name: string;
</script>

<main>
  <h1 class="text-3xl font-bold underline">Hello world!</h1>

  <h1>Hello {name}!</h1>
  <p>
    Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn
    how to build Svelte apps.
  </p>
  <br />

  <h1>Counter: {$counter}</h1>
  <button on:click={() => increase()} class="orange-button"> Increase </button>
  <button on:click={() => reset()} class="orange-button"> Reset </button>
</main>

<style lang="postcss" global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;

  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }

  .orange-button {
    @apply p-3 bg-white text-orange-600 border-2 rounded-lg border-orange-600;
  }

  .orange-button:focus {
    @apply border-orange-600;
  }

  .orange-button:hover {
    @apply bg-orange-600 text-white;
  }
</style>
