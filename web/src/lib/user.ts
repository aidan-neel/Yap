import { persistedState } from 'svelte-persisted-state';
import { v4 as uuidv4 } from 'uuid';
export const userId = persistedState('identifier', uuidv4());
