import {localStore} from "$lib/store/storage-store";
import type {Writable} from "svelte/store";

export type AppSetting = {
    theme: string
    leftSidebarHidden: boolean
}

export const appSettingStore:Writable<AppSetting> = localStore<AppSetting>('app-setting', {
    theme: 'dark',
    leftSidebarHidden: false
});
