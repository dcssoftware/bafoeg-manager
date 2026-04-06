import { convertStructTypes } from "$lib/functions/convertStructTypes";
import type { User } from "./types/user-model";

export async function getUsers(): Promise<User[] | undefined> {
  try {
    const response = await fetch(`/api/v1/user-management/`);
    if (!response.ok) {
      throw new Error('Failed to fetch users data');
    }

    const data = convertStructTypes(await response.json())

    return data;
  } catch (error) {

    console.error('Failed to fetch user data', error);
    return undefined
  }
}
