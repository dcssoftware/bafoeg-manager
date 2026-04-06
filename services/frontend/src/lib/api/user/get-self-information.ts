import ky from "ky";
import type { SelfInformation } from "./types/self-information";

export async function getSelfInformation(): Promise<SelfInformation | null> {
    return (await ky.get('/api/v1/self', { credentials: 'include' }).json<SelfInformation>());
}