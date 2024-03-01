import type {Writable} from "svelte/store";
import {localStore} from "$lib/store/storage-store";

export type EditorData = {
    workspaceName: string
    contextName: string
    namespaceName: string
    resourceName: string
    settings: Record<string, string>
    line: string[]
}

export const appEditorDataStore:Writable<Record<string, EditorData>> =
    localStore<Record<string, EditorData>>('editor-data', {});
