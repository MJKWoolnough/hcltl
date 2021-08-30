import type {DOMBind} from './dom.js';
import {createHTML} from './dom.js';

export {createHTML};

export const [br, button, div, h1, img, input, label, span, slot, style, table, tbody, td, th, thead, tr] = "br button div h1 img input label span slot style table tbody td th thead tr".split(" ").map(e => createHTML.bind(null, e)) as [DOMBind<HTMLElementTagNameMap["br"]>, DOMBind<HTMLElementTagNameMap["button"]>, DOMBind<HTMLElementTagNameMap["div"]>, DOMBind<HTMLElementTagNameMap["h1"]>, DOMBind<HTMLElementTagNameMap["img"]>, DOMBind<HTMLElementTagNameMap["input"]>, DOMBind<HTMLElementTagNameMap["label"]>, DOMBind<HTMLElementTagNameMap["span"]>, DOMBind<HTMLElementTagNameMap["slot"]>, DOMBind<HTMLElementTagNameMap["style"]>, DOMBind<HTMLElementTagNameMap["table"]>, DOMBind<HTMLElementTagNameMap["tbody"]>, DOMBind<HTMLElementTagNameMap["td"]>, DOMBind<HTMLElementTagNameMap["th"]>, DOMBind<HTMLElementTagNameMap["thead"]>, DOMBind<HTMLElementTagNameMap["tr"]>];
