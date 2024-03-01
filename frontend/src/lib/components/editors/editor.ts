import {EditorView, keymap} from "@codemirror/view";
import {Compartment, EditorState} from "@codemirror/state";
import {basicSetup} from "codemirror";
import {defaultKeymap} from "@codemirror/commands";
import {javascript} from "@codemirror/lang-javascript";
import themes from "$lib/components/editors/codemirror-themes";
import {appSettingStore} from "$lib/store/app-setting-store";

export function codeMirrorEditor(editorContainer:HTMLElement, {value}: {value:string}) {
    let themeConfig = new Compartment()
    let readOnly = true
    let activeTheme = 'Basic Light'
    let view: EditorView;
    let isScrollToBottom = true;

    console.log('codeMirrorEditor: ', value)
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
    });

    const fixedHeightEditor = EditorView.theme({
        "&": {height: "10px"},
        ".cm-scroller": {overflow: "auto"},
        ".cm-content, .cm-gutter": {minHeight: "200px"}
    }, {
        dark: true
    });

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
    });

    return {
        destroy() {
            // editorContainer.removeEventListener('keydown', handleKeydown);
            // previous?.focus();
        },
        update(data:{value:string}){
            const docLength = view.state.doc.length;
            // console.log('codeMirrorEditor: ', data)
            const scrollToBtm = {
                selection: {
                    anchor: docLength + 1,
                    head: docLength + 1
                },
                scrollIntoView: true,
            }

            let dispatchObj = {
                changes: {
                    from: docLength,
                    to: docLength,
                    // insert: data.value,
                    data: data.value
                }
            };

            // If isScrollToBottom is true, add scrollToBtm to the dispatch object
            if (isScrollToBottom) {
                dispatchObj = {
                    ...dispatchObj,
                    ...scrollToBtm
                };
            }
            view.dispatch(dispatchObj)
        }
    };
}