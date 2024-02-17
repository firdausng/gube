import type {Writable} from "svelte/store";
import {localStore} from "$lib/store/storage-store";
import type {AppSetting} from "$lib/store/app-setting-store";

export type WorkspaceContext = {
    name: string
    id: string
    active: boolean
}

export type Workspace = {
    name: string
    id: string
    contexts: WorkspaceContext[]
}

export type AppData = {
    activeWorkspace: Workspace | null
}

export const appDataStore:Writable<AppData> = localStore<AppData>('app-data', {
    activeWorkspace: null,
});
