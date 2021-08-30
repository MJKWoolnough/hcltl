import type {DOMBind} from './dom.js';
import {createSVG} from './dom.js';

export {createSVG};

export const [circle, defs, g, line, path, pattern, rect, svg, text, use] = "circle defs g line path pattern rect svg text use".split(" ").map(e => createSVG.bind(null, e)) as [DOMBind<SVGElementTagNameMap["circle"]>,DOMBind<SVGElementTagNameMap["defs"]>,DOMBind<SVGElementTagNameMap["g"]>,DOMBind<SVGElementTagNameMap["line"]>,DOMBind<SVGElementTagNameMap["path"]>,DOMBind<SVGElementTagNameMap["pattern"]>,DOMBind<SVGElementTagNameMap["rect"]>,DOMBind<SVGElementTagNameMap["svg"]>,DOMBind<SVGElementTagNameMap["text"]>,DOMBind<SVGElementTagNameMap["use"]>];
