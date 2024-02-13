<script lang="ts">
	import { onMount } from 'svelte';
	import {Compartment, EditorState} from '@codemirror/state';
	import {EditorView, keymap, ViewUpdate} from "@codemirror/view"
	import {defaultKeymap} from "@codemirror/commands"
	import { basicSetup } from 'codemirror';
	import { javascript } from '@codemirror/lang-javascript';
	import themes from '$lib/components/editors/codemirror-themes';
	import { appSettingStore } from '$lib/store/app-setting-store';

	const themeConfig = new Compartment()
	let value = `I1118 09:09:50.015690       1 node.go:141] Successfully retrieved node IP: 192.168.65.3
I1118 09:09:50.015734       1 server_others.go:110] "Detected node IP" address="192.168.65.3 server_others.go:110] "Detected node IP" address="192.168.65.3 server_others.go:110] "Detected node IP" address="192.168.65.3 "Detected node IP" address="192.168.65.3 server_others.go:110] "Detected node IP" address="192.168.65.3 "Detected node IP" address="192.168.65.3 server_others.go:110] "Detected node IP" address="192.168.65.3"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
I1118 09:09:50.015757       1 server_others.go:551] "Using iptables proxy"
`;
	let readOnly = true
	let activeTheme = 'Basic Light'
	let view: EditorView;

	let editorContainer;

	onMount(async()=>{
		let state: EditorState = EditorState.create({
			doc: value,
			extensions: [
				basicSetup,
				EditorState.readOnly.of(readOnly),
				EditorView.lineWrapping,
				keymap.of(defaultKeymap),
				javascript(),
				themeConfig.of([themes[activeTheme]])
				// tabSize.of(EditorState.tabSize.of(8))
			]
		})

		const fixedHeightEditor = EditorView.theme({
			"&": {height: "10px"},
			".cm-scroller": {overflow: "auto"},
			".cm-content, .cm-gutter": {minHeight: "200px"}
		}, true)

		view = new EditorView({
			search: {config: {top: true}},
			viewportMargin: 10,
			state,
			extensions: fixedHeightEditor,
			parent: editorContainer
		});

		appSettingStore.subscribe(val => {
			if (val.theme === 'dark') {
				activeTheme = 'Basic Dark'
				view.dispatch({
					effects: themeConfig.reconfigure(themes[activeTheme])
				})
			} else {
				activeTheme = 'Basic Light'
				view.dispatch({
					effects: themeConfig.reconfigure(themes[activeTheme])
				})
			}
		})

	})
</script>

<div class="">
	<div bind:this="{editorContainer}"></div>
</div>
