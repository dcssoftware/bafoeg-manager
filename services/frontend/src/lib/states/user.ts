import type { SelfInformation } from "$lib/api/user/types/self-information";
import { writable } from "svelte/store";

export const userState = writable<SelfInformation | null>(null);