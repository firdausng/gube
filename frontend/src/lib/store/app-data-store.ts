import type {Writable} from "svelte/store";
import {localStore} from "$lib/store/storage-store";

export type Resource = {
    name: string
}

export type WorkspaceContext = {
    name: string
    id: string
    active: boolean
    activeResource: Resource
    workspaceList: Resource[]
}

export type Workspace = {
    name: string
    id: string
    contexts: WorkspaceContext[]
    activeContext: WorkspaceContext|null
}

export type AppData = {
    activeWorkspace: Workspace
    workspaces: Workspace[]
}


export const appDataStore:Writable<AppData> = localStore<AppData>('app-data', null!);
