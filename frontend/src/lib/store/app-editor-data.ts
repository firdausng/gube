import type {Writable} from "svelte/store";
import {localStore} from "$lib/store/storage-store";

export type EditorData = {
    contextName: string
    namespaceName: string
    resourceName: string
    settings: Record<string, string>
    line: string[]
}

export const appEditorDataStore:Writable<EditorData[]> = localStore<EditorData[]>('editor-data', []);
