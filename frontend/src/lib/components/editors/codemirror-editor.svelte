<script lang="ts">
	import { codeMirrorEditor } from './editor';
	// import { onMount } from 'svelte';
	// import {Compartment, EditorState} from '@codemirror/state';
	// import {EditorView, keymap} from "@codemirror/view"
	// import {defaultKeymap} from "@codemirror/commands"
	// import { basicSetup } from 'codemirror';
	// import { javascript } from '@codemirror/lang-javascript';
	// import themes from '$lib/components/editors/codemirror-themes';
	// import { appSettingStore } from '$lib/store/app-setting-store';
	// import {appDataStore, type TabItem} from "$lib/store/app-data-store";
	//
	// const themeConfig = new Compartment()
	// let value = ``;
	// let tabItem: TabItem|null;
	// appDataStore.subscribe(d =>{
	// 	if(d.activeWorkspace.activeContext){
	// 		tabItem = d.activeWorkspace.activeContext.tabData.activeTab;
	// 	}
	// })
	//
	// let readOnly = true
	// let activeTheme = 'Basic Light'
	// let view: EditorView;
	// let editorContainer:HTMLElement;
	//
	// onMount(async()=>{
	// 	let state: EditorState = EditorState.create({
	// 		doc: value,
	// 		extensions: [
	// 			basicSetup,
	// 			EditorState.readOnly.of(readOnly),
	// 			EditorView.lineWrapping,
	// 			keymap.of(defaultKeymap),
	// 			javascript(),
	// 			themeConfig.of([themes[activeTheme]])
	// 			// tabSize.of(EditorState.tabSize.of(8))
	// 		]
	// 	});
	// 	const fixedHeightEditor = EditorView.theme({
	// 		"&": {height: "10px"},
	// 		".cm-scroller": {overflow: "auto"},
	// 		".cm-content, .cm-gutter": {minHeight: "200px"}
	// 	}, {
	// 		dark: true
	// 	});
	// 	view = new EditorView({
	// 		search: {config: {top: true}},
	// 		viewportMargin: 10,
	// 		state,
	// 		extensions: fixedHeightEditor,
	// 		parent: editorContainer
	// 	});
	// 	appSettingStore.subscribe(val => {
	// 		if (val.theme === 'dark') {
	// 			activeTheme = 'Basic Dark'
	// 			view.dispatch({
	// 				effects: themeConfig.reconfigure(themes[activeTheme])
	// 			})
	// 		} else {
	// 			activeTheme = 'Basic Light'
	// 			view.dispatch({
	// 				effects: themeConfig.reconfigure(themes[activeTheme])
	// 			})
	// 		}
	// 	});
	// })
</script>

<div class="">
<!--	<div bind:this="{editorContainer}"></div>-->
	<div use:codeMirrorEditor={{value:'test\nabc'}}></div>
</div>
