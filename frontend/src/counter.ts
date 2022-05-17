// import writable store
import { writable } from 'svelte/store';

// create counter store
export const counter = writable(0);

// increase counter function
export const increase = () => {
    counter.update(n => n + 1);
};

// function reset counter
export const reset = () => {
    counter.set(0);
}

