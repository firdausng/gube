import type {Writable} from "svelte/store";
import {localStore} from "$lib/store/storage-store";

export type Resource = {
    name: string
}

export type TabItem = {
    /* eslint-disable  @typescript-eslint/no-explicit-any */
    // component: typeof SvelteComponent<any>,
    component: string,
    name: string
}

export type Tab = {
    activeTab:TabItem| null
    tabs: TabItem[]
}

export type WorkspaceContext = {
    name: string
    id: string
    active: boolean
    activeResource: Resource
    tabData: Tab
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
