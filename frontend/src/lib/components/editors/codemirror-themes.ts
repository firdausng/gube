import { basicLight } from 'cm6-theme-basic-light'
import { basicDark } from 'cm6-theme-basic-dark'
import type {Extension} from "@codemirror/state";

// const themes = [
//     {
//         extension: basicLight,
//         name: 'Basic Light'
//     },
//     {
//         extension: basicDark,
//         name: 'Basic Dark'
//     }
// ]
const themes: Record<string, Extension> = {
    'Basic Light': basicLight,
    'Basic Dark': basicDark
}
export default themes