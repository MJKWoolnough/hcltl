package main

// File automatically generated with ./embed.sh

const (
	htmlStart = "<html><head><title>HCL Timeline</title><link rel=\"shortcut icon\" sizes=\"any\" href=\"data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 256 256'%3E%3Crect x='15' y='15' height='45' width='175' rx='5' fill='%23f00' /%3E%3Crect x='15' y='75' height='45' width='90' rx='5' fill='%230f0' /%3E%3Crect x='160' y='75' height='45' width='85' rx='5' fill='%230f0' /%3E%3Crect x='70' y='135' height='45' width='125' rx='5' fill='%2300f' /%3E%3Crect x='50' y='195' height='45' width='70' rx='5' fill='%23000' /%3E%3Crect x='170' y='195' height='45' width='50' rx='5' fill='%23000' /%3E%3C/svg%3E\" /><script type=\"module\">"
	htmlEnd   = "</script><style type=\"text/css\">html,body{margin:0;background-color:#fff}text{text-anchor:middle}table{table-layout:fixed;border:1px solid #000;border-collapse:collapse;border-spacing:0;margin:0;white-space:nowrap;padding:0}td,th{padding:0;margin:0}thead td{border:1px solid #000;text-align:center}thead th{position:sticky;top:0;z-index:2}thead td svg{width:1em;cursor:pointer}thead th svg{background-color:#fff;border-bottom:1px solid #000}tbody th{position:sticky;left:0;z-index:2}tbody th div{z-index:4;position:sticky;left:0;background-color:#ff0;border-right:1px solid #000;height:calc(2em * var(--rows));display:flex;justify-content:center;align-items:center}tbody tr:nth-child(2n) th div{background-color:#ee0}tbody td{position:relative;margin:0;padding:0}tbody td div{background-color:currentColor;position:absolute;z-index:1;height:calc(2em - 6px);margin:0;padding:0;top:calc(3px + var(--row) * 2em);box-sizing:border-box;border:1px solid #000}tbody tr:nth-child(2n){background-color:#ddd}#mouseLine{border-left:1px solid red;position:absolute;top:1px;z-index:0;height:calc(2em * var(--h) + 21px)}windows-window{z-index:20;position:fixed}</style></head><body></body></html>"
	jsStart   = "const t=(e,o)=>{if(\"string\"==typeof o)e.appendChild(document.createTextNode(o));else if(Array.isArray(o))for(const i of o)t(e,i);else if(o instanceof Node)e.appendChild(o);else if(o instanceof NodeList)for(const t of o)e.appendChild(t)},e=e=>(o,i,s)=>{const n=\"string\"==typeof o?document.createElementNS(e,o):o instanceof Node?o:document.createDocumentFragment();if(!(\"string\"==typeof i||i instanceof Array||i instanceof NodeList||i instanceof Node)&&(\"object\"!=typeof s||s instanceof Array||s instanceof Node||s instanceof NodeList)||([i,s]=[s,i]),\"object\"==typeof i&&n instanceof Element)for(const[t,e]of Object.entries(i))if(e instanceof Function){const o={};let i=t;t:for(;;){switch(i.charAt(0)){case\"1\":o.once=!0;break;case\"C\":o.capture=!0;break;case\"P\":o.passive=!0;break;default:break t}i=i.slice(1)}i.startsWith(\"on\")&&n.addEventListener(i.substr(2),e,o)}else if(\"class\"===t&&(e instanceof Array||e instanceof DOMTokenList)&&e.length>0)n.classList.add(...e);else if(\"style\"===t&&\"object\"==typeof e&&(n instanceof HTMLElement||n instanceof SVGElement))for(const t in e)void 0===e[t]?n.style.removeProperty(t):n.style.setProperty(t,e[t]);else\"string\"==typeof e||\"number\"==typeof e?n.setAttribute(t,e):\"boolean\"==typeof e?n.toggleAttribute(t,e):void 0===e&&n.hasAttribute(t)&&n.removeAttribute(t);return\"string\"==typeof s?n.textContent=s:s&&(s instanceof Array||s instanceof Node||s instanceof NodeList)&&t(n,s),n},o=e(\"http://www.w3.org/1999/xhtml\"),i=e(\"http://www.w3.org/2000/svg\"),s=t=>{for(;null!==t.lastChild;)t.removeChild(t.lastChild);return t},n=(t,e=!0)=>(window.setTimeout((()=>{t.focus(),(t instanceof HTMLInputElement||t instanceof HTMLTextAreaElement)&&e&&t.select()}),0),t),[r,h,d,l,a,c,p,f,w,m,u,v,g,x,b,y]=\"br button div h1 img input label span slot style table tbody td th thead tr\".split(\" \").map((t=>o.bind(null,t))),[k,E,C,L,P,M,A,T,z,$]=\"circle defs g line path pattern rect svg text use\".split(\" \").map((t=>i.bind(null,t))),H=new WeakMap,W=function(t,e){const o=this.parentNode;if(X||!(o instanceof Y)||0!==e.button)return;X=!0,o.style.setProperty(\"user-select\",\"none\");const i=this.offsetLeft,s=this.offsetTop,n=this.offsetWidth,r=this.offsetHeight,h=e.clientX,d=e.clientY,l=e=>{const o=e.clientX-h,l=e.clientY-d;switch(t){case 0:case 1:case 2:{const t=r-l;t>(parseInt(this.style.getPropertyValue(\"min-height\")||\"\")||100)&&(this.style.setProperty(\"--window-top\",`${s+l}px`),this.style.setProperty(\"--window-height\",`${t}px`))}break;case 4:case 5:case 6:{const t=r+l;t>(parseInt(this.style.getPropertyValue(\"min-height\")||\"\")||100)&&this.style.setProperty(\"--window-height\",`${t}px`)}}switch(t){case 2:case 3:case 4:{const t=n+o;t>(parseInt(this.style.getPropertyValue(\"min-width\")||\"\")||100)&&this.style.setProperty(\"--window-width\",`${t}px`)}break;case 0:case 7:case 6:{const t=n-o;t>(parseInt(this.style.getPropertyValue(\"min-width\")||\"\")||100)&&(this.style.setProperty(\"--window-left\",`${i+o}px`),this.style.setProperty(\"--window-width\",`${t}px`))}}},a=t=>{0===t.button&&(o.removeEventListener(\"mousemove\",l),o.removeEventListener(\"mouseup\",a),o.style.removeProperty(\"user-select\"),X=!1,this.dispatchEvent(new CustomEvent(\"resized\")))};o.addEventListener(\"mousemove\",l),o.addEventListener(\"mouseup\",a)},N=function(t){const e=this.parentNode;if(X||!(e instanceof Y)||0!==t.button)return;X=!0,this.style.setProperty(\"user-select\",\"none\");const o=t.clientX-this.offsetLeft,i=t.clientY-this.offsetTop,s=t=>{const s=t.clientX-o,n=t.clientY-i,[r,h]=((t,e,o,i)=>{const s=parseInt(t.getAttribute(\"snap\")||\"0\"),{offsetLeft:n,offsetTop:r,offsetWidth:h,offsetHeight:d}=e,l=n+h,a=r+d,c=o+h,p=i+d,f=[0,0];return s>0&&(n>=0&&o<0&&o>=-s?f[0]=-o:l<=t.offsetWidth&&c>t.offsetWidth&&c<=t.offsetWidth+s&&(f[0]=t.offsetWidth-c),r>=0&&i<0&&i>=-s?f[1]=-i:a<=t.offsetHeight&&p>t.offsetHeight&&p<=t.offsetHeight+s&&(f[1]=t.offsetHeight-p),Array.from(t.childNodes).filter((t=>t instanceof F&&t!==e)).forEach((t=>{const e=t.offsetLeft,h=t.offsetTop,d=e+t.offsetWidth,w=h+t.offsetHeight;i<=w&&p>=h&&(l<=e&&c>=e&&c<=e+s?f[0]=e-c:n>=d&&o<=d&&o>=d-s&&(f[0]=d-o)),o<=d&&c>=e&&(a<=h&&p>=h&&p<=h+s?f[1]=h-p:r>=w&&i<=w&&i>=w-s&&(f[1]=w-i))}))),f})(e,this,s,n);this.style.setProperty(\"--window-left\",s+r+\"px\"),this.style.setProperty(\"--window-top\",n+h+\"px\")},n=t=>{0===t.button&&(e.removeEventListener(\"mousemove\",s),e.removeEventListener(\"mouseup\",n),e.style.removeProperty(\"user-select\"),X=!1,this.dispatchEvent(new CustomEvent(\"moved\")))};e.addEventListener(\"mousemove\",s),e.addEventListener(\"mouseup\",n)},S=new Map,B=new Map,I=(t,e,o,i)=>new Promise(((s,r)=>{const l=U({\"window-hide\":!0,\"window-icon\":i||R,\"window-title\":e,\"hide-maximise\":\"true\",onremove:()=>s(!1)},[d(o),d({style:\"text-align: center\"},n(h({onclick:()=>{s(!0),l.remove()}},\"Ok\")))]);t.addWindow(l)||r(new Error(\"invalid target\"))})),D=(t,e,o,i)=>new Promise(((s,r)=>{const l=U({\"window-hide\":!0,\"window-icon\":i||R,\"window-title\":e,\"hide-maximise\":\"true\",onremove:()=>s(!1)},[d(o),d({style:\"text-align: center\"},[n(h({onclick:()=>{s(!0),l.remove()}},\"Ok\")),h({onclick:()=>l.remove()},\"Cancel\")])]);t.addWindow(l)||r(new Error(\"invalid target\"))})),O=(t,e,o,i,s)=>new Promise(((r,l)=>{const a=h({onclick:()=>{r(p.value),f.remove()}},\"Ok\"),p=n(c({value:i||\"\",onkeydown:t=>{switch(t.key){case\"Enter\":a.click();break;case\"Escape\":f.remove()}}})),f=U({\"window-hide\":!0,\"window-icon\":s||R,\"window-title\":e,\"hide-maximise\":\"true\",onremove:()=>r(null)},[d(o),p,d({style:\"text-align: center\"},a)]);t.addWindow(f)||l(new Error(\"invalid target\"))}));let V=null,X=!1;class Y extends HTMLElement{constructor(){super(),new.target===Y&&o(this.attachShadow({mode:\"closed\"}),[m({type:\"text/css\"},\":host{display:block;position:relative;overflow:hidden;width:var(--shell-width,100%);height:var(--shell-height,100%)}::slotted(windows-window){--taskmanager-on:none}::slotted(windows-window:last-of-type){--overlay-on:none}\"),w({name:\"desktop\"}),d(w())])}alert(t,e,o){return I(this,t,e,o)}confirm(t,e,o){return D(this,t,e,o)}prompt(t,e,o,i){return O(this,t,e,o,i)}addWindow(t){return this.appendChild(t),!0}realignWindows(){const{offsetWidth:t,offsetHeight:e}=this;Array.from(this.childNodes).filter((t=>t instanceof F)).forEach((o=>{const{offsetLeft:i,offsetTop:s,offsetWidth:n,offsetHeight:r}=o;i+n>t?o.style.setProperty(\"--window-left\",Math.max(t-n,0)+\"px\"):i<0&&o.style.setProperty(\"--window-left\",\"0px\"),s+r>e?o.style.setProperty(\"--window-top\",Math.max(e-r)+\"px\"):s<0&&o.style.setProperty(\"--window-top\",\"0px\")}))}}class j extends HTMLElement{constructor(){super(),this.setAttribute(\"slot\",\"desktop\"),o(this.attachShadow({mode:\"closed\"}),[m({type:\"text/css\"},\":host{position:absolute;top:0;left:0;bottom:0;right:0}\"),w({slot:\"desktop\"})])}}class F extends HTMLElement{constructor(){super();const t=f({part:\"title\"}),e=a({part:\"icon\",src:R}),i=f(),s=()=>this.focus();H.set(this,{titleElement:t,iconElement:e,extraButtons:i}),o(this.attachShadow({mode:\"closed\"}),[m({type:\"text/css\"},':host{position:absolute;display:block;background-color:#fff;color:#000;border:1px solid #000;width:var(--window-width,auto);height:var(--window-height,auto);min-width:100px;min-height:100px;top:var(--window-top,0);left:var(--window-left,0);list-style:none;padding:0;z-index:0}:host([maximised]){left:0;right:0;top:0;bottom:0;width:auto;height:auto}:host([resizable]){border-width:0}:host([resizable])>div:nth-child(2){display:block}:host([minimised]),:host>div:nth-child(2),:host([hide-titlebar])>div:nth-child(3),:host([hide-close])>div:nth-child(3)>div>button:nth-of-type(1),:host([hide-maximise])>div:nth-child(3)>div>button:nth-of-type(2),:host([hide-minimise])>div:nth-child(3)>div>button:nth-of-type(3),:host([window-hide])>div:nth-child(3)>div>button:nth-of-type(3){display:none}:host>div:nth-child(3){white-space:nowrap;height:calc(1em + 6px);background-color:#aaa;overflow:hidden;user-select:none;display:flex;align-items:center}:host>div:nth-child(3)>span{margin-right:calc(3em + 24px)}:host>div:nth-child(3)>img{height:1em;pointer-events:none}:host>div:nth-child(3)>div{position:absolute;right:0;top:0}:host>div:nth-child(3) button{padding:0;border-width:2px;float:right;background-repeat:no-repeat;background-position:center;background-color:#eee;background-size:1em 1em;width:calc(1em + 8px);height:calc(1em + 8px)}:host>div:nth-child(3)>div>button:nth-of-type(1){background-image:url(\\'data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cpath d=\"M1,1 L9,9 M9,1 L1,9\" stroke=\"black\" /%3E%3C/svg%3E\\')}:host>div:nth-child(3)>div>button:nth-of-type(2){background-image:url(\\'data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cpath d=\"M9,3 h-8 v-1 h8 v-1 h-8 v8 h8 v-8\" stroke=\"black\" fill=\"none\" /%3E%3C/svg%3E\\')}:host([maximised])>div:nth-child(3)>div>button:nth-of-type(2){background-image:url(\\'data:image/svg+xml,%3Csvg viewBox=\"0 0 15 13\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cpath d=\"M1,5 h8 v-1 h-8 v8 h8 v-8 m-3,0 v-3 h8 v8 h-5 m5,-7 h-8\" stroke=\"%23000\" fill=\"none\" /%3E%3C/svg%3E\\')}:host>div:nth-child(3)>div>button:nth-of-type(3){display:var(--taskmanager-on,block);background-image:url(\\'data:image/svg+xml,%3Csvg viewBox=\"0 0 10 10\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cline x1=\"1\" y1=\"9\" x2=\"9\" y2=\"9\" stroke=\"black\" /%3E%3C/svg%3E\\')}:host>div:nth-child(4){user-select:contain}:host([resizable])>div:nth-child(4){overflow:auto;position:absolute;bottom:0;left:0;right:0;top:calc(1em + 6px)}:host>div:nth-child(5){pointer-events:none;display:var(--overlay-on,block);position:absolute;background-color:rgba(0,0,0,0.1);top:0;left:0;bottom:0;right:0}:host([resizeable])>div:nth-child(5){top:var(calc(--window-resize), -2px);left:var(calc(--window-resize), -2px);bottom:var(calc(--window-resize), -2px);right:var(calc(--window-resize), -2px)}:host>div:nth-child(2)>div{position:absolute;border-color:currentColor;border-style:solid;border-width:0;z-index:-1}:host>div:nth-child(2)>div:nth-child(1){top:-2px;left:-2px;width:10px;height:10px;cursor:nwse-resize;border-left-width:3px;border-top-width:3px}:host>div:nth-child(2)>div:nth-child(2){top:-2px;left:8px;right:8px;border-top-width:3px;cursor:ns-resize}:host>div:nth-child(2)>div:nth-child(3){top:-2px;right:-2px;width:10px;height:10px;border-top-width:3px;border-right-width:3px;cursor:nesw-resize}:host>div:nth-child(2)>div:nth-child(4){top:8px;right:-2px;bottom:8px;border-right-width:3px;cursor:ew-resize}:host>div:nth-child(2)>div:nth-child(5){bottom:-2px;right:-2px;width:10px;height:10px;border-right-width:3px;border-bottom-width:3px;cursor:nwse-resize}:host>div:nth-child(2)>div:nth-child(6){bottom:-2px;left:8px;right:8px;border-bottom-width:3px;cursor:ns-resize}:host>div:nth-child(2)>div:nth-child(7){bottom:-2px;left:-2px;width:10px;height:10px;border-left-width:3px;border-bottom-width:3px;cursor:nesw-resize}:host>div:nth-child(2)>div:nth-child(8){top:8px;left:-2px;bottom:8px;border-left-width:3px;cursor:ew-resize}'),d(Array.from({length:8},((t,e)=>d({onmousedown:W.bind(this,e)})))),d({part:\"titlebar\",onmousedown:N.bind(this),ondblclick:t=>{t.target instanceof HTMLButtonElement||this.hasAttribute(\"hide-maximise\")||this.toggleAttribute(\"maximised\")}},[e,t,d({part:\"controls\"},[h({part:\"close\",onclick:()=>this.close()}),h({part:\"maximise\",onclick:()=>this.toggleAttribute(\"maximised\")}),h({part:\"minimise\",onclick:()=>this.toggleAttribute(\"minimised\")}),i])]),d(w()),d({onclick:s})]),this.addEventListener(\"mousedown\",s,{capture:!0})}connectedCallback(){V===this&&(V=null)}disconnectedCallback(){if(V===this)return;const t=B.get(this),e=S.get(this);t&&(S.delete(t),B.delete(this)),e&&(e.remove(),S.delete(this)),this.dispatchEvent(new CustomEvent(\"remove\",{cancelable:!1}))}attributeChangedCallback(t,e,o){const i=H.get(this);switch(t){case\"window-title\":i.titleElement.innerText=o;break;case\"window-icon\":i.iconElement.setAttribute(\"src\",o)}}static get observedAttributes(){return[\"window-title\",\"window-icon\"]}alert(t,e,o){return I(this,t,e,o)}confirm(t,e,o){return D(this,t,e,o)}prompt(t,e,o,i){return O(this,t,e,o,i)}addWindow(t){return!!this.parentNode&&(S.has(this)?(S.get(this).addWindow(t),!0):(S.set(this,t),B.set(t,this),this.parentNode.appendChild(t),!0))}addControlButton(t,e,o){const i=H.get(this).extraButtons.appendChild(h({style:{\"background-image\":`url(${t})`},onclick:()=>e.call(this),title:o}));return()=>i.remove()}focus(){const t=S.get(this);t?t.focus():(this.parentNode&&this.nextElementSibling&&(V=this,this.parentNode.appendChild(this)),super.focus())}close(){if(S.has(this))S.get(this).focus();else if(this.dispatchEvent(new CustomEvent(\"close\",{cancelable:!0})))return this.remove(),!0;return!1}}customElements.define(\"windows-shell\",Y),customElements.define(\"windows-desktop\",j),customElements.define(\"windows-window\",F);const U=(t,e)=>o(new F,t,e),R='data:image/svg+xml,%3Csvg viewBox=\"0 0 14 18\" xmlns=\"http://www.w3.org/2000/svg\"%3E%3Cpath d=\"M9,1 h-8 v16 h12 v-12 Z v4 h4\" stroke=\"black\" fill=\"none\" /%3E%3Cpath d=\"M3,8 h8 m-8,3 h8 m-8,3 h8\" stroke=\"gray\" /%3E%3C/svg%3E',G=["
	jsMid     = "],Z=["
	jsEnd     = "];let _=20;const q=Array.from({length:Z.length},(()=>!0)),J=[\"#29f\",\"#f43\",\"#fe3\",\"#4a5\",\"#92b\",\"#0bd\",\"#f52\",\"#cd3\",\"#35b\",\"#e16\",\"#098\"],K=d({id:\"mouseLine\"}),Q=d({style:{\"background-color\":\"#fff\",border:\"1px solid #000\",position:\"absolute\",top:0,\"z-index\":10}},\"TIME\"),tt=d(),et=o(new Y,void 0,void 0),ot=t=>{const e=t+\"\";return 1==e.length?\"0\"+e:e},it=t=>{const e=new Date(1e3*t);return`${e.getFullYear()}-${ot(e.getMonth()+1)}-${e.getDate()} ${ot(e.getHours())}:${ot(e.getMinutes())}:${ot(e.getSeconds())}`},st=g(T({viewBox:\"0 0 20 20\",onclick:()=>et.addWindow(nt)},[E(P({id:\"spoke\",d:\"M1,7 v2 a1,1 0,0,1 -2,0 v-2 z\",fill:\"#aaa\"})),C({transform:\"translate(10, 10)\"},[k({r:5.5,fill:\"none\",stroke:\"#aaa\",\"stroke-width\":4.5}),Array.from({length:12},((t,e)=>30*e)).map((t=>$({href:\"#spoke\",transform:`rotate(${t})`})))])])),nt=U({\"window-icon\":\"data:image/svg+xml,\"+encodeURIComponent('<svg xmlns=\"http://www.w3.org/2000/svg\"'+st.outerHTML.slice(4).replaceAll(\"#aaa\",\"#000\")),\"window-title\":\"Settings\"},[l(\"Settings\"),p({for:\"scale\"},\"Timeline Scale (pixels per minute): \"),c({id:\"scale\",type:\"number\",min:1,value:_,onchange:function(){_=parseInt(this.value)}}),r(),Z.map(((t,e)=>[p({for:`user_${e}`},`${t}: `),c({type:\"checkbox\",checked:!0,onclick:function(){q[e]=this.checked}}),r()])),h({onclick:()=>rt(G)},\"Rebuild\")]),rt=t=>{const e=new Map,i=v(),n=t=>{const e=t.offsetX+(t.target instanceof HTMLDivElement?t.target.offsetLeft:0);K.style.setProperty(\"left\",st.offsetWidth+e+\"px\"),Q.style.setProperty(\"left\",e-1+\"\"),Q.innerText=it(r+60*(e+1)/_)};let r=1/0,h=-1/0,l=0;for(const o of t){const[t,s,a]=o;if(!q[t])continue;let c,p=!1;if(e.has(t))[,,c]=e.get(t);else{const o=d(Z[t]),s=g({style:{color:J[e.size%J.length]},onmousemove:n});e.set(t,[o,s,c=[]]),i.appendChild(y([x(o),s]))}t:for(const t of c){for(const[,e,o]of t)if(s<o&&a>e)continue t;p=!0,t.push(o);break}p||(c.push([o]),l++),s<r&&(r=s),a>h&&(h=a)}K.style.setProperty(\"--h\",l+\"\"),r=60*Math.floor(r/60),h=60*Math.ceil(h/60);for(const[,[t,o,i]]of e){let e=0,s=0,n=0;for(const t of i){for(const[,i,h]of t)o.appendChild(d({title:it(i)+\" ⟶ \"+it(h),style:{width:_*(h-i)/60+1+\"px\",left:_*(i-r)/60-2+\"px\",\"--row\":e}})),s++,n+=h-i;e++}t.style.setProperty(\"--rows\",e+\"\"),t.setAttribute(\"title\",`Total Calls: ${s}\\nTotal Call Time: ${Math.floor(n/3600)}:${ot(Math.floor(n%3600/60))}:${ot(n%60)}`)}const a=[];for(let t=3600*Math.ceil(r/3600);t<=h;t+=3600){const e=new Date(1e3*t);a.push(z({x:_*(t-r)/60,y:15},`${ot(e.getHours())}:00`))}Q.style.setProperty(\"left\",\"-1000px\"),o(s(tt),u([b(y([st,x({style:{width:_*(h-r)/60+\"px\"}},[T({width:_*(h-r)/60,height:20,viewBox:`0 0 ${_*(h-r)/60} 20`},[E(M({id:\"ticks\",patternUnits:\"userSpaceOnUse\",width:60*_,height:20,x:-r%3600/60*_-2},[L({y2:20,stroke:\"#000\"}),L({x1:15*_,x2:15*_,y2:10,stroke:\"#000\"}),L({x1:30*_,x2:30*_,y2:10,stroke:\"#000\"}),L({x1:45*_,x2:45*_,y2:10,stroke:\"#000\"})])),A({width:\"100%\",height:\"100%\",fill:\"url(#ticks)\"}),a]),Q])])),i]))};(\"complete\"==document.readyState?Promise.resolve():new Promise((t=>globalThis.addEventListener(\"load\",t,{once:!0})))).then((()=>{o(s(document.body),[tt,K,et]),rt(G)}));"
)
