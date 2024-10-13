<script lang="ts">
  import { onMount } from 'svelte';

  interface Santri {
    name: string;
    age: number;
    className: string;
  }

  let santris: Santri[] = [];
  let name: string = '';
  let age: number = 0;
  let className: string = '';

  // Fetch santris from the backend
  const getSantris = async () => {
    const res = await fetch('/api/santris');
    santris = await res.json();
  };

  // Add new Santri
  const addSantri = async () => {
    const res = await fetch('/api/santris', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, age, className }),
    });
    if (res.ok) {
      getSantris();
      name = '';
      age = 0;
      className = '';
    }
  };

  // Use onMount to call getSantris on client-side only
  onMount(() => {
    getSantris();
  });
</script>

<!-- HTML Part -->
<h1>Santri List</h1>
<ul>
  {#each santris as santri}
    <li>{santri.name} - {santri.age} years old, Class {santri.className}</li>
  {/each}
</ul>

<h2>Add New Santri</h2>
<input type="text" bind:value={name} placeholder="Name" />
<input type="number" bind:value={age} placeholder="Age" />
<input type="text" bind:value={className} placeholder="Class" />
<button on:click={addSantri}>Add Santri</button>
