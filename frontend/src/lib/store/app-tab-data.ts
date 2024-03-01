import type {Writable} from "svelte/store";
import {localStore} from "$lib/store/storage-store";

export type TabItem = {
    /* eslint-disable  @typescript-eslint/no-explicit-any */
    // component: typeof SvelteComponent<any>,
    component: string,
    resourceName: string,
    resourceNamespace: string,
    name: string
}

export type Tab = {
    activeTab:TabItem| null
    tabs: TabItem[]
}

export const bottomTabDataStore:Writable<Record<string, Tab>> =
    localStore<Record<string, Tab>>('bottom-tab-data', {});
