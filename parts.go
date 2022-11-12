package main

// File automatically generated with ./embed.sh

const (
	start = "<html lang=\"en\"><head><title>HCL Timeline</title><meta charset=\"UTF-8\" /><link rel=\"shortcut icon\" sizes=\"any\" href=\"data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 256 256'%3E%3Crect x='15' y='15' height='45' width='175' rx='5' fill='%23f00' /%3E%3Crect x='15' y='75' height='45' width='90' rx='5' fill='%230f0' /%3E%3Crect x='160' y='75' height='45' width='85' rx='5' fill='%230f0' /%3E%3Crect x='70' y='135' height='45' width='125' rx='5' fill='%2300f' /%3E%3Crect x='50' y='195' height='45' width='70' rx='5' fill='%23000' /%3E%3Crect x='170' y='195' height='45' width='50' rx='5' fill='%23000' /%3E%3C/svg%3E\" /><script type=\"module\">const t=(e,o=[])=>{if(e instanceof d){const t=new Text(e+\"\");e[r](t),o.push(t)}else if(\"string\"==typeof e)o.push(e);else if(Array.isArray(e))for(const i of e)t(i,o);else if(e instanceof Node)o.push(e);else if(e instanceof NodeList||e instanceof HTMLCollection)for(const t of Array.from(e))o.push(t);return o},e=t=>t instanceof Object&&t.handleEvent instanceof Function,o=t=>t instanceof Function||e(t)&&!(t instanceof h)||t instanceof h&&o(t.value),i=t=>o(t)||t instanceof Array&&3===t.length&&o(t[0])&&t[1]instanceof Object&&\"boolean\"==typeof t[2],s=t=>t instanceof Object,n=t=>t instanceof CSSStyleDeclaration||t instanceof Object,l=t=>\"string\"==typeof t||t instanceof Array||t instanceof NodeList||t instanceof HTMLCollection||t instanceof Node||t instanceof d,r=Symbol(\"setNode\"),a=Symbol(\"update\"),c=Symbol(\"remove\");class d{#t=new Set;[r](t){this.#t.add(new WeakRef(t))}[a](){const t=this+\"\";for(const e of this.#t){const o=e.deref();o?o instanceof d?o[a]():o.textContent=t:this.#t.delete(e)}}[c](t){for(const e of this.#t){const o=e.deref();o&&o!==t||this.#t.delete(e)}}}class h extends d{#e;constructor(t){super(),this.#e=t,t instanceof d&&t[r](this)}get value(){return this.#e instanceof h?this.#e.value:this.#e}set value(t){this.#e!==t&&(this.#e instanceof d&&this.#e[c](this),this.#e=t,t instanceof d&&t[r](this)),this[a]()}handleEvent(t){const o=this.value;o instanceof Function?o.call(t.currentTarget,t):e(o)&&o.handleEvent(t)}toString(){return this.value.toString()}}const f=(e,o,a)=>{if(o&&l(o))a=o;else if(o instanceof NamedNodeMap&&e instanceof Element)for(const t of o)e.setAttributeNode(t.cloneNode());else if(e&&\"object\"==typeof o){const t=!!((c=e).style&&c.classList&&c.getAttributeNode&&c.removeAttribute&&c.setAttribute&&c.toggleAttribute);for(const l in o){const a=o[l];if(i(a)&&l.startsWith(\"on\")){const t=a instanceof Array;e[t&&a[2]?\"removeEventListener\":\"addEventListener\"](l.slice(2),t?a[0]:a,!!t&&a[1])}else if(t)if(\"boolean\"==typeof a)e.toggleAttribute(l,a);else if(void 0===a)e.removeAttribute(l);else if(a instanceof Array||a instanceof DOMTokenList){if(\"class\"===l&&a.length)for(let t of a){const o=t.slice(0,1),i=\"!\"!==o&&(\"~\"!==o||void 0);e.classList.toggle(i?t:t.slice(1),i)}}else if(\"class\"===l&&s(a))for(const t in a)e.classList.toggle(t,a[t]??void 0);else if(\"style\"===l&&n(a))for(const[t,o]of a instanceof CSSStyleDeclaration?Array.from(a,(t=>[t,a.getPropertyValue(t)])):Object.entries(a))void 0===o?e.style.removeProperty(t):e.style.setProperty(t,o.toString());else if(e.setAttribute(l,a),a instanceof d){const t=e.getAttributeNode(l);t&&a[r](t)}}}var c;return e instanceof Node&&(\"string\"!=typeof a||e.firstChild?a&&(e instanceof Element||e instanceof DocumentFragment?e.append(...t(a)):e.appendChild(m(a))):e.textContent=a),e},p=(t,e)=>Object.defineProperty(((o,i)=>f(document.createElementNS(t,e),o,i)),\"name\",{value:e}),u=(t,e,o)=>[t,{once:!!(1&e),capture:!!(2&e),passive:!!(4&e),signal:o},!!(8&e)],m=e=>{const o=document.createDocumentFragment();return\"string\"==typeof e?o.textContent=e:void 0!==e&&o.append(...t(e)),o},w=(e,o,i)=>{if(!e)return e;if(o&&l(o)&&(o=void(i=o)),\"string\"==typeof i)i=void(e.textContent=i);else if(i&&e instanceof Element)i=void e.replaceChildren(...t(i));else for(;null!==e.lastChild;)e.lastChild.remove();return f(e,o,i)},g=(t,e=!0)=>(window.setTimeout((()=>{t.focus(),(t instanceof HTMLInputElement||t instanceof HTMLTextAreaElement)&&e&&t.select()}),0),t),x=\"http://www.w3.org/1999/xhtml\",[v,b,y,k,E,C,S,$,M,A,L,T,N,H,I,W]=\"br button div h1 img input label slot span style table tbody td th thead tr\".split(\" \").map((t=>p(x,t))),z=\"complete\"===document.readyState?Promise.resolve():new Promise((t=>window.addEventListener(\"load\",t,{once:!0}))),O=\"http://www.w3.org/2000/svg\",[R,P,D,B,j,F,X,V,Y,_]=\"circle defs g line path pattern rect svg text use\".split(\" \").map((t=>p(O,t)));class K extends CSSStyleSheet{#o;#i;constructor(t=\"\",e=0){super(),this.#o=Z.test(t)?t:\"_\",this.#i=e}#s(t,e){const o=[];let i=\"\";for(const s in e){const n=e[s];J(n)?o.push(...this.#s(G(t,s),n)):i+=`${s}:${n};`}return i&&o.unshift(t+\"{\"+i+\"}\"),o}add(t,e){if(t.trim())for(const o of this.#s(t,e))this.insertRule(o,this.cssRules.length);return this}at(t,e){if(t.trim())if(e){let o=\"\";for(const t in e)for(const i of this.#s(t,e[t]))o+=i;this.insertRule(t+\"{\"+o+\"}\")}else this.insertRule(t);return this}id(){return this.#o+this.#i++}ids(t){return Array.from({length:t},(()=>this.id()))}toString(){let t=\"\";for(const e of this.cssRules)t+=e.cssText;return t}render(){return A({type:\"text/css\"},this+\"\")}}const U=t=>{const e=[],o=[];let i=0;for(let s=0;s<t.length;s++){const n=t.charAt(s);if('\"'===n||\"'\"===n)for(s++;s<t.length;s++){const e=t.charAt(s);if(\"\\\\\"===e)s++;else if(e===n)break}else\",\"!==n||e.length?n===e.at(-1)?e.pop():\"[\"===n?e.push(\"]\"):\"(\"===n&&e.push(\")\"):(o.push(t.slice(i,s)),i=s+1)}return o.push(t.slice(i,t.length)),o},G=(t,e)=>{const o=U(e);let i=\"\";for(const e of U(t))for(const t of o)i+=(i.length?\",\":\"\")+e+t;return i},J=t=>Object.getPrototypeOf(t)===Object.prototype,Z=/^\\-?[_a-z\\240-\\377][_a-z0-9\\-\\240-\\377]*$/i,q=(new K,{CANCEL:\"Cancel\",CLOSE:\"Close\",MAXIMISE:\"Maximise\",MINIMISE:\"Minimise\",OK:\"Ok\",RESTORE:\"Restore\"}),Q=(new K).add(\":host\",{display:\"block\",position:\"relative\",overflow:\"hidden\",width:\"var(--shell-width, 100%)\",height:\"var(--shell-height, 100%)\"}).add(\"::slotted(windows-window)\",{\"--taskmanager-on\":\"none\"}).add(\"::slotted(windows-window:last-of-type)\",{\"--overlay-on\":\"none\"}),tt=(new K).add(\":host\",{position:\"absolute\",top:0,left:0,bottom:0,right:0}),et=(new K).add(\":host\",{position:\"absolute\",display:\"block\",\"background-color\":\"#fff\",color:\"#000\",border:\"1px solid #000\",width:\"var(--window-width, auto)\",height:\"var(--window-height, auto)\",\"min-width\":\"100px\",\"min-height\":\"100px\",top:\"var(--window-top, 0)\",left:\"var(--window-left, 0)\",\"list-style\":\"none\",padding:0,\"z-index\":0,\"([maximised])\":{left:0,right:0,top:0,bottom:0,width:\"auto\",height:\"auto\",\">div:nth-child(2)>div>button:nth-of-type(2)\":{\"background-image\":`url('data:image/svg+xml,%3Csvg viewBox=\"0 0 15 13\" xmlns=\"${O}\"%3E%3Cpath d=\"M1,5 h8 v-1 h-8 v8 h8 v-8 m-3,0 v-3 h8 v8 h-5 m5,-7 h-8\" stroke=\"%23000\" fill=\"none\" /%3E%3C/svg%3E')`}},\"([resizable])\":{\"border-width\":0,\">div\":{\":nth-child(1)\":{display:\"block\"},\":nth-child(3)\":{overflow:\"auto\",position:\"absolute\",bottom:0,left:0,right:0,top:\"calc(1em + 6px)\"},\":nth-child(4)\":{top:\"var(calc(--window-resize), -2px)\",left:\"var(calc(--window-resize), -2px)\",bottom:\"var(calc(--window-resize), -2px)\",right:\"var(calc(--window-resize), -2px)\"}}},\">div\":{\":nth-child(1)>div\":{position:\"absolute\",\"border-color\":\"currentColor\",\"border-style\":\"solid\",\"border-width\":0,\"z-index\":-1,\":nth-child(1)\":{top:\"-2px\",left:\"-2px\",width:\"10px\",height:\"10px\",cursor:\"nwse-resize\",\"border-left-width\":\"3px\",\"border-top-width\":\"3px\"},\":nth-child(2)\":{top:\"-2px\",left:\"8px\",right:\"8px\",\"border-top-width\":\"3px\",cursor:\"ns-resize\"},\":nth-child(3)\":{top:\"-2px\",right:\"-2px\",width:\"10px\",height:\"10px\",\"border-top-width\":\"3px\",\"border-right-width\":\"3px\",cursor:\"nesw-resize\"},\":nth-child(4)\":{top:\"8px\",right:\"-2px\",bottom:\"8px\",\"border-right-width\":\"3px\",cursor:\"ew-resize\"},\":nth-child(5)\":{bottom:\"-2px\",right:\"-2px\",width:\"10px\",height:\"10px\",\"border-right-width\":\"3px\",\"border-bottom-width\":\"3px\",cursor:\"nwse-resize\"},\":nth-child(6)\":{bottom:\"-2px\",left:\"8px\",right:\"8px\",\"border-bottom-width\":\"3px\",cursor:\"ns-resize\"},\":nth-child(7)\":{bottom:\"-2px\",left:\"-2px\",width:\"10px\",height:\"10px\",\"border-left-width\":\"3px\",\"border-bottom-width\":\"3px\",cursor:\"nesw-resize\"},\":nth-child(8)\":{top:\"8px\",left:\"-2px\",bottom:\"8px\",\"border-left-width\":\"3px\",cursor:\"ew-resize\"}},\":nth-child(2)\":{\"white-space\":\"nowrap\",height:\"calc(1em + 6px)\",\"background-color\":\"#aaa\",overflow:\"hidden\",\"user-select\":\"none\",display:\"flex\",\"align-items\":\"center\",\">span\":{\"margin-right\":\"calc(3em + 24px)\"},\">img\":{height:\"1em\",\"pointer-events\":\"none\"},\">div\":{position:\"absolute\",right:0,top:0,\">button\":{\":nth-of-type(1)\":{\"background-image\":`url('data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"${O}\"%3E%3Cpath d=\"M1,1 L9,9 M9,1 L1,9\" stroke=\"black\" /%3E%3C/svg%3E')`},\":nth-of-type(2)\":{\"background-image\":`url('data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"${O}\"%3E%3Cpath d=\"M9,3 h-8 v-1 h8 v-1 h-8 v8 h8 v-8\" stroke=\"black\" fill=\"none\" /%3E%3C/svg%3E')`},\":nth-of-type(3)\":{display:\"var(--taskmanager-on, block)\",\"background-image\":`url('data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"${O}\"%3E%3Cline x1=\"1\" y1=\"9\" x2=\"9\" y2=\"9\" stroke=\"black\" /%3E%3C/svg%3E')`}}},\" button\":{padding:0,\"border-width\":\"2px\",float:\"right\",\"background-repeat\":\"no-repeat\",\"background-position\":\"center\",\"background-color\":\"#eee\",\"background-size\":\"1em 1em\",width:\"calc(1em + 8px)\",height:\"calc(1em + 8px)\"}},\":nth-child(3)\":{\"user-select\":\"contain\",\":not(.hasChild)+div:nth-child(5)\":{\"pointer-events\":\"none\"}},\":nth-child(4)\":{display:\"var(--overlay-on, block)\",position:\"absolute\",\"background-color\":\"RGBA(0, 0, 0, 0.1)\",top:0,left:0,bottom:0,right:0}},\"([minimised]),>div:nth-child(1),([hide-titlebar])>div:nth-child(2),([hide-close])>div:nth-child(2)>div>button:nth-of-type(1),([hide-maximise])>div:nth-child(2)>div>button:nth-of-type(2),([hide-minimise])>div:nth-child(2)>div>button:nth-of-type(3),([window-hide])>div:nth-child(2)>div>button:nth-of-type(3)\":{display:\"none\"}});let ot=null,it=!1;class st extends HTMLElement{alert(t,e,o){return new Promise(((i,s)=>{const n=dt({\"window-hide\":!0,\"window-icon\":o,\"window-title\":t,\"hide-maximise\":\"true\",onremove:()=>i(!1)},[y(e),y({style:\"text-align: center\"},g(b({onclick:()=>{i(!0),n.remove()}},q.OK)))]);this.addWindow(n)||s(new Error(\"invalid target\"))}))}confirm(t,e,o){return new Promise(((i,s)=>{const n=dt({\"window-hide\":!0,\"window-icon\":o,\"window-title\":t,\"hide-maximise\":\"true\",onremove:()=>i(!1)},[y(e),y({style:\"text-align: center\"},[g(b({onclick:()=>{i(!0),n.remove()}},q.OK)),b({onclick:()=>n.remove()},q.CANCEL)])]);this.addWindow(n)||s(new Error(\"invalid target\"))}))}prompt(t,e,o=\"\",i){return new Promise(((s,n)=>{const l=()=>{s(r.value),a.remove()},r=g(C({value:o,onkeydown:t=>{switch(t.key){case\"Enter\":l();break;case\"Escape\":a.remove()}}})),a=dt({\"window-hide\":!0,\"window-icon\":i,\"window-title\":t,\"hide-maximise\":\"true\",onremove:()=>s(null)},[y(e),r,y({style:\"text-align: center\"},b({onclick:l},q.OK))]);this.addWindow(a)||n(new Error(\"invalid target\"))}))}}class nt extends st{constructor(){super(),new.target===nt&&(f(this.attachShadow({mode:\"closed\"}),[$({name:\"desktop\"}),y($())]).adoptedStyleSheets=[Q])}addWindow(t){return t.parentNode!==this?f(this,t):t.nextElementSibling&&t.focus(),!0}realignWindows(){const{offsetWidth:t,offsetHeight:e}=this;for(const o of this.childNodes)if(o instanceof rt){const{offsetLeft:i,offsetTop:s,offsetWidth:n,offsetHeight:l}=o;i+n>t?f(o,{style:{\"--window-left\":Math.max(t-n,0)+\"px\"}}):i<0&&f(o,{style:{\"--window-left\":\"0px\"}}),s+l>e?f(o,{style:{\"--window-top\":Math.max(e-l)+\"px\"}}):s<0&&f(o,{style:{\"--window-top\":\"0px\"}})}}}class lt extends HTMLElement{constructor(){super(),setTimeout(f,0,this,{slot:\"desktop\"}),f(this.attachShadow({mode:\"closed\"}),$({slot:\"desktop\"})).adoptedStyleSheets=[tt]}attributeChangedCallback(t,e,o){\"slot\"===t&&\"desktop\"!==o&&f(this,{slot:\"desktop\"})}static get observedAttributes(){return[\"slot\"]}}class rt extends st{#n;#l;#r;#a;#c=null;#d=null;#h;constructor(){super();const t=()=>this.focus();f(this.attachShadow({mode:\"closed\"}),[y(Array.from({length:8},((t,e)=>y({onmousedown:t=>((t,e,o)=>{const i=t.parentNode;if(it||!(i instanceof nt)||0!==o.button)return;it=!0,f(i,{style:{\"user-select\":\"none\"}});const s=t.offsetLeft,n=t.offsetTop,l=t.offsetWidth,r=t.offsetHeight,a=o.clientX,c=o.clientY,d=new AbortController,{signal:h}=d;f(i,{onmousemove:u((o=>{const i=o.clientX-a,d=o.clientY-c;switch(e){case 0:case 1:case 2:{const e=r-d;e>(parseInt(t.style.getPropertyValue(\"min-height\")||\"\")||100)&&f(t,{style:{\"--window-top\":`${n+d}px`,\"--window-height\":`${e}px`}})}break;case 4:case 5:case 6:{const e=r+d;e>(parseInt(t.style.getPropertyValue(\"min-height\")||\"\")||100)&&f(t,{style:{\"--window-height\":`${e}px`}})}}switch(e){case 2:case 3:case 4:{const e=l+i;e>(parseInt(t.style.getPropertyValue(\"min-width\")||\"\")||100)&&f(t,{style:{\"--window-width\":`${e}px`}})}break;case 0:case 7:case 6:{const e=l-i;e>(parseInt(t.style.getPropertyValue(\"min-width\")||\"\")||100)&&f(t,{style:{\"--window-left\":`${s+i}px`,\"--window-width\":`${e}px`}})}}}),0,h),onmouseup:u((e=>{0===e.button&&(f(i,{style:{\"user-select\":void 0}}),d.abort(),it=!1,t.dispatchEvent(new CustomEvent(\"resized\")))}),0,h)})})(this,e,t)})))),y({part:\"titlebar\",onmousedown:t=>((t,e)=>{const o=t.parentNode;if(it||!(o instanceof nt)||0!==e.button)return;it=!0,f(o,{style:{\"user-select\":\"none\"}});const i=e.clientX-t.offsetLeft,s=e.clientY-t.offsetTop,n=new AbortController,{signal:l}=n;f(o,{onmousemove:u((e=>{const n=parseInt(o.getAttribute(\"snap\")||\"0\"),{offsetLeft:l,offsetTop:r,offsetWidth:a,offsetHeight:c}=t,d=l+a,h=r+c,p=e.clientX-i,u=e.clientY-s,m=p+a,w=u+c;let g=0,x=0;if(n>0){l>=0&&p<0&&p>=-n?g=-p:d<=o.offsetWidth&&m>o.offsetWidth&&m<=o.offsetWidth+n&&(g=o.offsetWidth-m),r>=0&&u<0&&u>=-n?x=-u:h<=o.offsetHeight&&w>o.offsetHeight&&w<=o.offsetHeight+n&&(x=o.offsetHeight-w);for(const e of o.childNodes)if(e instanceof rt&&e!==t){const t=e.offsetLeft,o=e.offsetTop,i=t+e.offsetWidth,s=o+e.offsetHeight;u<=s&&w>=o&&(d<=t&&m>=t&&m<=t+n?g=t-m:l>=i&&p<=i&&p>=i-n&&(g=i-p)),p<=i&&m>=t&&(h<=o&&w>=o&&w<=o+n?x=o-w:r>=s&&u<=s&&u>=s-n&&(x=s-u))}}f(t,{style:{\"--window-left\":p+g+\"px\",\"--window-top\":u+x+\"px\"}})}),0,l),onmouseup:u((e=>{0===e.button&&(f(o,{style:{\"user-select\":void 0}}),n.abort(),it=!1,t.dispatchEvent(new CustomEvent(\"moved\")))}),0,l)})})(this,t),ondblclick:t=>{t.target instanceof HTMLButtonElement||this.hasAttribute(\"hide-maximise\")||this.toggleAttribute(\"maximised\")}},[this.#l=E({part:\"icon\",src:ht}),this.#n=M({part:\"title\"}),y({part:\"controls\"},[b({part:\"close\",title:q.CLOSE,onclick:()=>this.close()}),this.#h=b({part:\"maximise\",title:q.MAXIMISE,onclick:()=>this.toggleAttribute(\"maximised\")}),b({part:\"minimise\",title:q.MINIMISE,onclick:()=>this.toggleAttribute(\"minimised\")}),this.#r=M()])]),this.#a=y($()),y({onclick:t})]).adoptedStyleSheets=[et],setTimeout(f,0,this,{onmousedown:u(t,2)})}connectedCallback(){ot===this&&(ot=null)}disconnectedCallback(){if(ot===this)return;const t=this.#d,e=this.#c;t&&(f(t.#a,{class:[\"!hasChild\"]}),t.#c=null,this.#d=null),e&&(e.remove(),this.#c=null),this.dispatchEvent(new CustomEvent(\"remove\"))}attributeChangedCallback(t,e,o){switch(t){case\"window-title\":w(this.#n,{title:o??\"\"},o??\"\");break;case\"window-icon\":f(this.#l,{src:o??ht});break;case\"maximised\":f(this.#h,{title:null===o?q.MAXIMISE:q.RESTORE})}}static get observedAttributes(){return[\"maximised\",\"window-icon\",\"window-title\"]}addWindow(t){return!!this.parentNode&&(this.#c?(this.#c.addWindow(t),!0):(this.#c=t,t.#d=this,f(this.#a,{class:[\"hasChild\"]}),f(this.parentNode,t),!0))}addControlButton(t,e,o){const i=b({style:{\"background-image\":`url(${JSON.stringify(t)})`},onclick:()=>e.call(this),title:o});return f(this.#r,i),()=>i.remove()}focus(){this.toggleAttribute(\"minimised\",!1);const t=this.#c;if(t)t.focus();else{if(this.parentNode&&this.nextElementSibling){ot=this;const t=[],e=document.createNodeIterator(this,NodeFilter.SHOW_ELEMENT);for(let o=this.#a;o;o=e.nextNode()){const{scrollTop:e,scrollLeft:i}=o;(e||i)&&t.push([o,e,i])}f(this.parentNode,this);for(const[e,o,i]of t)e.scrollTop=o,e.scrollLeft=i}super.focus()}}close(){if(this.#c)this.#c.focus();else if(this.dispatchEvent(new CustomEvent(\"close\",{cancelable:!0})))return this.remove(),!0;return!1}}customElements.define(\"windows-shell\",nt),customElements.define(\"windows-desktop\",lt),customElements.define(\"windows-window\",rt);const at=p(x,\"windows-shell\"),ct=p(x,\"windows-desktop\"),dt=p(x,\"windows-window\");let ht=`data:image/svg+xml,%3Csvg viewBox=\"0 0 14 18\" xmlns=\"${O}\"%3E%3Cpath d=\"M9,1 h-8 v16 h12 v-12 Z v4 h4\" stroke=\"black\" fill=\"none\" /%3E%3Cpath d=\"M3,8 h8 m-8,3 h8 m-8,3 h8\" stroke=\"gray\" /%3E%3C/svg%3E`;const ft=["
	mid   = "],pt=["
	mid2  = "],ut=["
	mid3  ="],mt=["
	mid4  = "],wt=["
	end   = "];let gt=20,xt=0;const vt=Array.from({length:pt.length},(()=>!0)),bt=Array.from({length:mt.length},(()=>!1)),yt=[[30,\"#0f0\"],[80,\"#ff0\"],[240,\"#fa0\"],[1/0,\"#f00\"]],kt=y({id:\"mouseLine\"}),Et=y({id:\"mouseLine2\"}),Ct=y({style:{\"background-color\":\"#fff\",border:\"1px solid #000\",position:\"absolute\",top:0,\"z-index\":10}}),St=y({class:\"sticky\"}),$t=at(ct(St)),Mt=t=>(t<10?\"0\":\"\")+t,At=t=>{const e=new Date(1e3*t);return`${e.getFullYear()}-${Mt(e.getMonth()+1)}-${e.getDate()} ${Mt(e.getHours())}:${Mt(e.getMinutes())}:${Mt(e.getSeconds())}`},Lt=t=>`${Math.floor(t/3600)}:${Mt(Math.floor(t%3600/60))}:${Mt(t%60)}`,Tt=(t,e)=>{const o=\"ID_\"+xt++;return[S({for:o},t),f(e,{id:o})]},Nt=(new Intl.Collator).compare,Ht=(t,e)=>Nt(t[0],e[0]),It=V({viewBox:\"0 0 20 20\",onclick:()=>{$t.addWindow(zt),zt.focus()}},[P(j({id:\"spoke\",d:\"M1,7 v2 a1,1 0,0,1 -2,0 v-2 z\",fill:\"#aaa\"})),D({transform:\"translate(10, 10)\"},[R({r:5.5,fill:\"none\",stroke:\"#aaa\",\"stroke-width\":4.5}),Array.from({length:12},((t,e)=>30*e)).map((t=>_({href:\"#spoke\",transform:`rotate(${t})`})))])]),Wt=H({rowspan:2},y(It)),zt=dt({\"window-icon\":(Rt=It,\"data:image/svg+xml,\"+encodeURIComponent('<svg xmlns=\"'+O+'\"'+(Rt instanceof SVGSVGElement?Rt.outerHTML.slice(4):Rt.outerHTML.slice(7,-7)+\"svg>\"))).replaceAll(\"aaa\",\"000\"),\"window-title\":\"Settings\",tabindex:-1,onkeydown:t=>{\"Escape\"===t.key&&zt.close()}},[k(\"Settings\"),Tt(\"Timeline Scale (pixels per minute): \",C({type:\"number\",min:1,value:gt,onchange:function(){gt=parseInt(this.value)}})),v(),Tt(\"Floating Logtime: \",C({type:\"checkbox\",checked:!0,onchange:function(){f(St,{class:this.checked?\"sticky\":void 0})}})),v(),v(),L([W([H(y({style:{\"text-decoration\":\"underline\"}},\"Toggle Users:\")),H(y({style:{\"text-decoration\":\"underline\"}},\"Highlight Lines:\"))]),W([N(pt.map(((t,e)=>[t,e])).sort(Ht).map((([t,e])=>[Tt(`${t}: `,C({type:\"checkbox\",checked:!0,onclick:function(){vt[e]=this.checked}})),v()]))),N(mt.map(((t,e)=>[t,e])).sort(Ht).map((([t,e])=>[Tt(`${t}: `,C({type:\"checkbox\",onclick:function(){bt[e]=this.checked}})),v()])))])]),b({onclick:()=>Ot(ft)},\"Rebuild\")]),Ot=t=>{const e=new Map,o=T(),i=[],s=t=>{const e=t.offsetX+(t.target instanceof HTMLDivElement?t.target.offsetLeft:0);f(kt,{style:{left:Wt.offsetWidth+e+\"px\"}}),f(Et,{style:{left:e+\"px\"}}),w(Ct,{style:{left:Wt.offsetWidth+e}},At(l+60*(e+2)/gt))},n=[];let l=1/0,r=-1/0,a=0;for(const n of t){const[t,a,c,d]=n;if(!vt[t])continue;let h,p,u,m=!1;if(e.has(t))[,,h,p,u]=e.get(t);else{const i=H(y(pt[t])),n=N({onmousemove:s});e.set(t,[i,n,h=[],p=[-1,0],u=[]]),f(o,W([i,n]))}a>p[0]?(p[0]>-1&&u.push([p[0],a]),p[1]+=c-a,p[0]=c):c>p[0]&&(p[1]+=c-p[0],p[0]=c);t:for(const t of h){for(const[,e,o]of t)if(a<o&&c>e)continue t;m=!0,t.push(n);break}if(m||h.push([n]),m=!1,d>0){t:for(const t of i){for(const[,e,,o]of t)if(d<e&&a>o)continue t;m=!0,t.push(n);break}m||i.push([n]),d<l&&(l=d)}a<l&&(l=a),c>r&&(r=c)}l=60*Math.floor(l/60),r=60*Math.ceil(r/60);for(const[,[t,o,i,[,s],n]]of e){let e=0,r=0,c=0;for(const[t,e]of n)f(o,y({class:\"wait\",title:`Down Time: ${Lt(e-t)}`,style:{left:gt*(t-l)/60-2+\"px\",width:gt*(e-t)/60+1+\"px\"}}));for(const t of i){for(const[,i,s,,n,a,d]of t)f(o,y({class:{h:bt[n]},title:`${At(i)} ⟶  ${At(s)}\\nCall Time: ${Lt(s-i)}${d?`\\n${ut[d]}`:\"\"}\\n${wt[a]}`,style:{width:gt*(s-i)/60+1+\"px\",left:gt*(i-l)/60-2+\"px\",\"--row\":e}})),r++,c+=s-i;e++}f(t,{title:`Total Calls: ${r}\\nTotal Call Time: ${Lt(c)}\\nNon-concurrent Call Time: ${Lt(s)}`}),a=Math.max(i.length,a)}f(kt,{style:{\"--h\":a*e.size+i.length}}),f(document.body,{style:{\"--rows\":a}});for(let t=3600*Math.ceil(l/3600);t<=r;t+=3600){const e=new Date(1e3*t);n.push(Y({x:gt*(t-l)/60,y:15},`${Mt(e.getHours())}:00`))}f(Ct,{style:{left:\"-1000px\"}}),w(St,L([I({style:{\"--rows\":i.length}},[W([Wt,N({style:{width:gt*(r-l)/60+\"px\"}},[V({width:gt*(r-l)/60,height:20,viewBox:`0 0 ${gt*(r-l)/60} 20`},[P(F({id:\"ticks\",patternUnits:\"userSpaceOnUse\",width:60*gt,height:20,x:-l%3600/60*gt-2},[B({y2:20,stroke:\"#000\"}),B({x1:15*gt,x2:15*gt,y2:10,stroke:\"#000\"}),B({x1:30*gt,x2:30*gt,y2:10,stroke:\"#000\"}),B({x1:45*gt,x2:45*gt,y2:10,stroke:\"#000\"})])),X({width:\"100%\",height:\"100%\",fill:\"url(#ticks)\"}),n]),Ct])]),W(N({onmousemove:s},[Et,i.map(((t,e)=>t.map((([t,o,,i,s,,n])=>y({class:{h:bt[s]},title:`${pt[t]}\\n${At(i)} ⟶  ${At(o)}\\nWait Time: ${Lt(o-i)}${n?`\\n${ut[n]}`:\"\"}`,style:{width:gt*(o-i)/60+1+\"px\",left:gt*(i-l)/60-2+\"px\",\"--row\":e,color:yt.find((([t])=>t>o-i))[1]}})))))]))]),o]))};var Rt;ft.sort((([,t],[,e])=>t-e)),z.then((()=>{w(document.body,[$t,kt]),Ot(ft)}));</script><style type=\"text/css\">html,body{margin:0;background-color:#fff;user-select:none}text{text-anchor:middle}table{table-layout:fixed;border-collapse:collapse;border-spacing:0;margin:0;white-space:nowrap;padding:0}td,th{padding:0;margin:0}thead th{text-align:center;background-color:#fff}.sticky thead{position:sticky;top:0;z-index:30}thead{background-color:#fff}thead th svg{width:1em;cursor:pointer}thead td:first-child{background-color:#eee}thead td:first-child,tbody td{position:relative;height:calc(var(--rows) * 2em)}thead td:first-child div,tbody td div{position:absolute;top:calc(var(--row) * 2em + 3px);height:calc(2em - 6px);background-color:currentColor;z-index:10}tbody td div{border:1px solid #000;box-sizing:border-box}tbody tr th{background-color:#ff0}tbody tr:nth-child(2n) th{background-color:#ee0}tbody tr:nth-child(2n){background-color:#ddd}tbody td{color:#8cf}th{position:sticky;left:0;z-index:20}.wait{top:0;height:calc(var(--rows) * 2em);border-width:0;background-color:transparent}#mouseLine,#mouseLine2{border-left:1px solid red;position:absolute;top:1px;z-index:0;left:-10px;height:calc(2em * var(--h) + 21px)}#mouseLine2{height:calc(2em * var(--rows))}windows-shell{position:static;overflow:visible}windows-window{z-index:40;position:fixed;outline:0}windows-window table th:first-child{padding-right:5em}tbody .h{background-color:#a2f}thead .h{border:1px solid #000;box-sizing:border-box;z-index:11 !important}</style></head><body></body></html>"
)
