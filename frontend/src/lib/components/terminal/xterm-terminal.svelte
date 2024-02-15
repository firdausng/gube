<script lang="ts">
	import 'xterm/css/xterm.css';
	import { Terminal } from 'xterm';
	import { FitAddon } from 'xterm-addon-fit';
	import { CanvasAddon } from 'xterm-addon-canvas';
	import { WebglAddon } from 'xterm-addon-webgl';
	import { LigaturesAddon } from 'xterm-addon-ligatures';
	import { Unicode11Addon } from 'xterm-addon-unicode11';
	import { WebLinksAddon } from 'xterm-addon-web-links';
	import { SearchAddon } from 'xterm-addon-search';
	import { SearchBarAddon } from 'xterm-addon-search-bar';
	import { onMount } from 'svelte';
	import { AdventureTime, Atom, AlienBlood, Afterglow, Batman, Blazer } from 'xterm-theme';
	import { RunCommand } from '$lib/wailsjs/go/backend/App';

	let terminalElement: HTMLElement;
	let termFit;
	let term: Terminal;

	function initializeXterm() {
		term = new Terminal({
			cursorBlink: true,
			allowProposedApi: true,
			// allowTransparency: true,
			macOptionIsMeta: true,
			macOptionClickForcesSelection: true,
			scrollback: 0,
			fontSize: 13,
			fontFamily: 'Consolas,Liberation Mono,Menlo,Courier,monospace',
			// rows: 30,
			theme: Afterglow,
			windowOptions: {
				fullscreenWin: true,
				getIconTitle: true,
				maximizeWin: true,
			}
		});

		termFit = new FitAddon();
		term.loadAddon(termFit);

		setRenderingMode();

		term.loadAddon(new Unicode11Addon());
		// activate the new version
		term.unicode.activeVersion = '11';

		const webLinksAddon = new WebLinksAddon((evt, uri) => {
			evt.preventDefault();
			open(uri);
		});
		term.loadAddon(webLinksAddon);

		const searchAddon = new SearchAddon()
		term.loadAddon(new SearchBarAddon({ searchAddon }));

		termFit.fit();
		term.focus()

		term.open(terminalElement);
		term.loadAddon(new LigaturesAddon())
	}

	let cmd;

	const newline = '\n\r';
	const startLineStr = '# ';
	let cmdStr = startLineStr;
	onMount(async () => {
		initializeXterm();
		term.write(startLineStr);
		term.onData((data) => {
			console.log('onData', JSON.stringify(data));

			if (data === '\x7F') {
				// Handle backspace
				cmdStr = cmdStr.slice(0, cmdStr.length - 1);
				term.write('\b \b');
			} else if (data === '\r') {
				// Handle Enter for command execution
				let lines = cmdStr.split(newline);

				let inputCmd = lines[lines.length - 1].substring(2);
				console.log(cmdStr, inputCmd);
				RunCommand(cmdStr).then(async cmd => {
					cmdStr = cmd.Data + newline;
					term.writeln(cmdStr);
					cmd = await RunCommand(inputCmd);
					console.log(cmd);
					if (cmd.ErrorMessage.length === 0) {
						term.write(cmd.Data);
					} else {
						term.write(cmd.ErrorMessage);
					}
					term.writeln("");
					term.write(startLineStr)
				});
			} else {
				// Add entered characters to command string
				cmdStr += data;
				console.log(cmdStr);
				term.write(data);
			}
		});

		window.onresize = () => {
			termFit?.fit();
		};

		term.onResize((event) => {
			var rows = event.rows;
			var cols = event.cols;
			// App.SetTTYSize(rows, cols);
		});
	});

	const setRenderingMode = async () => {

		//TODO enable this again when found webgl supported by wails
		// const settings = (await invoke('get_settings')) as string
		//
		// if (settings['useWebGL']) {
		// 	console.log('Trying to use WebGL')
		//
		// 	if(webglIsSupported()) {
		// 		addons['webglAddon'] = new WebglAddon()
		// 		addons['webglAddon'].onContextLoss(e => {
		// 			addons['webglAddon'].dispose();
		// 		})
		// 		terminalInterface.loadAddon(addons['webglAddon'])
		// 		return
		// 	}
		//
		// 	alert('WebGL is not supported. Falling back to canvas rendering. \n You can turn off WebGL in settings.')
		// }

		const canvasAddon = new CanvasAddon();
		term.loadAddon(canvasAddon);
	};
</script>

<div
	id="terminal"
	class="w-full"
	bind:this="{terminalElement}"
/>
